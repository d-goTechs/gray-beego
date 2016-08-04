// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gray-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//http://127.0.0.1:9001/v1/grey_configs
		beego.NSNamespace("/grey_configs",
			beego.NSInclude(
				&controllers.GreyConfigsController{},
			),
		),
		//http://127.0.0.1:9001/v1/grey_servers
		beego.NSNamespace("/grey_servers",
			beego.NSInclude(
				&controllers.GreyServersController{},
			),
		),
		//http://127.0.0.1:9001/v1/sub_systems
		beego.NSNamespace("/sub_systems",
			beego.NSInclude(
				&controllers.SubSystemsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
