package store

import (
	"os"
	"testing"

	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
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
}
