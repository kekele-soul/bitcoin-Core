package main

import (
	"bitcoin-Core/bitcoinServices"
	_ "bitcoin-Core/routers"
	"fmt"
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
