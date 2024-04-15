package keys

import (
	"encoding/binary"
	"fmt"
	"sync"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

// Sample represents a unique key for a specific sample
//
//	8 bytes for namespace
//	4 bytes for resource
//	4 bytes for prefix which defaults to v1.PREFIX_DATA
//	8 bytes for sample id
//
// This object is pooled so all uses must start by calling New and ensure
// Release is called after use.
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

type Meta [8 + 4 + 4 + +8]byte

func NewMeta() *Meta {
	return metaPool.Get().(*Meta)
}

func (m *Meta) Encode() []byte {
	return m[:]
}

func (m *Meta) Info() *Meta {
	binary.LittleEndian.PutUint32(m[8:], uint32(v1.RESOURCE_META_INFO))
	return m
}

func (m *Meta) Data() *Meta {
	binary.LittleEndian.PutUint32(m[8:], uint32(v1.RESOURCE_META))
	return m
}

func (m *Meta) WithID(id uint64) *Meta {
	binary.LittleEndian.PutUint64(m[8+4+4:], id)
	return m
}

func (m *Meta) WithRESOURCE(res v1.RESOURCE) *Meta {
	binary.LittleEndian.PutUint32(m[8+4+8:], uint32(res))
	return m
}

func (m *Meta) WithNS(id uint64) *Meta {
	binary.LittleEndian.PutUint64(m[0:], id)
	return m
}

func (m *Meta) Reset() *Meta {
	clear(m[:])
	binary.LittleEndian.PutUint32(m[8:], uint32(v1.RESOURCE_META))
	return m
}

func (m *Meta) String() string {
	return fmt.Sprintf("%d:%d:%d:%d",
		binary.LittleEndian.Uint64(m[:]),
		binary.LittleEndian.Uint32(m[8:]),
		binary.LittleEndian.Uint64(m[8+4:]),
		binary.LittleEndian.Uint32(m[8+4+8:]),
	)
}

func (m *Meta) Release() {
	metaPool.Put(m)
}

var metaPool = &sync.Pool{New: func() any {
	var m Meta
	binary.LittleEndian.PutUint32(m[8:], uint32(v1.RESOURCE_META))
	return &m
}}
