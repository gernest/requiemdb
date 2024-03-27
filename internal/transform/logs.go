package transform

import (
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

func (c *Context) transformLogs(rm *logsv1.ResourceLogs) {
	if rm.SchemaUrl != "" {
		c.label(
			labels.NewBytes(v1.RESOURCE_LOGS, v1.PREFIX_RESOURCE_SCHEMA).Value(rm.SchemaUrl),
		)
	}
	if rm.Resource != nil {
		c.attributes(v1.RESOURCE_LOGS, v1.PREFIX_RESOURCE_ATTRIBUTES, rm.Resource.Attributes)
	}
	for _, sm := range rm.ScopeLogs {
		if sm.SchemaUrl != "" {
			c.label(
				labels.NewBytes(v1.RESOURCE_LOGS, v1.PREFIX_SCOPE_SCHEMA).Value(sm.SchemaUrl),
			)
		}
		if sc := sm.Scope; sc != nil {
			if sc.Name != "" {
				c.label(
					labels.NewBytes(v1.RESOURCE_LOGS, v1.PREFIX_SCOPE_NAME).Value(sc.Name),
				)
			}
			if sc.Version != "" {
				c.label(
					labels.NewBytes(v1.RESOURCE_LOGS, v1.PREFIX_SCOPE_VERSION).Value(sc.Version),
				)
			}
			c.attributes(v1.RESOURCE_LOGS, v1.PREFIX_SCOPE_ATTRIBUTES, sc.Attributes)
		}

		for _, r := range sm.LogRecords {
			c.Timestamp(r.TimeUnixNano)
			if r.SeverityText != "" {
				c.label(
					labels.NewBytes(v1.RESOURCE_LOGS, v1.PREFIX_LOGS_LEVEL).Value(r.SeverityText),
				)
			}
			c.attributes(v1.RESOURCE_LOGS, v1.PREFIX_ATTRIBUTES, r.Attributes)
			if r.SpanId != nil {
				c.label(
					labels.NewBytes(v1.RESOURCE_LOGS, v1.PREFIX_SPAN_ID).ValueBytes(r.SpanId),
				)
			}
			if r.TraceId != nil {
				c.label(
					labels.NewBytes(v1.RESOURCE_LOGS, v1.PREFIX_TRACE_ID).ValueBytes(r.TraceId),
				)
			}
		}
	}
}
