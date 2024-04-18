package store

import (
	"testing"

	"github.com/gernest/rbf"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/samples"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
)

func BenchmarkStore(b *testing.B) {

	store := testStore(b)

	data, err := test.MetricsSamples()
	ls := samples.Get()
	defer ls.Release()

	for n, v := range data {
		ls.Items = append(ls.Items, &v1.Sample{
			Id:   uint64(n),
			Data: v,
		})
	}

	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := store.SaveSamples(ls)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func testStore(t testing.TB) *Storage {
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
	rb := rbf.NewDB(t.TempDir(), nil)
	require.NoError(t, rb.Open())
	t.Cleanup(func() {
		rb.Close()
	})
	tree, err := lsm.New(db, seq)
	require.NoError(t, err)
	store, err := NewStore(db, rb, seq, tree)
	require.NoError(t, err)
	t.Cleanup(func() {
		store.Close()
	})
	return store
}
