package api

import (
	"page/constant/rqs"
	"page/constant/rsp"
	"page/constant/status"
	"page/controller"
	"page/storage"
)

// GetArticle 获取编辑的文章
// url: /api/article/get?article_id=xxx
// method: get
// header:
// X-ACCESS-TOKEN : access-token
// response: {"article_id":"","title":"","content":""...}
func (handler *ArticleEditorHandler) GetArticle() {
	//	var auth *model.Auth
	auth, ok := handler.CheckAuth()
	if !ok {
		return
	}
	response := controller.NewResponse()

	articleID, err := handler.GetInt64(rqs.BodyArticleId, -1)
	if err != nil {
		response.SetStatus(status.IllegalReqParam)
		response.SetMessage(err.Error())
		handler.HandleJsonResponse(response)
		return
	}

	if articleID == -1 {
		response.SetStatus(status.IllegalReqParam)
		response.SetMessage(err.Error())
		handler.HandleJsonResponse(response)
		return
	}

	article, err := storage.GetArticleById(articleID)
	if err != nil {
		if err == storage.ErrNoRows {
			response.SetStatus(status.SourceNotFound)
			response.SetMessage(status.Text(status.SourceNotFound))
			handler.HandleJsonResponse(response)
			return
		}
		response.SetStatus(status.ServerInternalErr)
		response.SetMessage(err.Error())
		handler.HandleJsonResponse(response)
		return
	}

	if article.UserId != auth.Uid {
		response.SetStatus(status.AccessDenied)
		handler.HandleJsonResponse(response)
		return
	}
	response.SetStatus(status.OK)
	response.SetData(rsp.BodyArticle, article)
	handler.HandleJsonResponse(response)
	return

}
