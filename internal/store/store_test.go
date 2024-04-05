package store

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/gernest/requiemdb/internal/x"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestMetrics(t *testing.T) {
	db, err := test.DB()
	require.NoError(t, err)
	defer db.Close()

	tree, err := lsm.New(db)
	require.NoError(t, err)
	store, err := NewStore(db, tree)
	require.NoError(t, err)
	defer store.Close()

	data, err := test.MetricsSamples()
	require.NoError(t, err)
	for _, v := range data {
		err = store.Save(v)
		require.NoError(t, err)
	}
	t.Run("Before compaction", func(t *testing.T) {
		meta := tree.GetBuffer()
		require.Equal(t, len(data), len(meta))

		for i, m := range meta {
			require.NotZero(t, m.MinTs)
			require.NotZero(t, m.MaxTs)
			require.Less(t, m.MinTs, m.MaxTs)
			require.Equal(t, uint64(i), m.Id)
		}
	})
	t.Run("After compaction", func(t *testing.T) {
		meta := tree.GetBuffer()
		require.NoError(t, tree.Compact())
		require.Zero(t, len(tree.GetBuffer()))
		var parts []*lsm.Part
		tree.Iter(func(p *lsm.Part) error {
			parts = append(parts, p)
			return nil
		})
		require.Equal(t, 1, len(parts))
		p := parts[0]
		require.Equal(t, meta[0].MinTs, p.MinTS)
		require.Equal(t, meta[len(meta)-1].MaxTs, p.MaxTS)
		require.Equal(t, tree.Size(), p.Size)
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
			func(lbl *labels.Label, sample *lsm.Samples) {
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
