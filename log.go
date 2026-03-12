package tlog

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func log(fn func(string, ...any), o Task, msg string, v ...any) {
	fn(msg, append([]any{"task", o}, v...)...)
}

func appendError(err error, v ...any) []any {
	return append([]any{"error", err}, v...)
}

func (t Task) Info(msg string, v ...any) {
	log(logger.Info, t, msg, v...)
}

func (t Task) Debug(msg string, v ...any) {
	log(logger.Debug, t, msg, v...)
}

func (t Task) Warn(msg string, v ...any) {
	log(logger.Warn, t, msg, v...)
}

func (t Task) Error(err error, msg string, v ...any) {
	log(logger.Error, t, msg, appendError(err, v...)...)
}
