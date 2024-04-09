package lsm

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/compute"
	"github.com/apache/arrow/go/v16/arrow/ipc"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/arrow/scalar"
	"github.com/apache/arrow/go/v16/arrow/util"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/protoarrow"
)

const (
	IDColumn       = 0
	MinTSColumn    = 1
	MaxTSColumn    = 2
	ResourceColumn = 3
)

type Part struct {
	Record arrow.Record
	Size   uint64
	MinTS  uint64
	MaxTS  uint64
}

type Tree struct {
	root *Node[*Part]
	size atomic.Uint64

	bufferMu sync.RWMutex
	build    *protoarrow.Build

	nodes   []*Node[*Part]
	records []arrow.Record

	db  *badger.DB
	key [8 + 4]byte
}

func New(db *badger.DB) (*Tree, error) {
	t := &Tree{
		root:  &Node[*Part]{},
		build: protoarrow.New(memory.DefaultAllocator, &v1.Meta{}),
		db:    db,
	}
	binary.LittleEndian.PutUint32(t.key[8:], uint32(v1.RESOURCE_META))
	return t, t.load()
}

func (t *Tree) Iter(f func(*Part) error) {
	t.root.Iterate(func(n *Node[*Part]) error {
		if n.value == nil {
			return nil
		}
		return f(n.value)
	})
}

func (t *Tree) Append(meta *v1.Meta) {
	t.bufferMu.Lock()
	defer t.bufferMu.Unlock()
	t.build.Append(meta)
}

func (t *Tree) Flush() {
	t.bufferMu.Lock()
	defer t.bufferMu.Unlock()
	r := t.build.NewRecord()
	if r.NumRows() == 0 {
		r.Release()
		return
	}
	min := r.Column(MinTSColumn).(*array.Uint64).Uint64Values()
	max := r.Column(MaxTSColumn).(*array.Uint64).Uint64Values()
	t.add(&Part{
		Record: r,
		Size:   uint64(util.TotalRecordSize(r)),
		MinTS:  min[0],
		MaxTS:  max[len(max)-1],
	})
}

func (t *Tree) Start(ctx context.Context) {
	slog.Info("Starting compaction loop")
	defer func() {
		slog.Info("Exit compaction loop")
	}()
	ts := time.NewTicker(time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ts.C:
			err := t.Compact()
			if err != nil {
				slog.Error("failed compaction", "err", err)
			}
		}
	}
}

func (t *Tree) Size() uint64 {
	return t.size.Load()
}

func (t *Tree) add(part *Part) {
	t.size.Add(part.Size)
	t.root.Prepend(part)
}

func (t *Tree) Close() error {
	err := t.Compact()
	if err != nil {
		return err
	}
	t.Iter(func(p *Part) error {
		p.Record.Release()
		return nil
	})
	t.root = &Node[*Part]{}
	return nil
}

func (t *Tree) Compact() error {
	t.Flush()
	t.root.Iterate(func(n *Node[*Part]) error {
		if n.value == nil {
			return nil
		}
		t.nodes = append(t.nodes, n)
		t.records = append(t.records, n.value.Record)
		return nil
	})

	defer func() {
		clear(t.nodes)
		clear(t.records)
		t.nodes = t.nodes[:0]
		t.records = t.records[:0]
	}()

	if len(t.nodes) == 0 {
		return nil
	}
	if len(t.nodes) == 1 {
		err := t.save(t.records[0])
		if err != nil {
			t.records[0].Release()
			return err
		}
		return nil
	}
	r, err := protoarrow.Merge(t.records)
	if err != nil {
		return err
	}
	err = t.save(r)
	if err != nil {
		r.Release()
		return err
	}
	for i := range t.records {
		t.records[i].Release()
	}
	node := t.findNode(t.nodes[0])
	part := &Part{
		Record: r,
		Size:   uint64(util.TotalRecordSize(r)),
		MinTS:  t.nodes[0].value.MinTS,
		MaxTS:  t.nodes[len(t.nodes)-1].value.MaxTS,
	}
	x := &Node[*Part]{
		value: part,
	}
	for !node.next.CompareAndSwap(t.nodes[0], x) {
		node = t.findNode(t.nodes[0])
	}
	t.size.Add(-part.Size)
	return nil
}

func (t *Tree) load() error {
	err := t.db.View(func(txn *badger.Txn) error {
		it, err := txn.Get(t.key[:])
		if err != nil {
			return err
		}
		return it.Value(func(val []byte) error {
			rd, err := ipc.NewReader(bytes.NewReader(val))
			if err != nil {
				return err
			}
			defer rd.Release()
			r, err := rd.Read()
			if err != nil {
				return err
			}
			r.Retain()
			minTs := r.Column(MinTSColumn).(*array.Uint64).Uint64Values()
			maxTs := r.Column(MaxTSColumn).(*array.Uint64).Uint64Values()
			t.add(&Part{
				Record: r,
				Size:   uint64(util.TotalRecordSize(r)),
				MinTS:  minTs[0],
				MaxTS:  maxTs[len(maxTs)-1],
			})
			return nil
		})
	})
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (t *Tree) save(r arrow.Record) error {
	var b bytes.Buffer
	w := ipc.NewWriter(&b,
		ipc.WithSchema(r.Schema()),
		ipc.WithZstd(),
	)
	err := w.Write(r)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return t.db.Update(func(txn *badger.Txn) error {
		return txn.Set(t.key[:], b.Bytes())
	})
}

func (t *Tree) Scan(resource v1.RESOURCE, start, end uint64) (*Samples, error) {
	t.Flush()

	samples := NewSamples()
	err := t.root.Iterate(func(n *Node[*Part]) error {
		if n.value == nil {
			return nil
		}
		if acceptRange(n.value.MinTS, n.value.MaxTS, start, end) {
			n.value.Record.Retain()
			defer n.value.Record.Release()
			ids, err := ComputeSample(n.value.Record, uint64(resource), start, end)
			if err != nil {
				return err
			}
			for i := range ids {
				samples.Add(ids[i])
			}
			if n.value.MaxTS < end {
				return nil
			}
		}
		return io.EOF
	})

	if err != nil {
		samples.Release()
		return nil, err
	}
	return samples, nil
}

func acceptRange(minTs, maxTs uint64, start, end uint64) bool {
	return contains(minTs, maxTs, start) ||
		containsUp(minTs, maxTs, end) ||
		containsUp(start, end, minTs) ||
		containsUp(start, end, maxTs)
}

func contains(start, end, slot uint64) bool {
	return slot >= start && slot <= end
}

func containsUp(start, end, slot uint64) bool {
	return slot > start && slot < end
}

// ComputeSample returns all sample id for resource that are within start and
// end range.
func ComputeSample(r arrow.Record, resource, start, end uint64) (ids []uint64, err error) {
	ctx := context.Background()

	rsc, err := compute.CallFunction(ctx, "equal", nil, compute.NewDatumWithoutOwning(
		r.Column(ResourceColumn),
	),
		&compute.ScalarDatum{Value: &scalar.Uint64{
			Value: resource,
		}},
	)
	if err != nil {
		return nil, err
	}
	defer rsc.Release()
	if rsc.Len() == 0 {
		return []uint64{}, nil
	}
	contains, err := compute00(ctx, r, start, end)
	if err != nil {
		return nil, err
	}
	defer contains.Release()

	and, err := compute.CallFunction(ctx, "and", nil, rsc, contains)
	if err != nil {
		return nil, err
	}
	defer and.Release()

	filter := and.(*compute.ArrayDatum).MakeArray().(*array.Boolean)
	defer filter.Release()

	b := array.NewUint32Builder(memory.DefaultAllocator)
	b.Reserve(filter.Len())
	for i := 0; i < filter.Len(); i++ {
		if filter.Value(i) {
			b.Append(uint32(i))
		}
	}
	defer b.Release()
	a := b.NewArray()
	defer a.Release()

	rs, err := compute.TakeArray(ctx, r.Column(IDColumn), a)
	if err != nil {
		return nil, err
	}
	defer rs.Release()
	return rs.(*array.Uint64).Uint64Values(), nil
}

func compute00(ctx context.Context, r arrow.Record, start, end uint64) (compute.Datum, error) {
	case01, err := compute01(ctx, r, start)
	if err != nil {
		return nil, err
	}
	defer case01.Release()
	case02, err := compute02(ctx, r, start, end)
	if err != nil {
		return nil, err
	}
	defer case02.Release()
	case03, err := compute03(ctx, r, start, end)
	if err != nil {
		return nil, err
	}
	defer case03.Release()

	base, err := compute.CallFunction(ctx, "or", nil, case01, case02)
	if err != nil {
		return nil, err
	}
	defer base.Release()
	return compute.CallFunction(ctx, "or", nil, base, case03)
}

// [minTs,[start],maxTs]
//
//	minTs <= start && maxTs > start
func compute01(ctx context.Context, r arrow.Record, start uint64) (compute.Datum, error) {
	value := &compute.ScalarDatum{Value: scalar.MakeScalar(start)}
	lo, err := compute.CallFunction(ctx, "less_equal", nil,
		compute.NewDatumWithoutOwning(r.Column(MinTSColumn)), value)
	if err != nil {
		return nil, err
	}
	defer lo.Release()
	hi, err := compute.CallFunction(ctx, "greater", nil,
		compute.NewDatumWithoutOwning(r.Column(MaxTSColumn)), value)
	if err != nil {
		return nil, err
	}
	defer hi.Release()
	return compute.CallFunction(ctx, "and", nil, lo, hi)
}

// [minTs,[start...end],maxTs]
//
//	(start >= minTs && start <=maxTs) && (end > minTs && end <maxTs)
func compute02(ctx context.Context, r arrow.Record, start, end uint64) (compute.Datum, error) {
	lo, err := containsStart(ctx, r, start)
	if err != nil {
		return nil, err
	}
	defer lo.Release()
	hi, err := containsEnd(ctx, r, end)
	if err != nil {
		return nil, err
	}
	defer hi.Release()
	return compute.CallFunction(ctx, "and", nil, lo, hi)
}

func containsStart(ctx context.Context, r arrow.Record, start uint64) (compute.Datum, error) {
	value := &compute.ScalarDatum{Value: scalar.MakeScalar(start)}
	lo, err := compute.CallFunction(ctx, "less_equal", nil,
		compute.NewDatumWithoutOwning(r.Column(MinTSColumn)), value)
	if err != nil {
		return nil, err
	}
	defer lo.Release()
	hi, err := compute.CallFunction(ctx, "greater_equal", nil,
		compute.NewDatumWithoutOwning(r.Column(MaxTSColumn)),
		value)
	if err != nil {
		return nil, err
	}
	defer hi.Release()
	return compute.CallFunction(ctx, "and", nil, lo, hi)
}

func containsEnd(ctx context.Context, r arrow.Record, end uint64) (compute.Datum, error) {
	value := &compute.ScalarDatum{Value: scalar.MakeScalar(end)}
	lo, err := compute.CallFunction(ctx, "less", nil,
		compute.NewDatumWithoutOwning(r.Column(MinTSColumn)), value)
	if err != nil {
		return nil, err
	}
	defer lo.Release()
	hi, err := compute.CallFunction(ctx, "greater", nil,
		compute.NewDatumWithoutOwning(r.Column(MaxTSColumn)),
		value)
	if err != nil {
		return nil, err
	}
	defer hi.Release()
	return compute.CallFunction(ctx, "and", nil, lo, hi)
}

// [start,[minTs...maxTs],end]
//
//	minTs >= start  && maxTs < end
func compute03(ctx context.Context, r arrow.Record, start, end uint64) (compute.Datum, error) {
	value := &compute.ScalarDatum{Value: scalar.MakeScalar(start)}
	lo, err := compute.CallFunction(ctx, "greater_equal", nil,
		compute.NewDatumWithoutOwning(r.Column(MinTSColumn)), value)
	if err != nil {
		return nil, err
	}
	defer lo.Release()
	hi, err := compute.CallFunction(ctx, "less", nil,
		compute.NewDatumWithoutOwning(r.Column(MaxTSColumn)),
		&compute.ScalarDatum{Value: scalar.MakeScalar(end)})
	if err != nil {
		return nil, err
	}
	defer hi.Release()
	return compute.CallFunction(ctx, "and", nil, lo, hi)
}

func (t *Tree) findNode(node *Node[*Part]) (list *Node[*Part]) {
	t.root.Iterate(func(n *Node[*Part]) error {
		if n.next.Load() == node {
			list = n
			return io.EOF
		}
		return nil
	})
	return
}

type Samples struct {
	roaring64.Bitmap
}

func NewSamples() *Samples {
	return samplesPool.Get().(*Samples)
}

func (s *Samples) Release() {
	s.Clear()
	samplesPool.Put(s)
}

var samplesPool = &sync.Pool{New: func() any {
	return new(Samples)
}}
