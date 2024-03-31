package snippets

import (
	"errors"

	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto"
	"github.com/dop251/goja"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/compile"
	"github.com/requiemdb/requiemdb/internal/compress"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrSnippetNotFound = errors.New("snippet not found")
)

type Snippets struct {
	db    *badger.DB
	cache *ristretto.Cache
}

func New(db *badger.DB, cacheBudget int64) (*Snippets, error) {
	if cacheBudget == 0 {
		cacheBudget = 26 << 20
	}
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     cacheBudget,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &Snippets{db: db, cache: cache}, nil
}

func (s *Snippets) Close() error {
	s.cache.Close()
	return nil
}

func (s *Snippets) Rename(old, new string) error {
	txn := s.db.NewTransaction(true)
	defer txn.Discard()
	oldKey := buildKey(old)
	it, err := txn.Get(oldKey)
	if err != nil {
		return err
	}
	value, err := it.ValueCopy(nil)
	if err != nil {
		return err
	}
	err = txn.Set(buildKey(new), value)
	if err != nil {
		return err
	}

	// clear cache
	s.cache.Del(xxhash.Sum64String(old))
	return txn.Commit()

}

func (s *Snippets) List() (*v1.SnippetInfo_List, error) {
	prefix := buildKey("")
	var ls []*v1.SnippetInfo
	err := s.db.View(func(txn *badger.Txn) error {
		o := badger.DefaultIteratorOptions
		o.Prefix = prefix
		it := txn.NewIterator(o)
		defer it.Close()
		var snippet v1.Snippet

		for it.Rewind(); it.ValidForPrefix(prefix); it.Next() {
			err := it.Item().Value(func(val []byte) error {
				return proto.Unmarshal(val, &snippet)
			})
			if err != nil {
				return err
			}
			ls = append(ls, &v1.SnippetInfo{
				Name:        snippet.Name,
				Description: snippet.Description,
				CreatedAt:   proto.Clone(snippet.CreatedAt).(*timestamppb.Timestamp),
				UpdatedAt:   proto.Clone(snippet.UpdatedAt).(*timestamppb.Timestamp),
			})

		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &v1.SnippetInfo_List{Snippets: ls}, nil
}

func (s *Snippets) GetProgram(name string) (*goja.Program, error) {
	hash := xxhash.Sum64String(name)
	if o, ok := s.cache.Get(hash); ok {
		return o.(*goja.Program), nil
	}
	key := buildKey(name)
	var program *goja.Program
	var cost int
	err := s.db.View(func(txn *badger.Txn) error {
		it, err := txn.Get(key)
		if err != nil {
			return err
		}
		var code v1.Snippet
		err = it.Value(func(val []byte) error {
			return proto.Unmarshal(val, &code)
		})
		if err != nil {
			return err
		}
		data, err := compress.Decompress(code.Compiled)
		if err != nil {
			return err
		}
		program, err = goja.Compile(code.Name, string(data), true)
		cost = len(data)
		return err
	})
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, ErrSnippetNotFound
		}
		return nil, err
	}
	s.cache.Set(hash, program, int64(cost))
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
	// We normally create/update then execute snippets. Put the compiled program in cache
	// for faster loading
	s.cache.Set(xxhash.Sum64String(name), progam, int64(len(compiled)))
	return
}
