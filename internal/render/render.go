package render

import (
	"io"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
)

func Result(w io.Writer, r *v1.Result) error {
	if r == nil {
		return nil
	}
	return nil
}
