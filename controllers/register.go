package controllers

import (
	"bitcoin-Core/models/user"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get()  {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	var user user.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("用户解析失败")
		return
	}

	err = user.SaveUser()
	if err != nil {
		fmt.Println(err.Error())
		r.Data["err"] = fmt.Sprintf("注册失败! %s", err)
		r.Ctx.WriteString("注册失败")
		return
	}
	r.TplName= "index.html"
}
