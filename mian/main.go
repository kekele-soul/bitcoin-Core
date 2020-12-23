package main

import (
	_ "bitcoin-Core/routers"
	"github.com/astaxie/beego"
)

func main() {


	str := bitcoinServices.GetBC().GetBestBlockHash()
	fmt.Println(str)

	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/views", "views")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run()
}
