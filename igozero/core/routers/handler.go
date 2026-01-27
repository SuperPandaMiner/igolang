package routers

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type HandlerFunc func(r *http.Request) (any, error)

func Route(method string, path string, handler HandlerFunc) rest.Route {
	return rest.Route{
		Method:  method,
		Path:    path,
		Handler: Handler(handler),
	}
}

func Handler(handle HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := handle(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err, errorHandlerFunc)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
