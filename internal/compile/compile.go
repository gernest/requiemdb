package compile

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/requiemdb/requiemdb/packages/rq"
)

const (
	pattern = "rq-cor-*"
)

func Compile(data []byte) ([]byte, error) {
	dir, err := writeRq(data)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	slog.Info("Compiling file", "path", dir)
	result := api.Build(api.BuildOptions{
		Bundle: true,
		Alias: map[string]string{
			rq.Module: "./rq",
		},
		AbsWorkingDir: dir,
		EntryPoints: []string{
			"index.ts",
		},
	})
	if len(result.Errors) > 0 {
		return nil, errors.New(result.Errors[0].Text)
	}
	return result.OutputFiles[0].Contents, nil
}

func writeRq(data []byte) (string, error) {
	f, err := os.MkdirTemp("", pattern)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filepath.Join(f, "rq.js"), rq.RQ, 0600)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filepath.Join(f, "index.ts"), data, 0600)
	if err != nil {
		return "", err
	}
	return f, nil
}
