package protoarrow

import (
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Build struct {
	r *array.RecordBuilder
}

func New(mem memory.Allocator, a proto.Message) *Build {
	return &Build{
		r: array.NewRecordBuilder(mem, Schema(a)),
	}
}

func (b *Build) Append(v proto.Message) {
	build(b.r, v)
}

func (b *Build) Release() {
	b.r.Release()
}

func (b Build) NewRecord() arrow.Record {
	return b.r.NewRecord()
}

func build(r *array.RecordBuilder, msg proto.Message) {
	p := msg.ProtoReflect()
	fields := p.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		write(r.Field(i), p.Get(fields.Get(i)))
	}
}

func write(b array.Builder, value protoreflect.Value) {
	switch e := b.(type) {
	case *array.Uint64Builder:
		e.Append(value.Uint())
	case *array.BinaryBuilder:
		e.Append(value.Bytes())
	case *array.ListBuilder:
		ls := value.List()
		if !ls.IsValid() {
			e.AppendNull()
			return
		}
		e.Append(true)
		vb := e.ValueBuilder()
		vb.Reserve(ls.Len())
		for i := 0; i < ls.Len(); i++ {
			write(vb, ls.Get(i))
		}
	case *array.StructBuilder:
		msg := value.Message()
		if !msg.IsValid() {
			e.AppendNull()
			return
		}
		e.Append(true)
		fields := msg.Descriptor().Fields()
		for i := 0; i < fields.Len(); i++ {
			write(e.FieldBuilder(i), msg.Get(fields.Get(i)))
		}
	default:
		panic(fmt.Sprintf("%T is not handled", e))
	}
}
