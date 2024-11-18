package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"iconfig"
	"iecho/routers"
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

	engine := echo.New()
	engine.Use(middleware.Recover())

	// request logger
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} \n",
		Output: ilogger.LoggerWriter(),
	}))
	// echo logger
	engine.Logger.SetHeader("${time_rfc3339} ${level} ${short_file}:${line}")
	engine.Logger.SetOutput(ilogger.LoggerWriter())

	if IsModeProd() {
		engine.HideBanner = true
		engine.Logger.SetLevel(log.WARN)
	} else {
		// 默认 false，会设置 echo logger 等级为 debug
		engine.Debug = true
	}

	routers.Router(engine)

	ShutDown = func() {
		// 禁止接收新请求并等待处理未完成的请求
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := engine.Shutdown(ctx); err != nil {
			ilogger.Error(err.Error())
		}
	}

	addr := fmt.Sprintf("%s:%s", iconfig.Server.Addr, iconfig.Server.Port)
	go func() {
		if err := engine.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
