package logger

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"ibeego/conf"
	"testing"
)

func init() {
	config.InitGlobalInstance("ini", "../conf/app.conf")
	conf.Init()
	Init()
}

func TestLog(t *testing.T) {
	logs.Debug("i am debug")
	logs.Info("i am info")
	logs.Warn("i am warn")
	logs.Error("i am error")
	logs.Critical("oh, crash")
}
