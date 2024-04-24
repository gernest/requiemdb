package translate

import (
	"bytes"
	"encoding/binary"
	"errors"
	"sync"

	"github.com/dgraph-io/badger/v4"
)

type Translate struct {
	db  *badger.DB
	seq *badger.Sequence

	keys, ids []byte
}

func New(db *badger.DB, prefix []byte) (*Translate, error) {
	seq, err := db.GetSequence(append(prefix, []byte("seq")...), 1<<10)
	if err != nil {
		return nil, err
	}
	return &Translate{db: db, seq: seq,
		keys: append(prefix, []byte("keys")...),
		ids:  append(prefix, []byte("ids")...),
	}, nil
}

func (t *Translate) Close() error {
	return t.seq.Release()
}

func (b *Translate) TranslateID(id uint64) (k string, err error) {
	err = b.db.View(func(txn *badger.Txn) error {
		g := get()
		g.Write(b.ids)
		g.Write(u64tob(id))
		defer put(g)
		it, err := txn.Get(g.Bytes())
		if err != nil {
			return err
		}
		return it.Value(func(val []byte) error {
			k = string(val)
			return nil
		})
	})
	return
}

func (b *Translate) TranslateBulkID(ids []uint64, f func(key []byte) error) (err error) {
	err = b.db.View(func(txn *badger.Txn) error {
		g := get()
		defer put(g)
		var buf [8]byte
		for _, id := range ids {
			g.Reset()
			g.Write(b.ids)
			binary.BigEndian.PutUint64(buf[:], id)
			g.Write(buf[:])
			it, err := txn.Get(g.Bytes())
			if err != nil {
				return err
			}
			err = it.Value(f)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func (b *Translate) TranslateKey(key []byte) (n uint64, err error) {
	err = b.db.Update(func(txn *badger.Txn) error {
		g := get()
		g.Write(b.keys)
		g.Write(key)
		defer put(g)
		it, err := txn.Get(g.Bytes())
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				n, err = b.seq.Next()
				if err != nil {
					return err
				}
				x := u64tob(n)
				k := bytes.Clone(g.Bytes())
				g.Reset()
				g.Write(b.ids)
				g.Write(x)
				ik := bytes.Clone(g.Bytes())
				return errors.Join(
					txn.Set(k, x),
					txn.Set(ik, key),
				)
			}
			return err
		}
		return it.Value(func(val []byte) error {
			n = btou64(val)
			return nil
		})
	})
	return
}

func get() *bytes.Buffer {
	return kb.Get().(*bytes.Buffer)
}

func put(b *bytes.Buffer) {
	b.Reset()
	kb.Put(b)
}

var kb = &sync.Pool{New: func() any { return new(bytes.Buffer) }}

// Find finds translated key. Returns 0 if no key was found
func (b *Translate) Find(key []byte) (n uint64, err error) {
	err = b.db.Update(func(txn *badger.Txn) error {
		k := append(b.keys, key...)
		it, err := txn.Get(k)
		if err != nil {
			return err
		}
		return it.Value(func(val []byte) error {
			n = btou64(val)
			return nil
		})
	})
	return
}

// u64tob encodes v to big endian encoding.
func u64tob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// btou64 decodes b from big endian encoding.
func btou64(b []byte) uint64 { return binary.BigEndian.Uint64(b) }
