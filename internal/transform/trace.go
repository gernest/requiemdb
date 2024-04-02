package transform

import (
	"encoding/hex"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/labels"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func (c *Context) transformTrace(rm *tracev1.ResourceSpans) {
	if rm.SchemaUrl != "" {
		c.Label(func(lbl *labels.Label) {
			lbl.WithResource(v1.RESOURCE_TRACES).
				WithPrefix(v1.PREFIX_RESOURCE_SCHEMA).
				WithKey(rm.SchemaUrl)

		})
	}
	if rm.Resource != nil {
		c.attributes(v1.RESOURCE_TRACES, v1.PREFIX_RESOURCE_ATTRIBUTES, rm.Resource.Attributes)
	}
	for _, sm := range rm.ScopeSpans {
		if sm.SchemaUrl != "" {
			c.Label(func(lbl *labels.Label) {
				lbl.WithResource(v1.RESOURCE_TRACES).
					WithPrefix(v1.PREFIX_SCOPE_SCHEMA).
					WithKey(sm.SchemaUrl)

			})
		}
		if sc := sm.Scope; sc != nil {
			if sc.Name != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_TRACES).
						WithPrefix(v1.PREFIX_SCOPE_NAME).
						WithKey(sc.Name)
				})
			}
			if sc.Version != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_TRACES).
						WithPrefix(v1.PREFIX_SCOPE_VERSION).
						WithKey(sc.Version)
				})
			}
			c.attributes(v1.RESOURCE_TRACES, v1.PREFIX_SCOPE_ATTRIBUTES, sc.Attributes)
		}
		for _, span := range sm.Spans {
			if span.Name != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_TRACES).
						WithPrefix(v1.PREFIX_NAME).
						WithKey(span.Name)
				})
			}
			if span.TraceId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_TRACES).
						WithPrefix(v1.PREFIX_TRACE_ID).
						WithKey(hex.EncodeToString(span.TraceId))
				})
			}
			if span.SpanId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_TRACES).
						WithPrefix(v1.PREFIX_SPAN_ID).
						WithKey(hex.EncodeToString(span.SpanId))
				})
			}
			if span.ParentSpanId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_TRACES).
						WithPrefix(v1.PREFIX_PARENT_SPAN_ID).
						WithKey(hex.EncodeToString(span.ParentSpanId))
				})
			}
			c.Range(span.StartTimeUnixNano, span.EndTimeUnixNano)
		}
	}
}
