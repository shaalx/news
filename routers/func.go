package routers

import (
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	beego.AddFuncMap("Newsid", Newsid)
	beego.AddFuncMap("TimeStamp", TimeStamp)
}

func Newsid(idi interface{}) string {
	id := idi.(float64)
	iid := int64(id)
	return fmt.Sprintf("%v", iid)
}

func TimeStamp(date string) string {
	beego.Debug(date)
	return "1398780001"
}
