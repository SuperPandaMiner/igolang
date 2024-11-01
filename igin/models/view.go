package models

import "net/http"

type ModelView struct {
	// id
	ID int64 `json:"id,string"` // string 防止 ctx.json 精度丢失
	// 创建时间
	CreatedAt int64 `json:"createdAt"`
	// 更新时间
	UpdatedAt int64 `json:"updatedAt"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdateBy string `json:"updateBy"`
}

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
