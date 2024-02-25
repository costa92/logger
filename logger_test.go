package logger

import (
	"context"
	"testing"
)

func show() {
	Info("info message")
	Infow("info message", "da", "123")
}

func TestLevel(t *testing.T) {
	show()
}

func TestNewOptions(t *testing.T) {
	opts := NewOptions()
	opts.Level = "info"
	opts.Format = "json"
	opts.EnableColor = true
	opts.EnableCaller = true
	opts.OutputPaths = []string{"stdout"}
	opts.ErrorOutputPaths = []string{"stderr"}
	opts.Development = false
	opts.DisableCaller = true
	opts.DisableStacktrace = true
	opts.FieldPair = []interface{}{
		FieldPair{"service", "client_string"},
	}
	t.Log(opts)

	errs := opts.Validate()

	for _, err := range errs {
		t.Log(err)
	}
	Init(opts)
	With("trace_id", "274ac2bbf9d5")
	show()
	Debug("debug message")
	Errorw("error message", "da", "123")
}

func Test_Log_WithTraceID(t *testing.T) {
	opts := NewOptions()
	opts.Level = "info"
	opts.Format = "json"
	opts.EnableColor = true
	opts.EnableCaller = true
	opts.OutputPaths = []string{"stdout"}
	opts.ErrorOutputPaths = []string{"stderr"}
	opts.Development = false
	opts.DisableCaller = true
	opts.DisableStacktrace = true
	opts.FieldPair = []interface{}{
		"service", "client_string",
	}
	Init(opts)
	ctx := context.Background()
	//WithTraceID(ctx)
	log := WithTraceID(ctx, "trace_id", "274ac2bbf9d5")
	log.Info("info message")

	//show()
}

//func TestColorLogger(t *testing.T) {
//	config := NewDevelopmentConfig()
//	config.Level = DebugLevel
//	config.DisableStacktrace = true
//	config.DisableCaller = true
//	SetConfig(config)
//
//	show()
//}
//
//func TestProdLogger(t *testing.T) {
//	config := NewProductionConfig()
//	config.Level = DebugLevel
//	config.DisableStacktrace = true
//	config.DisableCaller = true
//	SetConfig(config)
//	var l Level
//	l.Set("debug")
//	SetLevel(l)
//
//	With("trace_id", "274ac2bbf9d5")
//	With("span_id", "383d60f1")
//
//	show()
//}
//
//func TestProdLoggerMap(t *testing.T) {
//	config := NewProductionConfig(FieldPair{"service", "client_string"})
//	config.Level = DebugLevel
//	config.DisableStacktrace = true
//	config.DisableCaller = true
//	SetConfig(config)
//	var l Level
//	l.Set("debug")
//	SetLevel(l)
//
//	With("trace_id", "274ac2bbf9d5")
//	With("span_id", "383d60f1")
//
//	show()
//}
//
//func TestNewDefaultConfig(t *testing.T) {
//	lc := NewDefaultConfig()
//	lc.DisableStacktrace = true
//	lc.EnableColor = true
//	lc.Encoding = "console"
//	lc.InitialFields = map[string]interface{}{
//		"service": "client_string",
//	}
//
//	SetConfig(lc)
//
//	show()
//}
