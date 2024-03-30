package data

import (
	"sort"

	"github.com/cespare/xxhash/v2"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricsV1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
	"google.golang.org/protobuf/proto"
)

func CollapseMetrics(ds ...*metricsV1.MetricsData) *metricsV1.MetricsData {
	buf := make([]byte, 0, 4<<10)
	h := new(xxhash.Digest)
	resourceMetrics := make(map[uint64]*metricsV1.ResourceMetrics)
	scopedMetrics := make(map[uint64]*metricsV1.ScopeMetrics)
	metrics := make(map[uint64]*metricsV1.Metric)
	order := make(map[uint64]int)
	var idx int
	for _, md := range ds {
		for _, rm := range md.ResourceMetrics {
			h.WriteString(rm.SchemaUrl)
			buf = hashAttributes(h, buf, rm.Resource.GetAttributes())
			rh := h.Sum64()
			resource, ok := resourceMetrics[rh]
			if !ok {
				resource = &metricsV1.ResourceMetrics{
					SchemaUrl: rm.SchemaUrl,
				}
				if rm.Resource != nil {
					resource.Resource = proto.Clone(rm.Resource).(*resourcev1.Resource)
				}
				idx++
				resourceMetrics[rh] = resource
				order[rh] = idx
			}
			for _, sm := range rm.ScopeMetrics {
				h.WriteString(rm.SchemaUrl)
				if sc := sm.Scope; sc != nil {
					h.WriteString(sc.Name)
					h.WriteString(sc.Version)
					buf = hashAttributes(h, buf, sc.Attributes)
				}
				sh := h.Sum64()

				scope, ok := scopedMetrics[sh]
				if !ok {
					scope := &metricsV1.ScopeMetrics{
						SchemaUrl: sm.SchemaUrl,
					}
					if sm.Scope != nil {
						scope.Scope = proto.Clone(sm.Scope).(*commonv1.InstrumentationScope)
					}
					scopedMetrics[sh] = scope
					resource.ScopeMetrics = append(resource.ScopeMetrics, scope)
				}

				for _, m := range sm.Metrics {
					h.WriteString(m.Name)
					mh := h.Sum64()

					om, ok := metrics[mh]
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

	// We sort resources before returning them
	rs := make([]uint64, 0, len(order))
	id := make([]int, 0, len(order))
	for r := range order {
		rs = append(rs, r)
	}
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return order[rs[id[i]]] < order[rs[id[j]]]
	})
	o := &metricsV1.MetricsData{
		ResourceMetrics: make([]*metricsV1.ResourceMetrics, 0, len(rs)),
	}
	for _, i := range id {
		o.ResourceMetrics = append(o.ResourceMetrics,
			resourceMetrics[rs[i]],
		)
	}
	return o
}
