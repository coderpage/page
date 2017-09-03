package api

import (
	"encoding/json"
	"io/ioutil"
	"page/constant/status"
	"page/controller"
	"page/model"
	"page/storage"
	"page/tool/secure"
	"strconv"
	"time"
)

type LoginHandler struct {
	controller.BaseController
}

// 登录接口
// url: /api/login
// method: post
// body: {"email":"xxx", "pwd":"xxx", "auth_duration":"x", "web":"xxx"}
// response:
// {"status":"", "msg":"xx", "user":{}, "auth_token":{}}
// error:
// 10001: 无法读取请求内容
// 10006: 用户未激活
// 10007: 用户名或密码错误
// 10011: 请求参数有误
// 500: 服务器出错
// -1: 未知错误
// @router /login [post]
func (this *LoginHandler) Login() {
	resp := this.NewResponse()

	body := this.Ctx.Request.Body
	defer body.Close()
	// 读取 body 内容
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		resp.SetStatus(status.UnprocessableEntity).SetMessage("read body failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	rspBd := make(map[string]string)
	err = json.Unmarshal(bodyBytes, &rspBd)
	if err != nil {
		resp.SetStatus(status.UnprocessableEntity).SetMessage("parse body failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	email := rspBd["email"]
	pwd := rspBd["pwd"]
	web := rspBd["web"]
	duration := rspBd["auth_duration"]

	durationInt, err := strconv.Atoi(duration)
	if err != nil {
		resp.SetStatus(status.IllegalReqParam)
		resp.SetMessage("auth_duration must type of int")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	if email == "" || pwd == "" {
		resp.SetStatus(status.IllegalReqParam)
		resp.SetMessage("miss email or password")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	user := &model.User{Email: email, Password: pwd}

	// 检查邮箱、密码
	err = storage.CheckEmailPwd(user)
	if err == nil {
		if user.Group == model.UserGroupNoActived {
			resp.SetStatus(status.UserNotActivated)
			resp.SetMessage(status.Text(status.UserNotActivated))
			this.Data["json"] = resp
			this.ServeJSON()
			return
		}

		token := secure.GenerateToken(32)
		expiry := time.Now().Add(time.Duration(durationInt) * time.Hour)
		auth := &model.Auth{Uid: user.Id, Token: token, Server: web, Status: "ok", Type: model.AuthTypeUserSignIn, ExpiryDate: expiry}
		_, err = storage.AddNewAuth(auth)
		if err != nil {
			resp.SetStatus(status.ServerInternalErr)
			resp.SetMessage("save token failed")
			this.Data["json"] = resp
			this.ServeJSON()
			return
		}

		resp.SetStatus(status.OK)
		resp.SetMessage("OK")
		resp.SetData("user", user)
		resp.SetData("auth_token", &model.AuthToken{Value: token, Expire: expiry})
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	if err == storage.ErrNoRows {
		resp.SetStatus(status.WrongUserNameOrPwd)
		resp.SetMessage("email or password is wrong")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp.SetStatus(status.UkownError)
	resp.SetMessage("login failed")
	this.Data["json"] = resp
	this.ServeJSON()
}
