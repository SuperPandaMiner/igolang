package logger

import (
	"github.com/gin-gonic/gin"
	"ilogger"
	"ilogger/izap"
)

func Init() {
	// 注册 izap
	izap.Register()
	ilogger.Init()

	gin.DefaultWriter = ilogger.LoggerWriter()
	gin.DefaultErrorWriter = ilogger.LoggerWriter()
}
