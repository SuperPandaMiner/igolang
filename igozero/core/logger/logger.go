package logger

import (
	"github.com/igolang/igozero/core/config"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func Init() {
	config.RestConfig.Log.Encoding = "plain"
	config.RestConfig.Log.TimeFormat = time.DateTime
	config.RestConfig.Log.Path = "./log"
	config.RestConfig.Log.FileTimeFormat = time.DateTime

	logx.MustSetup(config.RestConfig.Log)
}
