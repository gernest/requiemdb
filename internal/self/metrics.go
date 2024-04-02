package self

import (
	"context"

	"github.com/gernest/requiemdb/internal/self/metrics/transform"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	collector_metrics "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type Metrics struct {
	base
	Collector collector_metrics.MetricsServiceServer
}

var _ metric.Exporter = (*Metrics)(nil)

func (m *Metrics) Export(ctx context.Context, rm *metricdata.ResourceMetrics) error {

	tlpRm, err := transform.ResourceMetrics(rm)
	if err != nil {
		return err
	}
	_, err = m.Collector.Export(ctx, &collector_metrics.ExportMetricsServiceRequest{
		ResourceMetrics: []*metricsv1.ResourceMetrics{
			tlpRm,
		},
	})
	return err
}

func (m *Metrics) Temporality(kind metric.InstrumentKind) metricdata.Temporality {
	return metric.DefaultTemporalitySelector(kind)
}

func (m *Metrics) Aggregation(kind metric.InstrumentKind) metric.Aggregation {
	return metric.DefaultAggregationSelector(kind)
}

type base struct{}

func (base) ForceFlush(context.Context) error {
	return nil
}

func (base) Shutdown(context.Context) error {
	return nil
}
