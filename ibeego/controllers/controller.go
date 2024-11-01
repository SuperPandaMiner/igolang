package controllers

import (
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"ibeego/models"
	"strings"
)

const errorHandlerMsgKey = "error"

type BaseController struct {
	web.Controller
}

func (ctrl *BaseController) valid(obj interface{}) {
	ok, errors := ctrl.validWithoutAbort(obj)
	if !ok {
		ctrl.abort400(errors)
	}
}

func (ctrl *BaseController) validWithoutAbort(obj interface{}) (bool, string) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		ctrl.abort(err.Error())
	}
	if !ok {
		var validErrors []string
		for _, err := range valid.Errors {
			validErrors = append(validErrors, err.Message)
		}
		return false, strings.Join(validErrors, "; ")
	}
	return ok, ""
}

func (ctrl *BaseController) abort(err string) {
	ctrl.Data[errorHandlerMsgKey] = err
	ctrl.Abort("500")
}

func (ctrl *BaseController) abort400(err string) {
	ctrl.Data[errorHandlerMsgKey] = err
	ctrl.Abort("400")
}

func (ctrl *BaseController) abort401() {
	ctrl.Abort("401")
}

func (ctrl *BaseController) abort403() {
	ctrl.Abort("403")
}

func (ctrl *BaseController) writeSuccess(data interface{}) {
	_ = ctrl.Ctx.JSONResp(models.OkResponse(data))
}

type ErrorController struct {
	web.Controller
}

func (ctrl *ErrorController) Error400() {
	ctrl.error(400, "bad request")
}

func (ctrl *ErrorController) Error401() {
	ctrl.error(401, "unauthorized")
}

func (ctrl *ErrorController) Error403() {
	ctrl.error(403, "forbidden")
}

func (ctrl *ErrorController) Error404() {
	ctrl.error(404, "page not found")
}

func (ctrl *ErrorController) Error500() {
	ctrl.error(500, "internal server error")
}

func (ctrl *ErrorController) error(code int, err string) {
	if errorString, ok := ctrl.Data[errorHandlerMsgKey].(string); ok {
		err = errorString
	}
	_ = ctrl.Ctx.JSONResp(models.ErrorResponseWithCode(code, err))
}
