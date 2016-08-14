package web

import (
	"page/constant/rsp"
)

func (handler *EditorHandler) GetEditorPage() {
	handler.Data[rsp.PageTitle] = "编辑"
	handler.TplName = "editor.html"
}
