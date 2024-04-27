package dataframe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColumns(t *testing.T) {
	columns := NewColumns()
	require.Equal(t, 97, len(columns.Metrics))
	require.Equal(t, 38, len(columns.Traces))
	require.Equal(t, 21, len(columns.Logs))
}
