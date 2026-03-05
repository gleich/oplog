package oplog

import (
	"time"
)

func logSince(fn func(string, ...any), opt Opt, msg string, start time.Time, v ...any) {
	log(fn, opt, msg, append([]any{"duration", formatDuration(time.Since(start))}, v...)...)
}

func (opt Opt) InfoSince(msg string, start time.Time, v ...any) {
	logSince(logger.Info, opt, msg, start, v...)
}

func (opt Opt) DebugSince(msg string, start time.Time, v ...any) {
	logSince(logger.Debug, opt, msg, start, v...)
}

func (opt Opt) WarnSince(msg string, start time.Time, v ...any) {
	logSince(logger.Warn, opt, msg, start, v...)
}

func (opt Opt) ErrorSince(msg string, start time.Time, v ...any) {
	logSince(logger.Error, opt, msg, start, v...)
}
