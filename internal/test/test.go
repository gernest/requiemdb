package test

import (
	"embed"
	"io"
	"path"

	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

//go:embed samples
var samples embed.FS

func MetricsSamples() ([]*v1.Data, error) {
	m := "samples/metrics"
	entries, err := samples.ReadDir(m)
	if err != nil {
		return nil, err
	}
	var ls []*v1.Data
	for _, e := range entries {
		f, err := samples.Open(path.Join(m, e.Name()))
		if err != nil {
			return nil, err
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		var o v1.Data
		err = protojson.Unmarshal(b, &o)
		if err != nil {
			return nil, err
		}
		ls = append(ls, &o)
	}
	return ls, nil
}

func DB() (*badger.DB, error) {
	return badger.Open(
		badger.DefaultOptions("").
			WithInMemory(true).
			WithLogger(nil),
	)
}
