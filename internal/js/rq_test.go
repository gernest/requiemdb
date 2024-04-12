package js

import (
	"bytes"
	"os"
	"testing"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/compile"
	"github.com/gernest/requiemdb/internal/data"
	"github.com/stretchr/testify/require"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func TestConsole(t *testing.T) {
	var b bytes.Buffer
	r := New().WithOutput(&b)
	defer r.Release()
	_, err := r.Runtime.RunString(`console.log("hello,world")`)
	require.NoError(t, err)
	require.Equal(t, "hello,world\n", b.String())
}

func TestRequire(t *testing.T) {
	var b bytes.Buffer
	r := New().WithOutput(&b)
	defer r.Release()
	src, err := os.ReadFile("testdata/require.ts")
	require.NoError(t, err)
	data, err := compile.Compile(src)
	require.NoError(t, err)
	_, err = r.Runtime.RunString(string(data))
	require.NoError(t, err)
	require.Equal(t, "map[resourceMetrics:[]]\n", b.String())
}

func TestMetrics_query(t *testing.T) {
	program := setup(t, `import { Metrics,render } from "@requiemdb/rq";
	render((new Metrics()).query());
	`)
	vm := New().WithData(data.Zero(v1.RESOURCE_METRICS))
	defer vm.Release()
	err := vm.Run(program)
	require.NoError(t, err)
	require.NotNil(t, vm.Export)
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
