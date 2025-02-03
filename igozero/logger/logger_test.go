package logger

import (
	"github.com/zeromicro/go-zero/core/logx"
	"igozero/config"
	"testing"
)

func Test(t *testing.T) {
	config.Init("../config.yml")

	Init()

	logx.Debug("i am debug")
	logx.Infof("%s, i am info", "hi")
	logx.Error("i am error")
	logx.Severe("i am severe")
	logx.Stat("i am stat")
	logx.Slow("i am slow")

	//for i := 1; i < 10000; i++ {
	//	logx.Infof("%s, i am info", "hi")
	//}
}
