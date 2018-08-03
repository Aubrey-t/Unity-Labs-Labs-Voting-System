package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/update`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/univote/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/univote/controllers:UserController"],
		beego.ControllerComments{
			Method: "ValidateToken",
			Router: `/validate_token`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
