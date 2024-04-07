package logger

import (
	"log/slog"
	"os"
)

func Setup(logLevel string) slog.Level {
	var level slog.Level
	level.UnmarshalText([]byte(logLevel))
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: level,
				},
			),
		),
	)
	return level
}
