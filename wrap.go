package logger

import (
	"context"
	"go.uber.org/zap"
)

func Debug(args ...interface{}) {
	std.logger.Debug(args...)
}

func Info(args ...interface{}) {
	std.logger.Info(args...)
}

func Warn(args ...interface{}) {
	std.logger.Warn(args...)
}

func Error(args ...interface{}) {
	std.logger.Error(args...)
}

func DPanic(args ...interface{}) {
	std.logger.DPanic(args...)
}

func Panic(args ...interface{}) {
	std.logger.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.logger.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	std.logger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	std.logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	std.logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	std.logger.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	std.logger.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	std.logger.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	std.logger.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.logger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.logger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	std.logger.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	std.logger.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.logger.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func Sync() error {
	return std.logger.Sync()
}

func With(keyValues ...interface{}) Logger {
	return std.With(keyValues...)
}

func WithTraceID(ctx context.Context, keyValues ...interface{}) Logger {
	return std.WithTraceID(ctx, keyValues...)
}

func GetZapLogger() *zap.Logger {
	return std.GetZapLogger()
}
