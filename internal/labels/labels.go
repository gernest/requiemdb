package labels

import (
	"bytes"
	"errors"
	"fmt"
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

func (l *Labels) Debug() (o string) {
	o = "["
	for i, v := range l.ls {
		if i != 0 {
			o += ", "
		}
		o += v.String()
	}
	o += "]"
	return
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
}

func NewLabel() *Label {
	return labelPool.Get().(*Label)
}

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
	b := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		b.Reset()
		bufferPool.Put(b)
	}()
	fmt.Fprintf(b, "%d:%d:%d:%s", l.Namespace, l.Resource, l.Prefix, l.Key)
	if l.Value != "" {
		fmt.Fprintf(b, "=%s", l.Value)
	}
	return b.String()
}

var bufferPool = &sync.Pool{New: func() any { return new(bytes.Buffer) }}

var ErrKeyTooShort = errors.New("labels: Key too short")

func (l *Label) Reset() *Label {
	l.Namespace = 0
	l.Resource = 0
	l.Prefix = 0
	l.Key = ""
	l.Value = ""
	return l
}

func (l *Label) Release() {
	labelPool.Put(l.Reset())
}

var labelPool = &sync.Pool{New: func() any { return new(Label) }}
