package data

import (
	"sync"

	metricsV1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func CollapseMetrics(ds []*metricsV1.MetricsData) *metricsV1.MetricsData {
	xm := newMetricsSorter()
	defer xm.Release()
	var idx int
	for _, md := range ds {
		for _, rm := range md.ResourceMetrics {
			rh := xm.hashResource(rm.SchemaUrl, rm.Resource)
			resource, ok := xm.resource[rh]
			if !ok {
				resource = rm
				idx++
				xm.resource[rh] = resource
				xm.order[rh] = idx

				// add all scopes for this resource
				for _, sm := range rm.ScopeMetrics {
					sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
					xm.scope[sh] = sm
					for _, m := range sm.Metrics {
						mh := xm.hashMetrics(sh, m.Name)
						xm.metrics[mh] = m
					}
				}
				continue
			}
			for _, sm := range rm.ScopeMetrics {
				sh := xm.hashScope(rh, sm.SchemaUrl, sm.Scope)
				scope, ok := xm.scope[sh]
				if !ok {
					xm.scope[sh] = sm
					resource.ScopeMetrics = append(resource.ScopeMetrics, sm)
					for _, m := range sm.Metrics {
						mh := xm.hashMetrics(sh, m.Name)
						xm.metrics[mh] = m
					}
					continue
				}
				for _, m := range sm.Metrics {
					mh := xm.hashMetrics(sh, m.Name)
					om, ok := xm.metrics[mh]
					if !ok {
						scope.Metrics = append(scope.Metrics, m)
						xm.metrics[mh] = m
						continue
					}
					if gauge := m.GetGauge(); gauge != nil {
						om.GetGauge().DataPoints = append(om.GetGauge().DataPoints, gauge.DataPoints...)
					}
					if sum := m.GetSum(); sum != nil {
						om.GetSum().DataPoints = append(om.GetSum().DataPoints, sum.DataPoints...)
					}
					if hist := m.GetHistogram(); hist != nil {
						om.GetHistogram().DataPoints = append(om.GetHistogram().DataPoints, hist.DataPoints...)
					}
					if ehist := m.GetExponentialHistogram(); ehist != nil {
						om.GetExponentialHistogram().DataPoints = append(om.GetExponentialHistogram().DataPoints,
							ehist.DataPoints...)
					}
					if sum := m.GetSummary(); sum != nil {
						om.GetSummary().DataPoints = append(om.GetSummary().DataPoints, sum.DataPoints...)
					}
				}
			}
		}
	}
	xm.Sort()
	return &metricsV1.MetricsData{ResourceMetrics: xm.Result()}
}

type metricsSorter struct {
	*Sorter[*metricsV1.Metric, *metricsV1.ResourceMetrics, *metricsV1.ScopeMetrics]
}

func newMetricsSorter() *metricsSorter {
	return metricsPool.Get().(*metricsSorter)
}

func (ms *metricsSorter) Release() {
	ms.Reset()
	metricsPool.Put(ms)
}

var metricsPool = &sync.Pool{New: func() any {
	return &metricsSorter{
		Sorter: newSorter[*metricsV1.Metric, *metricsV1.ResourceMetrics, *metricsV1.ScopeMetrics](),
	}
}}
