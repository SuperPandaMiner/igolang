package logger

import (
	"github.com/zeromicro/go-zero/core/logx"
	"igozero/config"
	"time"
)

func Init() {
	config.Config.Log.Encoding = "plain"
	config.Config.Log.TimeFormat = time.DateTime
	config.Config.Log.Path = "./log"
	config.Config.Log.FileTimeFormat = time.DateTime

	logx.MustSetup(config.Config.Log)
}
