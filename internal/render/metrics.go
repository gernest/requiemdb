package render

import (
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	v1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type MetricsFormatOption struct {
	Resource bool
	Scope    bool
	Metrics  bool
}

func MetricsData(md *v1.MetricsData, o MetricsFormatOption) string {
	b := get()
	defer put(b)
	for i, rm := range md.ResourceMetrics {
		if i != 0 {
			b.WriteByte('\n')
		}
		b.WriteString(resource(rm, &o))
	}
	return b.String()
}

func resource(rm *v1.ResourceMetrics, o *MetricsFormatOption) string {
	b := get()
	defer put(b)

	fmt.Fprintf(b, "=== %s\n", rm.SchemaUrl)
	if o.Resource {
		w := tabwriter.NewWriter(b, 0, 0, 1, ' ', tabwriter.AlignRight)
		fmt.Fprintf(w, "resource.schema \t%s\n", rm.GetSchemaUrl())
		fmt.Fprintf(w, "resource.attributes \t%s\n", attr(rm.GetResource().GetAttributes()))
		w.Flush()
	}

	for i, s := range rm.ScopeMetrics {
		if i != 0 {
			b.WriteByte('\n')
		}
		b.WriteString(scope(s, o))
	}
	return b.String()
}

func scope(sc *v1.ScopeMetrics, o *MetricsFormatOption) string {
	b := get()
	defer put(b)

	fmt.Fprintf(b, "--> %s\n", sc.GetScope().GetName())
	if o.Scope {
		w := tabwriter.NewWriter(b, 0, 0, 1, ' ', tabwriter.AlignRight)
		fmt.Fprintf(w, "scope.name \t%s\n", sc.GetScope().GetName())
		fmt.Fprintf(w, "scope.schema \t%s\n", sc.SchemaUrl)
		fmt.Fprintf(w, "scope.version \t%s\n", sc.GetScope().GetVersion())
		fmt.Fprintf(w, "scope.attributes \t%s\n", attr(sc.GetScope().GetAttributes()))
		w.Flush()
	}
	for i, m := range sc.Metrics {
		if i != 0 {
			b.WriteByte('\n')
		}
		b.WriteString(metrics(m, o))
	}
	return b.String()
}

var sep = []byte("----------\n")

func metrics(m *v1.Metric, o *MetricsFormatOption) string {
	b := get()
	defer put(b)

	unit := m.Unit
	if unit == "" && strings.Contains(m.Description, "nanoseconds") {
		unit = "ns"
	}

	fmt.Fprintf(b, "%s\n", m.Name)
	if o.Metrics {
		w := tabwriter.NewWriter(b, 0, 0, 1, ' ', 0)
		fmt.Fprintf(w, "description \t%s\t\n", m.Description)
		fmt.Fprintf(w, "unit \t%s\t\n", unit)
		w.Flush()
	}
	switch e := m.Data.(type) {
	case *v1.Metric_Gauge:
		b.WriteString(numeric(e.Gauge.DataPoints, unit))
		return b.String()
	case *v1.Metric_Sum:
		b.WriteString(numeric(e.Sum.DataPoints, unit))
		return b.String()
	default:
		return b.String()
	}
}

func numeric(points []*v1.NumberDataPoint, unit string) string {
	b := get()
	defer put(b)
	w := tabwriter.NewWriter(b, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "TIMESTAMP \tVALUE \tATTRIBUTES \t")
	for _, d := range points {
		fmt.Fprintf(w, "%s \t%s \t%s \t\n",
			formatTime(d.TimeUnixNano),
			formatUnit(numericDataValue(d), unit),
			attr(d.Attributes))
	}
	w.Flush()
	return b.String()
}

func formatTime(ns uint64) string {
	return time.Unix(0, int64(ns)).UTC().Format(time.DateTime)
}

func numericDataValue(d *v1.NumberDataPoint) float64 {
	switch e := d.Value.(type) {
	case *v1.NumberDataPoint_AsInt:
		return float64(e.AsInt)
	default:
		return d.GetAsDouble()
	}
}
