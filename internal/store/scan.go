package store

import (
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/visit"
)

func CompileFilters(fs []*v1.Scan_Filter) (o visit.All) {
	for _, f := range fs {
		switch e := f.Value.(type) {
		case *v1.Scan_Filter_Base:
			switch e.Base.Prop {
			case v1.Scan_RESOURCE_SCHEMA:
				o.SetResourceSchema(e.Base.Value)
			case v1.Scan_SCOPE_SCHEMA:
				o.SetScopeSchema(e.Base.Value)
			case v1.Scan_SCOPE_NAME:
				o.SetScopeName(e.Base.Value)
			case v1.Scan_SCOPE_VERSION:
				o.SetScopVersion(e.Base.Value)
			case v1.Scan_NAME:
				o.SetName(e.Base.Value)
			case v1.Scan_TRACE_ID:
				o.SetTraceID(e.Base.Value)
			case v1.Scan_SPAN_ID:
				o.SetSpanID(e.Base.Value)
			case v1.Scan_PARENT_SPAN_ID:
				o.SetParentSpanID(e.Base.Value)
			case v1.Scan_LOGS_LEVEL:
				o.SetLogLevel(e.Base.Value)
			}
		case *v1.Scan_Filter_Attr:
			switch e.Attr.Prop {
			case v1.Scan_RESOURCE_ATTRIBUTES:
				o.SetResourceAttr(e.Attr.Key, e.Attr.Value)
			case v1.Scan_SCOPE_ATTRIBUTES:
				o.SetScopeAttr(e.Attr.Key, e.Attr.Value)
			case v1.Scan_ATTRIBUTES:
				o.SetScopeAttr(e.Attr.Key, e.Attr.Value)
			}
		}
	}
	return
}
