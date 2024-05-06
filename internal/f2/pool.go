package f2

import (
	"sync"

	storev1 "github.com/gernest/requiemdb/gen/go/store/v1"
	"google.golang.org/protobuf/proto"
)

var (
	metricPool = NewPool(func() *storev1.Metric { return new(storev1.Metric) })
	histPool   = NewPool(func() *storev1.HistogramDataPoint { return new(storev1.HistogramDataPoint) })
	exHistPool = NewPool(func() *storev1.ExponentialHistogramDataPoint { return new(storev1.ExponentialHistogramDataPoint) })
	sumtPool   = NewPool(func() *storev1.SummaryDataPoint { return new(storev1.SummaryDataPoint) })
	logPool    = NewPool(func() *storev1.LogRecord { return new(storev1.LogRecord) })
	spanPool   = NewPool(func() *storev1.Span { return new(storev1.Span) })
)

type Pool[T proto.Message] struct {
	sync sync.Pool
}

func NewPool[T proto.Message](init func() T) *Pool[T] {
	p := &Pool[T]{}
	p.sync.New = func() any { return init() }
	return p
}

func (p *Pool[T]) Get() T {
	return p.sync.Get().(T)
}

func (p *Pool[T]) Put(v T) {
	proto.Reset(v)
	p.sync.Put(v)
}
