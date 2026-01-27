package logger

import (
	"github.com/igolang/ilogger"
	"github.com/igolang/ilogger/izap"
)

func Init() {
	// 注册 izap
	izap.Register()
	ilogger.Init()
}
