package controller

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"page/constant/rqs"
	"page/constant/status"
	"page/model"
	"page/storage"
)

type BaseController struct {
	beego.Controller
}

// ReadJsonBody 读取 json 格式的 body 数据
func (controller *BaseController) ReadJsonBody() (body map[string]interface{}, err error) {
	httpBody := controller.Ctx.Request.Body
	defer httpBody.Close()
	bodyBytes, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return
	}
	fmt.Printf("body:", string(bodyBytes))
	body = make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &body)
	return
}

func (handler *BaseController) HandleJsonResponse(response Response) {
	handler.Data["json"] = response
	handler.ServeJSON()
}

// CheckAuth 检查是否合法的请求，
// 读取 http 请求头信息中的 X-ACCESS-TOKEN，查询 token 是否有效
// 若验证失败，http 请求直接返回错误
func (handler *BaseController) CheckAuth() (auth *model.Auth, ok bool) {
	// 检查 ACCESS-TOKEN
	token := handler.Ctx.Input.Header(rqs.HeaderAccessToken)
	if token == "" {
		response := NewResponse()
		response.SetStatus(status.AccessDenied)
		response.SetMessage("miss header:" + rqs.HeaderAccessToken)
		handler.HandleJsonResponse(response)
		return
	}

	auth, err := storage.FindAuthByToken(token)

	if err != nil {
		response := NewResponse()
		response.SetStatus(status.AccessDenied)
		if err == storage.ErrNoRows {
			response.SetMessage("token is not exist")
		} else {
			response.SetMessage(status.Text(status.AccessDenied))
		}
		handler.HandleJsonResponse(response)
		return
	}
	if !auth.IsValid() {
		response := NewResponse()
		response.SetStatus(status.AccessDenied)
		response.SetMessage("token expired")
		handler.HandleJsonResponse(response)
		return
	}
	return auth, true
}
