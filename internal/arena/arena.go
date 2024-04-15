package arena

import (
	"io"
	"sync"

	"github.com/gernest/requiemdb/internal/compress"
	"google.golang.org/protobuf/proto"
)

type Arena struct {
	data   []byte
	offset int
}

func (a *Arena) NewWriter() io.Writer {
	a.offset = len(a.data)
	return a
}

func (a *Arena) Bytes() []byte {
	return a.data[a.offset:]
}

func (a *Arena) Write(p []byte) (int, error) {
	a.data = append(a.data, p...)
	return len(p), nil
}

func (a *Arena) Compress(msg proto.Message) ([]byte, error) {
	data, err := a.Marshal(msg)
	if err != nil {
		return nil, err
	}
	err = compress.To(a.NewWriter(), data)
	if err != nil {
		a.data = a.data[:a.offset]
		return nil, err
	}
	return a.Bytes(), nil
}

func (a *Arena) Marshal(msg proto.Message) (b []byte, err error) {
	a.offset = len(a.data)
	a.data, err = proto.MarshalOptions{}.MarshalAppend(a.data, msg)
	if err != nil {
		a.data = a.data[:a.offset]
		return nil, err
	}
	b = a.data[a.offset:]
	return
}

func (a *Arena) Release() {
	clear(a.data)
	a.data = a.data[:0]
	a.offset = 0
	arenaPool.Put(a)
}

func New() *Arena {
	return arenaPool.Get().(*Arena)
}

var arenaPool = &sync.Pool{New: func() any { return new(Arena) }}
