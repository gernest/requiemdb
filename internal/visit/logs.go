package visit

import (
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

type Logs struct{}

var _ Visitor[*logsv1.LogsData] = (*Logs)(nil)

func (Logs) Visit(data *logsv1.LogsData, visitor Visit) *logsv1.LogsData {
	var resources []*logsv1.ResourceLogs
	start, end := visitor.TimeRange()
	for _, rm := range data.ResourceLogs {
		if !AcceptResource(rm, visitor) {
			continue
		}
		// we avoid initializing this in case we skip selecting this resource
		var a *logsv1.ResourceLogs

		for _, sm := range rm.ScopeLogs {
			if !AcceptScope(sm.SchemaUrl, sm.Scope, visitor) {
				continue
			}
			var scope *logsv1.ScopeLogs
			// We have the right scope now we need to select data points
			for _, ms := range sm.LogRecords {
				if !visitor.AcceptLogLevel(ms.SeverityText) {
					continue
				}
				if !AcceptDataPoint(ms, start, end, visitor) {
					continue
				}
				if scope == nil {
					scope = &logsv1.ScopeLogs{
						Scope: sm.Scope,
					}
				}
				scope.LogRecords = append(scope.LogRecords, ms)
			}
			if scope != nil {
				if a == nil {
					a = &logsv1.ResourceLogs{
						Resource:  rm.Resource,
						SchemaUrl: rm.SchemaUrl,
					}
				}
				a.ScopeLogs = append(a.ScopeLogs, scope)
			}
		}
		if a != nil {
			resources = append(resources, a)
		}
	}
	return &logsv1.LogsData{ResourceLogs: resources}
}
