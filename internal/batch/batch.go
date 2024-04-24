package batch

import (
	"fmt"
	"slices"
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

type Batch struct {
	ids    []uint64
	labels []*bitmaps.Bitmap
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
	b.labels = b.labels[:0]
	for _, t := range b.times {
		t.Release()
	}
	clear(b.times)
	b.times = b.times[:0]

	for _, t := range b.labels {
		t.Release()
	}
	clear(b.labels)
	b.labels = b.labels[:0]
}

func (b *Batch) addTs(ts time.Time) {
	q := NeqQuantumTime()
	q.Set(ts)
	b.times = append(b.times, q)
}

func (b *Batch) Add(id uint64, ts time.Time, labels *bitmaps.Bitmap) {
	b.ids = append(b.ids, id)
	b.labels = append(b.labels, labels)
	b.addTs(ts)

}

func (b *Batch) Build() (Fragments, error) {
	o, err := b.makeFragments()
	if err != nil {
		return nil, err
	}
	b.Reset()
	return o, nil
}

func (b *Batch) makeFragments() (Fragments, error) {
	shardWidth := b.shardWidth()
	for j := range b.ids {
		if b.labels[j] == nil {
			continue
		}
		it := b.labels[j].Iterator()
		views, err := b.times[j].Views(Quantum)
		if err != nil {
			return nil, errors.Wrap(err, "calculating views")
		}
		shard := ^uint64(0)
		for it.HasNext() {
			col := it.Next()
			if col/shardWidth != shard {
				shard = col / shardWidth
			}
			for _, view := range views {
				b.frags.GetOrCreate(shard, view).
					DirectAdd(b.ids[j]*shardWidth + (col % shardWidth))
			}
		}
	}
	return b.frags, nil
}

func (b *Batch) shardWidth() uint64 {
	return shardwidth.ShardWidth
}

type Fragments map[uint64]map[string]*roaring.Bitmap

func (f Fragments) String() string {
	keys := make([]uint64, 0, len(f))
	for k := range f {
		keys = append(keys, k)
	}
	slices.Sort(keys)
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
			fmt.Fprintf(&b, "%d:%s => %v\n", k, view, r)
		}
	}
	return b.String()
}

func (f Fragments) GetOrCreate(shard uint64, view string) *roaring.Bitmap {
	key := shard
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
