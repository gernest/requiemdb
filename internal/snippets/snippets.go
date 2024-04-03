package snippets

import (
	"errors"

	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto"
	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/compile"
	"github.com/gernest/requiemdb/internal/compress"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrSnippetNotFound = errors.New("snippet not found")
)

type Snippets struct {
	db     *badger.DB
	hashed *ristretto.Cache
}

func New(db *badger.DB, cacheBudget int64) (*Snippets, error) {
	if cacheBudget == 0 {
		cacheBudget = 26 << 20
	}

	hashed, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     8 << 20,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &Snippets{db: db, hashed: hashed}, nil
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

func buildKey(name string) []byte {
	key := make([]byte, 9)
	key[len(key)-1] = byte(v1.RESOURCE_SNIPPETS)
	key = append(key, []byte(name)...)
	return key
}

func (s *Snippets) Upsert(name string, data []byte) error {
	raw, compiled, err := s.build(name, data)
	if err != nil {
		return err
	}
	now := timestamppb.Now()

	key := buildKey(name)

	err = s.db.Update(func(txn *badger.Txn) error {
		var code *v1.Snippet
		it, err := txn.Get(key)
		if err != nil {
			if !errors.Is(err, badger.ErrKeyNotFound) {
				return err
			}
			code = &v1.Snippet{
				Name:      name,
				Raw:       raw,
				Compiled:  compiled,
				CreatedAt: now,
				UpdatedAt: now,
			}
		} else {
			code = &v1.Snippet{}
			err = it.Value(func(val []byte) error {
				return proto.Unmarshal(val, code)
			})
			if err != nil {
				return err
			}
			code.Raw = raw
			code.Compiled = compiled
			code.UpdatedAt = now
		}
		o, err := proto.Marshal(code)
		if err != nil {
			return err
		}
		return txn.Set(key, o)
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Snippets) build(name string, data []byte) (raw, compiled []byte, err error) {
	compiled, err = compile.Compile(data)
	if err != nil {
		return
	}
	// make sure it is a valid goja program
	progam, err := goja.Compile(name, string(compiled), true)
	if err != nil {
		return nil, nil, err
	}

	raw, err = compress.Compress(data)
	if err != nil {
		return
	}
	compiled, err = compress.Compress(compiled)
	if err != nil {
		return
	}
	s.hashed.Set(xxhash.Sum64String(name), progam, int64(len(compiled)))
	return
}
