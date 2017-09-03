package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"page/constant/status"
	"page/controller"
	"page/model"
	"page/storage"
	"page/tool/mail"
	"page/tool/secure"
	"time"
)

type RegisterHandler struct {
	controller.BaseController
}

// 请求注册验证码
// body : {"email":""}
// response : {"status":xx, "message":"", "register_id":"xxx"}
// @router /register/verifycode [post]
func (this *RegisterHandler) RegisterCode() {

	resp := this.NewResponse()
	// 读取 request body
	body := this.Ctx.Request.Body
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		// 读取 request body 出错，err: 10001
		resp.SetStatus(status.UnprocessableEntity).SetMessage("read body failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	// 解析请求 body
	rqsBd := make(map[string]string)
	err = json.Unmarshal(bodyBytes, &rqsBd)
	if err != nil {
		// 解析 request body 出错，err: 10001
		resp.SetStatus(status.UnprocessableEntity)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	email := rqsBd["email"]
	if email == "" {
		resp.SetStatus(status.IllegalReqParam).SetMessage("miss params: email")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	userExist := storage.IsUserExistByEmail(email)
	if userExist {
		resp.SetStatus(status.UserExist).SetMessage(status.Text(status.UserExist))
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	activeCode := secure.GenerateDigitActiveCode(6)
	err = sendVerifyCodeByEmail(email, activeCode)
	if err != nil {
		resp.SetStatus(status.SendEmailFailed).SetMessage("send email failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp.SetStatus(status.OK).SetMessage(status.Text(status.OK))
	this.Data["json"] = resp
	this.ServeJSON()
}

// sendActiveEmail 发送激活邮件
// 注册码
// code  验证码
func sendVerifyCodeByEmail(email string, verifyCode string) (err error) {

	expire := time.Now().Add(30 * time.Minute)
	auth := &model.Auth{Key: email, Token: verifyCode, Type: model.AuthTypeUserActive, ExpiryDate: expire}

	_, err = storage.AddNewAuth(auth)
	if err != nil {
		return errors.New("Add Auth Err:" + err.Error())
	}

	mailer, err := mail.NewServiceMailer()
	if err != nil {
		return err
	}

	body := fmt.Sprintf(`%s，您好！
<br>
<br>
为确保是您本人操作，您已选择通过该邮件地址获取验证码验证身份。请在邮件验证码输入框输入下方验证码：
<br>
<br>
<font color="red" size="5">%s</font>
<br>
<br>
勿向任何人泄露您收到的验证码。验证码会在邮件发送30分钟后失效。
<br>
<br>
Mine.Page 帐号
<br>
<br>
`, email, verifyCode)
	err = mailer.SendMail(email, "Mine.Page", "Mine 账号邮件验证码", "html", body)

	return err
}

// 注册接口
// url: /api/register
// method: post
// body : {"email":"xxx", "pwd":"xxx", "verify_code":""}
// err:
// 10001: 读取请求实体出错
// 10004: 用户已存在
// 10011: 请求参数有错
// -1: 未知错误
// 200: OK
// @router /register [post]
func (this *RegisterHandler) Register() {
	resp := this.NewResponse()

	// 读取 request body
	body := this.Ctx.Request.Body
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		// 读取 request body 出错，err: 10001
		resp.SetStatus(status.UnprocessableEntity).SetMessage("Read Body Failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	rqsBd := make(map[string]string)
	err = json.Unmarshal(bodyBytes, &rqsBd)
	if err != nil {
		// 解析 request body 出错，err: 10001
		resp.SetStatus(status.UnprocessableEntity)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	email := rqsBd["email"]
	pwd := rqsBd["pwd"]
	verifyCode := rqsBd["verify_code"]
	if email == "" || pwd == "" || verifyCode == "" {
		resp.SetStatus(status.IllegalReqParam)
		resp.SetMessage("miss params: email,password,verify_code ")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	isUserExist := storage.IsUserExistByEmail(email)
	if isUserExist {
		resp.SetStatus(status.UserExist).SetMessage(status.Text(status.UserExist))
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	// check verify code
	auth, err := storage.FindAuthByKeyLatest(email)
	if err != nil {
		resp.SetStatus(status.SourceNotFound).SetMessage("code not found")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	fmt.Println("now time ", time.Now().Unix())
	fmt.Println("auth time ", auth.ExpiryDate.Unix())
	if time.Now().After(auth.ExpiryDate) {
		resp.SetStatus(status.TokenExpired).SetMessage("code was expired")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	if auth.Token != verifyCode {
		resp.SetStatus(status.TokenIncorrect)
		resp.SetMessage("code incorrect")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	user := &model.User{}
	user.Email = email
	user.Password = pwd
	user.Group = model.UserGroupCommon
	err = storage.CreateUser(user)
	if err != nil {
		resp.SetStatus(status.UkownError)
		resp.SetMessage("create user failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp.SetStatus(status.OK)
	resp.SetMessage(status.Text(status.OK))
	resp.SetData("user", user)
	this.Data["json"] = resp
	this.ServeJSON()
}

// sendActiveEmail 发送激活邮件
// redirectUrl  验证后重定向的地址
// func sendActiveEmail(user *model.User, redirectUrl string) (err error) {
//
// 	activeToken := secure.GenerateToken(32)
// 	expire := time.Now().Add(24 * time.Hour)
// 	auth := &model.Auth{Uid: user.Id, Token: activeToken, Type: model.AuthTypeUserActive, Redirect: redirectUrl, ExpiryDate: expire}
//
// 	_, err = storage.AddNewAuth(auth)
// 	if err != nil {
// 		return errors.New("Add Auth Err:" + err.Error())
// 	}
//
// 	mailer, err := mail.NewServiceMailer()
// 	if err != nil {
// 		return err
// 	}
//
// 	activeUrl := fmt.Sprintf("%s/active/byemail?%s=%s&%s=%s", conf.ServerBaseURL, rqs.UrlParamToken, activeToken, rqs.UrlParamRedirect, redirectUrl)
// 	body := fmt.Sprintf(`	尊敬的 %s 您好！
// <br>
// 点击 <a href="%s">链接</a> 可激活您的的账号！
// <br>
// 为保障您的帐号安全，请在24小时内点击该链接，您也可以将链接复制到浏览器地址栏访问。如果您并未尝试激活邮箱，请忽略本邮件，由此给您带来的不便请谅解。
// <br>
// <br>
// 本邮件由系统自动发出，请勿直接回复！
// <br>
// <br>
// `, user.Email, activeUrl)
// 	err = mailer.SendMail(user.Email, "Page", "请激活账号", "html", body)
//
// 	return err
// }
