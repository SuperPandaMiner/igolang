package routers

import (
	"context"
	"github.com/igolang/igozero/core/models"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var HandlerRegisterFunc = func(engine *rest.Server) {}

func Router(engine *rest.Server) {
	// response
	httpx.SetOkHandler(func(ctx context.Context, resp any) any {
		return models.OkResponse(resp)
	})
	// 替换全局错误处理
	//httpx.SetErrorHandlerCtx(errorHandler)

	// cors
	engine.Use(rest.ToMiddleware(cors))

	// api
	HandlerRegisterFunc(engine)
}
