package controllers

import (
	"encoding/json"
	"errors"
	"time"

	"../models"
	"../util"
	"github.com/astaxie/beego/logs"
)

//UserLoginController 登陆数据体
type UserLoginController struct {
	BaseController
	_username string
	_password string
}

//Admin 用于接收请求参数的结构体
type Admin struct {
	Username string
	Password string
}

//登陆错误状态码(ret)
const (
	UserNotLogin        = 10001 //用户未登录
	UserLoginParamNull  = 10002 //参数为空
	UserLoginParamError = 10003 //参数错误
	UserNotExist        = 10004 //用户不存在
)

var userLoginRedisCacheSeconds = 3600 //登陆缓存时长-1h

//参数校验
func (l *UserLoginController) validParams() error {

	var admin Admin
	json.Unmarshal(l.Ctx.Input.RequestBody, &admin)

	if username := admin.Username; username != "" {
		l._username = username
	} else {
		l.Report(l.Ctx.Input.IP(), "0", "POST", "1", "validParams", models.GetLoginAdminUserName(), "参数为空(username)", time.Now())
		l.OutPut(UserLoginParamNull, "参数为空(username)")
		return errors.New("参数为空")
	}

	if password := admin.Password; password != "" {
		l._password = password
	} else {
		l.Report(l.Ctx.Input.IP(), "0", "POST", "1", "validParams", models.GetLoginAdminUserName(), "参数为空(password)", time.Now())
		l.OutPut(UserLoginParamNull, "参数为空(password)")
		return errors.New("参数为空")
	}

	logs.Debug("login username:" + l._username + ",password:" + l._password)
	return nil
}

//登陆校验:超管同一时间只能有一人登陆
//第一步：先检查缓存中是否存在登陆信息,有则验证是否是本人,是则不用重复登陆,不是则重新登陆
//第二步：若缓存没有,查询数据库,若数据库没有,直接返回
//第三步：若数据库存在,则校验数据,正确则写缓存,不正确直接返回
func (l *UserLoginController) checkLogin() error {

	loginUserNameKey := "login_admin_username"
	logs.Debug("---- user login cache redis key1:" + loginUserNameKey)
	if val, _, err := models.GetByKey(loginUserNameKey); err == nil && val != "" {
		if val != l._username {
		} else {
			l.Report(l.Ctx.Input.IP(), "0", "POST", "0", "checkLogin", models.GetLoginAdminUserName(), "已登陆！", time.Now())
			l.OutPut(200, "已登陆！")
			return errors.New("已登录")
		}
	}

	password, _ := models.QueryPwdByUserName(l._username)
	if password == "" {
		l.Report(l.Ctx.Input.IP(), "0", "POST", "1", "checkLogin", models.GetLoginAdminUserName(), "用户不存在", time.Now())
		l.OutPut(UserNotExist, "用户不存在!")
		return errors.New("用户不存在！")
	}
	if password != util.MD5(l._password) {
		l.Report(l.Ctx.Input.IP(), "0", "POST", "1", "checkLogin", models.GetLoginAdminUserName(), "密码错误", time.Now())
		l.OutPut(UserLoginParamError, "密码错误!")
		return errors.New("密码错误")
	}

	models.SetByKey(loginUserNameKey, l._username, userLoginRedisCacheSeconds)

	return nil
}

//Post 请求入口
func (l *UserLoginController) Post() {

	logs.Debug("----- 进入controller -----")

	err := l.validParams()
	if err != nil {
		return
	}

	err = l.checkLogin()
	if err != nil {
		return
	}

	l.Report(l.Ctx.Input.IP(), "0", "POST", "0", "checkLogin", models.GetLoginAdminUserName(), "登陆成功", time.Now())
	l.OutPut(200, "登陆成功")
}
