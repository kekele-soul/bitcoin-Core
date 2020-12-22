package main

import (
	"github.com/astaxie/beego"
)

func main() {
	//dbMysql.ConnectDB()
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
