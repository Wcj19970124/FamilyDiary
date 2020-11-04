package models

import (
	"time"

	"../util"
	"github.com/astaxie/beego/logs"
)

//User 后台用户结构体
type User struct {
	Id         int
	Username   string
	Password   string
	Head       string
	Gender     string
	Remark     string
	CreateUser string
	UpdateUser string
	Status     string
	CreateTime time.Time
	UpdateTime time.Time
}

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
