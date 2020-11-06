package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"../common"
	"../models"
)

//UserController 后台用户数据体
type UserController struct {
	BaseController
}

//用户操作错误码
const (
	UserAddErr          = 30001 //添加用户失败
	UserDelErr          = 30002 //删除用户失败
	UserUpdateErr       = 30003 //更新用户失败
	UserQueryErr        = 30004 //查询用户失败
	PermissionDenied    = 30005 //无权限
	UserNameExists      = 30006 //用户名已存在
	UserAllocateRoleErr = 30007 //用户角色分配失败
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
			u.Report(u.Ctx.Input.IP(), "1", "POST", "1", "GetUserByUserName", models.GetLoginAdminUserName(), "用户名已存在", time.Now())
			u.OutPut(UserNameExists, "用户名已存在!")
			return
		}
		if err := models.AddUser(user); err != nil {
			u.Report(u.Ctx.Input.IP(), "1", "POST", "1", "AddUser", models.GetLoginAdminUserName(), "添加用户失败！", time.Now())
			u.OutPut(UserAddErr, "添加用户失败！")
			return
		}

		u.Report(u.Ctx.Input.IP(), "1", "POST", "0", "AddUser", models.GetLoginAdminUserName(), "添加用户成功", time.Now())
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
			u.Report(u.Ctx.Input.IP(), "1", "DELETE", "1", "DelUser", models.GetLoginAdminUserName(), "删除用户失败", time.Now())
			u.OutPut(UserDelErr, "删除用户失败!")
			return
		}
		u.Report(u.Ctx.Input.IP(), "1", "DELETE", "0", "DelUser", models.GetLoginAdminUserName(), "删除用户成功", time.Now())
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
			u.Report(u.Ctx.Input.IP(), "1", "PUT", "1", "GetUserByUserName", models.GetLoginAdminUserName(), "该用户名已存在,无法修改!", time.Now())
			u.OutPut(UserNameExists, "该用户名已存在,无法修改!")
			return
		}
		if err := models.UpdateUser(user); err != nil {
			u.Report(u.Ctx.Input.IP(), "1", "PUT", "1", "UpdateUser", models.GetLoginAdminUserName(), "用户信息更新失败", time.Now())
			u.OutPut(UserUpdateErr, "用户信息更新失败!")
			return
		}
		u.Report(u.Ctx.Input.IP(), "1", "PUT", "0", "UpdateUser", models.GetLoginAdminUserName(), "用户信息更新成功", time.Now())
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
			u.Report(u.Ctx.Input.IP(), "1", "GET", "1", "GetUserByID", models.GetLoginAdminUserName(), "查询用户失败！", time.Now())
			u.OutPut(UserQueryErr, "查询用户失败！")
			return
		}
		u.Report(u.Ctx.Input.IP(), "1", "GET", "0", "QueryUser", models.GetLoginAdminUserName(), "查询用户成功", time.Now())
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
		fmt.Println(u.Ctx.Input.IP())
		data, err := models.GetUsers(page)
		if err != nil {
			u.Report(u.Ctx.Input.IP(), "1", "POST", "1", "GetUsers", models.GetLoginAdminUserName(), "用户列表查询失败", time.Now())
			u.OutPut(UserQueryErr, "用户列表查询失败!")
			return
		}
		u.Report(u.Ctx.Input.IP(), "1", "POST", "0", "QueryUsers", models.GetLoginAdminUserName(), "用户列表查询成功", time.Now())
		u.OutPutList(200, "用户列表查询成功!", data)
	}
}

//AllocateRoles 为用户分配角色 1-n的关系
func (u *UserController) AllocateRoles() {
	//TODD:基础校验
	if ok := u.verificate(); !ok {
		return
	}
	//TODD：分配角色
	param := make(map[string]interface{})
	if json.Unmarshal(u.Ctx.Input.RequestBody, &param) == nil {
		fmt.Println(param)
		if err := models.AllocateRoles(param); err != nil {
			u.Report(u.Ctx.Input.IP(), "1", "POST", "1", "AllocateRoles", models.GetLoginAdminUserName(), "用户分配角色失败！", time.Now())
			u.OutPut(UserAllocateRoleErr, "用户分配角色失败！")
			return
		}
		u.Report(u.Ctx.Input.IP(), "1", "POST", "0", "AllocateRoles", models.GetLoginAdminUserName(), "用户分配角色成功", time.Now())
		u.OutPut(200, "用户分配角色成功!")
	}
}
