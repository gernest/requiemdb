package batch

import (
	"cmp"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gernest/rbf/quantum"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/roaring"
	"github.com/gernest/roaring/shardwidth"
	"github.com/pkg/errors"
)

const (
	Quantum = quantum.TimeQuantum("YMDH")
)

const (
	Timestamp uint64 = iota
	Labels
)

type Batch struct {
	ids    []uint64
	ts     []uint64
	labels [][]uint64
	frags  Fragments
	times  []*QuantizedTime
}

func New() *Batch {
	return &Batch{
		frags: make(Fragments),
	}
}

func (b *Batch) Reset() {
	b.frags = make(Fragments)
	b.ids = b.ids[:0]
	b.ts = b.ts[:0]
	b.labels = b.labels[:0]
	for _, t := range b.times {
		t.Release()
	}
	clear(b.times)
	b.times = b.times[:0]
}

func (b *Batch) addTs(ts time.Time) {
	q := NeqQuantumTime()
	q.Set(ts)
	b.times = append(b.times, q)
	b.ts = append(b.ts, uint64(ts.UnixMilli()))
}

func (b *Batch) addLabels(r *bitmaps.Bitmap) {
	rowIDs := make([]uint64, 0, r.GetCardinality())
	it := r.Iterator()
	for it.HasNext() {
		rowID := it.Next()
		rowIDs = append(rowIDs, rowID)

	}
	b.labels = append(b.labels, rowIDs)
}

func (b *Batch) Add(id uint64, ts time.Time, labels *bitmaps.Bitmap) {
	b.ids = append(b.ids, id)
	b.addTs(ts)
	b.addLabels(labels)
}

func (b *Batch) Build() (Fragments, error) {
	o, err := b.makeFragments()
	if err != nil {
		return nil, err
	}
	b.frags = make(Fragments)
	return o, nil
}

func (b *Batch) makeFragments() (Fragments, error) {
	shardWidth := b.shardWidth()
	tsShard := ^uint64(0) // impossible sentinel value for shard.

	rowIDSets := b.labels
	labelShard := ^uint64(0) // impossible sentinel value for shard.

	for j := range b.ids {
		tsCol := b.ids[j]
		if tsCol/shardWidth != tsShard {
			tsShard = tsCol / shardWidth
		}

		labelCol, rowIDs := b.ids[j], rowIDSets[j]
		if labelCol/shardWidth != labelShard {
			labelShard = labelCol / shardWidth
		}

		views, err := b.times[j].Views(Quantum)
		if err != nil {
			return nil, errors.Wrap(err, "calculating views")
		}
		for _, view := range views {
			b.frags.GetOrCreate(tsShard, Timestamp, view).
				DirectAdd(b.ts[j]*shardWidth + (tsCol % shardWidth))
			for _, labelRow := range rowIDs {
				b.frags.GetOrCreate(labelShard, Labels, view).
					DirectAdd(labelRow*shardWidth + (labelCol % shardWidth))
			}
		}
	}
	return b.frags, nil
}

func (b *Batch) shardWidth() uint64 {
	return shardwidth.ShardWidth
}

type Fragments map[FragmentKey]map[string]*roaring.Bitmap

func (f Fragments) String() string {
	keys := make([]FragmentKey, 0, len(f))
	for k := range f {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Compare(&keys[j]) == -1
	})
	var b strings.Builder
	var ls []string
	for i := range keys {
		k := keys[i]
		v := f[k]
		ls = ls[:0]
		for view := range v {
			ls = append(ls, view)
		}
		sort.Strings(ls)
		for _, view := range ls {
			r := v[view]
			fmt.Fprintf(&b, "%d:%d:%s => %v\n", k.Field, k.Shard, view, r)
		}
	}
	return b.String()
}

type FragmentKey struct {
	Shard uint64
	Field uint64
}

func (f *FragmentKey) Compare(o *FragmentKey) int {
	n := cmp.Compare(f.Field, o.Field)
	if n == 0 {
		return cmp.Compare(f.Shard, o.Shard)
	}
	return n
}

func (f Fragments) GetOrCreate(shard uint64, field uint64, view string) *roaring.Bitmap {
	key := FragmentKey{shard, field}
	viewMap, ok := f[key]
	if !ok {
		viewMap = make(map[string]*roaring.Bitmap)
		f[key] = viewMap
	}
	bm, ok := viewMap[view]
	if !ok {
		bm = roaring.NewBTreeBitmap()
		viewMap[view] = bm
	}
	return bm
}

func (f Fragments) GetViewMap(shard uint64, field uint64) map[string]*roaring.Bitmap {
	key := FragmentKey{shard, field}
	viewMap, ok := f[key]
	if !ok {
		return nil
	}
	// Remove any views which have an empty bitmap.
	// TODO: Ideally we would prevent allocating the empty bitmap to begin with,
	// but the logic is a bit tricky, and since we don't want to spend too much
	// time on it right now, we're leaving that for a future exercise.
	for k, v := range viewMap {
		if v.Count() == 0 {
			delete(viewMap, k)
		}
	}
	return viewMap
}

func (f Fragments) DeleteView(shard uint64, field uint64, view string) {
	vm := f.GetViewMap(shard, field)
	if vm == nil {
		return
	}
	delete(vm, view)
}
