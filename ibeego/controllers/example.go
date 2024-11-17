package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"ibeego/models"
)

type ExampleController struct {
	BaseController
}

func (ctrl ExampleController) Success() {
	ctrl.writeSuccess("success")
}

func (ctrl ExampleController) Option() {
	ctrl.Ctx.WriteString("option")
}

func (ctrl ExampleController) BadRequest() {
	ctrl.abort400("400")
}

func (ctrl ExampleController) Authentication() {
	ctrl.abort401()
}

func (ctrl ExampleController) Forbidden() {
	ctrl.abort403()
}

func (ctrl ExampleController) InternalServerError() {
	ctrl.abort("internal server error")
}

func (ctrl ExampleController) Panic() {
	panic("panic")
}

func (ctrl ExampleController) Query() {
	query := make(map[string]interface{})
	query["string"] = ctrl.GetString("string")
	i, _ := ctrl.GetInt("int")
	query["int"] = i
	b, _ := ctrl.GetBool("bool")
	query["bool"] = b
	ctrl.writeSuccess(query)
}

func (ctrl ExampleController) Path() {
	ctrl.writeSuccess(ctrl.Ctx.Input.Param(":id"))
}

func (ctrl ExampleController) Body() {
	type param struct {
		String string `json:"string" valid:"Required"`
		Int    int    `json:"int" valid:"Range(1, 10)"`
		Bool   bool   `json:"bool"`
		Name   string `json:"name" valid:"MaxSize(5)"`
	}

	body := param{}
	err := ctrl.BindJSON(&body)
	if err != nil {
		ctrl.abort(err.Error())

	}
	ok, errors := ctrl.validWithoutAbort(&body)
	if !ok {
		fmt.Println(errors)
	}
	ctrl.valid(&body)
	ctrl.writeSuccess(body)
}

func TokenFilter(ctx *context.Context) {
	token := ctx.Input.Header("token")
	if token == "" {
		ctx.Output.SetStatus(401)
		_ = ctx.JSONResp(models.ErrorResponseWithCode(401, "authorization required"))
	} else {
		ctx.Input.SetData("token", token)
	}
}
