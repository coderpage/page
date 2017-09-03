package app

import (
	"page/constant/rsp"
	"page/controller"
)

type HomeHandler struct {
	controller.BaseController
}

// @router / [get]
func (handler *HomeHandler) GetHomePage() {
	handler.Data[rsp.PageTitle] = "首页"
	handler.TplName = "home.html"
}
