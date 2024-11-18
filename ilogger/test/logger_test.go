package test

import (
	"ilogger"
	"ilogger/izap"
	"testing"
)

func Test(t *testing.T) {
	// izap 注册 LoggerRegisterFunc
	izap.Register()
	// 初始化
	ilogger.Init()

	ilogger.Debug("i am debug")
	ilogger.Info("i am info")
	ilogger.Warn("i am warn")
	ilogger.Error("i am error")
}
