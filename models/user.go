package models

import (
	"time"

	"../util"
	"github.com/astaxie/beego/logs"
)

//AddUser 添加用户
func AddUser(user User) error {

	sql := "insert into fd_user(username,password,head,gender,remark,create_user,create_time,update_user,update_time,status) values(?,?,?,?,?,?,?,?,?,?)"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, user.Username, util.MD5(user.Password), user.Head, user.Gender, user.Remark, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0).Exec()
	if err != nil {
		logs.Error("--- add user failed,err:" + err.Error())
		return err
	}

	logs.Debug("insert user sql:" + sql)
	return nil
}

//DelUser 删除用户
func DelUser(id int) error {

	sql := "update fd_user set status = 1 where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, id).Exec()
	if err != nil {
		logs.Error("--- del user failed,err:" + err.Error())
		return err
	}

	logs.Debug("del user sql:" + sql)
	return nil
}

//UpdateUser 修改用户信息
func UpdateUser(user User) error {

	sql := "update fd_user set username=?,password=?,head=?,gender=?,remark=?,create_user=?,create_time=?,update_user=?,update_time=?,status=? where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, user.Username, util.MD5(user.Password), user.Head, user.Gender, user.Remark, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), user.Status, user.Id).Exec()
	if err != nil {
		logs.Error("--- update user failed,err:" + err.Error())
		return err
	}

	logs.Debug("--- update user sql:" + sql)
	return nil
}
