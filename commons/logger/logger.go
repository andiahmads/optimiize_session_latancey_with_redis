package logger

import (
	"log/slog"
	"os"
)

func Slogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return logger
}
