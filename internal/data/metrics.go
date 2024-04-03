package data

import (
	"slices"
	"sort"
	"sync"

	"github.com/cespare/xxhash/v2"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
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
			xm.hash.Reset()
			xm.hash.WriteString(rm.SchemaUrl)
			xm.attr(rm.Resource.Attributes)
			rhb := xm.hash.Sum(nil)
			rh := xm.hash.Sum64()
			resource, ok := xm.resource[rh]
			if !ok {
				resource = &metricsV1.ResourceMetrics{
					SchemaUrl: rm.SchemaUrl,
				}
				if rm.Resource != nil {
					resource.Resource = proto.Clone(rm.Resource).(*resourcev1.Resource)
				}
				idx++
				xm.resource[rh] = resource
				xm.order[rh] = idx
			}
			for _, sm := range rm.ScopeMetrics {
				xm.hash.Reset()
				xm.hash.Write(rhb)
				xm.hash.WriteString(rm.SchemaUrl)
				if sc := sm.Scope; sc != nil {
					xm.hash.WriteString(sc.Name)
					xm.hash.WriteString(sc.Version)
					xm.attr(sc.Attributes)
				}
				shb := xm.hash.Sum(nil)
				sh := xm.hash.Sum64()

				scope, ok := xm.scope[sh]
				if !ok {
					scope = &metricsV1.ScopeMetrics{
						SchemaUrl: sm.SchemaUrl,
					}
					if sm.Scope != nil {
						scope.Scope = proto.Clone(sm.Scope).(*commonv1.InstrumentationScope)
					}
					xm.scope[sh] = scope
					resource.ScopeMetrics = append(resource.ScopeMetrics, scope)
				}

				for _, m := range sm.Metrics {
					xm.hash.Reset()
					xm.hash.Write(shb)
					xm.hash.WriteString(m.Name)
					mh := xm.hash.Sum64()

					om, ok := xm.metrics[mh]
					if !ok {
						om = &metricsV1.Metric{
							Name:        m.Name,
							Description: m.Description,
							Unit:        m.Unit,
						}
						scope.Metrics = append(scope.Metrics, om)
					}
					switch e := om.Data.(type) {
					case *metricsV1.Metric_Gauge:
						om.Data = &metricsV1.Metric_Gauge{
							Gauge: &metricsV1.Gauge{
								DataPoints: append(om.GetGauge().GetDataPoints(),
									e.Gauge.GetDataPoints()...),
							},
						}
					case *metricsV1.Metric_Sum:
						om.Data = &metricsV1.Metric_Sum{
							Sum: &metricsV1.Sum{
								AggregationTemporality: e.Sum.GetAggregationTemporality(),
								DataPoints: append(om.GetSum().GetDataPoints(),
									e.Sum.GetDataPoints()...),
							},
						}
					case *metricsV1.Metric_Histogram:
						om.Data = &metricsV1.Metric_Histogram{
							Histogram: &metricsV1.Histogram{
								AggregationTemporality: e.Histogram.GetAggregationTemporality(),
								DataPoints: append(om.GetHistogram().GetDataPoints(),
									e.Histogram.GetDataPoints()...),
							},
						}
					case *metricsV1.Metric_ExponentialHistogram:
						om.Data = &metricsV1.Metric_ExponentialHistogram{
							ExponentialHistogram: &metricsV1.ExponentialHistogram{
								AggregationTemporality: e.ExponentialHistogram.GetAggregationTemporality(),
								DataPoints: append(om.GetExponentialHistogram().GetDataPoints(),
									e.ExponentialHistogram.GetDataPoints()...),
							},
						}
					case *metricsV1.Metric_Summary:
						om.Data = &metricsV1.Metric_Summary{
							Summary: &metricsV1.Summary{
								DataPoints: append(om.GetSummary().GetDataPoints(),
									e.Summary.GetDataPoints()...),
							},
						}
					}
				}
			}
		}
	}
	xm.Sort()
	return xm.Result()
}

func metricsFrom(ls []*v1.Data) []*metricsV1.MetricsData {
	o := make([]*metricsV1.MetricsData, len(ls))
	for i := range ls {
		o[i] = ls[i].GetMetrics()
	}
	return o
}

type RM struct {
	hash     xxhash.Digest
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
		r.hash.Write(r.buf)
	}
}
