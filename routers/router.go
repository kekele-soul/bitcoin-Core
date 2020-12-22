package routers

import (
	"bitcoin-Core/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //登录页面接口
    beego.Router("/index",&controllers.LoginController{})
    //用于注册页面，注册后跳转登录页面
    beego.Router("/register",&controllers.RegisterController{})
}
