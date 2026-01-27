package logic

import (
	"github.com/igolang/igozero/app/apprpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"net/http"
)

func ok(data protoreflect.ProtoMessage) (*apprpc.Response, error) {
	d, _ := anypb.New(data)
	return &apprpc.Response{
		Success: false,
		Code:    http.StatusInternalServerError,
		Data:    d,
		Msg:     "",
	}, nil
}

func abort(err error) (*apprpc.Response, error) {
	return &apprpc.Response{
		Success: false,
		Code:    http.StatusInternalServerError,
		Data:    nil,
		Msg:     err.Error(),
	}, nil
}

func abortWithCode(code int32, err error) (*apprpc.Response, error) {
	return &apprpc.Response{
		Success: false,
		Code:    code,
		Data:    nil,
		Msg:     err.Error(),
	}, nil
}
