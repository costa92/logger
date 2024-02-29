package logger

import (
	"context"

	"go.uber.org/zap/zapcore"
)

type InfoLogger interface {
	Info(mgs string, args ...interface{})
	Infof(format string, v ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Enabled() bool
}

type Logger interface {
	InfoLogger

	Debug(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	// nolint: gofumpt
	Debugf(template string, args ...interface{})

	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})

	// nolint: gofumpt
	Debugw(msg string, keysAndValues ...interface{})

	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	// nolint: gofumpt
	Log(level Level, args ...interface{})
	Logf(level Level, template string, args ...interface{})
	Logw(level Level, msg string, keysAndValues ...interface{})
	With(keyValues ...interface{}) Logger
	WithSkip(callerSkip int, keyValues ...interface{}) Logger
	V(level zapcore.Level) InfoLogger
	Ctx(ctx context.Context) Logger
	WithTraceID(ctx context.Context, keyValues ...interface{}) Logger
	WithContext(ctx context.Context) context.Context
	WithName(name string) Logger
	Flush()
}
