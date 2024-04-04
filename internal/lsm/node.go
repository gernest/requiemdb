package lsm

import (
	"errors"
	"io"
	"sync/atomic"
)

type Node[T any] struct {
	next  atomic.Pointer[Node[T]]
	value T
}

func (n *Node[T]) Iterate(f func(*Node[T]) error) error {
	if err := f(n); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	node := n.next.Load()
	for {
		if node == nil {
			return nil
		}
		if err := f(node); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		node = node.next.Load()
	}
}

func (n *Node[T]) Prepend(part T) *Node[T] {
	return n.prepend(&Node[T]{value: part})
}

func (n *Node[T]) prepend(node *Node[T]) *Node[T] {
	for {
		next := n.next.Load()
		node.next.Store(next)
		if n.next.CompareAndSwap(next, node) {
			return node
		}
	}
}
