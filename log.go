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

func Info(opt Operation, msg string, v ...any) {
	log(logger.Info, opt, msg, v...)
}

func Debug(opt Operation, msg string, v ...any) {
	log(logger.Debug, opt, msg, v...)
}

func Warn(opt Operation, msg string, v ...any) {
	log(logger.Warn, opt, msg, v...)
}

func Error(err error, opt Operation, msg string, v ...any) {
	log(logger.Error, opt, msg, append([]any{"error", err}, v...)...)
}
