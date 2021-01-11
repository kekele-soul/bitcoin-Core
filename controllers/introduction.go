package controllers

import (
	"github.com/astaxie/beego"
)

type IntroductionController struct {
	beego.Controller
}

func (i *IntroductionController) Get() {
	i.TplName = "introduction.html"
}
