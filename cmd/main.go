package main

import (
	"log/slog"
	"os"

	"github.com/KiriKovr1/logger/pkg/handler/local"
)

func main() {
	opts := local.LocalHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := local.New(os.Stdout, opts)

	logger := slog.New(handler)

	l := logger.With(slog.String("a", "a"))

	l.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
}
