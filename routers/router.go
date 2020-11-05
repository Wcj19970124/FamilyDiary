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
	beego.Router("/get/user", &controllers.UserController{}, "get:QueryUser")
	beego.Router("/get/users", &controllers.UserController{}, "post:QueryUsers")

	//角色管理
	beego.Router("/add/role", &controllers.RoleController{}, "post:AddRole")
	beego.Router("/del/role", &controllers.RoleController{}, "delete:DelRole")
	beego.Router("/update/role", &controllers.RoleController{}, "put:UpdateRole")
	beego.Router("/get/role", &controllers.RoleController{}, "get:QueryRole")
	beego.Router("/get/roles", &controllers.RoleController{}, "post:QueryRoles")

	//权限管理
	beego.Router("/add/permission", &controllers.PermissionController{}, "post:AddPermission")
	beego.Router("/del/permission", &controllers.PermissionController{}, "delete:DelPermission")
	beego.Router("/update/permission", &controllers.PermissionController{}, "put:UpdatePermission")
	beego.Router("/get/permission", &controllers.PermissionController{}, "Get:QueryPermission")
	beego.Router("/get/permissions", &controllers.PermissionController{}, "post:QueryPermissions")
	beego.Router("/get/permissionsTree", &controllers.PermissionController{}, "get:QueryPermissionsTree")
}
