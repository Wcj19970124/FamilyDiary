package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"../common"
	"../models"
)

//RoleController 角色数据体
type RoleController struct {
	BaseController
}

//角色操作错误状态码
const (
	RoleAddErr                 = 30001 //角色添加失败
	RoleDelErr                 = 30002 //角色删除失败
	RoleUpdateErr              = 30003 //角色修改失败
	RoleQueryErr               = 30004 //角色查询失败
	RoleAlreadyExists          = 30006 //角色已存在
	RoleAllocatePermissionsErr = 30007 //角色分配权限
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
			r.Report(r.Ctx.Input.IP(), "1", "POST", "1", "AddRole", models.GetLoginAdminUserName(), "角色已存在", time.Now())
			r.OutPut(RoleAlreadyExists, "角色已存在!")
			return
		}
		if err := models.AddRole(role); err != nil {
			r.Report(r.Ctx.Input.IP(), "1", "POST", "1", "AddRole", models.GetLoginAdminUserName(), "角色添加失败", time.Now())
			r.OutPut(RoleAddErr, "角色添加失败!")
			return
		}
		r.Report(r.Ctx.Input.IP(), "1", "POST", "0", "AddRole", models.GetLoginAdminUserName(), "角色添加成功", time.Now())
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
			r.Report(r.Ctx.Input.IP(), "1", "DELETE", "1", "DelRole", models.GetLoginAdminUserName(), "角色删除失败", time.Now())
			r.OutPut(RoleDelErr, "角色删除失败!")
			return
		}
		r.Report(r.Ctx.Input.IP(), "1", "DELETE", "0", "DelRole", models.GetLoginAdminUserName(), "角色删除成功", time.Now())
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
			r.Report(r.Ctx.Input.IP(), "1", "PUT", "1", "UpdateRole", models.GetLoginAdminUserName(), "角色已存在", time.Now())
			r.OutPut(RoleAlreadyExists, "角色已存在!")
			return
		}
		if err := models.UpdateRole(role); err != nil {
			r.Report(r.Ctx.Input.IP(), "1", "PUT", "1", "UpdateRole", models.GetLoginAdminUserName(), "角色更新失败", time.Now())
			r.OutPut(RoleUpdateErr, "角色更新失败!")
			return
		}
		r.Report(r.Ctx.Input.IP(), "1", "PUT", "0", "UpdateRole", models.GetLoginAdminUserName(), "角色更新成功！", time.Now())
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
			r.Report(r.Ctx.Input.IP(), "1", "GET", "1", "QueryRole", models.GetLoginAdminUserName(), "角色查询失败", time.Now())
			r.OutPut(RoleQueryErr, "角色查询失败")
			return
		}
		r.Report(r.Ctx.Input.IP(), "1", "GET", "0", "QueryRole", models.GetLoginAdminUserName(), "角色查询成功", time.Now())
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
			r.Report(r.Ctx.Input.IP(), "1", "POST", "1", "QueryRoles", models.GetLoginAdminUserName(), "角色列表查询失败", time.Now())
			r.OutPut(RoleQueryErr, "角色列表查询失败!")
			return
		}
		r.Report(r.Ctx.Input.IP(), "1", "POST", "0", "QueryRoles", models.GetLoginAdminUserName(), "角色列表查询成功", time.Now())
		r.OutPutList(200, "角色列表查询成功!", data)
	}
}

//AllocatePermissions 为角色分配权限
func (r *RoleController) AllocatePermissions() {
	//TODD：基础校验
	if ok := r.verificate(); !ok {
		return
	}
	//TODD：分配权限
	param := make(map[string]interface{})
	if json.Unmarshal(r.Ctx.Input.RequestBody, &param) == nil {
		if err := models.AllocatePermissions(param); err != nil {
			r.Report(r.Ctx.Input.IP(), "1", "POST", "1", "AllocatePermissions", models.GetLoginAdminUserName(), "分配权限失败", time.Now())
			r.OutPut(RoleAllocatePermissionsErr, "分配权限失败!")
			return
		}
		r.Report(r.Ctx.Input.IP(), "1", "POST", "0", "AllocatePermissions", models.GetLoginAdminUserName(), "分配权限成功", time.Now())
		r.OutPut(200, "分配权限成功!")
	}
}
