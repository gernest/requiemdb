package visit

import (
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

type Trace struct{}

var _ Visitor[*tracev1.TracesData] = (*Trace)(nil)

func (Trace) Visit(data *tracev1.TracesData, visitor Visit) *tracev1.TracesData {
	var resources []*tracev1.ResourceSpans
	start, end := visitor.TimeRange()
	for _, rm := range data.ResourceSpans {
		if !AcceptResource(rm, visitor) {
			continue
		}
		// we avoid initializing this in case we skip selecting this resource
		var a *tracev1.ResourceSpans

		for _, sm := range rm.ScopeSpans {
			if !AcceptScope(sm.SchemaUrl, sm.Scope, visitor) {
				continue
			}
			var scope *tracev1.ScopeSpans
			// We have the right scope now we need to select data points
			for _, ms := range sm.Spans {
				if !(visitor.AcceptName(ms.Name) &&
					visitor.AcceptTraceID(ms.TraceId) &&
					visitor.AcceptParentSpanID(ms.SpanId) &&
					visitor.AcceptAttributes(ms.Attributes)) {
					continue
				}
				ts := ms.StartTimeUnixNano
				if !(ts >= start && ts < end) {
					continue
				}
				if scope == nil {
					scope = &tracev1.ScopeSpans{
						Scope: sm.Scope,
					}
				}
				scope.Spans = append(scope.Spans, ms)
			}
			if scope != nil {
				if a == nil {
					a = &tracev1.ResourceSpans{
						Resource:  rm.Resource,
						SchemaUrl: rm.SchemaUrl,
					}
				}
				a.ScopeSpans = append(a.ScopeSpans, scope)
			}
		}
		if a != nil {
			resources = append(resources, a)
		}
	}
	return &tracev1.TracesData{ResourceSpans: resources}
}
