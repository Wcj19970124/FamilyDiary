package models

import (
	"../util"
	"github.com/astaxie/beego/logs"
)

//QueryPwdByUserName 根据用户名查询用户密码
func QueryPwdByUserName(username string) (string, error) {

	sql := "select password from fd_user where username = ? and status = 0"

	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return "", err
	}

	var user User
	err = dbProxy.Raw(sql, username).QueryRow(&user)
	if err != nil {
		logs.Error("---- db query failed,err:" + err.Error() + " ----")
		return "", err
	}

	logs.Debug("query sql:" + sql + ",query result:" + util.ConvertStructToString(&user))
	return user.Password, nil
}

//GetLoginAdminUserName 获取登陆用户的用户名
func GetLoginAdminUserName() string {
	key := "login_admin_username"
	username, _, _ := GetByKey(key)
	return username
}
