package render

import (
	"testing"

	"github.com/stretchr/testify/require"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
)

func TestAnyValue(t *testing.T) {
	type T struct {
		a *v1.AnyValue
		w string
	}

	ts := []T{
		{
			a: &v1.AnyValue{},
		},
		{
			a: &v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: "hello"}},
			w: `"hello"`,
		},
		{
			a: &v1.AnyValue{Value: &v1.AnyValue_BoolValue{BoolValue: true}},
			w: "true",
		},
		{
			a: &v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: 2}},
			w: "2",
		},
		{
			a: &v1.AnyValue{Value: &v1.AnyValue_DoubleValue{DoubleValue: 2.5}},
			w: "2.5",
		},
		{
			a: &v1.AnyValue{Value: &v1.AnyValue_BytesValue{BytesValue: []byte("hello")}},
			w: "68656c6c6f",
		},
		{
			a: &v1.AnyValue{Value: &v1.AnyValue_ArrayValue{ArrayValue: &v1.ArrayValue{
				Values: []*v1.AnyValue{
					{Value: &v1.AnyValue_StringValue{StringValue: "hello"}},
					{Value: &v1.AnyValue_BoolValue{BoolValue: true}},
					{Value: &v1.AnyValue_IntValue{IntValue: 2}},
					{Value: &v1.AnyValue_DoubleValue{DoubleValue: 2.5}},
					{Value: &v1.AnyValue_BytesValue{BytesValue: []byte("hello")}},
				},
			}}},
			w: "[\"hello\", true, 2, 2.5, 68656c6c6f]",
		},
		{
			a: &v1.AnyValue{Value: &v1.AnyValue_KvlistValue{KvlistValue: &v1.KeyValueList{
				Values: []*v1.KeyValue{
					{
						Key: "string",
						Value: &v1.AnyValue{
							Value: &v1.AnyValue_StringValue{StringValue: "hello"},
						},
					},
					{
						Key: "bool",
						Value: &v1.AnyValue{
							Value: &v1.AnyValue_BoolValue{BoolValue: true},
						},
					},
					{
						Key: "int",
						Value: &v1.AnyValue{
							Value: &v1.AnyValue_IntValue{IntValue: 2},
						},
					},
					{
						Key: "float",
						Value: &v1.AnyValue{
							Value: &v1.AnyValue_DoubleValue{DoubleValue: 2.5},
						},
					},
					{
						Key: "bytes",
						Value: &v1.AnyValue{
							Value: &v1.AnyValue_BytesValue{BytesValue: []byte("hello")},
						},
					},
				},
			}},
			},
			w: "{ string = \"hello\", bool = true, int = 2, float = 2.5, bytes = 68656c6c6f }",
		},
	}

	b := get()
	defer put(b)

	for _, v := range ts {
		b.Reset()
		value(b, v.a)
		require.Equal(t, v.w, b.String())
	}
}
