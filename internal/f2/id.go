package f2

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/requiemdb/internal/logger"
)

type idGen struct {
	seq *badger.Sequence
}

func newID(db *badger.DB, prefix []byte) (*idGen, error) {
	seq, err := db.GetSequence(append(prefix, []byte("id")...), 4<<10)
	if err != nil {
		return nil, err
	}
	return &idGen{seq: seq}, nil
}

func (i *idGen) Next() uint64 {
	n, err := i.seq.Next()
	if err != nil {
		logger.Fail("failed generating id", "err", err)
	}
	return n
}

func (i *idGen) Close() error {
	return i.seq.Release()
}
