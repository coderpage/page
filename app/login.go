package app

import (
	"page/constant/rsp"
	"page/controller"
)

type LoginHandler struct {
	controller.BaseController
}

// @router /login [get]
func (handler *LoginHandler) GetLoginPage() {
	handler.Data[rsp.PageTitle] = "登录"
	handler.TplName = "login.html"
}
