package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:project`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:pid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_cmdb/controllers:InfoIniController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
