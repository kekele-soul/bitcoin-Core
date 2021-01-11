package controllers

import "github.com/astaxie/beego"

type MemberController struct {
	beego.Controller
}

func (m *MemberController) Get() {
	m.TplName = "member.html"
}