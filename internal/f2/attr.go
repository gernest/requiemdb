package f2

import (
	"bytes"
	"sync"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	resouncev1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

type Attr struct {
	resource, scope, attr []*v1.KeyValue
	rb, sb, ab            roaring64.Bitmap
	tr                    Translate
	b                     bytes.Buffer
	hash                  xxhash.Digest
}

func newAttr() *Attr {
	return attrPool.Get().(*Attr)
}

func (a *Attr) Resource(r *resouncev1.Resource) {
	a.resource = a.resource[:0]
	a.resource = append(a.resource, r.GetAttributes()...)
	a.do(&a.rb, a.resource)
}

func (a *Attr) Scope(scope *v1.InstrumentationScope) {
	a.scope = a.scope[:0]
	a.scope = append(a.scope, scope.GetAttributes()...)
	a.do(&a.sb, a.scope)
}

func (a *Attr) Attr(name []byte, kv []*v1.KeyValue) ([]*v1.KeyValue, *roaring64.Bitmap, uint64) {
	a.ab.Clear()
	a.ab.Or(&a.rb)
	a.ab.Or(&a.sb)
	a.do(&a.ab, kv)
	a.name(name)
	a.attr = append(a.attr[:0], a.resource...)
	a.attr = append(a.attr, a.scope...)
	a.attr = append(a.attr, kv...)
	a.hash.Reset()
	a.ab.WriteTo(&a.hash)
	return a.attr, &a.ab, a.hash.Sum64()
}

func (a *Attr) Reset(tr Translate) *Attr {
	a.resource = a.resource[:0]
	a.scope = a.scope[:0]
	a.attr = a.attr[:0]
	a.rb.Clear()
	a.sb.Clear()
	a.ab.Clear()
	a.b.Reset()
	a.hash.Reset()
	a.tr = tr
	return a
}

func (a *Attr) Release() {
	a.Reset(nil)
	attrPool.Put(a)
}

func (a *Attr) name(name []byte) {
	a.b.Reset()
	a.b.Write(nameKey)
	a.b.Write(equal)
	a.b.Write(name)
	a.ab.Add(a.tr.Tr(a.b.Bytes()))
}

func (a *Attr) do(b *roaring64.Bitmap, kv []*v1.KeyValue) {
	b.Clear()
	for _, v := range kv {
		a.b.Reset()
		a.b.WriteString(v.Key)
		a.b.Write(equal)
		if toString(&a.b, v.Value) {
			b.Add(a.tr.Tr(a.b.Bytes()))
		}
	}
}

func toString(b *bytes.Buffer, a *v1.AnyValue) bool {
	switch e := a.Value.(type) {
	case *v1.AnyValue_StringValue:

		b.WriteString(e.StringValue)
		return true
	default:
		return false
	}
}

var attrPool = &sync.Pool{New: func() any { return new(Attr) }}
