package routers

import (
	"bitcoin-Core/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//登录页面接口
	beego.Router("/index", &controllers.LoginController{})
	//邮箱登录接口
	beego.Router("/emil_index", &controllers.Emil_indexControllers{})
	//用于注册页面，注册后跳转登录页面
	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/bitcoin", &controllers.RpcCommunication{})

	beego.Router("/introduction", &controllers.IntroductionController{})

	beego.Router("/home", &controllers.HomeController{})

	beego.Router("/about_us", &controllers.About_usControllers{})

	beego.Router("/member", &controllers.MemberController{})
}
