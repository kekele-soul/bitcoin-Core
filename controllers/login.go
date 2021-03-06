package controllers

import (
	"bitcoin-Core/models/user"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "index.html"
}

func (l *LoginController) Post() {
	var user user.User

	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("用户解析失败")
		return
	}

	if user.UserName != "" {
		b, err := user.QueryUserByUserName()
		if err != nil {
			fmt.Println(err.Error())
			l.Ctx.WriteString("用户信息查询失败，请检查账户信息！")
			return
		}
		if !b {
			l.Ctx.WriteString("用户或密码错误请重试！")
			return
		}
		l.TplName = "home.html"
		return
	}
	if user.MailBox != "" {
		b, err := user.QueryUserByEmail()
		if err != nil {
			fmt.Println(err.Error())
			l.Ctx.WriteString("用户信息查询失败，请检查账户信息！")
			return
		}
		if !b {
			l.Ctx.WriteString("邮箱或密码错误请重试！")
			return
		}
		l.TplName = "home.html"
		return
	}
}