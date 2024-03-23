package transform

import (
	v1 "github.com/requiemdb/requiemdb/gen/go/samples/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func (c *Context) transformTrace(rm *tracev1.ResourceSpans) {
	if rm.SchemaUrl != "" {
		c.label(
			labels.NewBytes(v1.PREFIX_RESOURCE_SCHEMA).Value(rm.SchemaUrl),
		)
	}
	if rm.Resource != nil {
		c.attributes(v1.PREFIX_RESOURCE_ATTRIBUTES, rm.Resource.Attributes)
	}
	for _, sm := range rm.ScopeSpans {
		if sm.SchemaUrl != "" {
			c.label(
				labels.NewBytes(v1.PREFIX_SCOPE_SCHEMA).Value(sm.SchemaUrl),
			)
		}
		if sc := sm.Scope; sc != nil {
			if sc.Name != "" {
				c.label(
					labels.NewBytes(v1.PREFIX_SCOPE_NAME).Value(sc.Name),
				)
			}
			if sc.Version != "" {
				c.label(
					labels.NewBytes(v1.PREFIX_SCOPE_VERSION).Value(sc.Version),
				)
			}
			c.attributes(v1.PREFIX_SCOPE_ATTRIBUTES, sc.Attributes)
		}
		for _, span := range sm.Spans {
			if span.Name != "" {
				c.label(labels.NewBytes(v1.PREFIX_SPAN_NAME).Value(span.Name))
			}
			if span.TraceId != nil {
				c.label(labels.NewBytes(v1.PREFIX_SPAN_TRACE_ID).ValueBytes(span.TraceId))
			}
			if span.SpanId != nil {
				c.label(labels.NewBytes(v1.PREFIX_SPAN_SPAN_ID).ValueBytes(span.SpanId))
			}
			if span.ParentSpanId != nil {
				c.label(labels.NewBytes(v1.PREFIX_SPAN_PARENT_SPAN_ID).ValueBytes(span.ParentSpanId))
			}
			c.Range(span.StartTimeUnixNano, span.EndTimeUnixNano)
		}
	}
}
