package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	. "github.com/toukii/news/models"

	"fmt"
	"html/template"
)

type NewsVController struct {
	beego.Controller
}

// @router / [get]
func (c *NewsVController) VHome() {
	c.Data["title"] = "Home"
	c.TplNames = "index.html"
}

// @router /hot [get]
func (c *NewsVController) VHot() {
	hot_url := "http://news-at.zhihu.com/api/3/news/hot"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["title"] = "Hot"
	c.Data["content"] = v
	c.TplNames = "hot.html"
}

// @router /latest [get]
func (c *NewsVController) VLatest() {
	hot_url := "http://news-at.zhihu.com/api/3/news/latest"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["title"] = "Latest"
	c.Data["content"] = v
	c.TplNames = "latest.html"
}

// @router /before/:date:string [get]
func (c *NewsVController) VNewsBefore() {
	date := c.Ctx.Input.Param(":date")
	hot_url := fmt.Sprintf("http://news.at.zhihu.com/api/4/news/before/%s", date)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["title"] = date

	c.Data["content"] = v
	c.TplNames = "latest.html"
}

// @router /themes [get]
func (c *NewsVController) Themes() {
	hot_url := "http://news-at.zhihu.com/api/4/themes"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["content"] = v
	c.Data["title"] = "Themes"
	c.TplNames = "themes.html"
}

// @router /theme/:id:int [get]
func (c *NewsVController) Theme() {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/theme/%s", id)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["title"] = id
	c.Data["content"] = v
	c.TplNames = "theme.html"
}

// @router /sections [get]
func (c *NewsVController) Sections() {
	hot_url := "http://news-at.zhihu.com/api/3/sections"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["title"] = "Sections"
	c.Data["content"] = v
	c.TplNames = "sections.html"
}

// @router /section/:id:int [get]
func (c *NewsVController) Section() {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/3/section/%s", id)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["title"] = id
	c.Data["sid"] = id
	c.Data["content"] = v
	c.TplNames = "section.html"
}

// @router /bsection/:sid:int/:date:string [get]
func (c *NewsVController) BSection() {
	sid := c.Ctx.Input.Param(":sid")
	date := c.Ctx.Input.Param(":date")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/section/%s/before/%s", sid, date)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["title"] = sid
	c.Data["sid"] = sid
	c.Data["content"] = v
	c.TplNames = "section.html"
}

// @router /news/:id:string [get]
func (c *NewsVController) News() {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/news/%s", id)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		c.Abort("400")
	}
	c.Data["title"] = "News"
	c.Data["id"] = id
	m, ok := v.(map[string]interface{})
	if ok {
		vv := m["body"]
		c.Data["content"] = template.HTML(fmt.Sprintf("%v", vv))
		c.Data["story"] = c.story(id)
		c.TplNames = "news.html"
		return
	}
	c.Data["title"] = "News"
	c.Data["content"] = v
	c.TplNames = "news.html"
}

// @router /story/:id:string [get]
func (c *NewsVController) Story() {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/story-extra/%s", id)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /lcom/:id:string [get]
func (c *NewsVController) LongCom() {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/story/%s/long-comments", id)
	// hot_url := "http://news-at.zhihu.com/api/4/story/4232852/long-comments"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

// @router /scom/:id:string [get]
func (c *NewsVController) ShortCom() {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/story/%s/short-comments", id)
	// hot_url := "http://news-at.zhihu.com/api/4/story/4232852/short-comments"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return
	}
	c.Data["json"] = v
	c.ServeJson()
}

func (c *NewsVController) story(id string) interface{} {
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/story-extra/%s", id)
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return nil
	}
	return v
}

func (c *NewsVController) longCom() interface{} {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/story/%s/long-comments", id)
	// hot_url := "http://news-at.zhihu.com/api/4/story/4232852/long-comments"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return nil
	}
	return v
}

func (c *NewsVController) shortCom() interface{} {
	id := c.Ctx.Input.Param(":id")
	hot_url := fmt.Sprintf("http://news-at.zhihu.com/api/4/story/%s/short-comments", id)
	// hot_url := "http://news-at.zhihu.com/api/4/story/4232852/short-comments"
	resp := httplib.Get(hot_url)

	var v interface{}
	err := resp.ToJson(&v)
	if CheckErr(err) {
		return nil
	}
	return v
}
