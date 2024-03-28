package lsm

import "sync/atomic"

type Node[T any] struct {
	next  atomic.Pointer[Node[T]]
	value T
}

func (n *Node[T]) Iterate(f func(*Node[T]) bool) {
	if !(f(n)) {
		return
	}
	node := n.next.Load()
	for {
		if node == nil {
			return
		}
		if !f(node) {
			return
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
