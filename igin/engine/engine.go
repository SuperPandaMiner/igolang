package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"igin/config/iconfig"
	"igin/routers"
	"ilogger"
	"net/http"
	"time"
)

const (
	dev  = "dev"
	prod = "prod"
)

var mode = dev

var ShutDown func()

func Run() {
	mode = iconfig.Server.Mode

	if IsModeProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := routers.Router()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", iconfig.Server.Addr, iconfig.Server.Port),
		Handler: engine,
	}

	ShutDown = func() {
		// 禁止接收新请求并等待处理未完成的请求
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			ilogger.Error(err.Error())
		}
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ilogger.Info("listen: %s\n", err)
		}
	}()
}

func Mode() string {
	return mode
}

func IsModeProd() bool {
	return mode != dev
}
