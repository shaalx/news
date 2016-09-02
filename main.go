package main

import (
	"github.com/astaxie/beego"
	"github.com/toukii/news/controllers"
	_ "github.com/toukii/news/routers"
)

func main() {
	ns := beego.NewNamespace("/js",
		beego.NSInclude(&controllers.NewsController{}),
	)
	beego.AddNamespace(ns)
	app := beego.Include(&controllers.NewsVController{})
	app.Run()
}
