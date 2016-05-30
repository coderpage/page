package common

import (
	"page/constant/rsp"
	"page/controller"
)

type HomeHandler struct {
	controller.BaseController
}

func (handler *HomeHandler) GetHomePage() {
	handler.Data[rsp.PageTitle] = "首页"
	handler.TplName = "home.html"
}
