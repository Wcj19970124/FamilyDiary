package models

import (
	"strconv"
	"time"

	"../common"
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

//GetUserByUserName 根据用户名查询用户是否已经存在
func GetUserByUserName(username string) bool {

	sql := "select id from fd_user where username = ? and status = 0"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return false
	}

	var user User
	err = dbProxy.Raw(sql, username).QueryRow(&user)
	if err == nil && strconv.Itoa(user.Id) != "" {
		return false
	}

	logs.Debug("--- query sql:" + sql)
	return true
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

//GetUserByID 根据用户id查询用户信息
func GetUserByID(id int) (map[string]interface{}, error) {

	sql := "select id,username,password,head,gender,remark from fd_user where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var user User
	err = dbProxy.Raw(sql, id).QueryRow(&user)
	if err != nil {
		logs.Error("---- query user failed,err:" + err.Error())
		return nil, err
	}

	m := make(map[string]interface{})
	m["user"] = user

	return m, nil
}

//GetUsers 分页查询用户列表
func GetUsers(p common.Page) (map[string]interface{}, error) {

	sql1 := "select id,username,password,head,gender,remark,status from fd_user limit ?,?"
	sql2 := "select count(id) as totalCount from fd_user"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var users []User
	p.SetStartNo() //设置查询数据起始索引
	_, err = dbProxy.Raw(sql1, p.StartNo, p.PageSize).QueryRows(&users)
	if err != nil {
		logs.Error("---- query users(page) failed,err:" + err.Error())
		return nil, err
	}

	err = dbProxy.Raw(sql2).QueryRow(&p.TotalCount)
	if err != nil {
		logs.Error("---- query users totalCount failed,err:" + err.Error())
		return nil, err
	}

	p.List = users
	p.SetTotalPage() //设置总页数
	m := make(map[string]interface{})
	m["page"] = p

	return m, nil
}
