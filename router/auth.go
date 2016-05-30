package router

import (
	"page/auth"

	"github.com/astaxie/beego"
)

func registerAuths() {
	beego.Router("/register", &auth.SignUpHandler{}, "get:GetRegisterPage")
	beego.Router("/api/register", &auth.SignUpHandler{}, "post:Register")
	beego.Router("/active/byemail", &auth.UserActiveHandler{}, "get:ActiveFromEmail")
	beego.Router("api/login", &auth.LoginHandler{}, "post:Login")
	beego.Router("/login", &auth.LoginHandler{}, "get:GetLoginPage")
}

//beego.Router("/uauth/signup", &controllers.SignUpHandler{}, "post:SignUp")
//	beego.Router("/uauth/signin", &controllers.SignInHandler{}, "post:SignIn")
//	beego.Router("/uauth/user/active", &controllers.UserActiveHandler{}, "get:ActiveFromEmail")
//	beego.Router("/uauth/user/active/sendemail", &controllers.UserActiveHandler{}, "post:ResendActivateEmail")
//	beego.Router("/uauth/find/user/withtk", &controllers.UserDataHandler{}, "post:FindUserWithAuthToken")
//	beego.Router("/uauth/user/fpwd/email", &controllers.ResetPwdHandler{}, "post:FindPwdByEmail")
//	beego.Router("/uauth/user/fpwd/email", &controllers.ResetPwdHandler{}, "get:AuthResetAction")
//	beego.Router("/uauth/user/resetpwd", &controllers.ResetPwdHandler{}, "post:ResetPwd")
