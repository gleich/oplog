package oplog

import (
	"time"
)

func logSince(fn func(string, ...any), opt Operation, msg string, start time.Time, v ...any) {
	log(fn, opt, msg, append([]any{"duration", formatDuration(time.Since(start))}, v...)...)
}

func InfoSince(opt Operation, msg string, start time.Time, v ...any) {
	logSince(logger.Info, opt, msg, start, v...)
}

func DebugSince(opt Operation, msg string, start time.Time, v ...any) {
	logSince(logger.Debug, opt, msg, start, v...)
}

func WarnSince(opt Operation, msg string, start time.Time, v ...any) {
	logSince(logger.Warn, opt, msg, start, v...)
}

func ErrorSince(opt Operation, msg string, start time.Time, v ...any) {
	logSince(logger.Error, opt, msg, start, v...)
}
