package loggers

import (
	"log/slog"
	"os"
)

var slogger *slog.Logger
var slevel slog.Level

func Init(version string) {
	switch version {
	case "dev":
		slevel = slog.LevelDebug
	default:
		slevel = slog.LevelInfo
	}

	slogger = slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slevel,
				AddSource: false,
			},
		),
	)

	slog.SetDefault(slogger)

	slog.Info("Logger initialized", slog.String("level", slevel.String()))
}

func GetLevel() slog.Level {
	return slevel
}
