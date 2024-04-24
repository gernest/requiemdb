package translate

import (
	"encoding/binary"
	"errors"

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
		it, err := txn.Get(append(b.ids, u64tob(id)...))
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

func (b *Translate) TranslateKey(key []byte) (n uint64, err error) {
	err = b.db.Update(func(txn *badger.Txn) error {
		k := append(b.keys, key...)
		it, err := txn.Get(k)
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				n, err = b.seq.Next()
				if err != nil {
					return err
				}
				x := u64tob(n)
				return errors.Join(
					txn.Set(k, x),
					txn.Set(append(b.ids, x...), []byte(key)),
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
