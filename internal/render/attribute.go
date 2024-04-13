package render

import (
	"bytes"
	"fmt"
	"strconv"

	v1 "go.opentelemetry.io/proto/otlp/common/v1"
)

func attr(kv []*v1.KeyValue) string {
	b := get()
	defer put(b)
	value(b, &v1.AnyValue{Value: &v1.AnyValue_KvlistValue{
		KvlistValue: &v1.KeyValueList{Values: kv},
	}})
	return b.String()
}

func value(b *bytes.Buffer, v *v1.AnyValue) {
	switch e := v.Value.(type) {
	case *v1.AnyValue_StringValue:
		fmt.Fprintf(b, "%q", e.StringValue)
	case *v1.AnyValue_BoolValue:
		b.WriteString(strconv.FormatBool(e.BoolValue))
	case *v1.AnyValue_IntValue:
		b.WriteString(strconv.FormatInt(e.IntValue, 10))
	case *v1.AnyValue_DoubleValue:
		fmt.Fprintf(b, "%v", e.DoubleValue)
	case *v1.AnyValue_ArrayValue:
		b.WriteByte('[')
		for i, a := range e.ArrayValue.Values {
			if i != 0 {
				b.WriteByte(',')
				b.WriteByte(' ')
			}
			value(b, a)
		}
		b.WriteByte(']')
	case *v1.AnyValue_KvlistValue:
		b.WriteByte('{')
		for i, a := range e.KvlistValue.Values {
			if i != 0 {
				b.WriteByte(',')
				b.WriteByte(' ')
			} else {
				b.WriteByte(' ')
			}
			b.WriteString(a.Key)
			b.WriteString(" = ")
			value(b, a.Value)
		}
		if len(e.KvlistValue.Values) > 0 {
			b.WriteByte(' ')
		}
		b.WriteByte('}')
	case *v1.AnyValue_BytesValue:
		fmt.Fprintf(b, "%x", e.BytesValue)
	}
}
