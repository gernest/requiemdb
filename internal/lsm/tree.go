package lsm

import (
	"sync"
	"sync/atomic"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/arrow/util"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/protoarrow"
)

// BufferSize minimum number of v1.Meta before building arrow.Record
const BufferSize = 4 << 10

type Part struct {
	Record  arrow.Record
	Size    uint64
	MinDate uint64
	MaxDate uint64
}

type Tree struct {
	root *Node[*Part]
	size atomic.Uint64

	// It is wasteful to build records of single Meta message. We buffer them and
	// search using binary search.
	buffer     []*v1.Meta
	bufferSize int
	bufferMu   sync.RWMutex

	build *protoarrow.Build

	nodes   []*Node[*Part]
	records []arrow.Record
}

func New() *Tree {
	return &Tree{
		root:       &Node[*Part]{},
		buffer:     make([]*v1.Meta, 0, BufferSize),
		bufferSize: BufferSize,
		build:      protoarrow.New(memory.DefaultAllocator, &v1.Meta{}),
	}
}

func (t *Tree) Append(meta *v1.Meta) {
	t.bufferMu.Lock()
	defer t.bufferMu.Unlock()
	t.buffer = append(t.buffer, meta)
	if len(t.buffer) < t.bufferSize {
		return
	}
	t.unsafeSaveBuffer()
}

func (t *Tree) Flush() {
	t.bufferMu.Lock()
	defer t.bufferMu.Unlock()
	t.unsafeSaveBuffer()
}

func (t *Tree) unsafeSaveBuffer() {
	if len(t.buffer) == 0 {
		return
	}
	for _, m := range t.buffer {
		t.build.Append(m)
	}
	t.buffer = t.buffer[:0]
	r := t.build.NewRecord()
	t.add(&Part{
		Record:  r,
		Size:    uint64(util.TotalRecordSize(r)),
		MinDate: t.buffer[0].Date,
		MaxDate: t.buffer[len(t.buffer)-1].Date,
	})
}

func (t *Tree) add(part *Part) {
	t.size.Add(part.Size)
	t.root.Prepend(part)
}

func (t *Tree) Compact() error {
	t.Flush()
	t.root.Iterate(func(n *Node[*Part]) bool {
		if n.value == nil {
			return true
		}
		t.nodes = append(t.nodes, n)
		t.records = append(t.records, n.value.Record)
		return true
	})
	if len(t.nodes) <= 1 {
		return nil
	}
	defer func() {
		clear(t.nodes)
		clear(t.records)
		t.nodes = t.nodes[:0]
		t.records = t.records[:0]
	}()

	r, err := protoarrow.Merge(t.records)
	if err != nil {
		return err
	}
	for i := range t.records {
		t.records[i].Release()
	}
	node := t.findNode(t.nodes[0])
	part := &Part{
		Record:  r,
		Size:    uint64(util.TotalRecordSize(r)),
		MinDate: t.nodes[0].value.MinDate,
		MaxDate: t.nodes[len(t.nodes)-1].value.MaxDate,
	}
	x := &Node[*Part]{
		value: part,
	}
	for !node.next.CompareAndSwap(t.nodes[0], x) {
		node = t.findNode(t.nodes[0])
	}
	t.size.Add(-part.Size)
	return nil
}

func (t *Tree) findNode(node *Node[*Part]) (list *Node[*Part]) {
	t.root.Iterate(func(n *Node[*Part]) bool {
		if n.next.Load() == node {
			list = n
			return false
		}
		return true
	})
	return
}
