package controllers

import "github.com/astaxie/beego"

type Emil_indexControllers struct {
	beego.Controller
}

func (e *Emil_indexControllers) Get()  {
	e.TplName = "emil_index.html"
}