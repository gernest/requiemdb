package x

import (
	"github.com/requiemdb/requiemdb/internal/compress"
	"google.golang.org/protobuf/proto"
)

func Compress(data []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	return compress.Compress(data)
}

func Decompress(msg proto.Message, size *int64) func(data []byte) error {
	return func(data []byte) error {
		data, err := compress.Decompress(data)
		if err != nil {
			return err
		}
		*size = int64(len(data))
		return proto.Unmarshal(data, msg)
	}
}
