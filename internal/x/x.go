package x

import (
	"github.com/gernest/requiemdb/internal/compress"
	"google.golang.org/protobuf/proto"
)

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

func HighBits(v uint64) uint64 { return v >> 16 }
func LowBits(v uint64) uint16  { return uint16(v & 0xFFFF) }
