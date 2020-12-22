package controllers

import (
	"bitcoin-Core/models"
	"fmt"
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
		return
	}



	b,err := user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("用户信息查询失败，请检查账户信息！")
		return
	}
	if b == false {
		l.Ctx.WriteString("用户或密码错误请重试！")
		return
	}
	if b== true {
		l.TplName = "home.html"

	}


}