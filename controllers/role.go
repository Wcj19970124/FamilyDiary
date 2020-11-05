package controllers

import (
	"encoding/json"
	"fmt"

	"../common"
	"../models"
)

//RoleController 角色数据体
type RoleController struct {
	BaseController
}

//角色操作错误状态码
const (
	RoleAddErr        = 30001 //角色添加失败
	RoleDelErr        = 30002 //角色删除失败
	RoleUpdateErr     = 30003 //角色修改失败
	RoleQueryErr      = 30004 //角色查询失败
	RoleAlreadyExists = 30006 //角色已存在
)

//登录态和权限验证
func (r *RoleController) verificate() bool {
	return r.Verificate()
}

//AddRole 添加角色
func (r *RoleController) AddRole() {
	//TODD：基础校验
	if ok := r.verificate(); !ok {
		return
	}
	//TODD：添加角色
	var role models.Role
	if json.Unmarshal(r.Ctx.Input.RequestBody, &role) == nil {
		fmt.Println(role)
		if ok := models.GetRoleByRoleName(role.RoleName); !ok {
			r.OutPut(RoleAlreadyExists, "角色已存在!")
			return
		}
		if err := models.AddRole(role); err != nil {
			r.OutPut(RoleAddErr, "角色添加失败!")
			return
		}
		r.OutPut(200, "角色添加成功!")
	}
}

//DelRole 删除角色
func (r *RoleController) DelRole() {
	//TODD：基础校验
	if ok := r.verificate(); !ok {
		return
	}
	//TODD:删除角色
	if id, err := r.GetInt("id"); err == nil {
		if err := models.DelRole(id); err != nil {
			r.OutPut(RoleDelErr, "角色删除失败!")
			return
		}
		r.OutPut(200, "角色删除成功!")
	}
}

//UpdateRole 更新角色
func (r *RoleController) UpdateRole() {
	//TODD：基础校验
	if ok := r.verificate(); !ok {
		return
	}
	//TODD：更新角色
	var role models.Role
	if json.Unmarshal(r.Ctx.Input.RequestBody, &role) == nil {
		if ok := models.GetRoleByRoleName(role.RoleName); !ok {
			r.OutPut(RoleAlreadyExists, "角色已存在!")
			return
		}
		if err := models.UpdateRole(role); err != nil {
			r.OutPut(RoleUpdateErr, "角色更新失败!")
			return
		}
		r.OutPut(200, "角色更新成功！")
	}
}

//QueryRole 根据id查询角色信息
func (r *RoleController) QueryRole() {
	//TODD：基础校验
	if ok := r.verificate(); !ok {
		return
	}
	//TODD：查询
	if id, err := r.GetInt("id"); err == nil {
		role, err := models.QueryRole(id)
		if err != nil {
			r.OutPut(RoleQueryErr, "角色查询失败")
		}
		r.OutPutList(200, "角色查询成功!", role)
	}
}

//QueryRoles 分页查询角色列表
func (r *RoleController) QueryRoles() {
	//TODD：基础校验
	if ok := r.verificate(); !ok {
		return
	}
	//TODD：分页查询
	var page common.Page
	if json.Unmarshal(r.Ctx.Input.RequestBody, &page) == nil {
		data, err := models.QueryRoles(page)
		if err != nil {
			r.OutPut(RoleQueryErr, "角色列表查询失败!")
			return
		}
		r.OutPutList(200, "角色列表查询成功!", data)
	}
}
