package dataframe

import (
	"strings"

	"github.com/apache/arrow/go/v17/parquet/schema"
)

type Column struct {
	ID int
	*schema.Column
}

// Column registers all columns belonging to Metrics, Traces and Logs. This
// allows  only reading columns belong to the resource we are interested in.
type Columns struct {
	Metrics []*Column
	Traces  []*Column
	Logs    []*Column
}

func NewColumns() *Columns {
	schema := get().Parquet()
	o := &Columns{}
	for i := 0; i < schema.NumColumns(); i++ {
		col := schema.Column(i)
		path := col.Path()
		switch {
		case strings.HasPrefix(path, "metrics."):
			o.Metrics = append(o.Metrics, &Column{
				ID:     i,
				Column: col,
			})
		case strings.HasPrefix(path, "logs."):
			o.Logs = append(o.Logs, &Column{
				ID: i, Column: col,
			})
		case strings.HasPrefix(path, "traces."):
			o.Traces = append(o.Traces, &Column{
				ID: i, Column: col,
			})
		}
	}
	return o
}
