package logger

import (
	"log/slog"
	"time"
)

// Level represents different logging levels.
type Level = slog.Level

// Set of possible logging levels.
const (
	LevelDebug       = Level(slog.LevelDebug)
	LevelInfo        = Level(slog.LevelInfo)
	LevelWarn        = Level(slog.LevelWarn)
	LevelError Level = slog.LevelError
)

// /////////////////////////////////////////////////////////////////

// Record represents data that is being logged.
type Record struct {
	Time       time.Time
	Message    string
	Level      Level
	Attributes map[string]any
}
