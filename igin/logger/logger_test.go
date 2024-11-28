package logger

import (
	"igin/config"
	"ilogger"
	"testing"
)

func Test(t *testing.T) {
	config.Init("../../iconfig/config.yml")

	Init()

	ilogger.Debug("i am debug")
	ilogger.Info("i am info")
	ilogger.Warn("i am warn")
	ilogger.Error("i am error")
}
