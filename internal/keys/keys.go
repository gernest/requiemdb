package keys

import (
	"fmt"
	"sync"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

type Sample struct {
	Resource v1.RESOURCE
	ID       uint64
	b        []byte
}

func New() *Sample {
	return samplePool.Get().(*Sample)
}

var samplePool = &sync.Pool{New: func() any {
	var s Sample
	return s.Reset()
}}

func (s *Sample) WithResource(r v1.RESOURCE) *Sample {
	s.Resource = r
	return s
}

func (s *Sample) Reset() *Sample {
	s.b = s.b[:0]
	s.Resource = 0
	s.ID = 0
	return s
}

func (s *Sample) Release() {
	s.Reset()
	samplePool.Put(s)
}

func (s *Sample) WithID(id uint64) *Sample {
	s.ID = id
	return s
}

func (s *Sample) Encode() []byte {
	s.b = fmt.Appendf(s.b[:0], "%d:%d", s.Resource, s.ID)
	return s.b
}

func (s *Sample) String() string {
	return string(s.Encode())
}

func (s *Sample) Prefix() []byte {
	s.b = fmt.Appendf(s.b[:0], "%d:", s.Resource)
	return s.b
}
