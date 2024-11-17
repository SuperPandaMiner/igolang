package models

import "net/http"

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func OkResponse(object interface{}) *Response {
	return &Response{
		Success: true,
		Code:    http.StatusOK,
		Msg:     "success",
		Data:    object,
	}
}

func ErrorResponse(msg string) *Response {
	return &Response{
		Success: false,
		Code:    http.StatusInternalServerError,
		Msg:     msg,
		Data:    nil,
	}
}

func ErrorResponseWithCode(code int, msg string) *Response {
	return &Response{
		Success: false,
		Code:    code,
		Msg:     msg,
		Data:    nil,
	}
}
