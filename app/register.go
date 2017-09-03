package app

import (
	"page/constant/rsp"
	"page/controller"
)

type RegisterHandler struct {
	controller.BaseController
}

// @router /register [get]
func (this *RegisterHandler) GetRegisterPage() {
	this.Data[rsp.PageTitle] = "注册"
	this.TplName = "register.html"
}
