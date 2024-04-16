package visit

import (
	metricsV1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type Metrics struct{}

func (Metrics) Visit(data *metricsV1.MetricsData, visitor *All) *metricsV1.MetricsData {
	var resources []*metricsV1.ResourceMetrics
	for _, rm := range data.ResourceMetrics {
		if !AcceptResource(rm, visitor) {
			continue
		}
		// we avoid initializing this in case we skip selecting this resource
		var a *metricsV1.ResourceMetrics

		for _, sm := range rm.ScopeMetrics {
			if !AcceptScope(sm.SchemaUrl, sm.Scope, visitor) {
				continue
			}
			var scope *metricsV1.ScopeMetrics

			// We have the right scope now we need to select data points
			for _, ms := range sm.Metrics {
				if !visitor.AcceptName(ms.Name) {
					continue
				}
				if v := ms.GetGauge(); v != nil {
					var points []*metricsV1.NumberDataPoint
					for _, d := range v.DataPoints {
						if !AcceptDataPoint(d, visitor) {
							continue
						}
						points = append(points, d)

					}
					if len(points) > 0 {
						if scope == nil {
							scope = &metricsV1.ScopeMetrics{
								Scope: sm.Scope,
							}
						}
						scope.Metrics = append(scope.Metrics, &metricsV1.Metric{
							Name: ms.Name,
							Data: &metricsV1.Metric_Gauge{
								Gauge: &metricsV1.Gauge{
									DataPoints: points,
								},
							},
						})
					}
				}
				if v := ms.GetSum(); v != nil {
					var points []*metricsV1.NumberDataPoint
					for _, d := range v.DataPoints {
						if !AcceptDataPoint(d, visitor) {
							continue
						}
						points = append(points, d)
					}
					if len(points) > 0 {
						if scope == nil {
							scope = &metricsV1.ScopeMetrics{
								Scope: sm.Scope,
							}
						}
						scope.Metrics = append(scope.Metrics, &metricsV1.Metric{
							Name: ms.Name,
							Data: &metricsV1.Metric_Sum{
								Sum: &metricsV1.Sum{
									DataPoints:             points,
									AggregationTemporality: v.AggregationTemporality,
									IsMonotonic:            v.IsMonotonic,
								},
							},
						})
					}
				}
				if v := ms.GetHistogram(); v != nil {
					var points []*metricsV1.HistogramDataPoint
					for _, d := range v.DataPoints {
						if !AcceptDataPoint(d, visitor) {
							continue
						}
						points = append(points, d)

					}
					if len(points) > 0 {
						if scope == nil {
							scope = &metricsV1.ScopeMetrics{
								Scope: sm.Scope,
							}
						}
						scope.Metrics = append(scope.Metrics, &metricsV1.Metric{
							Name: ms.Name,
							Data: &metricsV1.Metric_Histogram{
								Histogram: &metricsV1.Histogram{
									DataPoints:             points,
									AggregationTemporality: v.AggregationTemporality,
								},
							},
						})
					}
				}
				if v := ms.GetExponentialHistogram(); v != nil {
					var points []*metricsV1.ExponentialHistogramDataPoint
					for _, d := range v.DataPoints {
						if !AcceptDataPoint(d, visitor) {
							continue
						}
						points = append(points, d)

					}
					if len(points) > 0 {
						if scope == nil {
							scope = &metricsV1.ScopeMetrics{
								Scope: sm.Scope,
							}
						}
						scope.Metrics = append(scope.Metrics, &metricsV1.Metric{
							Name: ms.Name,
							Data: &metricsV1.Metric_ExponentialHistogram{
								ExponentialHistogram: &metricsV1.ExponentialHistogram{
									DataPoints:             points,
									AggregationTemporality: v.AggregationTemporality,
								},
							},
						})
					}
				}
				if v := ms.GetSummary(); v != nil {
					var points []*metricsV1.SummaryDataPoint
					for _, d := range v.DataPoints {
						if !AcceptDataPoint(d, visitor) {
							continue
						}
						points = append(points, d)

					}
					if len(points) > 0 {
						if scope == nil {
							scope = &metricsV1.ScopeMetrics{
								Scope: sm.Scope,
							}
						}
						scope.Metrics = append(scope.Metrics, &metricsV1.Metric{
							Name: ms.Name,
							Data: &metricsV1.Metric_Summary{
								Summary: &metricsV1.Summary{
									DataPoints: points,
								},
							},
						})
					}
				}

			}
			if scope != nil {
				if a == nil {
					a = &metricsV1.ResourceMetrics{
						Resource:  rm.Resource,
						SchemaUrl: rm.SchemaUrl,
					}
				}
				a.ScopeMetrics = append(a.ScopeMetrics, scope)
			}
		}
		if a != nil {
			resources = append(resources, a)
		}
	}
	return &metricsV1.MetricsData{ResourceMetrics: resources}
}
