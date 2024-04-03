package data

import (
	"sync"

	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func CollapseTrace(ds []*tracev1.TracesData) *tracev1.TracesData {
	xm := newtraceSorter()
	defer xm.Release()
	var idx int
	for _, md := range ds {
		for _, rm := range md.ResourceSpans {
			rh := xm.hashResource(rm.SchemaUrl, rm.Resource)
			resource, ok := xm.resource[rh]
			if !ok {
				resource = rm
				idx++
				xm.resource[rh] = resource
				xm.order[rh] = idx

				// add all scopes for this resource
				for _, sm := range rm.ScopeSpans {
					sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
					xm.scope[sh] = sm
				}
				continue
			}
			for _, sm := range rm.ScopeSpans {
				sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
				scope, ok := xm.scope[sh]
				if !ok {
					xm.scope[sh] = sm
					resource.ScopeSpans = append(resource.ScopeSpans, sm)
					continue
				}
				scope.Spans = append(scope.Spans, sm.Spans...)
			}
		}
	}
	xm.Sort()
	return &tracev1.TracesData{ResourceSpans: xm.Result()}
}

type traceSorter struct {
	*Sorter[*tracev1.Span, *tracev1.ResourceSpans, *tracev1.ScopeSpans]
}

func newtraceSorter() *traceSorter {
	return tracePool.Get().(*traceSorter)
}

func (l *traceSorter) Release() {
	l.Reset()
	tracePool.Put(l)
}

var tracePool = &sync.Pool{New: func() any {
	return &traceSorter{
		Sorter: newSorter[*tracev1.Span, *tracev1.ResourceSpans, *tracev1.ScopeSpans](),
	}
}}
