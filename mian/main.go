package main

import (
	"bitcoin-Core/bitcoinServices"
	_ "bitcoin-Core/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {



	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
