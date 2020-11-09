package controllers

import (
	"time"

	"errors"

	"../models"
)

//UserLogoutController 登出数据体
type UserLogoutController struct {
	BaseController
}

//用户登出操作错误状态码
const (
	UserLogoutErr = 30010 //用户登出失败
)

//退出登陆
//将redis中的用户登陆信息置空
func (l *UserLogoutController) logout() error {

	key := "login_admin_username"
	_, err := models.DelKey(key)
	if err != nil {
		l.Report(l.Ctx.Input.IP(), "1", "POST", "1", "logout", models.GetLoginAdminUserName(), "删除登陆用户缓存信息失败!", time.Now())
		l.OutPut(UserLogoutErr, "用户登出失败")
		return errors.New("用户登出失败")
	}

	return nil
}

//Post Post请求入口
func (l *UserLogoutController) Post() {

	err := l.logout()
	if err != nil {
		return
	}

	l.Report(l.Ctx.Input.IP(), "1", "POST", "0", "logout", models.GetLoginAdminUserName(), "登出成功!!", time.Now())
	l.OutPut(200, "登出成功!")
}
