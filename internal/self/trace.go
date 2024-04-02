package self

import (
	"context"

	"github.com/gernest/requiemdb/internal/self/trace/tracetransform"
	"go.opentelemetry.io/otel/sdk/trace"
	collector_trace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

type Trace struct {
	base
	Collector collector_trace.TraceServiceServer
}

var _ trace.SpanExporter = (*Trace)(nil)

func (m *Trace) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	rs := tracetransform.Spans(spans)
	_, err := m.Collector.Export(ctx, &collector_trace.ExportTraceServiceRequest{
		ResourceSpans: rs,
	})
	return err
}
