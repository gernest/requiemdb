package data

import (
	"sync"

	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

func CollapseLogs(ds []*logsv1.LogsData) *logsv1.LogsData {
	xm := newLogsSorter()
	defer xm.Release()
	var idx int
	for _, md := range ds {
		for _, rm := range md.ResourceLogs {
			rh := xm.hashResource(rm.SchemaUrl, rm.Resource)
			resource, ok := xm.resource[rh]
			if !ok {
				resource = rm
				idx++
				xm.resource[rh] = resource
				xm.order[rh] = idx

				// add all scopes for this resource
				for _, sm := range rm.ScopeLogs {
					sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
					xm.scope[sh] = sm
				}
				continue
			}
			for _, sm := range rm.ScopeLogs {
				sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
				scope, ok := xm.scope[sh]
				if !ok {
					xm.scope[sh] = sm
					resource.ScopeLogs = append(resource.ScopeLogs, sm)
					continue
				}
				scope.LogRecords = append(scope.LogRecords, sm.LogRecords...)
			}
		}
	}
	xm.Sort()
	return &logsv1.LogsData{ResourceLogs: xm.Result()}
}

type logsSorter struct {
	*Sorter[*logsv1.LogRecord, *logsv1.ResourceLogs, *logsv1.ScopeLogs]
}

func newLogsSorter() *logsSorter {
	return logsPool.Get().(*logsSorter)
}

func (l *logsSorter) Release() {
	l.Reset()
	logsPool.Put(l)
}

var logsPool = &sync.Pool{New: func() any {
	return &logsSorter{
		Sorter: newSorter[*logsv1.LogRecord, *logsv1.ResourceLogs, *logsv1.ScopeLogs](),
	}
}}
