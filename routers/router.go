package routers

import (
	"page/api"
	"page/app"

	"github.com/astaxie/beego"
)

func init() {
	apiNameSpace := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&api.RegisterHandler{},
				&api.LoginHandler{})))

	beego.Include(&app.HomeHandler{}, &app.LoginHandler{}, &app.RegisterHandler{})

	beego.AddNamespace(apiNameSpace)
}

// Register des
func Register() {
	// beego.Router("/api/user/register/verfiycode", &api.RegisterHandler{}, "post:RegisterCode")
	// beego.Router("/api/user/register", &api.RegisterHandler{}, "post:Register")
	//
	// beego.Router("/user/register", &app.RegisterHandler{}, "get:GetRegisterPage")

}

// import (
// 	"page/auth"
//
// 	"github.com/astaxie/beego"
// )

// func registerAuths() {
// 	beego.Router("/register", &auth.SignUpHandler{}, "get:GetRegisterPage")
// 	beego.Router("/api/register", &auth.SignUpHandler{}, "post:Register")
// 	beego.Router("/active/byemail", &auth.UserActiveHandler{}, "get:ActiveFromEmail")
// 	beego.Router("api/login", &auth.LoginHandler{}, "post:Login")
// 	beego.Router("/login", &auth.LoginHandler{}, "get:GetLoginPage")
// }

// Register 注册路由
// func Register() {
// registerAuths()
// registerCommons()
//	beego.Router("/", &controller.HomeHandler{}, "get:HomePage")
//	beego.Router("/login", &controller.LoginHandler{}, "get:Login")
//	beego.Router("/login", &controller.LoginHandler{}, "post:LoginByEmail")
//	beego.Router("/write", &controller.EditorHandler{}, "get:EditorPage")
//	beego.Router("/write/save", &controller.EditorHandler{}, "post:EditorSave")
//	beego.Router("/write/publish", &controller.EditorHandler{}, "post:EditorPublish")

// }
