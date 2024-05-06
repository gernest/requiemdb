package f2

import (
	"sync"

	"google.golang.org/protobuf/proto"
)

type Pool[T proto.Message] struct {
	sync sync.Pool
}

func NewPool[T proto.Message](init func() T) *Pool[T] {
	p := &Pool[T]{}
	p.sync.New = func() any { return init() }
	return p
}

func (p *Pool[T]) Get() T {
	return p.sync.Get().(T)
}

func (p *Pool[T]) Put(v T) {
	proto.Reset(v)
	p.sync.Put(v)
}
