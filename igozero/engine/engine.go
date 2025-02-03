package engine

import (
	"github.com/zeromicro/go-zero/rest"
	"igozero/config"
	"igozero/routers"
)

var ShutDown func()

func Run() {
	engine := rest.MustNewServer(config.Config.RestConf, rest.WithNotFoundHandler(routers.NotFoundHandler()))

	routers.Router(engine)

	ShutDown = func() {
		engine.Stop()
	}

	go func() {
		engine.Start()
	}()
}
