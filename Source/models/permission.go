package models

import (
	"strconv"
	"time"

	"../common"
	"github.com/astaxie/beego/logs"
)

//Permission 后台权限数据体
type Permission struct {
	Id              int
	PermissionName  string
	ParentId        int
	Type            string
	Url             string
	CreateUser      string
	UpdateUser      string
	Status          string
	CreateTime      time.Time
	UpdateTime      time.Time
	ChildPermission []*Permission
}

//GetPermissionByPermissionName 根据权限名判断权限是否已存在
func GetPermissionByPermissionName(permissionName string) bool {

	sql := "select id from fd_permission where permission_name = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return false
	}

	var permission Permission
	err = dbProxy.Raw(sql, permissionName).QueryRow(&permission)
	if err == nil && strconv.Itoa(permission.Id) != "" {
		return false
	}

	logs.Debug("---- query sql:" + sql)
	return true
}

//AddPermission 添加权限
func AddPermission(permission Permission) error {

	sql := "insert into fd_permission(permission_name,parent_id,type,url,create_user,create_time,update_user,update_time,status) values(?,?,?,?,?,?,?,?,?)"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, permission.PermissionName, permission.ParentId, permission.Type, permission.Url, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0).Exec()
	if err != nil {
		logs.Error("---- insert permission failed,err:" + err.Error())
		return err
	}

	return nil
}

//DelPermission 删除权限
func DelPermission(id int) error {

	sql := "update fd_permission set status = 1 where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, id).Exec()
	if err != nil {
		logs.Error("---- delete permission failed,err:" + err.Error())
		return err
	}

	return nil
}

//UpdatePermission 更新权限
func UpdatePermission(permission Permission) error {

	sql := "update fd_permission set permission_name = ?,parent_id =?,type = ?,url = ?,create_user=?,create_time=?,update_user=?,update_time=?,status=? where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, permission.PermissionName, permission.ParentId, permission.Type, permission.Url, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0, permission.Id).Exec()
	if err != nil {
		logs.Error("---- update permission failed,err:" + err.Error())
		return err
	}

	return nil
}

//QueryPermission 根据id查询权限信息
func QueryPermission(id int) (map[string]interface{}, error) {

	sql := "select id,permission_name,parent_id,type,url,status from fd_permission where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var permission Permission
	err = dbProxy.Raw(sql, id).QueryRow(&permission)
	if err != nil {
		logs.Error("---- query permission by id failed,err:" + err.Error())
		return nil, err
	}

	m := make(map[string]interface{})
	m["permission"] = permission

	return m, nil
}

//QueryPermissions 分页查询权限列表
func QueryPermissions(p common.Page) (map[string]interface{}, error) {

	sql1 := "select count(id) from fd_permission"
	sql2 := "select id,permission_name,parent_id,type,url,status from fd_permission limit ?,?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var permissions []Permission
	p.SetStartNo() //设置查询数据的起始索引
	_, err = dbProxy.Raw(sql2, p.StartNo, p.PageSize).QueryRows(&permissions)
	if err != nil {
		logs.Error("---- query permission(page) failed,err:" + err.Error())
		return nil, err
	}

	err = dbProxy.Raw(sql1).QueryRow(&p.TotalCount)
	if err != nil {
		logs.Error("---- query permission totalCount failed,err:" + err.Error())
		return nil, err
	}

	p.SetTotalPage()
	p.List = permissions
	m := make(map[string]interface{})
	m["page"] = p

	return m, nil
}

//QueryPermissionsTree 查询权限列表，返回权限树
func QueryPermissionsTree() (map[string]interface{}, error) {

	sql := "select id,permission_name,parent_id,type,url,status from fd_permission"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var permissions []*Permission
	_, err = dbProxy.Raw(sql).QueryRows(&permissions)
	if err != nil {
		logs.Error("---- query permission(tree) failed,err:" + err.Error())
		return nil, err
	}

	m := convertPermissionsToPermissionTree(permissions)

	return m, nil
}

//将权限转换为权限树 --- 递归转换
func convertPermissionsToPermissionTree(permissions []*Permission) map[string]interface{} {

	root := make(map[string]interface{})
	permissionTree := []*Permission{}

	//提取父级权限
	for _, permission := range permissions {
		if permission.ParentId == 0 {
			permissionTree = append(permissionTree, permission)
		}
	}

	convertChildPermissionToPermissionTree(permissions, permissionTree)
	root["permissionTree"] = permissionTree
	return root
}

//将子权限放入权限树
func convertChildPermissionToPermissionTree(permissions []*Permission, parentPermissions []*Permission) {

	if len(parentPermissions) == 0 {
		return
	}

	//添加子权限进入父权限中
	for _, parentPermission := range parentPermissions {
		for _, permission := range permissions {
			if permission.ParentId == parentPermission.Id {
				parentPermission.ChildPermission = append(parentPermission.ChildPermission, permission)
			}
		}
	}

	//递归添加子权限
	for _, parentPermission := range parentPermissions {
		convertChildPermissionToPermissionTree(permissions, parentPermission.ChildPermission)
	}
}
