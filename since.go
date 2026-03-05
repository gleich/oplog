package oplog

import (
	"log/slog"
	"time"
)

func logSince(fn func(string, ...any), msg string, operation string, start time.Time, v ...any) {
	fn(
		msg,
		append(
			[]any{"operation", operation, "duration", formatDuration(time.Since(start))},
			v...)...)
}

func InfoSince(msg string, operation string, start time.Time, v ...any) {
	logSince(slog.Info, msg, operation, start, v...)
}

func DebugSince(msg string, operation string, start time.Time, v ...any) {
	logSince(slog.Debug, msg, operation, start, v...)
}

func WarnSince(msg string, operation string, start time.Time, v ...any) {
	logSince(slog.Warn, msg, operation, start, v...)
}

func ErrorSince(msg string, operation string, start time.Time, v ...any) {
	logSince(slog.Error, msg, operation, start, v...)
}
