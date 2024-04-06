package render

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"

	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
)

func attr(attr []*commonv1.KeyValue) string {
	b := get()
	defer put(b)
	b.WriteByte('{')
	for i, kv := range attr {
		if i != 0 {
			b.WriteByte(',')
		}
		b.WriteString(kv.Key)
		switch e := kv.Value.Value.(type) {
		case *commonv1.AnyValue_StringValue:
			fmt.Fprintf(b, "%q", e.StringValue)
		case *commonv1.AnyValue_BoolValue:
			b.WriteString(strconv.FormatBool(e.BoolValue))
		case *commonv1.AnyValue_DoubleValue:
			fmt.Fprintf(b, "%v", e.DoubleValue)
		case *commonv1.AnyValue_IntValue:
			fmt.Fprintf(b, "%v", e.IntValue)
		}
	}
	b.WriteByte('}')
	return b.String()
}

func get() *bytes.Buffer {
	return bytesPool.Get().(*bytes.Buffer)
}

func put(b *bytes.Buffer) {
	b.Reset()
	bytesPool.Put(b)
}

var bytesPool = &sync.Pool{New: func() any { return new(bytes.Buffer) }}
