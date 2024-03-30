package store

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	"github.com/requiemdb/requiemdb/internal/lsm"
	"google.golang.org/protobuf/proto"
)

type Scanner struct {
	it      *sampleIter
	missing roaring64.Bitmap
	config  *v1.Scan
	key     [8 + 1 + 8]byte
}

func NewScanner(db *badger.DB, scan *v1.Scan, samples *lsm.Samples) *Scanner {
	return &Scanner{
		it:     newIter(db.NewTransaction(false), scan, samples),
		config: scan,
	}
}

func (s *Scanner) Missing() *roaring64.Bitmap {
	return &s.missing
}

func (s *Scanner) Close() {
	s.it.Close()
}

func (s *Scanner) HasNext() bool {
	return s.it.HasNext()
}

func (s *Scanner) Next() (*v1.Data, error) {
	ns, id := s.it.Next()
	copy(s.key[:], ns[:])
	s.key[8] = byte(s.config.Scope)
	copy(s.key[9:], id)

	it, err := s.it.txn.Get(s.key[:])
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			s.missing.Add(s.it.sample)
			return nil, nil
		}
		return nil, err
	}
	var sample v1.Sample
	err = it.Value(func(val []byte) error {
		return proto.Unmarshal(val, &sample)
	})
	if err != nil {
		return nil, err
	}
	return sample.Data, nil
}

type sampleIter struct {
	txn     *badger.Txn
	b       *lsm.Iterator
	date    uint64
	r       *roaring64.Bitmap
	reverse bool
	it      roaring64.IntIterable64
	d       [8]byte
	sample  uint64
	id      [8]byte
	labels  *labels.Labels
	err     error
}

func newIter(txn *badger.Txn, scan *v1.Scan, samples *lsm.Samples) *sampleIter {
	i := &sampleIter{}
	i.txn = txn
	i.reverse = scan.Reverse
	i.labels = build(v1.RESOURCE(scan.Scope), scan.Filters)
	limit := scan.Limit
	if limit == 0 {
		limit = math.MaxUint64
	}
	it := samples.Iterator()
	if scan.Reverse {
		it = samples.ReverseIterator()
	}
	i.b = it
	return i
}

func (p *sampleIter) Close() {
	p.txn.Discard()
	p.labels.Release()
	p.labels = nil
}

func (p *sampleIter) Err() error { return p.err }

func (p *sampleIter) HasNext() bool {
	if p.err != nil {
		return false
	}
	if p.r == nil {
		p.date, p.r = p.b.Next()
		p.it = p.r.ReverseIterator()
		if !p.reverse {
			p.it = p.r.Iterator()
		}
		binary.LittleEndian.PutUint64(p.d[:], p.date)
		p.loadPartition()
	}
	if !p.it.HasNext() {
		if !p.b.HasNext() {
			// we are done
			return false
		}
		p.date, p.r = p.b.Next()
		p.it = p.r.ReverseIterator()
		if !p.reverse {
			p.it = p.r.Iterator()
		}
		binary.LittleEndian.PutUint64(p.d[:], p.date)
		p.loadPartition()
	}
	return p.it.HasNext()
}

func (p *sampleIter) Next() (date *[8]byte, id []byte) {
	p.sample = p.it.Next()
	binary.LittleEndian.PutUint64(p.id[:], p.sample)
	return &p.d, p.id[:]
}

func (p *sampleIter) loadPartition() {
	o := new(roaring64.Bitmap)
	for _, v := range p.labels.Values {
		err := loadLabel(p.txn, v.Namespaced(&p.d), o)
		if err != nil {
			p.r.Clear()
			if errors.Is(err, badger.ErrKeyNotFound) {
				// skip this partition
				return
			}
			p.err = err
			return
		}
		p.r.And(o)
		if p.r.IsEmpty() {
			return
		}
	}
}

func loadLabel(txn *badger.Txn, key []byte, o *roaring64.Bitmap) (err error) {
	it, err := txn.Get(key)
	if err != nil {
		return err
	}
	err = it.Value(func(val []byte) error {
		return o.UnmarshalBinary(val)
	})
	return
}

func build(resource v1.RESOURCE, filters []*v1.Scan_Filter) *labels.Labels {
	ls := labels.NewLabels()
	for _, f := range filters {
		switch e := f.Value.(type) {
		case *v1.Scan_Filter_Base:
			ls.Add(
				labels.NewBytes(resource, v1.PREFIX(e.Base.Prop)).
					Value(e.Base.Value),
			)
		case *v1.Scan_Filter_Attr:
			ls.Add(
				labels.NewBytes(resource, v1.PREFIX(e.Attr.Prop)).
					Add(e.Attr.Key).
					Value(e.Attr.Value),
			)
		}
	}
	return ls
}
