package transform

import (
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func (c *Context) transformMetrics(rm *metricsv1.ResourceMetrics) {
	if rm.SchemaUrl != "" {
		c.label(
			labels.NewBytes(v1.RESOURCE_METRICS, v1.PREFIX_RESOURCE_SCHEMA).Value(rm.SchemaUrl),
		)
	}
	if rm.Resource != nil {
		c.attributes(v1.RESOURCE_METRICS, v1.PREFIX_RESOURCE_ATTRIBUTES, rm.Resource.Attributes)
	}
	for _, sm := range rm.ScopeMetrics {
		if sm.SchemaUrl != "" {
			c.label(
				labels.NewBytes(v1.RESOURCE_METRICS, v1.PREFIX_SCOPE_SCHEMA).Value(sm.SchemaUrl),
			)
		}
		if sc := sm.Scope; sc != nil {
			if sc.Name != "" {
				c.label(
					labels.NewBytes(v1.RESOURCE_METRICS, v1.PREFIX_SCOPE_NAME).Value(sc.Name),
				)
			}
			if sc.Version != "" {
				c.label(
					labels.NewBytes(v1.RESOURCE_METRICS, v1.PREFIX_SCOPE_VERSION).Value(sc.Version),
				)
			}
			c.attributes(v1.RESOURCE_METRICS, v1.PREFIX_SCOPE_ATTRIBUTES, sc.Attributes)
		}
		for _, m := range sm.Metrics {
			c.label(
				labels.NewBytes(v1.RESOURCE_METRICS, v1.PREFIX_NAME).Value(m.Name),
			)
			switch e := m.Data.(type) {
			case *metricsv1.Metric_Gauge:
				transFormDataPoints(c, e.Gauge.DataPoints)
			case *metricsv1.Metric_Sum:
				transFormDataPoints(c, e.Sum.DataPoints)
			case *metricsv1.Metric_Histogram:
				for _, p := range e.Histogram.DataPoints {
					c.Timestamp(p.TimeUnixNano)
					c.attributes(v1.RESOURCE_METRICS, v1.PREFIX_ATTRIBUTES, p.Attributes)
				}
			case *metricsv1.Metric_ExponentialHistogram:
				for _, p := range e.ExponentialHistogram.DataPoints {
					c.Timestamp(p.TimeUnixNano)
					c.attributes(v1.RESOURCE_METRICS, v1.PREFIX_ATTRIBUTES, p.Attributes)
				}
			case *metricsv1.Metric_Summary:
				for _, p := range e.Summary.DataPoints {
					c.Timestamp(p.TimeUnixNano)
					c.attributes(v1.RESOURCE_METRICS, v1.PREFIX_ATTRIBUTES, p.Attributes)
				}
			}
		}
	}
}

func transFormDataPoints(ctx *Context, dp []*metricsv1.NumberDataPoint) {
	for _, p := range dp {
		ctx.Timestamp(p.TimeUnixNano)
		ctx.attributes(v1.RESOURCE_METRICS, v1.PREFIX_ATTRIBUTES, p.Attributes)
	}
}
