package main

import (
	"github.com/astaxie/beego"
	"github.com/shaalx/news/controllers"
	_ "github.com/shaalx/news/routers"
)

func main() {
	ns := beego.NewNamespace("/js",
		beego.NSInclude(&controllers.NewsController{}),
	)
	beego.AddNamespace(ns)
	app := beego.Include(&controllers.NewsVController{})
	app.Run()
}
