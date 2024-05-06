package f2

import (
	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/gernest/arrow3"
	v1 "github.com/gernest/requiemdb/gen/go/store/v1"
	"github.com/gernest/roaring/shardwidth"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type ID interface {
	Next() uint64
}

type Translate interface {
	Tr([]byte) uint64
}

var (
	attrPrefix = []byte("0:")
	equal      = []byte("=")
	nameKey    = []byte("__name__")
)

type Metrics struct {
	build     *arrow3.Schema[*v1.Metric]
	id        ID
	tr        Translate
	positions roaring64.Bitmap
}

func (m *Metrics) Append(data *metricsv1.MetricsData) {
	a := newAttr()
	defer a.Release()

	for _, rs := range data.ResourceMetrics {
		a.Resource(rs.Resource)
		for _, sc := range rs.ScopeMetrics {
			a.Scope(sc.Scope)

			for _, ms := range sc.Metrics {
				mx := metricPool.Get()

				mx.Name = ms.Name
				mx.Description = ms.Description
				mx.Unit = ms.Unit

				if v := ms.GetGauge(); v != nil {
					mx.Kind = v1.Metric_GAUGE
					m.gauge(mx, a, v)
				}
				if v := ms.GetSum(); v != nil {
					mx.Kind = v1.Metric_SUM
					m.sum(mx, a, v)
				}
				if v := ms.GetHistogram(); v != nil {
					mx.Kind = v1.Metric_HISTOGRAM
				}
				if v := ms.GetExponentialHistogram(); v != nil {
					mx.Kind = v1.Metric_EXPONENTIAL_HISTOGRAM
				}
				if v := ms.GetSummary(); v != nil {
					mx.Kind = v1.Metric_SUMMARY
				}
				metricPool.Put(mx)
			}

		}
	}
}

func (m *Metrics) gauge(mx *v1.Metric, a *Attr, g *metricsv1.Gauge) {
	name := []byte(mx.Name)
	var columns *roaring64.Bitmap
	for _, dp := range g.DataPoints {
		mx.StartTimeUnixNano = dp.StartTimeUnixNano
		mx.TimeUnixNano = dp.TimeUnixNano
		mx.Exemplars = dp.Exemplars
		switch e := dp.Value.(type) {
		case *metricsv1.NumberDataPoint_AsDouble:
			mx.Gauge = &e.AsDouble
		case *metricsv1.NumberDataPoint_AsInt:
			f := float64(e.AsInt)
			mx.Gauge = &f
		}
		mx.Attributes, columns, mx.Hash = a.Attr(name, dp.Attributes)
		id := m.id.Next()
		mx.Id = id
		it := columns.Iterator()
		for it.HasNext() {
			m.positions.Add(pos(id, it.Next()))
		}
		m.build.Append(mx)
	}
}

func (m *Metrics) sum(mx *v1.Metric, a *Attr, g *metricsv1.Sum) {
	name := []byte(mx.Name)
	mx.AggregationTemporality = g.AggregationTemporality
	mx.IsMonotonic = g.IsMonotonic
	var columns *roaring64.Bitmap
	for _, dp := range g.DataPoints {
		mx.StartTimeUnixNano = dp.StartTimeUnixNano
		mx.TimeUnixNano = dp.TimeUnixNano
		mx.Exemplars = dp.Exemplars
		switch e := dp.Value.(type) {
		case *metricsv1.NumberDataPoint_AsDouble:
			mx.Gauge = &e.AsDouble
		case *metricsv1.NumberDataPoint_AsInt:
			f := float64(e.AsInt)
			mx.Gauge = &f
		}
		mx.Attributes, columns, mx.Hash = a.Attr(name, dp.Attributes)
		id := m.id.Next()
		mx.Id = id
		it := columns.Iterator()
		for it.HasNext() {
			m.positions.Add(pos(id, it.Next()))
		}
		m.build.Append(mx)
	}
}

func pos(row, col uint64) uint64 {
	return row*shardwidth.ShardWidth + (col % shardwidth.ShardWidth)
}
