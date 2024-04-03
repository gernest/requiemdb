package data

import (
	"os"
	"testing"

	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestCollapse(t *testing.T) {
	data, err := test.MetricsSamples()
	require.NoError(t, err)
	r := Collapse(data)
	b, err := protojson.MarshalOptions{Multiline: true}.Marshal(r)
	require.NoError(t, err)
	// os.WriteFile("testdata/collapsed_metrics.json", b, 0600)
	want, err := os.ReadFile("testdata/collapsed_metrics.json")
	require.NoError(t, err)
	require.JSONEq(t, string(want), string(b))
}
