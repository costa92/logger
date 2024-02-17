package main

import "github.com/costa92/logger"

func main() {
	opts := &logger.Options{
		Level:            "debug",
		Format:           "console",
		EnableColor:      true,
		EnableCaller:     true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
	}

	// 初始化全局logger
	logger.Init(opts)
	defer logger.Flush()
	logger.Debug("1")
}
