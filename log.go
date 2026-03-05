package oplog

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type Operation string

func log(fn func(string, ...any), opt Operation, msg string, v ...any) {
	fn(msg, append([]any{"opt", opt}, v...)...)
}

func (opt Operation) Info(msg string, v ...any) {
	log(logger.Info, opt, msg, v...)
}

func (opt Operation) Debug(msg string, v ...any) {
	log(logger.Debug, opt, msg, v...)
}

func (opt Operation) Warn(msg string, v ...any) {
	log(logger.Warn, opt, msg, v...)
}

func (opt Operation) Error(err error, msg string, v ...any) {
	log(logger.Error, opt, msg, append([]any{"error", err}, v...)...)
}
