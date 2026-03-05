package oplog

import (
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
	logSince(logger.Info, msg, operation, start, v...)
}

func DebugSince(msg string, operation string, start time.Time, v ...any) {
	logSince(logger.Debug, msg, operation, start, v...)
}

func WarnSince(msg string, operation string, start time.Time, v ...any) {
	logSince(logger.Warn, msg, operation, start, v...)
}

func ErrorSince(msg string, operation string, start time.Time, v ...any) {
	logSince(logger.Error, msg, operation, start, v...)
}
