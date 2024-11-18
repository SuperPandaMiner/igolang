package logger

import (
	"ilogger"
	"ilogger/izap"
)

func Init() {
	// 注册 izap
	izap.Register()
	ilogger.Init()
}
