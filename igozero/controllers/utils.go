package controllers

import (
	"igozero/routers"
	"net/http"
)

func abort(err error) (any, error) {
	return nil, &routers.StatusError{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
	}
}

func abortWithCode(code int, err error) (any, error) {
	return nil, &routers.StatusError{
		Code: code,
		Msg:  err.Error(),
	}
}
