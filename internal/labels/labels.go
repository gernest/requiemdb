package labels

import (
	"bytes"
	"sync"

	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
)

type Labels struct {
	Values []*Bytes
}

func NewLabels() *Labels {
	return labelsPool.Get().(*Labels)
}

func (l *Labels) Add(value *Bytes) {
	l.Values = append(l.Values, value)
}

func (l *Labels) Reset() {
	for _, b := range l.Values {
		b.Release()
	}
	clear(l.Values)
	l.Values = l.Values[:0]
}

func (l *Labels) Release() {
	l.Reset()
	labelsPool.Put(l)
}

var labelsPool = &sync.Pool{New: func() any { return new(Labels) }}

type Bytes struct {
	bytes.Buffer
}

func (b *Bytes) Namespaced(ns *[8]byte) []byte {
	o := b.Bytes()
	copy(o, ns[:])
	return o
}

func (b *Bytes) Add(part string) *Bytes {
	if b.Len() > 0 {
		b.WriteByte('.')
	}
	b.WriteString(part)
	return b
}
func (b *Bytes) AddBytes(part []byte) *Bytes {
	if b.Len() > 0 {
		b.WriteByte('.')
	}
	b.Write(part)
	return b
}

func (b *Bytes) Value(value string) *Bytes {
	b.WriteByte('=')
	b.WriteString(value)
	return b
}

func (b *Bytes) ValueBytes(value []byte) *Bytes {
	b.WriteByte('=')
	b.Write(value)
	return b
}

func (b *Bytes) Release() {
	b.Reset()
	bytesPool.Put(b)
}

var namespace [8]byte

func NewBytes(kind v1.RESOURCE, prefix v1.PREFIX) *Bytes {
	b := bytesPool.Get().(*Bytes)
	// Reserve the first 8 bytes for namespace
	b.Write(namespace[:])
	b.WriteByte(byte(kind))
	b.WriteByte(byte(prefix))
	return b
}

var bytesPool = &sync.Pool{New: func() any { return new(Bytes) }}
