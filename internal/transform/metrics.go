package transform

import (
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/labels"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func (c *Context) transformMetrics(rm *metricsv1.ResourceMetrics) {
	c.addResource(rm)
	for _, sm := range rm.ScopeMetrics {
		c.addScope(sm)
		for _, m := range sm.Metrics {
			c.Label(func(lbl *labels.Label) {
				lbl.
					WithPrefix(v1.PREFIX_NAME).
					WithKey(m.Name)
			})
			switch e := m.Data.(type) {
			case *metricsv1.Metric_Gauge:
				for _, p := range e.Gauge.DataPoints {
					c.addDataBase(p)
				}
			case *metricsv1.Metric_Sum:
				for _, p := range e.Sum.DataPoints {
					c.addDataBase(p)
				}
			case *metricsv1.Metric_Histogram:
				for _, p := range e.Histogram.DataPoints {
					c.addDataBase(p)
				}
			case *metricsv1.Metric_ExponentialHistogram:
				for _, p := range e.ExponentialHistogram.DataPoints {
					c.addDataBase(p)

				}
			case *metricsv1.Metric_Summary:
				for _, p := range e.Summary.DataPoints {
					c.addDataBase(p)
				}
			}
		}
	}
}

type dataBase interface {
	GetTimeUnixNano() uint64
	GetAttributes() []*commonv1.KeyValue
}

func (c *Context) addDataBase(b dataBase) {
	c.Timestamp(b.GetTimeUnixNano())
	c.attributes(v1.PREFIX_ATTRIBUTES, b.GetAttributes())
}
