package arrow3

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/apache/arrow/go/v17/arrow/memory"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestRead(t *testing.T) {
	msg := &metricsv1.MetricsData{}
	b := build(msg.ProtoReflect())
	b.build(memory.DefaultAllocator)
	b.append(msg.ProtoReflect())
	msg.ResourceMetrics = []*metricsv1.ResourceMetrics{
		{ScopeMetrics: []*metricsv1.ScopeMetrics{
			{Metrics: []*metricsv1.Metric{
				{Name: "check", Data: &metricsv1.Metric_Gauge{
					Gauge: &metricsv1.Gauge{
						DataPoints: []*metricsv1.NumberDataPoint{
							{TimeUnixNano: 16, Value: &metricsv1.NumberDataPoint_AsInt{
								AsInt: 18,
							},
								Attributes: []*commonv1.KeyValue{
									{Key: "key", Value: &commonv1.AnyValue{
										Value: &commonv1.AnyValue_StringValue{
											StringValue: "value",
										},
									}},
								},
							},
						},
					},
				}},
			}},
		}},
	}
	b.append(msg.ProtoReflect())

	var o bytes.Buffer
	err := b.WriteParquet(&o)
	if err != nil {
		t.Fatal(err)
	}
	matchBytes(t, "testdata/otel_metrics_data.parquet", o.Bytes())

	f, err := os.Open("testdata/otel_metrics_data.parquet")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r, err := b.Read(context.Background(), f, nil)
	if err != nil {
		t.Fatal(err)
	}
	data, err := r.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	match(t, "testdata/otel_metrics_data_parquet_read.json", string(data))

	rd := unmarshal[*metricsv1.MetricsData](b.root, r, []int{1})

	data, err = protojson.Marshal(rd[0])
	if err != nil {
		t.Fatal(err)
	}
	match(t, "testdata/otel_metrics_data_parquet_read_decoded.json", string(data))
}

func TestWriteMultiple(t *testing.T) {
	msg := &metricsv1.MetricsData{}
	b := build(msg.ProtoReflect())
	b.build(memory.DefaultAllocator)
	b.append(msg.ProtoReflect())
	msg.ResourceMetrics = []*metricsv1.ResourceMetrics{
		{ScopeMetrics: []*metricsv1.ScopeMetrics{
			{Metrics: []*metricsv1.Metric{
				{Name: "check", Data: &metricsv1.Metric_Gauge{
					Gauge: &metricsv1.Gauge{
						DataPoints: []*metricsv1.NumberDataPoint{
							{TimeUnixNano: 16, Value: &metricsv1.NumberDataPoint_AsInt{
								AsInt: 18,
							},
								Attributes: []*commonv1.KeyValue{
									{Key: "key", Value: &commonv1.AnyValue{
										Value: &commonv1.AnyValue_StringValue{
											StringValue: "value",
										},
									}},
								},
							},
						},
					},
				}},
			}},
		}},
	}
	b.append(msg.ProtoReflect())

	var o bytes.Buffer
	r := b.NewRecord()
	err := b.WriteParquetRecords(&o, r, r)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkWriteParquet(tb *testing.B) {
	msg := &metricsv1.MetricsData{}
	b := build(msg.ProtoReflect())
	b.build(memory.DefaultAllocator)
	b.append(msg.ProtoReflect())
	msg.ResourceMetrics = []*metricsv1.ResourceMetrics{
		{ScopeMetrics: []*metricsv1.ScopeMetrics{
			{Metrics: []*metricsv1.Metric{
				{Name: "check", Data: &metricsv1.Metric_Gauge{
					Gauge: &metricsv1.Gauge{
						DataPoints: []*metricsv1.NumberDataPoint{
							{TimeUnixNano: 16, Value: &metricsv1.NumberDataPoint_AsInt{
								AsInt: 18,
							},
								Attributes: []*commonv1.KeyValue{
									{Key: "key", Value: &commonv1.AnyValue{
										Value: &commonv1.AnyValue_StringValue{
											StringValue: "value",
										},
									}},
								},
							},
						},
					},
				}},
			}},
		}},
	}
	b.append(msg.ProtoReflect())
	r := b.NewRecord()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		err := b.WriteParquetRecords(io.Discard, r, r)
		if err != nil {
			tb.Fatal(err)
		}
	}
}
