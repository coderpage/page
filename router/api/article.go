package api

import (
	"page/api"

	"github.com/astaxie/beego"
)

func RegisterArticles() {
	beego.Router("/api/article/create", &api.ArticleEditorHandler{}, "post:NewArticle")
	beego.Router("/api/article/get", &api.ArticleEditorHandler{}, "get:GetArticle")
	beego.Router("/api/article/publish", &api.ArticleEditorHandler{}, "post:PublishArticle")
}
