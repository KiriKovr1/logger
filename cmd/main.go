package main

import (
	"log/slog"
	"os"

	"github.com/KiriKovr1/logger/internal/handler/local"
)

func main() {
	opts := local.LocalHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	logger := slog.New(local.New(os.Stdout, opts))
	logger.With(slog.String("op", "cmd.main"))

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message", slog.String("op", "cmd.main"))
}
