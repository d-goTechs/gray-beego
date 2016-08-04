package main

import (
	"github.com/astaxie/beego"
	"gray-beego/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	_ "gray-beego/routers"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/" +
		beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqlReg)
}

func main() {
	orm.Debug = true
	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")

	beego.SetStaticPath("css","static/css")
	beego.SetStaticPath("js","static/js")
	beego.SetStaticPath("font-awesome-4.1.0/css","static/font-awesome-4.1.0/css")
	beego.SetStaticPath("font-awesome-4.1.0/fonts","static/font-awesome-4.1.0/fonts")

	beego.Run()
}
