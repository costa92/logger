package logger

func Debug(args ...interface{}) {
	std.zapLogger.Sugar().Debug(args...)
}

func Info(args ...interface{}) {
	std.zapLogger.Sugar().Info(args...)
}

func Warn(args ...interface{}) {
	std.zapLogger.Sugar().Warn(args...)
}

func Error(args ...interface{}) {
	std.zapLogger.Sugar().Error(args...)
}

func DPanic(args ...interface{}) {
	std.zapLogger.Sugar().DPanic(args...)
}

func Panic(args ...interface{}) {
	std.zapLogger.Sugar().Panic(args...)
}

func Fatal(args ...interface{}) {
	std.zapLogger.Sugar().Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	std.zapLogger.Sugar().Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	std.zapLogger.Sugar().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	std.zapLogger.Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	std.zapLogger.Sugar().Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	std.zapLogger.Sugar().DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	std.zapLogger.Sugar().Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	std.zapLogger.Sugar().Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func Sync() error {
	return std.zapLogger.Sugar().Sync()
}

//func With(keyValues ...interface{}) Logger {
//	return std.With(keysAndValues...)
//}

//func WithTraceID(ctx context.Context, keyValues ...interface{}) Logger {
//	return l.WithTraceID(ctx, keyValues...)
//}

//func Ctx(ctx context.Context) Logger {
//	return l.Ctx(ctx)
//}

//func GetZapLogger() *zap.Logger {
//	return l.GetZapLogger()
//}
