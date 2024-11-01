package logger

import (
	"github.com/gin-gonic/gin"
	"igin/config/iconfig"
	"ilogger"
	"izap"
)

func Init() {
	zapConfig := izap.Config{
		Out:   iconfig.Logger.Out,
		Level: iconfig.Logger.Level,
	}
	izap.Register(&zapConfig)

	ilogger.Init()

	gin.DefaultWriter = ilogger.LoggerWriter()
	gin.DefaultErrorWriter = ilogger.LoggerWriter()
}
