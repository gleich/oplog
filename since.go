package tlog

import (
	"time"
)

func logSince(fn func(string, ...any), task Task, msg string, start time.Time, v ...any) {
	log(fn, task, msg, append([]any{"duration", formatDuration(time.Since(start))}, v...)...)
}

func (t Task) InfoSince(msg string, start time.Time, v ...any) {
	logSince(logger.Info, t, msg, start, v...)
}

func (t Task) DebugSince(msg string, start time.Time, v ...any) {
	logSince(logger.Debug, t, msg, start, v...)
}

func (t Task) WarnSince(msg string, start time.Time, v ...any) {
	logSince(logger.Warn, t, msg, start, v...)
}

func (t Task) ErrorSince(msg string, start time.Time, v ...any) {
	logSince(logger.Error, t, msg, start, v...)
}
