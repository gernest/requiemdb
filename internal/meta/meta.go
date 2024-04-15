package meta

import (
	"slices"
	"sort"
	"sync"
	"unsafe"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
)

type meta struct {
	id  []uint64
	min []uint64
}

func (m *meta) Len() int {
	return len(m.id)
}

func (m *meta) Less(i, j int) bool {
	return m.min[i] < m.min[j]
}

func (m *meta) Swap(i, j int) {
	m.id[i], m.id[j] = m.id[j], m.id[i]
	m.min[i], m.min[j] = m.min[j], m.min[i]
}

func (m *meta) add(id, minTs uint64) {
	m.id = append(m.id, id)
	m.min = append(m.min, minTs)
}

func (m *meta) merge(o *meta) {
	m.id = append(m.id, o.id...)
	m.min = append(m.min, o.min...)
}

func (m *meta) Search(o *bitmaps.Bitmap, start, end uint64) {
	Find(o, m.id, m.min, start, end)
}

func Find(o *bitmaps.Bitmap, id, min []uint64, start, end uint64) {
	set := bitmaps.New()
	defer set.Release()
	from, _ := slices.BinarySearch(min, start)
	to, _ := slices.BinarySearch(min, end)
	for i := from; i < to; i++ {
		set.Add(id[i])
	}
	o.Or(&set.Bitmap)
}

func (m *meta) Reset() {
	m.id = m.id[:0]
	m.min = m.min[:0]
}

var baseMetaSize = int(unsafe.Sizeof(meta{}))

func (m *meta) Size() (n int) {
	n = baseMetaSize
	n += len(m.id) * 8
	n += len(m.min) * 8
	return
}

func (m *meta) Proto() *v1.Meta {
	return &v1.Meta{
		Id:    m.id,
		MinTs: m.min,
	}
}

type Meta struct {
	metrics meta
	traces  meta
	logs    meta
	min     uint64
	max     uint64
	id      uint64
	info    bool
}

func (r *Meta) IsInfo() bool {
	return r.info
}

func (r *Meta) Compacted() bool {
	return r.info
}

func (r *Meta) IsEmpty() bool {
	return r.min == 0 && r.max == 0
}

func New() *Meta {
	return pool.Get().(*Meta)
}

func FromInfo(m *v1.MetaInfo) *Meta {
	return &Meta{
		min:  m.MinTs,
		max:  m.MaxTs,
		id:   m.Id,
		info: true,
	}
}

var pool = &sync.Pool{New: func() any { return new(Meta) }}

func (r *Meta) Reset() {
	r.metrics.Reset()
	r.logs.Reset()
	r.traces.Reset()
}

var baseResourceSize = int(unsafe.Sizeof(Meta{}))

func (r *Meta) Size() (n int) {
	n = baseResourceSize
	n += r.metrics.Size()
	n += r.logs.Size()
	n += r.traces.Size()
	return
}

func (r *Meta) Min() uint64 {
	return r.min
}

func (r *Meta) Max() uint64 {
	return r.max
}

func (r *Meta) ID() uint64 {
	return r.id
}

func (r *Meta) bounds(lo, hi uint64) {
	if r.min == 0 {
		r.min = lo
	}
	r.min = min(r.min, lo)
	r.max = max(r.max, hi)
}

func (r *Meta) Sort() {
	sort.Sort(&r.metrics)
	sort.Sort(&r.logs)
	sort.Sort(&r.traces)
}

func (r *Meta) Release() {
	r.Reset()
	pool.Put(r)
}

func (r *Meta) Proto() (metrics, logs, traces *v1.Meta) {
	return r.metrics.Proto(), r.logs.Proto(), r.traces.Proto()
}

func (r *Meta) Info() *v1.MetaInfo {
	return &v1.MetaInfo{
		MinTs: r.min,
		MaxTs: r.max,
	}
}

func (r *Meta) Add(resource v1.RESOURCE, id, minTx, maxTs uint64) {
	switch resource {
	case v1.RESOURCE_METRICS:
		r.metrics.add(id, minTx)
	case v1.RESOURCE_LOGS:
		r.logs.add(id, minTx)
	case v1.RESOURCE_TRACES:
		r.traces.add(id, minTx)
	}
	r.bounds(minTx, maxTs)
}

func (r *Meta) Merge(o *Meta) {
	r.metrics.merge(&o.metrics)
	r.logs.merge(&o.logs)
	r.traces.merge(&o.traces)
	r.bounds(o.Min(), o.Max())
}

func (r *Meta) Search(o *bitmaps.Bitmap, resource v1.RESOURCE, start, end uint64) {
	switch resource {
	case v1.RESOURCE_METRICS:
		r.metrics.Search(o, start, end)
	case v1.RESOURCE_LOGS:
		r.logs.Search(o, start, end)
	case v1.RESOURCE_TRACES:
		r.traces.Search(o, start, end)
	}
}
