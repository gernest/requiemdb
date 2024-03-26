package compile

import (
	"errors"
	"fmt"
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
	dir, alias, err := createPackage(data)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	slog.Info("Compiling file", "path", dir)
	result := api.Build(api.BuildOptions{
		Bundle:        true,
		Alias:         alias,
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

func createPackage(data []byte) (string, map[string]string, error) {
	f, err := os.MkdirTemp("", pattern)
	if err != nil {
		return "", nil, err
	}
	alias, err := setup(f)
	if err != nil {
		return "", nil, err
	}
	err = os.WriteFile(filepath.Join(f, "index.ts"), data, 0600)
	if err != nil {
		return "", nil, err
	}
	return f, alias, nil
}

func setup(dir string) (alias map[string]string, err error) {
	alias = make(map[string]string)
	for name, data := range rq.PKG {
		path := filepath.Base(name) + ".js"
		fmt.Println(path)
		err = os.WriteFile(filepath.Join(dir, path), data, 0600)
		if err != nil {
			fmt.Println("====", err)
			return
		}
		// relative path to working directory
		alias[name] = "./" + path
	}
	return
}
