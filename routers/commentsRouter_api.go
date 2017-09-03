package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["page/api:LoginHandler"] = append(beego.GlobalControllerRouter["page/api:LoginHandler"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["page/api:RegisterHandler"] = append(beego.GlobalControllerRouter["page/api:RegisterHandler"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["page/api:RegisterHandler"] = append(beego.GlobalControllerRouter["page/api:RegisterHandler"],
		beego.ControllerComments{
			Method: "RegisterCode",
			Router: `/register/verifycode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
