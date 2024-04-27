package dataframe

import (
	"strings"

	"github.com/apache/arrow/go/v17/parquet/schema"
	"github.com/gernest/arrow3"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

type Column struct {
	ID int
	*schema.Column
}

// Column registers all columns belonging to Metrics, Traces and Logs. This
// allows  only reading columns belong to the resource we are interested in.
type Columns struct {
	Metrics []*Column
	all     struct {
		metrics []int
		traces  []int
		logs    []int
	}
	Traces []*Column
	Logs   []*Column
	schema *arrow3.Schema[*v1.Data]
}

func (c *Columns) Release() {
	put(c.schema)
	*c = Columns{}
}

func NewColumns() *Columns {
	a := get()
	schema := a.Parquet()
	o := &Columns{
		schema: a,
	}
	for i := 0; i < schema.NumColumns(); i++ {
		col := schema.Column(i)
		path := col.Path()
		switch {
		case strings.HasPrefix(path, "metrics."):
			o.Metrics = append(o.Metrics, &Column{
				ID:     i,
				Column: col,
			})
			o.all.metrics = append(o.all.metrics, i)
		case strings.HasPrefix(path, "logs."):
			o.Logs = append(o.Logs, &Column{
				ID: i, Column: col,
			})
			o.all.logs = append(o.all.logs, i)
		case strings.HasPrefix(path, "traces."):
			o.Traces = append(o.Traces, &Column{
				ID: i, Column: col,
			})
			o.all.traces = append(o.all.traces, i)
		}
	}
	return o
}

func (c *Columns) AllMetrics() []int {
	return c.all.metrics
}
func (c *Columns) AllTraces() []int {
	return c.all.traces
}

func (c *Columns) AllLogs() []int {
	return c.all.logs
}
