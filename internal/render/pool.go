package render

import (
	"bytes"
	"sync"
)

func get() *bytes.Buffer {
	return bytesPool.Get().(*bytes.Buffer)
}

func put(b *bytes.Buffer) {
	b.Reset()
	bytesPool.Put(b)
}

var bytesPool = &sync.Pool{New: func() any { return new(bytes.Buffer) }}
