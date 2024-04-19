package render

import (
	"os"
	"testing"

	"github.com/gernest/requiemdb/internal/test"
)

func TestNumericDataPoint(t *testing.T) {
	data := test.MetricsSamples(t)
	a := data[len(data)-1]
	o := MetricsData(a.GetMetrics(), MetricsFormatOption{})
	os.WriteFile("testdata/metrics.txt", []byte(o), 0600)
}
