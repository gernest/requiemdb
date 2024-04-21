package labels

import (
	"fmt"
	"sync"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

type Label struct {
	Resource v1.RESOURCE
	Prefix   v1.PREFIX
	Key      string
	Value    string
	buf      []byte
}

func NewLabel() *Label {
	return labelPool.Get().(*Label)
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
	l.buf = fmt.Appendf(l.buf, "%d:%d:%s", l.Resource, l.Prefix, l.Key)
	if l.Value != "" {
		l.buf = fmt.Appendf(l.buf, "=%s", l.Value)
	}
	return l.buf
}

func (l *Label) String() string {
	return string(l.Encode())
}

func (l *Label) Reset() *Label {
	l.Resource = 0
	l.Prefix = 0
	l.Key = ""
	l.Value = ""
	l.buf = l.buf[:0]
	return l
}

func (l *Label) Release() {
	labelPool.Put(l.Reset())
}

var labelPool = &sync.Pool{New: func() any { return new(Label) }}
