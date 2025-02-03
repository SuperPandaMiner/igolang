package controllers

import (
	"errors"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"igozero/config"
	"igozero/engine"
	"igozero/logger"
	"igozero/models"
	"igozero/routers"
	"net/http"
	"testing"
	"utils"
)

func init() {
	config.Init("../config.yml")

	logger.Init()

	routers.HandlerRegisterFunc = func(engine *rest.Server) {
		var routes []rest.Route
		routes = append(routes, routers.Route(http.MethodGet, "/success", success))
		routes = append(routes, routers.Route(http.MethodGet, "/400", badRequest))
		routes = append(routes, routers.Route(http.MethodGet, "/401", authentication))
		routes = append(routes, routers.Route(http.MethodGet, "/403", forbidden))
		routes = append(routes, routers.Route(http.MethodGet, "/500", internalServerError))
		routes = append(routes, routers.Route(http.MethodGet, "/panic", _panic))
		routes = append(routes, routers.Route(http.MethodGet, "/query", query))
		routes = append(routes, routers.Route(http.MethodGet, "/path/:id", path))
		routes = append(routes, routers.Route(http.MethodPost, "/post", post))
		routes = append(routes, routers.Route(http.MethodGet, "/any/:*", success))

		engine.AddRoutes(
			rest.WithMiddleware(tokenFilter, routes...),
			rest.WithPrefix("/igozero"),
		)
	}
}

func success(r *http.Request) (any, error) {
	return "success", nil
}

func badRequest(r *http.Request) (any, error) {
	return abortWithCode(400, errors.New("400"))
}

func authentication(r *http.Request) (any, error) {
	return abortWithCode(401, errors.New("401"))
}

func forbidden(r *http.Request) (any, error) {
	return abortWithCode(403, errors.New("403"))
}

func internalServerError(r *http.Request) (any, error) {
	return abort(errors.New("internal server error"))
}

func _panic(r *http.Request) (any, error) {
	panic("panic")
}

// http://127.0.0.1:8080/igozero/query?name=zhangsan&num=1&d=default&age=18
func query(r *http.Request) (any, error) {
	type Request struct {
		Name string `form:"name"`
		Num  int    `form:"num, range=[1:10)"`
		D    string `form:"d, default=d"`
		O    string `form:"o, optional"`
		Age  int    `form:"age, options=18|19"`
	}
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		return abort(err)
	}
	return req, nil
}

func path(r *http.Request) (any, error) {
	type Request struct {
		Id string `path:"id"`
	}
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		return abort(err)
	}
	return req.Id, nil
}

func post(r *http.Request) (any, error) {
	type Request struct {
		Name string `json:"name"`
		Num  int    `json:"num,range=[1:10)"`
		D    string `json:"d,default=d"`
		O    string `json:"o,optional"`
		Age  int    `json:"age,options=18|19"`
	}
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		return abort(err)
	}
	return req, nil
}

// Header Token
func tokenFilter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Token string `header:"token"`
		}
		var req Request
		err := httpx.ParseHeaders(r, &req)
		if err != nil || req.Token == "" {
			resp, _ := utils.ToJsonString(models.ErrorResponseWithCode(http.StatusUnauthorized, "token required"))
			routers.HttpError(w, http.StatusUnauthorized, resp)
			return
		}
		next(w, r)
	}
}

func Test(t *testing.T) {

	engine.Run()

	select {}
}
