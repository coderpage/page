package api

import (
	"page/constant/rsp"
	"page/constant/status"
	"page/controller"
	"page/model"
	"page/storage"
)

// NewArticle 创建新文章
// url: /api/article/create
// method: POST
// header:
// X-ACCESS-TOKEN : access-token
// body: {"title":"","content":"","tags":[]}
// errors:
// 10012 令牌有误，请求被拒绝
// 10001 无法解析 body 内容
// 500   服务器内部错误
func (handler *ArticleEditorHandler) NewArticle() {
	// 检查 token
	auth, ok := handler.CheckAuth()
	if !ok {
		return
	}

	response := controller.NewResponse()

	//读取 body
	article := &model.Article{}
	err := handler.ReadJsonBody(article)
	if err != nil {
		response.SetStatus(status.UnprocessableEntity)
		response.SetMessage(err.Error())
		handler.HandleJsonResponse(response)
		return
	}

	article.UserId = auth.Uid
	err = storage.AddNewArticle(article)
	// 数据库创建新文章失败
	if err != nil {
		response.SetStatus(status.ServerInternalErr)
		response.SetMessage(err.Error())
		handler.HandleJsonResponse(response)
		return
	}
	// 数据库创建新文章成功
	response.SetStatus(status.OK)
	response.SetMessage("create new article success")
	response.SetData(rsp.BodyArticleId, article.Id)
	handler.HandleJsonResponse(response)
	return
}
