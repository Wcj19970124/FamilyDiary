package models

import (
	"strconv"
	"time"

	"../common"
	"github.com/astaxie/beego/logs"
)

//Role 后台角色结构体
type Role struct {
	Id          int
	RoleName    string
	Description string
	CreateUser  string
	UpdateUser  string
	Status      string
	UpdateTime  time.Time
	CreateTime  time.Time
}

//GetRoleByRoleName 根据角色名查询角色是否已经存在
func GetRoleByRoleName(rolename string) bool {

	sql := "select id from fd_role where role_name = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return false
	}

	var role Role
	err = dbProxy.Raw(sql, rolename).QueryRow(&role)
	if err == nil && strconv.Itoa(role.Id) != "" {
		return false
	}

	logs.Debug("---- query sql: sql " + sql)
	return true
}

//AddRole 添加角色
func AddRole(role Role) error {

	sql := "insert into fd_role(role_name,description,create_user,create_time,update_user,update_time,status) values(?,?,?,?,?,?,?)"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, role.RoleName, role.Description, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0).Exec()
	if err != nil {
		logs.Error("---- insert role failed,err:" + err.Error())
		return err
	}

	return nil
}

//DelRole 删除角色
func DelRole(id int) error {

	sql := "update fd_role set status = 1 where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, id).Exec()
	if err != nil {
		logs.Error("---- delete role failed,err:" + err.Error())
		return err
	}

	return nil
}

//UpdateRole 更新角色
func UpdateRole(role Role) error {

	sql := "update fd_role set role_name = ?,description=?,create_user=?,create_time=?,update_user=?,update_time=?,status=? where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, role.RoleName, role.Description, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0, role.Id).Exec()
	if err != nil {
		logs.Error("---- update role failed,err:" + err.Error())
		return err
	}

	return nil
}

//QueryRole 查询角色
func QueryRole(id int) (map[string]interface{}, error) {

	sql := "select id,role_name,description from fd_role where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var role Role
	err = dbProxy.Raw(sql, id).QueryRow(&role)
	if err != nil {
		logs.Error("---- query role by id failed,err:" + err.Error())
		return nil, err
	}

	m := make(map[string]interface{})
	m["role"] = role

	return m, nil
}

//QueryRoles 分页查询角色列表
func QueryRoles(p common.Page) (map[string]interface{}, error) {

	sql1 := "select count(id) from fd_role"
	sql2 := "select id,role_name,description from fd_role limit ?,?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var roles []Role
	p.SetStartNo() //设置查询数据的起始索引
	_, err = dbProxy.Raw(sql2, p.StartNo, p.PageSize).QueryRows(&roles)
	if err != nil {
		logs.Error("---- query roles(page) failed,err:" + err.Error())
		return nil, err
	}

	err = dbProxy.Raw(sql1).QueryRow(&p.TotalCount)
	if err != nil {
		logs.Error("---- query roles totalCount failed,err:" + err.Error())
		return nil, err
	}

	p.SetTotalPage() //设置总页数
	p.List = roles
	m := make(map[string]interface{})
	m["page"] = p

	return m, nil
}

//AllocatePermissions 角色分配权限
func AllocatePermissions(param map[string]interface{}) error {

	sql := "insert into fd_role_permission(role_id,permission_id,create_user,create_time,update_user,update_time,status) values(?,?,?,?,?,?,?)"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	roleID := int(param["id"].(float64))
	permissionIds := param["permissionIds"].([]interface{})
	for _, permissionID := range permissionIds {
		_, err = dbProxy.Raw(sql, roleID, permissionID, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0).Exec()
		if err != nil {
			logs.Error("---- insert role permission info failed,err:" + err.Error())
			return err
		}
	}

	return nil
}
