package labels

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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

func (l *Labels) Debug() string {
	var b bytes.Buffer
	b.WriteByte('[')
	for i, lbl := range l.ls {
		if i != 0 {
			b.WriteByte(',')
			b.WriteByte(' ')
		}
		b.WriteString(lbl.String())
	}
	b.WriteByte(']')
	return b.String()
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

const ResourcePrefixSize = 8 + //namespace
	4 //resource

const StaticSize = ResourcePrefixSize +
	4 // prefix

var (
	ValueSep = []byte("=")
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

func (l *Label) String() string {
	return Debug(l.Encode())
}

func Debug(b []byte) string {
	return fmt.Sprintf("%d:%d:%d:%s",
		binary.LittleEndian.Uint64(b),
		binary.LittleEndian.Uint32(b[8:]),
		binary.LittleEndian.Uint32(b[8+4:]),
		string(b[StaticSize:]),
	)
}

func (l *Label) Encode() []byte {
	l.buffer = slices.Grow(l.buffer, StaticSize)[:StaticSize]
	binary.LittleEndian.PutUint64(l.buffer, l.Namespace)
	binary.LittleEndian.PutUint32(l.buffer[8:], uint32(l.Resource))
	binary.LittleEndian.PutUint32(l.buffer[8+4:], uint32(l.Prefix))
	l.buffer = append(l.buffer, []byte(l.Key)...)
	if l.Value != "" {
		l.buffer = append(l.buffer, ValueSep...)
		l.buffer = append(l.buffer, []byte(l.Value)...)
	}
	return l.buffer
}

var ErrKeyTooShort = errors.New("labels: Key too short")

func (l *Label) Decode(data []byte) error {
	if len(data) <= StaticSize {
		return ErrKeyTooShort
	}
	l.Namespace = binary.LittleEndian.Uint64(data)
	l.Resource = v1.RESOURCE(
		binary.LittleEndian.Uint32(data[8:]),
	)
	l.Prefix = v1.PREFIX(
		binary.LittleEndian.Uint32(data[8+4:]),
	)
	key, value, _ := bytes.Cut(data[StaticSize:], ValueSep)
	l.Key = string(key)
	l.Value = string(value)
	return nil
}

func (l *Label) Reset() *Label {
	l.Namespace = 0
	l.Resource = 0
	l.Prefix = 0
	l.Key = ""
	l.Value = ""
	l.buffer = l.buffer[:0]
	return l
}

func (l *Label) Release() {
	labelPool.Put(l.Reset())
}

var labelPool = &sync.Pool{New: func() any { return new(Label) }}
