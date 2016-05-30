package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"page/conf"
	"page/constant/rqs"
	"page/constant/rsp"
	"page/constant/status"
	"page/controller"
	"page/model"
	"page/storage"
	"page/tool/mail"
	"page/tool/secure"
	"time"
)

type SignUpHandler struct {
	controller.BaseController
}

func (this *SignUpHandler) GetRegisterPage() {
	this.Data[rsp.PageTitle] = "注册"
	this.TplName = "register.html"
}

// 注册接口
// url: /api/register
// method: post
// body : {"email":"xxx", "pwd":"xxx", "redirect":"xxxx"}
// err:
// 10001: 读取请求实体出错
// 10004: 用户已存在
// 10011: 请求参数有错
// -1: 未知错误
// 200: OK
func (this *SignUpHandler) Register() {
	resp := controller.NewResponse()

	// 读取 request body
	body := this.Ctx.Request.Body
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		// 读取 request body 出错，err: 10001
		resp.SetStatus(status.UnprocessableEntity)
		resp.SetMessage("Read Body Failed")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	user := &model.User{}
	rqsBd := make(map[string]string)
	err = json.Unmarshal(bodyBytes, &rqsBd)
	if err != nil {
		// 解析 request body 出错，err: 10001
		resp.SetStatus(status.UnprocessableEntity)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	user.Email = rqsBd[rqs.BodyEmail]
	user.Password = rqsBd[rqs.BodyPassword]
	// 注册激活回调链接
	activeRedirect := rqsBd[rqs.BodyRedirect]
	if user.Email == "" || user.Password == "" {
		resp.SetStatus(status.IllegalReqParam)
		resp.SetMessage("email or  password must not be empty")
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	err = storage.CreateUser(user)
	if err == nil {
		resp.SetStatus(status.OK)
		resp.SetMessage(status.Text(status.OK))
		this.Data["json"] = resp
		this.ServeJSON()

		sendActiveEmail(user, activeRedirect)
		return
	}

	if err == storage.ErrRowExist {
		resp.SetStatus(status.UserExist)
		resp.SetMessage(status.Text(status.UserExist))
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp.SetStatus(status.UkownError)
	resp.SetMessage("create user failed")
	this.Data["json"] = resp
	this.ServeJSON()
}

// sendActiveEmail 发送激活邮件
// redirectUrl  验证后重定向的地址
func sendActiveEmail(user *model.User, redirectUrl string) (err error) {

	activeToken := secure.GenerateToken(32)
	expire := time.Now().Add(24 * time.Hour)
	auth := &model.Auth{Uid: user.Id, Token: activeToken, Type: model.AuthTypeUserActive, Redirect: redirectUrl, ExpiryDate: expire}

	_, err = storage.AddNewAuth(auth)
	if err != nil {
		return errors.New("Add Auth Err:" + err.Error())
	}

	mailer, err := mail.NewServiceMailer()
	if err != nil {
		return err
	}

	activeUrl := fmt.Sprintf("%s/active/byemail?%s=%s&%s=%s", conf.ServerBaseURL, rqs.UrlParamToken, activeToken, rqs.UrlParamRedirect, redirectUrl)
	body := fmt.Sprintf(`	尊敬的 %s 您好！
<br>
点击 <a href="%s">链接</a> 可激活您的的账号！
<br>
为保障您的帐号安全，请在24小时内点击该链接，您也可以将链接复制到浏览器地址栏访问。如果您并未尝试激活邮箱，请忽略本邮件，由此给您带来的不便请谅解。
<br>
<br>
本邮件由系统自动发出，请勿直接回复！
<br>
<br>
`, user.Email, activeUrl)
	err = mailer.SendMail(user.Email, "Page", "请激活账号", "html", body)

	return err
}
