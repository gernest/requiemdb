package js

import (
	"testing"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/compile"
	"github.com/gernest/requiemdb/internal/data"
	"github.com/stretchr/testify/require"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func TestMetrics_query(t *testing.T) {
	program := setup(t, `import { Metrics,render } from "@requiemdb/rq";
	render((new Metrics()).query());
	`)
	vm := New().WithData(data.Zero(v1.RESOURCE_METRICS))
	defer vm.Release()
	err := vm.Run(program)
	require.NoError(t, err)
	require.NotNil(t, vm.Export)
	require.NotNil(t, vm.ScanRequest)
	_ = vm.Export.Export().(*metricsv1.MetricsData)
	require.True(t, vm.ExportOptions.JSON)
}

func setup(t *testing.T, data string) *goja.Program {
	t.Helper()
	b, err := compile.Compile([]byte(data))
	require.NoError(t, err)
	program, err := goja.Compile("index.js", string(b), true)
	require.NoError(t, err)
	return program
}
