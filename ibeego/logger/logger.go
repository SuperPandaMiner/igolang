package logger

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"ibeego/conf"
	"time"
)

const (
	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
)

func Init() {
	levelString := conf.Logger.Level
	level := logs.LevelInformational
	if levelString == levelDebug {
		level = logs.LevelDebug
	} else if levelString == levelWarn {
		level = logs.LevelWarning
	} else if levelString == levelError {
		level = logs.LevelError
	}

	output := conf.Logger.Out
	if output == logs.AdapterConsole {
		// beego 会自动初始化 console logger
		// err := logs.SetLogger(logs.AdapterConsole, fmt.Sprintf(`{"level":%d}`, level))
		logs.SetLevel(level)
	} else {
		filename := fmt.Sprintf("log.%d.log", time.Now().UnixMilli())
		maxsize := conf.Logger.Maxsize
		maxDays := conf.Logger.Maxdays
		output = logs.AdapterMultiFile

		err := logs.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(`{"level":%d,"filename":"%s","maxsize":%d,"maxdays":%d}`, level, filename, maxsize, maxDays))
		if err != nil {
			panic(err)
		}
	}

	logs.EnableFuncCallDepth(true)
}
