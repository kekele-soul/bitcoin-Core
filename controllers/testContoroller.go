package controllers

import (
	"bitcoin-Core/bitcoinServices"
	"github.com/astaxie/beego"
)

type Controllers struct {
	beego.Controller
}

func (c Controllers) Get()  {
	c.TplName = "getBlockCount.html"
}

func (c Controllers) Post()  {
	c.TplName = "getBlockCount.html"
	flo := bitcoinServices.GetBC().GetBlockCount
	c.Data["flo"] = flo

}
