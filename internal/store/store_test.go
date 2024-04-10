package store

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/gernest/requiemdb/internal/x"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestMetrics(t *testing.T) {
	store := testStore(t)
	data, err := test.MetricsSamples()
	require.NoError(t, err)
	for _, v := range data {
		err = store.Save(v)
		require.NoError(t, err)
	}
	t.Run("After compaction", func(t *testing.T) {
		require.NoError(t, store.tree.Compact())
		var parts []*lsm.Part
		store.tree.Iter(func(p *lsm.Part) error {
			parts = append(parts, p)
			return nil
		})
		require.Equal(t, 1, len(parts))
		p := parts[0]
		require.Equal(t, store.tree.Size(), p.Size)
		data, err := p.Record.MarshalJSON()
		require.NoError(t, err)
		// os.WriteFile("testdata/part.json", data, 0600)
		want, err := os.ReadFile("testdata/part.json")
		require.NoError(t, err)
		require.JSONEq(t, string(want), string(data))
	})
	t.Run("Must have data stored", func(t *testing.T) {
		txn := store.db.NewTransaction(false)
		defer txn.Discard()

		k := keys.New()
		defer k.Release()

		it, err := txn.Get(k.WithID(1).
			WithResource(v1.RESOURCE_METRICS).
			Encode())
		require.NoError(t, err)
		var size int64
		var o v1.Data
		err = it.Value(x.Decompress(&o, &size))
		require.NoError(t, err)
		require.True(t, proto.Equal(&o, data[1]))
	})

	t.Run("Generate bitmaps", func(t *testing.T) {
		var b bytes.Buffer
		lbl := labels.NewLabel()
		defer lbl.Release()
		err := listLabels(store.db,
			lbl.WithResource(v1.RESOURCE_METRICS),
			func(lbl *labels.Label, sample *bitmaps.Bitmap) {
				fmt.Fprintln(&b, lbl.String(), sample.String())
			},
		)
		require.NoError(t, err)
		// os.WriteFile("testdata/labels.txt", b.Bytes(), 0600)
		want, err := os.ReadFile("testdata/labels.txt")
		require.NoError(t, err)
		require.Equal(t, string(want), b.String())
	})
}

func BenchmarkStore(b *testing.B) {
	db, err := test.DB()
	require.NoError(b, err)
	defer db.Close()

	tree, err := lsm.New(db)
	require.NoError(b, err)
	store, err := NewStore(db, tree)
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

	tree, err := lsm.New(db)
	require.NoError(t, err)
	store, err := NewStore(db, tree)
	require.NoError(t, err)
	t.Cleanup(func() {
		store.Close()
	})
	return store
}
