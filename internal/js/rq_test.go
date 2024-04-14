package js

import (
	"bytes"
	"os"
	"testing"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/compile"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func TestMetrics(t *testing.T) {

	data := &v1.Data{
		Data: &v1.Data_Metrics{
			Metrics: &metricsv1.MetricsData{
				ResourceMetrics: []*metricsv1.ResourceMetrics{
					{
						SchemaUrl: "test.resource.schema",
						ScopeMetrics: []*metricsv1.ScopeMetrics{
							{
								Scope: &commonv1.InstrumentationScope{
									Name:    "test.scope",
									Version: "v0.0.1",
								},
								SchemaUrl: "test.scope.schema",
								Metrics: []*metricsv1.Metric{
									{
										Name:        "foo",
										Description: "Size",
										Unit:        "By",
										Data: &metricsv1.Metric_Gauge{
											Gauge: &metricsv1.Gauge{
												DataPoints: []*metricsv1.NumberDataPoint{
													{
														TimeUnixNano: uint64(test.Now().UnixNano()),
														Value: &metricsv1.NumberDataPoint_AsDouble{
															AsDouble: float64(1 << 10),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	var b bytes.Buffer
	vm := New().WithData(data).WithOutput(&b)
	defer vm.Release()

	run := func(t *testing.T, src string) {
		t.Helper()
		b.Reset()
		program := setup(t, src)
		err := vm.Run(program)
		require.NoError(t, err)
	}
	t.Run("json", func(t *testing.T) {
		run(t, `import { Metrics,render } from "@requiemdb/rq";
		Metrics.renderJSON((new Metrics()).query());
		`)
		// os.WriteFile("testdata/metrics/render_json.json", b.Bytes(), 0600)
		want, err := os.ReadFile("testdata/metrics/render_json.json")
		require.NoError(t, err)
		require.JSONEq(t, string(want), b.String())
	})
	t.Run("text", func(t *testing.T) {
		run(t, `import { Metrics,render } from "@requiemdb/rq";
		Metrics.render((new Metrics()).query());
		`)
		// os.WriteFile("testdata/metrics/render_text.txt", b.Bytes(), 0600)
		want, err := os.ReadFile("testdata/metrics/render_text.txt")
		require.NoError(t, err)
		require.Equal(t, string(want), b.String())
	})
}

func setup(t *testing.T, data string) *goja.Program {
	t.Helper()
	b, err := compile.Compile([]byte(data))
	require.NoError(t, err)
	program, err := goja.Compile("index.js", string(b), true)
	require.NoError(t, err)
	return program
}
