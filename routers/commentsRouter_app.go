package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["page/app:HomeHandler"] = append(beego.GlobalControllerRouter["page/app:HomeHandler"],
		beego.ControllerComments{
			Method: "GetHomePage",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["page/app:LoginHandler"] = append(beego.GlobalControllerRouter["page/app:LoginHandler"],
		beego.ControllerComments{
			Method: "GetLoginPage",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["page/app:RegisterHandler"] = append(beego.GlobalControllerRouter["page/app:RegisterHandler"],
		beego.ControllerComments{
			Method: "GetRegisterPage",
			Router: `/register`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
