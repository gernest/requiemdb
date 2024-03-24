package compile

import (
	"log/slog"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/requiemdb/requiemdb/internal/logger"
)

func Compile(log *slog.Logger, path string) []byte {
	log.Info("Compiling file", "path", path)
	result := api.Build(api.BuildOptions{
		External: []string{"rq"},
		EntryPoints: []string{
			path,
		},
	})
	if len(result.Errors) > 0 {
		for _, e := range result.Errors {
			log.Error(e.Text)
		}
		logger.Fail("Failed compiling snippet", "path", path)
	}
	return result.OutputFiles[0].Contents
}
