package controllers

import (
	"bitcoin-Core/models"
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
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("用户解析失败")
		return
	}

	err = user.SaveUser()
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("注册失败")
		return
	}
	r.TplName= "index.html"
}
