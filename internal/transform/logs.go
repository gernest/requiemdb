package transform

import (
	"encoding/hex"

	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

func (c *Context) transformLogs(rm *logsv1.ResourceLogs) {
	if rm.SchemaUrl != "" {
		c.Label(func(lbl *labels.Label) {
			lbl.WithResource(v1.RESOURCE_LOGS).
				WithPrefix(v1.PREFIX_RESOURCE_SCHEMA).
				WithKey(rm.SchemaUrl)

		})
	}
	if rm.Resource != nil {
		c.attributes(v1.RESOURCE_LOGS, v1.PREFIX_RESOURCE_ATTRIBUTES, rm.Resource.Attributes)
	}
	for _, sm := range rm.ScopeLogs {
		if sm.SchemaUrl != "" {
			c.Label(func(lbl *labels.Label) {
				lbl.WithResource(v1.RESOURCE_LOGS).
					WithPrefix(v1.PREFIX_SCOPE_SCHEMA).
					WithKey(sm.SchemaUrl)

			})
		}
		if sc := sm.Scope; sc != nil {
			if sc.Name != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_LOGS).
						WithPrefix(v1.PREFIX_SCOPE_NAME).
						WithKey(sc.Name)

				})
			}
			if sc.Version != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_LOGS).
						WithPrefix(v1.PREFIX_SCOPE_VERSION).
						WithKey(sc.Version)
				})
			}
			c.attributes(v1.RESOURCE_LOGS, v1.PREFIX_SCOPE_ATTRIBUTES, sc.Attributes)
		}

		for _, r := range sm.LogRecords {
			c.Timestamp(r.TimeUnixNano)
			if r.SeverityText != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_LOGS).
						WithPrefix(v1.PREFIX_LOGS_LEVEL).
						WithKey(r.SeverityText)
				})
			}
			c.attributes(v1.RESOURCE_LOGS, v1.PREFIX_ATTRIBUTES, r.Attributes)
			if r.SpanId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_LOGS).
						WithPrefix(v1.PREFIX_SPAN_ID).
						WithKey(hex.EncodeToString(r.SpanId))
				})
			}
			if r.TraceId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.WithResource(v1.RESOURCE_LOGS).
						WithPrefix(v1.PREFIX_TRACE_ID).
						WithKey(hex.EncodeToString(r.TraceId))
				})
			}
		}
	}
}
