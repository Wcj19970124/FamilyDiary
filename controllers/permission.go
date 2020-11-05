package controllers

import (
	"encoding/json"

	"../common"
	"../models"
)

//PermissionController 权限数据体
type PermissionController struct {
	BaseController
}

//权限操作错误状态码
const (
	PermissionAddErr        = 30001
	PermissionDelErr        = 30002
	PermissionUpdateErr     = 30003
	PermissionQueryErr      = 30004
	PermissionAlreadyExists = 300006
)

//登陆态和权限校验
func (p *PermissionController) verificate() bool {
	return p.Verificate()
}

//AddPermission 添加权限
func (p *PermissionController) AddPermission() {
	//TODD：基础校验
	if ok := p.verificate(); !ok {
		return
	}
	//TODD:添加权限
	var permission models.Permission
	if json.Unmarshal(p.Ctx.Input.RequestBody, &permission) == nil {
		if ok := models.GetPermissionByPermissionName(permission.PermissionName); !ok {
			p.OutPut(PermissionAlreadyExists, "权限已存在!")
			return
		}
		if err := models.AddPermission(permission); err != nil {
			p.OutPut(PermissionAddErr, "权限添加失败")
			return
		}
		p.OutPut(200, "权限添加成功!")
	}
}

//DelPermission 删除权限
func (p *PermissionController) DelPermission() {
	//TODD：基础校验
	if ok := p.verificate(); !ok {
		return
	}
	//TODD：删除权限
	if id, err := p.GetInt("id"); err == nil {
		if err := models.DelPermission(id); err != nil {
			p.OutPut(PermissionDelErr, "权限删除失败!")
			return
		}
		p.OutPut(200, "权限删除成功!")
	}
}

//UpdatePermission 更新权限
func (p *PermissionController) UpdatePermission() {
	//TODD：基础校验
	if ok := p.verificate(); !ok {
		return
	}
	//TODD:更新权限
	var permission models.Permission
	if json.Unmarshal(p.Ctx.Input.RequestBody, &permission) == nil {
		if ok := models.GetPermissionByPermissionName(permission.PermissionName); !ok {
			p.OutPut(PermissionAlreadyExists, "权限已存在!")
			return
		}
		if err := models.UpdatePermission(permission); err != nil {
			p.OutPut(PermissionUpdateErr, "权限更新失败!")
			return
		}
		p.OutPut(200, "权限更新成功!")
	}
}

//QueryPermission 根据id查询单条权限信息
func (p *PermissionController) QueryPermission() {
	//TODD：基础校验
	if ok := p.verificate(); !ok {
		return
	}
	//TODD:查询
	if id, err := p.GetInt("id"); err == nil {
		data, err := models.QueryPermission(id)
		if err != nil {
			p.OutPut(PermissionQueryErr, "权限查询失败!")
			return
		}
		p.OutPutList(200, "权限查询成功!", data)
	}
}

//QueryPermissions 分页查询权限信息
func (p *PermissionController) QueryPermissions() {
	//TODD：基础校验
	if ok := p.verificate(); !ok {
		return
	}
	//TODD：查询
	var page common.Page
	if json.Unmarshal(p.Ctx.Input.RequestBody, &page) == nil {
		data, err := models.QueryPermissions(page)
		if err != nil {
			p.OutPut(PermissionQueryErr, "权限列表查询失败!")
			return
		}
		p.OutPutList(200, "权限列表查询成功!", data)
	}
}

//QueryPermissionsTree 以树状图形式返回权限列表
func (p *PermissionController) QueryPermissionsTree() {
	//TODD：基础校验
	if ok := p.verificate(); !ok {
		return
	}
	//TODD：查询权限列表,以树状图显示
	data, err := models.QueryPermissionsTree()
	if err != nil {
		p.OutPut(PermissionQueryErr,"权限树查询失败!")
		return
	}
	p.OutPutList(200,"权限树查询成功!",data)
}
