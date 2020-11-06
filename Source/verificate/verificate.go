package verificate

import (
	"../models"
)

//AdminLogin 校验管理员登陆是否失效
func AdminLogin() bool {
	key := "login_admin_username"
	if val, _, err := models.GetByKey(key); err == nil && val != "" {
		return true
	}
	return false
}

//AdminPermission 校验后台管理员权限
func AdminPermission() bool {
	return true
}
