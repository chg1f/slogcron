package slogcron

import (
	"github.com/robfig/cron/v3"
	"github.com/sagikazarmark/slog-shim"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(l *slog.Logger) *Logger {
	return &Logger{Logger: l}
}

// Info logs routine messages about cron's operation.
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.Logger.Info(msg, keysAndValues...)
}

// Error logs an error condition.
func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.Logger.With(slog.String("error", err.Error())).Error(msg, keysAndValues...)
}

var _ cron.Logger = (*Logger)(nil)
