package logger

import (
	"io"
	"log/slog"
)

func Setup(logLevel string, out io.Writer) slog.Level {
	var level slog.Level
	level.UnmarshalText([]byte(logLevel))
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				out,
				&slog.HandlerOptions{
					Level: level,
				},
			),
		),
	)
	return level
}
