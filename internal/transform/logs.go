package transform

import (
	"encoding/hex"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/labels"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

func (c *Context) transformLogs(rm *logsv1.ResourceLogs) {
	c.addResource(rm)
	for _, sm := range rm.ScopeLogs {
		c.addScope(sm)

		for _, r := range sm.LogRecords {
			c.Timestamp(r.TimeUnixNano)
			if r.SeverityText != "" {
				c.Label(func(lbl *labels.Label) {
					lbl.
						WithPrefix(v1.PREFIX_LOGS_LEVEL).
						WithKey(r.SeverityText)
				})
			}
			c.attributes(v1.PREFIX_ATTRIBUTES, r.Attributes)
			if r.SpanId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.
						WithPrefix(v1.PREFIX_SPAN_ID).
						WithKey(hex.EncodeToString(r.SpanId))
				})
			}
			if r.TraceId != nil {
				c.Label(func(lbl *labels.Label) {
					lbl.
						WithPrefix(v1.PREFIX_TRACE_ID).
						WithKey(hex.EncodeToString(r.TraceId))
				})
			}
		}
	}
}
