package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	. "github.com/shaalx/news/models"

	"fmt"
	"html/template"
)

type NewsController struct {
	beego.Controller
}

// @router / [get]
func (c *NewsController) Home() {
	c.Data["title"] = "Home"
	c.TplNames = "jsindex.html"
}

// @router /hot [get]
func (c *NewsController) Hot() {
	hot_url := "http://news-at.zhihu.com/api/3/news/hot"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /latest [get]
func (c *NewsController) Latest() {
	hot_url := "http://news-at.zhihu.com/api/3/news/latest"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /before/:date:string [get]
func (c *NewsController) NewsBefore() {
	date := c.Ctx.Input.Param(":date")
	hot_url := fmt.Sprintf("http://news.at.zhihu.com/api/4/news/before/%s", date)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /themes [get]
func (c *NewsController) Themes() {
	hot_url := "http://news-at.zhihu.com/api/4/themes"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /theme [get]
func (c *NewsController) Theme() {
	hot_url := "http://news-at.zhihu.com/api/4/theme/11"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /sections [get]
func (c *NewsController) Sections() {
	hot_url := "http://news-at.zhihu.com/api/3/sections"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /section [get]
func (c *NewsController) Section() {
	hot_url := "http://news-at.zhihu.com/api/3/section/1"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /bsection [get]
func (c *NewsController) BSection() {
	hot_url := "http://news-at.zhihu.com/api/3/section/1/before/1398780001"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /news/:id:string [get]
func (c *NewsController) News() {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/news/%s", id)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		c.Abort("400")
	}
	m, ok := v.(map[string]interface{})
	if ok {
		vv := m["body"]
		c.Data["content"] = template.HTML(fmt.Sprintf("%v", vv))
		c.TplNames = "news.html"
		return
	}
	c.Data["title"] = "News"
	c.Data["content"] = v
	c.TplNames = "news.html"
}

// @router /story [get]
func (c *NewsController) Story() {
	hot_url := "http://news-at.zhihu.com/api/4/story-extra/3876052"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /lcom [get]
func (c *NewsController) LongCom() {
	hot_url := "http://news-at.zhihu.com/api/4/story/4232852/long-comments"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /scom [get]
func (c *NewsController) ShortCom() {
	hot_url := "http://news-at.zhihu.com/api/4/story/4232852/short-comments"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}
