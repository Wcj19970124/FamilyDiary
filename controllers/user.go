package controllers

import (
	"encoding/json"

	"../common"
	"../models"
)

//UserController 后台用户数据体
type UserController struct {
	BaseController
}

//用户操作错误码
const (
	UserAddErr       = 30001 //添加用户失败
	UserDelErr       = 30002 //删除用户失败
	UserUpdateErr    = 30003 //更新用户失败
	UserQueryErr     = 30004 //查询用户失败
	PermissionDenied = 30005 //无权限
	UserNameExists   = 30006 //用户名已存在
)

//登陆态和权限校验
func (u *UserController) verificate() bool {
	return u.Verificate()
}

//AddUser 添加用户
func (u *UserController) AddUser() {
	//TODD:基础校验
	if ok := u.verificate(); !ok {
		return
	}

	//TODD：添加用户
	var user models.User
	if json.Unmarshal(u.Ctx.Input.RequestBody, &user) == nil {
		if ok := models.GetUserByUserName(user.Username); !ok {
			u.OutPut(UserNameExists, "用户名已存在!")
			return
		}
		if err := models.AddUser(user); err != nil {
			u.OutPut(UserAddErr, "添加用户失败！")
			return
		}
		u.OutPut(200, "添加用户成功!")
	}
}

//DelUser 删除用户
func (u *UserController) DelUser() {
	//TODD:基础校验
	if ok := u.verificate(); !ok {
		return
	}
	//TODD：删除用户
	if id, err := u.GetInt("id"); err == nil {
		if err = models.DelUser(id); err != nil {
			u.OutPut(UserDelErr, "删除用户失败!")
			return
		}
		u.OutPut(200, "删除用户成功!")
	}
}

//UpdateUser 修改用户信息
func (u *UserController) UpdateUser() {

	//TODD:基础校验
	if ok := u.verificate(); !ok {
		return
	}
	//TODD：修改用户
	var user models.User
	if json.Unmarshal(u.Ctx.Input.RequestBody, &user) == nil {
		if ok := models.GetUserByUserName(user.Username); !ok {
			u.OutPut(UserNameExists, "该用户名已存在,无法修改!")
			return
		}
		if err := models.UpdateUser(user); err != nil {
			u.OutPut(UserUpdateErr, "用户信息更新失败!")
			return
		}
		u.OutPut(200, "用户信息更新成功!")
	}
}

//QueryUser 根据id查询单条用户信息
func (u *UserController) QueryUser() {
	//TODD:基础校验
	if ok := u.verificate(); !ok {
		return
	}
	//TODD:查询用户
	if id, err := u.GetInt("id"); err == nil {
		user, err := models.GetUserByID(id)
		if err != nil {
			u.OutPut(UserQueryErr, "查询用户失败！")
			return
		}
		u.OutPutList(200, "查询用户成功!", user)
	}
}

//QueryUsers 分页查询用户信息
func (u *UserController) QueryUsers() {
	//TODD:基础校验
	if ok := u.verificate(); !ok {
		return
	}
	//TODD:分页查询
	var page common.Page
	if json.Unmarshal(u.Ctx.Input.RequestBody, &page) == nil {
		data, err := models.GetUsers(page)
		if err != nil {
			u.OutPut(UserQueryErr, "用户列表查询失败!")
			return
		}
		u.OutPutList(200, "用户列表查询成功!", data)
	}
}
