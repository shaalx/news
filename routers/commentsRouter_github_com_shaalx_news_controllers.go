package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Home",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Hot",
			`/hot`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Latest",
			`/latest`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"NewsBefore",
			`/before/:date:string`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Themes",
			`/themes`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Theme",
			`/theme`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Sections",
			`/sections`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Section",
			`/section`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"BSection",
			`/bsection`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"News",
			`/news/:id:string`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"Story",
			`/story`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"LongCom",
			`/lcom`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsController"],
		beego.ControllerComments{
			"ShortCom",
			`/scom`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"VHome",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"VHot",
			`/hot`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"VLatest",
			`/latest`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"VNewsBefore",
			`/before/:date:string`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"Themes",
			`/themes`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"Theme",
			`/theme/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"Sections",
			`/sections`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"Section",
			`/section/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"BSection",
			`/bsection/:sid:int/:date:string`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"News",
			`/news/:id:string`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"Story",
			`/story/:id:string`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"LongCom",
			`/lcom/:id:string`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"] = append(beego.GlobalControllerRouter["github.com/toukii/news/controllers:NewsVController"],
		beego.ControllerComments{
			"ShortCom",
			`/scom/:id:string`,
			[]string{"get"},
			nil})

}
