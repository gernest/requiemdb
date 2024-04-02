package labels

import (
	"encoding/binary"
	"slices"
	"sync"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

type Labels struct {
	ls []*Label
}

func (l *Labels) New() *Label {
	lbl := NewLabel()
	l.ls = append(l.ls, lbl)
	return lbl
}

func (l *Labels) Iter(f func(lbl *Label) error) error {
	for _, v := range l.ls {
		err := f(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Labels) Reset() {
	for _, o := range l.ls {
		o.Release()
	}
	clear(l.ls)
	l.ls = l.ls[:0]
}

type Label struct {
	Namespace uint64
	Resource  v1.RESOURCE
	Prefix    v1.PREFIX
	Key       string
	Value     string
	buffer    []byte
}

func NewLabel() *Label {
	return labelPool.Get().(*Label)
}

const staticSize = 8 + 4 + 4

var (
	valueSep = []byte("=")
)

func (l *Label) WithNamespace(ns uint64) *Label {
	l.Namespace = ns
	return l
}

func (l *Label) WithResource(r v1.RESOURCE) *Label {
	l.Resource = r
	return l
}

func (l *Label) WithPrefix(p v1.PREFIX) *Label {
	l.Prefix = p
	return l
}

func (l *Label) WithKey(k string) *Label {
	l.Key = k
	return l
}

func (l *Label) WithValue(v string) *Label {
	l.Value = v
	return l
}

func (l *Label) Encode() []byte {
	l.buffer = slices.Grow(l.buffer, staticSize)[:staticSize]
	binary.LittleEndian.PutUint64(l.buffer, l.Namespace)
	binary.LittleEndian.PutUint32(l.buffer[8:], uint32(l.Resource))
	binary.LittleEndian.PutUint32(l.buffer[8+4:], uint32(l.Prefix))
	l.buffer = append(l.buffer, []byte(l.Key)...)
	if l.Value != "" {
		l.buffer = append(l.buffer, valueSep...)
		l.buffer = append(l.buffer, []byte(l.Value)...)
	}
	return l.buffer
}

func (l *Label) Reset() *Label {
	l.Namespace = 0
	l.Resource = 0
	l.Prefix = 0
	l.buffer = l.buffer[:0]
	return l
}

func (l *Label) Release() {
	l.Namespace = 0
	l.Resource = 0
	l.Prefix = 0
	l.buffer = l.buffer[:0]
	labelPool.Put(l)
}

var labelPool = &sync.Pool{New: func() any { return new(Label) }}
