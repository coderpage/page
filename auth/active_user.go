package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"page/constant/rqs"
	"page/constant/status"
	"page/controller"
	"page/model"
	"page/storage"
	"time"
)

type UserActiveHandler struct {
	controller.BaseController
}

// 激活用户
// 从邮箱链接中激活用户
// method: get
// url: http://host/path?tk=xxxx&redirect=xxxx
// 若 redirect 不为空，将激活结果重定向到 redirect 链接，后接参数 status & msg
// 若 redirect 为空，将激活结果通过 json 格式数据返回: {"status":"xxx","msg":"xxxx"}
// err:
// 10010 : token 不存在
// 10009 : token 过期
// 500 : 内部错误
// 200 : 激活成功
func (this *UserActiveHandler) ActiveFromEmail() {
	activeToken := this.GetString(rqs.UrlParamToken, "")
	redirect := this.GetString(rqs.UrlParamRedirect, "")

	response := controller.NewResponse()
	auth, err := storage.FindAuthByToken(activeToken)
	// 没有此 token
	if err != nil {
		if redirect != "" {
			redirect = fmt.Sprintf(redirect+"?status=%d&msg=%s", status.SourceNotFound, "token not found")
			this.Redirect(redirect, 302)
		} else {
			response.SetStatus(status.SourceNotFound)
			response.SetMessage("token not found")
			this.Data["json"] = response
			this.ServeJSON()
		}
		return
	}

	redirect = auth.Redirect
	// token 过期
	if time.Now().After(auth.ExpiryDate) {
		if redirect != "" {
			redirect = fmt.Sprintf(redirect+"?status=%d&msg=%s", status.TokenExpired, "token is expired")
			this.Redirect(redirect, 302)
		} else {
			response.SetStatus(status.TokenExpired)
			response.SetMessage("token was expired")
			this.Data["json"] = response
			this.ServeJSON()
		}
		return
	}

	uid := auth.Uid

	user := &model.User{Id: uid, Group: model.UserGroupCommon}

	err = storage.UpdateUser(user, "Group")
	if err != nil {
		if redirect != "" {
			redirect = fmt.Sprintf(redirect+"?status=%d&msg=%s", status.ServerInternalErr, "active failed")
			this.Redirect(redirect, 302)
		} else {
			response.SetStatus(status.ServerInternalErr)
			response.SetMessage("active failed")
			this.Data["json"] = response
			this.ServeJSON()
		}
		return
	}

	if redirect != "" {
		redirect = fmt.Sprintf(redirect+"?status=%d&msg=%s", status.OK, "user actived success")
		this.Redirect(redirect, 302)
	} else {
		response.SetStatus(status.OK)
		response.SetMessage("user actived success")
		this.Data["json"] = response
		this.ServeJSON()
	}

}

// ResendActivateEmail 重新发送激活邮件
// method : post
// url : /user/active/sendemail
// body : {"email":"xxx","redirect":"xxx"}
// err:
// 200 : OK
// 10001 : 请求 body 格式错误
// 10011 : 请求参数有错误
// 10005 : 用户不存在
func (this *UserActiveHandler) ResendActivateEmail() {
	resp := controller.NewResponse()

	body := this.Ctx.Request.Body
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		resp.SetStatus(status.UnprocessableEntity)
		resp.SetMessage("read body failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	rqsBd := make(map[string]string)
	err = json.Unmarshal(bodyBytes, &rqsBd)
	if err != nil {
		resp.SetStatus(status.UnprocessableEntity)
		resp.SetMessage("parse body failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	email := rqsBd[rqs.BodyEmail]
	redirect := rqsBd[rqs.BodyRedirect]
	if email == "" {
		resp.SetStatus(status.IllegalReqParam)
		resp.SetMessage("miss email")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	user, err := storage.FindUserByEmail(email)
	if err != nil {
		resp.SetStatus(status.UserNotExist)
		resp.SetMessage("user not exist")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	err = sendActiveEmail(user, redirect)
	if err != nil {
		resp.SetStatus(status.SendEmailFailed)
		resp.SetMessage("send activate user email failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp.SetStatus(status.OK)
	resp.SetMessage("send activate user email success")
	this.Data["json"] = resp
	this.ServeJSON()
}
