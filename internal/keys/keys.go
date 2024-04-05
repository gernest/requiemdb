package keys

import (
	"encoding/binary"
	"fmt"
	"sync"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

type Sample [8 + 4 + 4 + 8]byte

func New() *Sample {
	return samplePool.Get().(*Sample)
}

var samplePool = &sync.Pool{New: func() any {
	var s Sample
	return s.Reset()
}}

func (s *Sample) WithNamespace(ns uint64) *Sample {
	binary.LittleEndian.PutUint64(s[0:], ns)
	return s
}

func (s *Sample) WithResource(r v1.RESOURCE) *Sample {
	binary.LittleEndian.PutUint32(s[8:], uint32(r))
	return s
}

func (s *Sample) Reset() *Sample {
	clear(s[:])
	return s.WithPrefix(v1.PREFIX_DATA)
}

func (s *Sample) Release() {
	s.Reset()
	samplePool.Put(s)
}

func (s *Sample) WithPrefix(r v1.PREFIX) *Sample {
	binary.LittleEndian.PutUint32(s[8+4:], uint32(r))
	return s
}

func (s *Sample) WithID(id uint64) *Sample {
	binary.LittleEndian.PutUint64(s[8+4+4:], id)
	return s
}

func (s *Sample) Encode() []byte {
	return s[:]
}

func (s *Sample) String() string {
	return fmt.Sprintf("%d:%d:%d:%d",
		binary.LittleEndian.Uint64(s[:]),
		binary.LittleEndian.Uint32(s[8:]),
		binary.LittleEndian.Uint32(s[8+4:]),
		binary.LittleEndian.Uint64(s[8+4+4:]),
	)
}
