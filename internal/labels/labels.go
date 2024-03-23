package labels

import (
	"bytes"
	"sync"
)

const (
	ResourceSchema     = "resource.schema"
	ResourceAttributes = "resource.attributes"
	ScopeSchema        = "scope.schema"
	ScopeName          = "scope.name"
	ScopeVersion       = "scope.version"
	ScopeAttributes    = "scope.attributes"
	MetricName         = "name"
	Attribute          = "attribute"
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

func (b *Bytes) Add(part string) *Bytes {
	if b.Len() > 0 {
		b.WriteByte('.')
	}
	b.WriteString(part)
	return b
}

func (b *Bytes) Value(value string) *Bytes {
	b.WriteByte('=')
	b.WriteString(value)
	return b
}

func (b *Bytes) Release() {
	b.Reset()
	bytesPool.Put(b)
}

func NewBytes() *Bytes {
	return bytesPool.Get().(*Bytes)
}

var bytesPool = &sync.Pool{New: func() any { return new(Bytes) }}
