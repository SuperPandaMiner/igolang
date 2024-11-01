package izap

import (
	"ilogger"
	"testing"
)

func Test(t *testing.T) {
	config := Config{
		Out:          ilogger.ConsoleLog,
		Level:        ilogger.DebugLevel,
		LoggerNumber: 1,
	}
	// 注册 LoggerRegisterFunc
	Register(&config)
	// 初始化
	ilogger.Init()

	ilogger.Debug("i am debug")
	ilogger.Info("i am info")
	ilogger.Warn("i am warn")
	ilogger.Error("i am error")
}
