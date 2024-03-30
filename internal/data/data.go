package data

import (
	"github.com/cespare/xxhash/v2"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	"google.golang.org/protobuf/proto"
)

func Collapse(ts []*v1.Data) *v1.Data {
	return nil
}

func hashAttributes(h *xxhash.Digest, buf []byte, kv []*commonv1.KeyValue) []byte {
	for _, v := range kv {
		buf, _ = proto.MarshalOptions{}.MarshalAppend(buf[:0], v)
		h.Write(buf)
	}
	return buf
}
