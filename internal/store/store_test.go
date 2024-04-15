package store

import (
	"testing"

	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
)

func BenchmarkStore(b *testing.B) {
	db, err := test.DB()
	require.NoError(b, err)
	defer db.Close()
	seq, err := seq.New(db)
	require.NoError(b, err)
	tree, err := lsm.New(db, seq)
	require.NoError(b, err)
	store, err := NewStore(db, seq, tree)
	require.NoError(b, err)
	defer store.Close()

	data, err := test.MetricsSamples()
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			store.Save(v)
		}
	}
}

func testStore(t *testing.T) *Storage {
	t.Helper()
	db, err := test.DB()
	require.NoError(t, err)
	t.Cleanup(func() {
		db.Close()
	})
	seq, err := seq.New(db)
	require.NoError(t, err)
	t.Cleanup(func() {
		seq.Release()
	})
	tree, err := lsm.New(db, seq)
	require.NoError(t, err)
	store, err := NewStore(db, seq, tree)
	require.NoError(t, err)
	t.Cleanup(func() {
		store.Close()
	})
	return store
}
