package logger

import (
	"log/slog"
	"os"
)

func ConfigureSlog() error {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))

	slog.SetDefault(logger)
	return nil
}
