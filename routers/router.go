// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
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
	)
	beego.AddNamespace(ns)
}
