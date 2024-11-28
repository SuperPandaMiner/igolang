package test

import (
	"iconfig"
	"iconfig/iviper"
	"ilogger"
	"ilogger/izerolog"
	"testing"
)

func Test(t *testing.T) {
	iviper.Register("../../iconfig/config.yml")
	iconfig.Init()

	// izap
	//izap.Register()
	// zerolog
	izerolog.Register()
	// 初始化
	ilogger.Init()

	ilogger.Debug("i am debug")
	ilogger.Info("i am info")
	ilogger.Warn("i am warn")
	ilogger.Error("i am error")
}
