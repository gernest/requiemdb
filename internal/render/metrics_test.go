package render

import (
	"os"
	"testing"

	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
)

func TestNumericDataPoint(t *testing.T) {
	data, err := test.MetricsSamples()
	require.NoError(t, err)
	a := data[len(data)-1]
	o := MetricsData(a.GetMetrics(), MetricsFormatOption{})
	os.WriteFile("testdata/metrics.txt", []byte(o), 0600)
}
