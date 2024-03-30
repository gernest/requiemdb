package data

import (
	"github.com/cespare/xxhash/v2"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	"google.golang.org/protobuf/proto"
)

func hashAttributes(h *xxhash.Digest, buf []byte, kv []*commonv1.KeyValue) []byte {
	for _, v := range kv {
		buf, _ = proto.MarshalOptions{}.MarshalAppend(buf[:0], v)
		h.Write(buf)
	}
	return buf
}
