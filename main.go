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


	//
	//
	//beego.Get("/get", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Eric get"))
	//})
	//beego.Post("/post", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Eric post"))
	//})
	//beego.Put("/put", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Eric put"))
	//})
	//beego.Head("/head", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Eric head"))
	//})
	//beego.Options("/option", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Eric option"))
	//})
	//beego.Delete("/delete", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Eric delete"))
	//})
	//beego.Any("/any", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Eric any"))
	//})

	beego.Run()
}
