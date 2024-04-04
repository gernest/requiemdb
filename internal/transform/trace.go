package transform

import (
	"encoding/hex"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/labels"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func (c *Context) transformTrace(rm *tracev1.ResourceSpans) {
	c.addResource(rm)
	for _, sm := range rm.ScopeSpans {
		c.addScope(sm)
		for _, span := range sm.Spans {
			if span.Name != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.
						WithPrefix(v1.PREFIX_NAME).
						WithKey(span.Name)
				})
			}
			if span.TraceId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.
						WithPrefix(v1.PREFIX_TRACE_ID).
						WithKey(hex.EncodeToString(span.TraceId))
				})
			}
			if span.SpanId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.
						WithPrefix(v1.PREFIX_SPAN_ID).
						WithKey(hex.EncodeToString(span.SpanId))
				})
			}
			if span.ParentSpanId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.
						WithPrefix(v1.PREFIX_PARENT_SPAN_ID).
						WithKey(hex.EncodeToString(span.ParentSpanId))
				})
			}
			c.Range(span.StartTimeUnixNano, span.EndTimeUnixNano)
		}
	}
}
