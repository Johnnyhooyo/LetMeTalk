package common

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

var DefaultLogger = slog.New(tint.NewHandler(os.Stderr, &tint.Options{
	Level:      slog.LevelInfo,
	TimeFormat: time.DateTime,
}))

var DebugLogger = slog.New(tint.NewHandler(os.Stderr, &tint.Options{
	Level:      slog.LevelDebug,
	TimeFormat: time.DateTime,
}))
