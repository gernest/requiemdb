package transform

import (
	"github.com/requiemdb/requiemdb/internal/labels"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func (c *Context) transformMetrics(rm *metricsv1.ResourceMetrics) {
	if rm.SchemaUrl != "" {
		c.label(
			labels.NewBytes().Add(labels.ResourceSchema).Value(rm.SchemaUrl),
		)
	}
	if rm.Resource != nil {
		c.attributes(labels.ResourceAttributes, rm.Resource.Attributes)
	}
	for _, sm := range rm.ScopeMetrics {
		if sm.SchemaUrl != "" {
			c.label(
				labels.NewBytes().Add(labels.ScopeSchema).Value(sm.SchemaUrl),
			)
		}
		if sc := sm.Scope; sc != nil {
			if sc.Name != "" {
				c.label(
					labels.NewBytes().Add(labels.ScopeName).Value(sc.Name),
				)
			}
			if sc.Version != "" {
				c.label(
					labels.NewBytes().Add(labels.ScopeVersion).Value(sc.Version),
				)
			}
			c.attributes(labels.ScopeAttributes, sc.Attributes)
		}

		for _, m := range sm.Metrics {
			c.label(
				labels.NewBytes().Add(labels.MetricName).Value(m.Name),
			)
			switch e := m.Data.(type) {
			case *metricsv1.Metric_Gauge:
				transFormDataPoints(c, e.Gauge.DataPoints)
			case *metricsv1.Metric_Sum:
				transFormDataPoints(c, e.Sum.DataPoints)
			case *metricsv1.Metric_Histogram:
				for _, p := range e.Histogram.DataPoints {
					c.Timestamp(p.TimeUnixNano)
					c.attributes(labels.Attribute, p.Attributes)
				}
			case *metricsv1.Metric_ExponentialHistogram:
				for _, p := range e.ExponentialHistogram.DataPoints {
					c.Timestamp(p.TimeUnixNano)
					c.attributes(labels.Attribute, p.Attributes)
				}
			case *metricsv1.Metric_Summary:
				for _, p := range e.Summary.DataPoints {
					c.Timestamp(p.TimeUnixNano)
					c.attributes(labels.Attribute, p.Attributes)
				}
			}
		}
	}
}

func transFormDataPoints(ctx *Context, dp []*metricsv1.NumberDataPoint) {
	for _, p := range dp {
		ctx.Timestamp(p.TimeUnixNano)
		ctx.attributes(labels.Attribute, p.Attributes)
	}
}
