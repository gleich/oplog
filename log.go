package oplog

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type Opt string

func log(fn func(string, ...any), opt Opt, msg string, v ...any) {
	fn(msg, append([]any{"opt", opt}, v...)...)
}

func (opt Opt) Info(msg string, v ...any) {
	log(logger.Info, opt, msg, v...)
}

func (opt Opt) Debug(msg string, v ...any) {
	log(logger.Debug, opt, msg, v...)
}

func (opt Opt) Warn(msg string, v ...any) {
	log(logger.Warn, opt, msg, v...)
}

func (opt Opt) Error(err error, msg string, v ...any) {
	log(logger.Error, opt, msg, append([]any{"error", err}, v...)...)
}
