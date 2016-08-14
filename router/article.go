package router

import (
	"page/editor"

	"github.com/astaxie/beego"
)

func registerArticles() {
	// APIs
	beego.Router("/api/article/create", &editor.ArticleEditorHandler{}, "post:NewArticle")
	beego.Router("/api/article/get", &editor.ArticleEditorHandler{}, "get:GetArticle")
	beego.Router("/api/article/publish", &editor.ArticleEditorHandler{}, "post:PublishArticle")
}
