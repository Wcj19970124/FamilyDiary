package routers

import (
	"../controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/login", &controllers.UserLoginController{})
}
