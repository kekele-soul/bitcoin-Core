package controllers

import (
	"bitcoin-Core/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get()  {
	l.TplName = "index.html"
}

func (l *LoginController) Post() {
	var user models.User

	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("用户解析失败")
	}

	err = user.QueryUser()
	if err != nil {
		l.Ctx.WriteString("用户信息查询失败，请检查账户信息！")
		return
	}

	l.TplName = "home.html"
}