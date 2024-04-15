package seq

import (
	"encoding/binary"
	"errors"

	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/logger"
)

type Seq struct {
	meta   *badger.Sequence
	sample *badger.Sequence
}

func New(db *badger.DB) (*Seq, error) {
	// first 8 is reserved for namespace
	seqKey := make([]byte, 8+4)
	binary.LittleEndian.PutUint32(seqKey[8:], uint32(v1.RESOURCE_SAMPLE_ID))

	sample, err := db.GetSequence(seqKey, 1<<20)
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint32(seqKey[8:], uint32(v1.RESOURCE_META))
	meta, err := db.GetSequence(seqKey, 1<<20)
	if err != nil {
		sample.Release()
		return nil, err
	}
	return &Seq{meta: meta, sample: sample}, nil
}

func (s *Seq) Release() error {
	return errors.Join(s.meta.Release(), s.sample.Release())
}

func (s *Seq) SampleID() uint64 {
	id, err := s.sample.Next()
	if err != nil {
		logger.Fail("failed generating sample id", "err", err)
	}
	return id
}

func (s *Seq) MetaID() uint64 {
	m, err := s.meta.Next()
	if err != nil {
		logger.Fail("failed generating meta id", "err", err)
	}
	return m
}
