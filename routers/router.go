// @APIVersion 0.0.1
// @Title 权限管理系统
// @Description 权限系统(SAM)
// @Contact twdp@gmail.com
// @TermsOfServiceUrl https://tianwei.pro
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"tianwei.pro/sam/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1/api",
		beego.NSNamespace("/portal",
			beego.NSInclude(
				&controllers.PortalController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
				),
		),
		beego.NSNamespace("/system",
			beego.NSInclude(
				&controllers.SystemController{},
				)),
	)
	beego.AddNamespace(ns)
}
