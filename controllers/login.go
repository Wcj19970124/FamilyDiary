package controllers

import (
	"errors"

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

//登陆错误状态码(ret)
const (
	UserNotLogin        = 10001 //用户未登录
	UserLoginParamNull  = 10002 //参数为空
	UserLoginParamError = 10003 //参数错误
	UserNotExist        = 10004 //用户不存在
)

var userLoginRedisKeyPrefix = "login_" //登陆缓存前缀
var userLoginRedisCacheSeconds = 3600  //登陆缓存时长

//参数校验
func (l *UserLoginController) validParams() error {

	if username := l.GetString("username"); username != "" {
		l._username = username
	} else {
		l.OutPut(UserLoginParamNull, "参数为空(username)")
		return errors.New("参数为空")
	}

	if password := l.GetString("password"); password != "" {
		l._password = password
	} else {
		l.OutPut(UserLoginParamNull, "参数为空(password)")
		return errors.New("参数为空")
	}

	logs.Debug("login username:" + l._username + ",password:" + l._password)
	return nil
}

//登陆校验
//第一步：先检查缓存中是否存在登陆信息,有则不用重复登陆
//第二步：若缓存没有,查询数据库,若数据库没有,直接返回
//第三步：若数据库存在,则校验数据,正确则写缓存,不正确直接返回
func (l *UserLoginController) checkLogin() error {

	loginKey := userLoginRedisKeyPrefix + l._username
	logs.Debug("---- user login cache redis key:" + loginKey)
	if val, _, err := models.GetByKey(loginKey); err == nil || val != "" {
		l.OutPut(200, "已登陆！")
		return errors.New("已登录")
	}

	password, _ := models.QueryPwdByUserName(l._username)
	if password == "" {
		l.OutPut(UserNotExist, "用户不存在!")
		return errors.New("用户不存在！")
	}
	if password != util.MD5(l._password) {
		l.OutPut(UserLoginParamError, "密码错误!")
		return errors.New("密码错误")
	}

	models.SetByKey(loginKey, password, userLoginRedisCacheSeconds)

	return nil
}

//Post 请求入口
func (l *UserLoginController) Post() {
	err := l.validParams()
	if err != nil {
		return
	}

	err = l.checkLogin()
	if err != nil {
		return
	}

	l.OutPut(200, "登陆成功")
}
