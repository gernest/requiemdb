package f2

import (
	"errors"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/store/v1"
	"github.com/gernest/requiemdb/internal/arrow3"
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
	metrics    = []byte("metrics/")
)

type Metrics struct {
	build     *arrow3.Schema[*v1.Metric]
	id        *idGen
	tr        *cacheTr
	positions roaring64.Bitmap
}

func NewMetrics(db *badger.DB) (*Metrics, error) {
	tr, err := newCacheTr(db, metrics)
	if err != nil {
		return nil, err
	}
	s, err := arrow3.New[*v1.Metric](memory.DefaultAllocator)
	if err != nil {
		tr.Close()
		return nil, err
	}
	id, err := newID(db, metrics)
	if err != nil {
		s.Release()
		tr.Close()
		return nil, err
	}
	return &Metrics{
		build: s,
		tr:    tr,
		id:    id,
	}, nil
}

func (m *Metrics) Close() error {
	return errors.Join(
		m.tr.Close(), m.id.Close(),
	)
}

func (m *Metrics) Append(data *metricsv1.MetricsData) {
	a := newAttr().Reset(m.tr)
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
					m.hist(mx, a, v)
				}
				if v := ms.GetExponentialHistogram(); v != nil {
					mx.Kind = v1.Metric_EXPONENTIAL_HISTOGRAM
					m.expHist(mx, a, v)
				}
				if v := ms.GetSummary(); v != nil {
					mx.Kind = v1.Metric_SUMMARY
					m.summary(mx, a, v)
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
func (m *Metrics) summary(mx *v1.Metric, a *Attr, g *metricsv1.Summary) {
	name := []byte(mx.Name)
	var columns *roaring64.Bitmap
	for _, dp := range g.DataPoints {
		mx.StartTimeUnixNano = dp.StartTimeUnixNano
		mx.TimeUnixNano = dp.TimeUnixNano
		mx.Summary = &v1.SummaryDataPoint{
			Count:          dp.Count,
			Sum:            dp.Sum,
			QuantileValues: dp.QuantileValues,
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

func (m *Metrics) hist(mx *v1.Metric, a *Attr, g *metricsv1.Histogram) {
	name := []byte(mx.Name)
	mx.AggregationTemporality = g.AggregationTemporality
	var columns *roaring64.Bitmap
	for _, dp := range g.DataPoints {
		mx.StartTimeUnixNano = dp.StartTimeUnixNano
		mx.TimeUnixNano = dp.TimeUnixNano
		mx.Exemplars = dp.Exemplars
		mx.Histogram = &v1.HistogramDataPoint{
			Count:          dp.Count,
			Sum:            dp.Sum,
			BucketCounts:   dp.BucketCounts,
			ExplicitBounds: dp.ExplicitBounds,
			Min:            dp.Min,
			Max:            dp.Max,
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

func (m *Metrics) expHist(mx *v1.Metric, a *Attr, g *metricsv1.ExponentialHistogram) {
	name := []byte(mx.Name)
	mx.AggregationTemporality = g.AggregationTemporality
	var columns *roaring64.Bitmap
	for _, dp := range g.DataPoints {
		mx.StartTimeUnixNano = dp.StartTimeUnixNano
		mx.TimeUnixNano = dp.TimeUnixNano
		mx.Exemplars = dp.Exemplars
		mx.ExponentialHistogram = &v1.ExponentialHistogramDataPoint{
			Count:         dp.Count,
			Sum:           dp.Sum,
			Scale:         dp.Scale,
			ZeroCount:     dp.ZeroCount,
			Positive:      dp.Positive,
			Negative:      dp.Negative,
			Min:           dp.Min,
			Max:           dp.Max,
			ZeroThreshold: dp.ZeroThreshold,
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
