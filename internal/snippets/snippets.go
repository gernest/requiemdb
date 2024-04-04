package snippets

import (
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/ristretto"
	"github.com/dop251/goja"
)

type Snippets struct {
	hashed *ristretto.Cache
}

func New() (*Snippets, error) {
	hashed, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     26 << 20,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &Snippets{hashed: hashed}, nil
}

func (s *Snippets) Close() {
	s.hashed.Close()
}

func (s *Snippets) GetProgramData(data []byte) (*goja.Program, error) {
	hash := xxhash.Sum64(data)
	if o, ok := s.hashed.Get(hash); ok {
		return o.(*goja.Program), nil
	}
	program, err := goja.Compile("index.js", string(data), true)
	if err != nil {
		return nil, err
	}
	cost := len(data)
	s.hashed.Set(hash, program, int64(cost))
	return program, nil
}
