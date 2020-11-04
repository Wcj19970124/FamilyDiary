package routers

import (
	"../controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/user/login", &controllers.UserLoginController{})

	//用户管理
	beego.Router("/add/user", &controllers.UserController{}, "post:AddUser")
	beego.Router("/del/user", &controllers.UserController{}, "delete:DelUser")
	beego.Router("/update/user", &controllers.UserController{}, "put:UpdateUser")
}
