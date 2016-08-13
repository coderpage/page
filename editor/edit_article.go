package editor

//import (
//	"encoding/json"
//	"io/ioutil"
//	"page/constant/rqs"
//	"page/constant/rsp"
//	"page/constant/status"
//	"page/controller"
//	"page/model"
//	"page/storage"
//	"strconv"
//	"time"
//)

//// GetArticle 获取编辑的文章
//// url: /api/article/get
//// method: post
//// header:
//// X-ACCESS-TOKEN : access-token
//// body: {"article_id":""}
//func (handler *ArticleEditorHandler) GetArticle() {
//	//	var auth *model.Auth
//	auth, ok := handler.checkAuth()
//	if !ok {
//		return
//	}
//	response := controller.NewResponse()

//HandleResponse:
//	handler.handleJsonResponse(response)

//	// 解析请求 body 数据
//	body := handler.Ctx.Request.Body
//	defer body.Close()
//	bodyBytes, err := ioutil.ReadAll(body)
//	if err != nil {
//		response.SetStatus(status.UnprocessableEntity)
//		response.SetMessage("read body failed")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}

//	rqsBd := make(map[string]string)
//	err = json.Unmarshal(bodyBytes, &rqsBd)
//	if err != nil {
//		response.SetStatus(status.UnprocessableEntity)
//		response.SetMessage("parse body failed")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}

//	articleId := rqsBd[rqs.BodyArticleId]
//	if articleId == "" {
//		response.SetStatus(status.IllegalReqParam)
//		response.SetMessage("miss reqest arg: article_id")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}

//	articleIdInt, err := strconv.ParseInt(articleId, 10, 64)
//	if err != nil {
//		response.SetStatus(status.IllegalReqParam)
//		response.SetMessage("article_id type wrong")
//		goto HandleResponse
//	}
//	// 查询文章，返回最新的版本，有可能是草稿类型
//	articles, err := storage.GetArticleByIdWithDraft(articleIdInt)
//	if articles != nil && len(articles) > 0 {
//		article := articles[0]
//		if article.UserId != auth.Uid {
//			response.SetStatus(status.AccessDenied)
//			response.SetMessage(status.Text(status.AccessDenied))
//			goto HandleResponse
//		}
//		response.SetStatus(status.OK)
//		response.SetData(rsp.BodyArticle, articles[0])
//		goto HandleResponse
//	} else {
//		response.SetStatus(status.SourceNotFound)
//		response.SetMessage(status.Text(status.SourceNotFound))
//		goto HandleResponse
//	}
//}

//// 保存文章草稿
//// url: /api/article/save
//// method: post
//// header:
//// X-ACCESS-TOKEN : access-token
//// body: {"article_id":"","title":"","content":"","tags":[]}
//func (handler *ArticleEditorHandler) SaveArticleDraft() {
//	response := controller.NewResponse()
//	// 检查 ACCESS-TOKEN
//	token := handler.GetString(rqs.HeaderAccessToken, "")
//	if token == "" {
//		response.SetStatus(status.AccessDenied)
//		response.SetMessage("miss header:" + rqs.HeaderAccessToken)
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}

//	auth, err := storage.FindAuthByToken(token)
//	if err != nil {
//		response.SetStatus(status.AccessDenied)
//		response.SetMessage(status.Text(status.AccessDenied))
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}
//	if !auth.IsValid() {
//		response.SetStatus(status.AccessDenied)
//		response.SetMessage(status.Text(status.AccessDenied))
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}

//	// 解析请求 body 数据
//	body := handler.Ctx.Request.Body
//	defer body.Close()
//	bodyBytes, err := ioutil.ReadAll(body)
//	if err != nil {
//		response.SetStatus(status.UnprocessableEntity)
//		response.SetMessage("read body failed")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}

//	rqsBd := make(map[string]string)
//	err = json.Unmarshal(bodyBytes, &rqsBd)
//	if err != nil {
//		response.SetStatus(status.UnprocessableEntity)
//		response.SetMessage("parse body failed")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}

//	articleId := rqsBd[rqs.BodyArticleId]
//	articleTitle := rqsBd[rqs.BodyArticleTitle]
//	articleContent := rqsBd[rqs.BodyArticleContent]
//	// 如果 article_id 为空，添加一篇草稿文章
//	if articleId == "" {
//		articleDraft := &model.Article{UserId: auth.Uid, Title: articleTitle, Content: articleContent, Status: model.ArticleStatusDraft, Created: time.Now()}
//		err = storage.AddNewArticle(articleDraft)
//		if err != nil {
//			response.SetStatus(status.ServerInternalErr)
//			response.SetMessage("save article draft failed")
//			handler.Data["json"] = response
//			handler.ServeJSON()
//			return
//		} else {
//			response.SetStatus(status.OK)
//			response.SetMessage("save article draft success")
//			response.SetData(rsp.BodyArticleId, articleDraft.Id)
//			handler.Data["json"] = response
//			handler.ServeJSON()

//			// TODO 保存 tags
//			return
//		}
//	}

//	// 如果 article_id 不为空，查询 article_id 对应的文章
//	articleIdInt, err := strconv.ParseInt(articleId, 10, 64)
//	article, err := storage.GetArticleById(articleIdInt)
//	// 查询失败，return error
//	if err != nil {
//		response.SetStatus(status.SourceNotFound)
//		response.SetMessage("save failed! can't find this article")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}
//	// 查询成功，更新草稿文章
//	// 如果为文章为草稿类型，更新
//	if article.Status == model.ArticleStatusDraft {
//		article.Title = articleTitle
//		article.Content = articleContent

//		err = storage.UpdateArticle(article, "Title", "Content")
//		if err != nil {
//			response.SetStatus(status.ServerInternalErr)
//			response.SetMessage("save failed!")
//			handler.Data["json"] = response
//			handler.ServeJSON()
//			return
//		}
//		response.SetStatus(status.OK)
//		response.SetMessage("ok")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}
//	// 如果文章为发布类型，创建一个草稿文章
//	newArticle := &model.Article{UserId: auth.Uid, Title: articleTitle, Content: articleContent, Parent: articleIdInt}
//	err = storage.AddNewArticle(newArticle)
//	if err != nil {
//		response.SetStatus(status.ServerInternalErr)
//		response.SetMessage("save failed!")
//		handler.Data["json"] = response
//		handler.ServeJSON()
//		return
//	}
//	response.SetStatus(status.OK)
//	response.SetMessage("OK")
//	handler.Data["json"] = response
//	handler.ServeJSON()
//	return
//}

//// PublishArticle 发布文章
//// url: /api/article/publish
//// method: post
//// header:
//// X-ACCESS-TOKEN : access-token
//// body: {"article_id":"","title":"","content":"","tags":[]}
//func (handler *ArticleEditorHandler) PublishArticle() {
//	//	auth, ok := handler.checkAuth()
//	//	if !ok {
//	//		return
//	//	}

//	//	response := controller.NewResponse()
//	//	// 读取 body
//	//	body, err := handler.ReadJsonBody()
//	//	if err != nil {
//	//		response.SetStatus(status.UnprocessableEntity)
//	//		response.SetMessage("parse body failed")
//	//		handler.handleJsonResponse(response)
//	//		return
//	//	}

//	//	articleID := (body[rqs.BodyArticleId]).(int64)
//	//	title := (body[rqs.BodyArticleTitle]).(string)
//	//	content := (body[rqs.BodyArticleContent]).(string)

//	//	article, err := storage.GetArticleById(articleID)
//	//	if err != nil {
//	//		response.SetStatus(status.SourceNotFound)
//	//		response.SetMessage("no article with ID:" + strconv.FormatInt(articleID, 10))
//	//		handler.handleJsonResponse(response)
//	//		return
//	//	}

//	//	if article.UserId != auth.Uid {
//	//		response.SetStatus(status.AccessDenied)
//	//		response.SetMessage("access denied")
//	//		handler.handleJsonResponse(response)
//	//		return
//	//	}
//}

//func (handler *ArticleEditorHandler) handleJsonResponse(response controller.Response) {
//	handler.Data["json"] = response
//	handler.ServeJSON()
//}

//// checkAuth 检查是否合法的请求，X-ACCESS-TOKEN 有效
//func (handler *ArticleEditorHandler) checkAuth() (auth *model.Auth, ok bool) {
//	// 检查 ACCESS-TOKEN
//	token := handler.GetString(rqs.HeaderAccessToken, "")
//	if token == "" {
//		response := controller.NewResponse()
//		response.SetStatus(status.AccessDenied)
//		response.SetMessage("miss header:" + rqs.HeaderAccessToken)
//		handler.handleJsonResponse(response)
//		return
//	}

//	auth, err := storage.FindAuthByToken(token)
//	if err != nil {
//		response := controller.NewResponse()
//		response.SetStatus(status.AccessDenied)
//		response.SetMessage(status.Text(status.AccessDenied))
//		handler.handleJsonResponse(response)
//		return
//	}
//	if !auth.IsValid() {
//		response := controller.NewResponse()
//		response.SetStatus(status.AccessDenied)
//		response.SetMessage(status.Text(status.AccessDenied))
//		handler.handleJsonResponse(response)
//		return
//	}
//	return auth, true
//}
