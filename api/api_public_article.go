package api

import (
	"fmt"
	"page/constant/status"
	"page/controller"
	"page/model"
	"page/storage"
)

type publishRequestBody struct {
	ArticleID      int64  `json:"article_id"`
	ArtileTitle    string `json:"title"`
	ArticleContent string `json:"content"`
}

// Publish 发布文章
// url: /api/article/publish
// method: POST
// header:
// Content-Type : application/json; charset=utf-8
// X-ACCESS-TOKEN : access-token
// body: {"article_id":id,"title":"","content":"","tags":[]}
// errors:
// 10012 令牌有误，请求被拒绝
// 10001 无法解析 body 内容
// 500   服务器内部错误
func (handler *ArticleEditorHandler) PublishArticle() {
	auth, ok := handler.CheckAuth()
	if !ok {
		return
	}
	response := controller.NewResponse()

	body := &publishRequestBody{}
	err := handler.ReadJsonBody(body)
	if err != nil {
		response.SetStatus(status.UnprocessableEntity)
		response.SetMessage(status.Text(status.UnprocessableEntity))
		handler.HandleJsonResponse(response)
		return
	}

	if body.ArticleID == 0 {
		response.SetStatus(status.IllegalReqParam)
		response.SetMessage("miss param: article_id")
		handler.HandleJsonResponse(response)
		return
	}

	article, err := storage.GetArticleById(body.ArticleID)
	if err != nil {
		if err == storage.ErrNoRows {
			response.SetStatus(status.SourceNotFound)
			response.SetMessage(fmt.Sprintf("can't find article with id:%d", body.ArticleID))
		} else {
			response.SetStatus(status.ServerInternalErr)
			response.SetMessage(err.Error())
		}
		handler.HandleJsonResponse(response)
		return
	}

	if article.UserId != auth.Uid {
		response.SetStatus(status.AccessDenied)
		response.SetMessage(status.Text(status.AccessDenied))
		handler.HandleJsonResponse(response)
		return
	}

	var columns []string
	if body.ArtileTitle != "" {
		article.Title = body.ArtileTitle
		columns = append(columns, "title")
	}
	if body.ArticleContent != "" {
		article.Content = body.ArticleContent
		columns = append(columns, "content")
	}
	article.Status = model.ArticleStatusPublish
	columns = append(columns, "status")

	err = storage.UpdateArticle(article, columns...)
	if err != nil {
		response.SetStatus(status.ServerInternalErr)
		response.SetMessage(err.Error())
		handler.HandleJsonResponse(response)
		return
	}

	response.SetStatus(status.OK)
	response.SetMessage("publish success")
	handler.HandleJsonResponse(response)
	return
}
