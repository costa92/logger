package logger

import "context"

type key int

const (
	logContextKey key = iota
)

func (l *logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, l)
}

func (l *logger) Ctx(ctx context.Context) Logger {
	return l.WithCallerSkip(ctx, defaultCallerSkip, TraceEvent)
}

func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		logger := ctx.Value(logContextKey)
		if logger != nil {
			return logger.(Logger)
		}
	}
	return WithName("Unknown-Context")
}
