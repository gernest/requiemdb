package logger

import (
	"log/slog"
	"os"
)

func Fail(msg string, args ...any) {
	slog.Default().Error(msg, args...)
	os.Exit(1)
}
