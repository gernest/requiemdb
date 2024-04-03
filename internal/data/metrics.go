package data

import (
	"encoding/binary"
	"slices"
	"sort"
	"sync"

	"github.com/cespare/xxhash/v2"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricsV1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
	"google.golang.org/protobuf/proto"
)

func CollapseMetrics(ds []*metricsV1.MetricsData) *metricsV1.MetricsData {
	xm := NewRM()
	defer xm.Release()
	var idx int
	for _, md := range ds {
		for _, rm := range md.ResourceMetrics {
			rh := xm.hashResource(rm.SchemaUrl, rm.Resource)
			resource, ok := xm.resource[rh]
			if !ok {
				resource = rm
				idx++
				xm.resource[rh] = resource
				xm.order[rh] = idx

				// add all scopes for this resource
				for _, sm := range rm.ScopeMetrics {
					sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
					xm.scope[sh] = sm
					for _, m := range sm.Metrics {
						mh := xm.hashMetrics(sh, m.Name)
						xm.metrics[mh] = m
					}
				}
				continue
			}
			for _, sm := range rm.ScopeMetrics {
				sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
				scope, ok := xm.scope[sh]
				if !ok {
					xm.scope[sh] = sm
					resource.ScopeMetrics = append(resource.ScopeMetrics, sm)
					for _, m := range sm.Metrics {
						mh := xm.hashMetrics(sh, m.Name)
						xm.metrics[mh] = m
					}
					continue
				}
				for _, m := range sm.Metrics {
					mh := xm.hashMetrics(sh, m.Name)
					om, ok := xm.metrics[mh]
					if !ok {
						scope.Metrics = append(scope.Metrics, m)
						xm.metrics[mh] = m
						continue
					}
					if gauge := m.GetGauge(); gauge != nil {
						om.GetGauge().DataPoints = append(om.GetGauge().DataPoints, gauge.DataPoints...)
					}
					if sum := m.GetSum(); sum != nil {
						om.GetSum().DataPoints = append(om.GetSum().DataPoints, sum.DataPoints...)
					}
					if hist := m.GetHistogram(); hist != nil {
						om.GetHistogram().DataPoints = append(om.GetHistogram().DataPoints, hist.DataPoints...)
					}
					if ehist := m.GetExponentialHistogram(); ehist != nil {
						om.GetExponentialHistogram().DataPoints = append(om.GetExponentialHistogram().DataPoints,
							ehist.DataPoints...)
					}
					if sum := m.GetSummary(); sum != nil {
						om.GetSummary().DataPoints = append(om.GetSummary().DataPoints, sum.DataPoints...)
					}
				}
			}
		}
	}
	xm.Sort()
	return xm.Result()
}

type RM struct {
	h        xxhash.Digest
	resource map[uint64]*metricsV1.ResourceMetrics
	scope    map[uint64]*metricsV1.ScopeMetrics
	metrics  map[uint64]*metricsV1.Metric
	order    map[uint64]int
	rs       []uint64
	id       []int
	buf      []byte
}

func NewRM() *RM {
	return rmPool.Get().(*RM)
}

func (r *RM) Release() {
	r.Reset()
	rmPool.Put(r)
}

func (r *RM) Reset() {
	clear(r.resource)
	clear(r.scope)
	clear(r.metrics)
	clear(r.order)
	r.rs = r.rs[:0]
	r.id = r.id[:0]
	r.buf = r.buf[:0]
}

var rmPool = &sync.Pool{New: func() any { return newRM() }}

func newRM() *RM {
	return &RM{
		resource: make(map[uint64]*metricsV1.ResourceMetrics),
		scope:    make(map[uint64]*metricsV1.ScopeMetrics),
		metrics:  make(map[uint64]*metricsV1.Metric),
		order:    make(map[uint64]int),
		buf:      make([]byte, 0, 4<<10),
	}
}

func (r *RM) hashMetrics(scope uint64, name string) uint64 {
	r.h.Reset()
	r.hashNum(scope)
	r.h.WriteString(name)
	return r.h.Sum64()
}

func (r *RM) hashNum(v uint64) {
	binary.BigEndian.PutUint64(r.buf[:8], v)
	r.h.Write(r.buf[:8])
	r.buf = r.buf[:0]
}

func (r *RM) hashResource(schema string, resource *resourcev1.Resource) uint64 {
	r.h.Reset()
	r.h.WriteString(schema)
	if resource != nil {
		r.attr(resource.Attributes)
	}
	return r.h.Sum64()
}

func (r *RM) hashScope(resource uint64, schema string, scope *commonv1.InstrumentationScope) uint64 {
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

func (r *RM) Len() int {
	return len(r.rs)
}

func (r *RM) Result() *metricsV1.MetricsData {
	o := &metricsV1.MetricsData{
		ResourceMetrics: make([]*metricsV1.ResourceMetrics, 0, len(r.rs)),
	}
	for _, i := range r.id {
		o.ResourceMetrics = append(o.ResourceMetrics,
			r.resource[r.rs[i]],
		)
	}
	return o
}

func (r *RM) Sort() {
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
func (r *RM) Less(i, j int) bool {
	return r.order[r.rs[r.id[i]]] < r.order[r.rs[r.id[j]]]
}

func (r *RM) Swap(i, j int) {
	r.id[i], r.id[j] = r.id[j], r.id[i]
}

func (r *RM) attr(kv []*commonv1.KeyValue) {
	for _, v := range kv {
		r.buf, _ = proto.MarshalOptions{}.MarshalAppend(r.buf[:0], v)
		r.h.Write(r.buf)
	}
}
