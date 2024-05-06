package f2

import (
	"github.com/gernest/arrow3"
	v1 "github.com/gernest/requiemdb/gen/go/store/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type ID interface {
	Next() uint64
}

type Metrics struct {
	build *arrow3.Schema[*v1.Metric]
	id    ID
}

func (m *Metrics) Append(data *metricsv1.MetricsData) {
	a := newAttr()
	defer a.Release()

	for _, rs := range data.ResourceMetrics {
		a.Resource(rs.Resource)
		for _, sc := range rs.ScopeMetrics {
			a.Scope(sc.Scope)

			for _, ms := range sc.Metrics {
				mx := metricPool.Get()

				mx.Name = ms.Name
				mx.Description = ms.Description
				mx.Unit = ms.Unit

				if v := ms.GetGauge(); v != nil {
					mx.Kind = v1.Metric_GAUGE
				}
				if v := ms.GetSum(); v != nil {
					mx.Kind = v1.Metric_SUM
				}
				if v := ms.GetHistogram(); v != nil {
					mx.Kind = v1.Metric_HISTOGRAM
				}
				if v := ms.GetExponentialHistogram(); v != nil {
					mx.Kind = v1.Metric_EXPONENTIAL_HISTOGRAM
				}
				if v := ms.GetSummary(); v != nil {
					mx.Kind = v1.Metric_SUMMARY
				}
				metricPool.Put(mx)
			}

		}
	}
}
