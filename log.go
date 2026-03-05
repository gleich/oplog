package oplog

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func log(fn func(string, ...any), msg string, opt string, v ...any) {
	fn(msg, append([]any{"operation", opt}, v...)...)
}

func Info(msg, opt string, v ...any) {
	log(logger.Info, msg, opt, v...)
}

func Debug(msg, opt string, v ...any) {
	log(logger.Debug, msg, opt, v...)
}

func Warn(msg, opt string, v ...any) {
	log(logger.Warn, msg, opt, v...)
}

func Error(err error, msg, opt string, v ...any) {
	log(logger.Error, msg, opt, append([]any{"error", err}, v...)...)
}
