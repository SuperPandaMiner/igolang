package routers

import (
	"context"
	"errors"
	"fmt"
	"igozero/models"
	"net/http"
	"utils"
)

type StatusError struct {
	Code int
	Msg  string
}

func (e StatusError) Error() string {
	return e.Msg
}

func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, _ := utils.ToJsonString(models.ErrorResponseWithCode(http.StatusNotFound, "page not found"))
		HttpError(w, http.StatusNotFound, resp)
	}
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// 使用 errorHandlerFunc，以便其他地方错误处理时能够执行源码判断 GrpcError 逻辑。httpx#responses#doHandleError()
func errorHandler(ctx context.Context, err error) (int, any) {
	var statusError *StatusError
	if errors.As(err, &statusError) {
		return statusError.Code, models.ErrorResponseWithCode(statusError.Code, statusError.Msg)
	} else {
		return http.StatusInternalServerError, models.ErrorResponse(err.Error())
	}
}

func errorHandlerFunc(w http.ResponseWriter, err error) {
	var statusError *StatusError
	var resp string
	var code int
	if errors.As(err, &statusError) {
		resp, _ = utils.ToJsonString(models.ErrorResponseWithCode(statusError.Code, statusError.Msg))
		code = statusError.Code
	} else {
		resp, _ = utils.ToJsonString(models.ErrorResponse(err.Error()))
		code = http.StatusInternalServerError
	}

	HttpError(w, code, resp)
}

func HttpError(w http.ResponseWriter, code int, resp any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, resp)
}
