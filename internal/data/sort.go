package data

import (
	"encoding/binary"
	"slices"
	"sort"

	"github.com/cespare/xxhash/v2"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
	"google.golang.org/protobuf/proto"
)

type Sorter[T any, R any, S any] struct {
	h        xxhash.Digest
	resource map[uint64]R
	scope    map[uint64]S
	metrics  map[uint64]T
	order    map[uint64]int
	rs       []uint64
	id       []int
	buf      []byte
}

func (r *Sorter[T, R, S]) Reset() {
	clear(r.resource)
	clear(r.scope)
	clear(r.metrics)
	clear(r.order)
	r.rs = r.rs[:0]
	r.id = r.id[:0]
	r.buf = r.buf[:0]
}

func newSorter[T any, R any, S any]() *Sorter[T, R, S] {
	return &Sorter[T, R, S]{
		resource: make(map[uint64]R),
		scope:    make(map[uint64]S),
		metrics:  make(map[uint64]T),
		order:    make(map[uint64]int),
		buf:      make([]byte, 0, 4<<10),
	}
}

func (r *Sorter[T, R, S]) hashMetrics(scope uint64, name string) uint64 {
	r.h.Reset()
	r.hashNum(scope)
	r.h.WriteString(name)
	return r.h.Sum64()
}

func (r *Sorter[T, R, S]) hashNum(v uint64) {
	binary.BigEndian.PutUint64(r.buf[:8], v)
	r.h.Write(r.buf[:8])
	r.buf = r.buf[:0]
}

func (r *Sorter[T, R, S]) hashResource(schema string, resource *resourcev1.Resource) uint64 {
	r.h.Reset()
	r.h.WriteString(schema)
	if resource != nil {
		r.attr(resource.Attributes)
	}
	return r.h.Sum64()
}

func (r *Sorter[T, R, S]) hashScope(resource uint64, schema string, scope *commonv1.InstrumentationScope) uint64 {
	r.h.Reset()
	r.hashNum(resource)
	r.h.WriteString(schema)
	if scope != nil {
		r.h.WriteString(scope.Name)
		r.h.WriteString(scope.Version)
		if scope.Attributes != nil {
			r.attr(scope.Attributes)
		}
	}
	return r.h.Sum64()
}

func (r *Sorter[T, R, S]) Len() int {
	return len(r.rs)
}

func (r *Sorter[T, R, S]) Result() []R {
	o := make([]R, 0, len(r.rs))
	for _, i := range r.id {
		o = append(o,
			r.resource[r.rs[i]],
		)
	}
	return o
}

func (r *Sorter[T, R, S]) Sort() {
	r.rs = slices.Grow(r.rs, len(r.order))
	for v := range r.order {
		r.rs = append(r.rs, v)
	}
	r.id = slices.Grow(r.id, len(r.order))
	for i := 0; i < len(r.order); i++ {
		r.id = append(r.id, i)
	}
	sort.Sort(r)
}
func (r *Sorter[T, R, S]) Less(i, j int) bool {
	return r.order[r.rs[r.id[i]]] < r.order[r.rs[r.id[j]]]
}

func (r *Sorter[T, R, S]) Swap(i, j int) {
	r.id[i], r.id[j] = r.id[j], r.id[i]
}

func (r *Sorter[T, R, S]) attr(kv []*commonv1.KeyValue) {
	for _, v := range kv {
		r.buf, _ = proto.MarshalOptions{}.MarshalAppend(r.buf[:0], v)
		r.h.Write(r.buf)
	}
}
