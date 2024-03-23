package protoarrow

import (
	"strings"

	"github.com/apache/arrow/go/v16/arrow"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type SchemaContext struct {
	Index  int
	Paths  []string
	Fields []Field
}

func (s *SchemaContext) emit(name string) {
	s.Fields = append(s.Fields, Field{
		Name:  strings.Join(append(s.Paths, name), "."),
		Index: s.Index,
	})
	s.Index++
}

func (s *SchemaContext) enterGroup(name string) {
	s.Paths = append(s.Paths, name)
}

func (s *SchemaContext) leaveGroup() {
	s.Paths = s.Paths[:len(s.Paths)-1]
}

type Field struct {
	Name  string
	Index int
}

func Schema(msg proto.Message, opts ...*SchemaContext) *arrow.Schema {
	ctx := &SchemaContext{}
	if len(opts) > 0 {
		ctx = opts[0]
	}
	return toSchema(msg.ProtoReflect(), ctx)
}

func toSchema(msg protoreflect.Message, ctx *SchemaContext) *arrow.Schema {
	fields := msg.Descriptor().Fields()
	a := make([]arrow.Field, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		a[i] = toField(fields.Get(i), ctx)
	}
	return arrow.NewSchema(a, nil)
}

func toField(f protoreflect.FieldDescriptor, ctx *SchemaContext) arrow.Field {
	fd := toFieldBase(f, ctx)
	if f.IsList() {
		fd.Type = arrow.ListOf(fd.Type)
		fd.Nullable = true
	}
	return fd
}

func toFieldBase(f protoreflect.FieldDescriptor, ctx *SchemaContext) arrow.Field {
	switch f.Kind() {
	case protoreflect.BoolKind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.FixedWidthTypes.Boolean,
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.PrimitiveTypes.Int32,
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.PrimitiveTypes.Uint32,
		}
	case protoreflect.Int64Kind,
		protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.PrimitiveTypes.Int64,
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.PrimitiveTypes.Uint64,
		}
	case protoreflect.FloatKind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.PrimitiveTypes.Float32,
		}
	case protoreflect.DoubleKind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.PrimitiveTypes.Float64,
		}
	case protoreflect.StringKind:
		ctx.emit(string(f.Name()))
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.BinaryTypes.String,
		}
	case protoreflect.BytesKind:
		return arrow.Field{
			Name: string(f.Name()),
			Type: arrow.BinaryTypes.Binary,
		}
	case protoreflect.MessageKind:
		ctx.enterGroup(string(f.Name()))
		fields := f.Message().Fields()
		a := make([]arrow.Field, fields.Len())
		for i := 0; i < fields.Len(); i++ {
			a[i] = toField(fields.Get(i), ctx)
		}
		ctx.leaveGroup()
		return arrow.Field{
			Name:     string(f.Name()),
			Type:     arrow.StructOf(a...),
			Nullable: true,
		}
	default:
		panic(f.Kind().String() + "is not supported")
	}
}
