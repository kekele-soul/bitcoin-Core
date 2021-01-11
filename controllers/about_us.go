package controllers

import "github.com/astaxie/beego"

type About_usControllers struct {
	beego.Controller
}

func (a *About_usControllers) Get() {
	a.TplName = "about_us.html"
}
