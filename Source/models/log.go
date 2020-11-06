package models

import (
	"time"

	"../common"
	"github.com/astaxie/beego/logs"
)

//SysLog 后台日志结构体
type SysLog struct {
	Id           int
	Ip           string
	Type         string
	Method       string
	Fail         string
	FunctionName string
	OperateUser  string
	Description  string
	OperateTime  time.Time
}

//AddSysLog 写入日志
func AddSysLog(ip string, t string, method string, fail string, function string, operateUser string, description string, operateTime time.Time) error {

	sql := "insert into fd_syslog(ip,type,method,fail,operate_time,function_name,operate_user,description) values(?,?,?,?,?,?,?,?)"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	logs.Debug(sql)
	_, err = dbProxy.Raw(sql, ip, t, method, fail, operateTime, function, operateUser, description).Exec()
	if err != nil {
		logs.Error("---- insert log failed,err:" + err.Error())
		return err
	}

	return nil
}

//DelSysLog 删除单条日志
func DelSysLog(id int) error {

	sql := "delete from fd_syslog where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, id).Exec()
	if err != nil {
		logs.Error("---- delete log failed,err:" + err.Error())
		return err
	}

	return nil
}

//DelSysLogs 日志批量删除
func DelSysLogs(param map[string]interface{}) error {

	sql := "delete from fd_syslog where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	sysLogIds := param["sysLogIds"].([]interface{})
	for _, sysLogID := range sysLogIds {
		_, err = dbProxy.Raw(sql, sysLogID).Exec()
		if err != nil {
			logs.Error("---- delete log failed,err:" + err.Error())
			return err
		}
	}

	return nil
}

//QuerySysLogs 分页查询日志列表
func QuerySysLogs(p common.Page) (map[string]interface{}, error) {

	sql1 := "select count(id) from fd_syslog"
	sql2 := "select id,ip,type,method,fail,function_name,operate_time,operate_user,description from fd_syslog limit ?,?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var sysLogs []SysLog
	p.SetStartNo() //设置搜索的起始索引
	_, err = dbProxy.Raw(sql2, p.StartNo, p.PageSize).QueryRows(&sysLogs)
	if err != nil {
		logs.Error("---- get syslog(page) failed,err:" + err.Error())
		return nil, err
	}

	err = dbProxy.Raw(sql1).QueryRow(&p.TotalCount)
	if err != nil {
		logs.Error("---- get syslog totalCount failed,err:" + err.Error())
		return nil, err
	}

	p.SetTotalPage() //设置总页数
	p.List = sysLogs
	m := make(map[string]interface{})
	m["page"] = p

	return m, nil
}
