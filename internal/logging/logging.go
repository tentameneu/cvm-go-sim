package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/tentameneu/cvm-go-sim/internal/config"
)

var logger *slog.Logger

func InitializeLogger(writer io.Writer) error {
	var level slog.Level

	if config.Conf != nil {
		switch config.LogLevel() {
		case "info":
			level = slog.LevelInfo
		case "debug":
			level = slog.LevelDebug
		case "deep":
			level = LevelDeep
		default:
			return fmt.Errorf("invalid logging level '%s'", config.LogLevel())
		}
	} else {
		level = slog.LevelInfo
	}

	options := &slog.HandlerOptions{
		Level: level,
	}

	handler := newHandler(writer, options)
	logger = slog.New(handler)

	return nil
}

func Logger() *slog.Logger {
	return logger
}

func LogDeep(msg string, args ...any) {
	logger.Log(context.Background(), LevelDeep, msg, args...)
}