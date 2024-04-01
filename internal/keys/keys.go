package keys

import (
	"encoding/binary"

	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
)

type Sample struct {
	Namespace uint64
	Partition uint64
	Resource  v1.RESOURCE
	ID        uint64
}

func (s *Sample) Encode() []byte {
	o := make([]byte, 8+8+4+8)
	binary.LittleEndian.PutUint64(o, s.Namespace)
	binary.LittleEndian.PutUint64(o[8:], s.Partition)
	binary.LittleEndian.PutUint32(o[8+8:], uint32(s.Resource))
	binary.LittleEndian.PutUint64(o[8+8+4:], s.ID)
	return o
}
