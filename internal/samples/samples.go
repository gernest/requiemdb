package samples

import (
	"sync"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

type List struct {
	Items []*v1.Sample
}

func Get() *List {
	return pool.Get().(*List)
}

func (l *List) Release() {
	clear(l.Items)
	l.Items = l.Items[:0]
}

var pool = &sync.Pool{New: func() any { return new(List) }}
