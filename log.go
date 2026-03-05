package oplog

import "log/slog"

func log(fn func(string, ...any), msg string, operation string, v ...any) {
	fn(msg, append([]any{"operation", operation}, v...)...)
}

func Info(msg string, operation string, v ...any) {
	log(slog.Info, msg, operation, v...)
}

func Debug(msg string, operation string, v ...any) {
	log(slog.Debug, msg, operation, v...)
}

func Warn(msg string, operation string, v ...any) {
	log(slog.Warn, msg, operation, v...)
}

func Error(err error, operation string, v ...any) {
	log(slog.Error, err.Error(), operation, v...)
}
