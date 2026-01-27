package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/igolang/ilogger"
	"github.com/igolang/ilogger/izap"
)

func Init() {
	// 注册 izap
	izap.Register()
	ilogger.Init()

	gin.DefaultWriter = ilogger.LoggerWriter()
	gin.DefaultErrorWriter = ilogger.LoggerWriter()
}
