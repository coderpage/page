package router

import (
	"page/common"

	"github.com/astaxie/beego"
)

func registerCommons() {
	beego.Router("/", &common.HomeHandler{}, "get:GetHomePage")
}
