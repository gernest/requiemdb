package self

import (
	"context"
	"errors"

	"go.opentelemetry.io/contrib/instrumentation/host"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	collector_metrics "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	collector_trace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

var Resource = resource.NewWithAttributes(
	semconv.SchemaURL,
	semconv.ServiceName("requiemdb"),
)

type Shutdown interface {
	Shutdown(context.Context) error
}

type closers []Shutdown

func (o closers) Shutdown(ctx context.Context) error {
	e := make([]error, len(o))
	for i := range o {
		e[i] = o[i].Shutdown(ctx)
	}
	return errors.Join(e...)
}

func Setup(ctx context.Context, m collector_metrics.MetricsServiceServer, t collector_trace.TraceServiceServer) (Shutdown, error) {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(Resource),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(&Trace{Collector: t}),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	met := &Metrics{
		Collector: m,
	}
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(met),
		),
		sdkmetric.WithResource(Resource),
	)

	otel.SetMeterProvider(provider)
	err := host.Start()
	if err != nil {
		tp.Shutdown(context.Background())
		provider.Shutdown(context.Background())
		return nil, err
	}
	err = runtime.Start()
	if err != nil {
		tp.Shutdown(context.Background())
		provider.Shutdown(context.Background())
		return nil, err
	}
	return closers{tp, provider}, nil
}
