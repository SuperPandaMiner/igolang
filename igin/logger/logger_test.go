package logger

import (
	"github.com/igolang/ilogger"
	"igin/config"
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
