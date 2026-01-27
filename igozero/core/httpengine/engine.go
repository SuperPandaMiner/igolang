package httpengine

import (
	"github.com/igolang/igozero/core/config"
	"github.com/igolang/igozero/core/routers"
	"github.com/zeromicro/go-zero/rest"
)

var ShutDown func()

func Run() {
	engine := rest.MustNewServer(config.RestConfig.RestConf, rest.WithNotFoundHandler(routers.NotFoundHandler()))

	routers.Router(engine)

	ShutDown = func() {
		engine.Stop()
	}

	go func() {
		engine.Start()
	}()
}
