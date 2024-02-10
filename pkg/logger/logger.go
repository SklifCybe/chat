package logger

import (
	"log/slog"
	"os"

	"github.com/SklifCybe/chat/pkg/logger/handlers/slogpretty"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		opts := slogpretty.PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}

		handler := opts.NewPrettyHandler(os.Stdout)

		log = slog.New(handler)
	case "prod":
		fallthrough
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
