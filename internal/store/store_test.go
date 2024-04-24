package store

import (
	"strconv"
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/samples"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/shards"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/gernest/requiemdb/internal/translate"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestSave(t *testing.T) {

	store := testStore(t)

	data := test.MetricsSamples(t)
	ls := samples.Get()
	defer ls.Release()

	for n, v := range data {
		ls.Items = append(ls.Items, &v1.Sample{
			Id:   uint64(n),
			Data: v,
		})
	}
	require.NoError(t, store.SaveSamples(ls))

	t.Run("instant without compile filters", func(t *testing.T) {
		r, err := store.Scan(&v1.Scan{
			Scope: v1.Scan_METRICS,
		})
		require.NoError(t, err)
		require.True(t, proto.Equal(data[len(data)-1], r))
	})
}

func BenchmarkStore(b *testing.B) {

	store := testStore(b)

	data := test.MetricsSamples(b)
	ls := samples.Get()
	defer ls.Release()

	for n, v := range data {
		ls.Items = append(ls.Items, &v1.Sample{
			Id:   uint64(n),
			Data: v,
		})
	}

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
	db := test.DB(t)
	tr, err := translate.New(db, []byte(strconv.FormatUint(uint64(v1.RESOURCE_TRANSLATE_ID), 10)))
	require.NoError(t, err)
	t.Cleanup(func() {
		tr.Close()
	})
	seq, err := seq.New(db)
	require.NoError(t, err)
	t.Cleanup(func() {
		seq.Release()
	})
	rb := shards.New(t.TempDir(), nil)
	t.Cleanup(func() {
		rb.Close()
	})
	require.NoError(t, err)
	store, err := NewStore(db, rb, tr, seq, test.Now)
	require.NoError(t, err)
	t.Cleanup(func() {
		store.Close()
	})
	return store
}
