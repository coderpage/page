package web

import (
	"page/web"

	"github.com/astaxie/beego"
)

func Register() {
	beego.Router("/m/editor", &web.EditorHandler{}, "get:GetEditorPage")
}
