package lsm

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/arena"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/meta"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/x"
	"google.golang.org/protobuf/proto"
)

const (
	PartSize = 16 << 20
)

type Tree struct {
	root *Node[*meta.Meta]
	size atomic.Uint64

	bufferMu sync.RWMutex
	build    *meta.Meta

	nodes []*Node[*meta.Meta]

	db  *badger.DB
	seq *seq.Seq
}

func New(db *badger.DB, seq *seq.Seq) (*Tree, error) {
	t := &Tree{
		root:  &Node[*meta.Meta]{},
		build: meta.New(),
		db:    db,
		seq:   seq,
	}
	return t, t.load()
}

func (t *Tree) Append(resource v1.RESOURCE, id, minTs, maxTs uint64) {
	t.bufferMu.Lock()
	defer t.bufferMu.Unlock()
	t.build.Add(resource, id, minTs, maxTs)
}

func (t *Tree) save(r *meta.Meta) (uint64, error) {
	id := t.seq.MetaID()
	key := keys.NewMeta().WithID(id)
	defer key.Release()
	a := arena.New()
	defer a.Release()
	metrics, logs, traces := r.Proto()

	md, err := a.Compress(metrics)
	if err != nil {
		return 0, err
	}
	ld, err := a.Compress(logs)
	if err != nil {
		return 0, err
	}
	td, err := a.Compress(traces)
	if err != nil {
		return 0, err
	}
	info, err := a.Compress(r.Info())
	if err != nil {
		return 0, err
	}
	return id, t.db.Update(func(txn *badger.Txn) error {
		return errors.Join(
			txn.Set(
				bytes.Clone(key.Reset().Info().
					WithID(id).
					Encode()),
				info,
			),
			txn.Set(
				bytes.Clone(key.Reset().Data().WithRESOURCE(v1.RESOURCE_METRICS).
					WithID(id).
					Encode()),
				md,
			),
			txn.Set(
				bytes.Clone(key.Reset().Data().WithRESOURCE(v1.RESOURCE_LOGS).
					WithID(id).
					Encode()),
				ld,
			),
			txn.Set(
				bytes.Clone(key.Reset().Data().WithRESOURCE(v1.RESOURCE_TRACES).
					WithID(id).
					Encode()),
				td,
			),
		)
	})
}

func (t *Tree) Flush() {
	t.bufferMu.Lock()
	defer t.bufferMu.Unlock()
	r := t.build
	t.build = meta.New()
	if r.Min() == 0 && r.Max() == 0 {
		r.Release()
		return
	}
	r.Sort()
	t.add(r)
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

func (t *Tree) add(part *meta.Meta) {
	t.size.Add(uint64(part.Size()))
	t.root.Prepend(part)
}

func (t *Tree) Close() error {
	err := t.Compact()
	if err != nil {
		return err
	}
	m := meta.New()
	defer m.Release()
	t.root.Iterate(func(n *Node[*meta.Meta]) error {
		if n.value == nil || n.value.Compacted() {
			n.value.Release()
			n.value = nil
			return nil
		}
		m.Merge(n.value)
		n.value.Release()
		n.value = nil
		return nil
	})
	t.root = nil // avoid using a closed tree
	if m.IsEmpty() {
		return nil
	}
	_, err = t.save(m)
	return err
}

func (t *Tree) Compact() error {
	t.Flush()
	t.root.Iterate(func(n *Node[*meta.Meta]) error {
		if n.value == nil || n.value.Compacted() {
			return nil
		}
		t.nodes = append(t.nodes, n)
		return nil
	})

	defer func() {
		clear(t.nodes)
		t.nodes = t.nodes[:0]
	}()

	if len(t.nodes) == 0 {
		return nil
	}
	if len(t.nodes) == 1 {
		return nil
	}
	defer func() {
		for i := range t.nodes {
			t.nodes[i].value.Release()
			t.nodes[i].value = nil
		}
	}()
	m := meta.New()
	for i := 0; i < len(t.nodes); i++ {
		m.Merge(t.nodes[i].value)
	}
	if m.Size() >= PartSize {
		id, err := t.save(m)
		if err != nil {
			return err
		}
		info := meta.FromInfo(&v1.MetaInfo{
			MinTs: m.Min(),
			MaxTs: m.Max(),
			Id:    id,
		})
		// we release m to be reused. After this compaction we will only have info
		// metadata in the tree
		m.Release()
		m = info
	}

	node := t.findNode(t.nodes[0])

	x := &Node[*meta.Meta]{value: m}
	for !node.next.CompareAndSwap(t.nodes[0], x) {
		node = t.findNode(t.nodes[0])
	}
	t.size.Add(-uint64(m.Size()))
	return nil
}

func (t *Tree) load() error {
	key := keys.NewMeta().Info()
	defer key.Release()
	prefix := key[:8+4]
	var i v1.MetaInfo
	return t.db.View(func(txn *badger.Txn) error {
		o := badger.DefaultIteratorOptions
		o.Prefix = prefix
		it := txn.NewIterator(o)
		defer it.Close()
		for it.Rewind(); it.ValidForPrefix(prefix); it.Next() {
			err := it.Item().Value(func(val []byte) error {
				return proto.Unmarshal(val, &i)
			})
			if err != nil {
				return err
			}
			t.add(meta.FromInfo(&i))
		}
		return nil
	})
}

func (t *Tree) Scan(resource v1.RESOURCE, start, end uint64) (*bitmaps.Bitmap, error) {
	t.Flush()

	samples := bitmaps.New()
	err := t.root.Iterate(func(n *Node[*meta.Meta]) error {
		if n.value == nil {
			return nil
		}
		if acceptRange(n.value.Min(), n.value.Max(), start, end) {
			if n.value.Compacted() {
				err := t.readCompacted(n.value, samples, resource, start, end)
				if err != nil {
					slog.Error("failed searching compacted meta",
						"id", n.value.ID(), "err", err)
				}
			} else {
				n.value.Search(samples, resource, start, end)
			}
			if n.value.Max() < end {
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

func (t *Tree) readCompacted(m *meta.Meta, o *bitmaps.Bitmap, resource v1.RESOURCE, start, end uint64) error {
	key := keys.NewMeta().Data().WithID(m.ID()).WithRESOURCE(resource)
	defer key.Release()
	return t.db.View(func(txn *badger.Txn) error {
		it, err := txn.Get(key.Encode())
		if err != nil {
			return err
		}
		var data v1.Meta
		var size int64
		err = it.Value(x.Decompress(&data, &size))
		if err != nil {
			return err
		}
		meta.SearchMeta(&data, o, start, end)
		return nil
	})
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

func (t *Tree) findNode(node *Node[*meta.Meta]) (list *Node[*meta.Meta]) {
	t.root.Iterate(func(n *Node[*meta.Meta]) error {
		if n.next.Load() == node {
			list = n
			return io.EOF
		}
		return nil
	})
	return
}
