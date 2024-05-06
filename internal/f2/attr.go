package f2

import (
	"sync"

	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	resouncev1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

type Attr struct {
	resource, scope, attr []*v1.KeyValue
}

func newAttr() *Attr {
	return attrPool.Get().(*Attr)
}

func (a *Attr) Resource(r *resouncev1.Resource) {
	a.resource = a.resource[:0]
	a.resource = append(a.resource, r.GetAttributes()...)
}

func (a *Attr) Scope(scope *v1.InstrumentationScope) {
	a.scope = a.scope[:0]
	if a.scope != nil {
		a.scope = append(a.scope, scope.GetAttributes()...)
	}
}

func (a *Attr) scoped() int {
	return len(a.resource) + len(a.scope)
}

func (a *Attr) Attr(kv []*v1.KeyValue) []*v1.KeyValue {
	if a.scoped() > 0 && len(a.attr) < a.scoped() {
		a.attr = append(a.attr, a.resource...)
		a.attr = append(a.attr, a.scope...)
	}
	a.attr = append(a.attr[:a.scoped()], kv...)
	return a.attr
}

func (a *Attr) Reset() {
	a.resource = a.resource[:0]
	a.scope = a.scope[:0]
	a.attr = a.attr[:0]
}

func (a *Attr) Release() {
	a.Reset()
	attrPool.Put(a)
}

var attrPool = &sync.Pool{New: func() any { return new(Attr) }}
