package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyConfigsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:GreyServersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"] = append(beego.GlobalControllerRouter["gray-beego/controllers:SubSystemsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
