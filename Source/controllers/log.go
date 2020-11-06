package controllers

import (
	"encoding/json"
	"time"

	"../common"
	"../models"
)

//LogController 日志数据体
type LogController struct {
	BaseController
}

//日志操作错误状态码
const (
	SysLogDelErr   = 30002 //删除日志失败
	SysLogQueryErr = 30004 //查询日志失败
)

//登录态和权限校验
func (l *LogController) verificate() bool {
	return l.Verificate()
}

//DelSysLog 单条日志删除
func (l *LogController) DelSysLog() {
	//基础校验
	if ok := l.verificate(); !ok {
		return
	}
	//TODD：删除
	if id, err := l.GetInt("id"); err == nil {
		if err := models.DelSysLog(id); err != nil {
			l.Report(l.Ctx.Input.IP(), "1", "DELETE", "1", "DelSysLog", models.GetLoginAdminUserName(), "日志删除失败", time.Now())
			l.OutPut(SysLogDelErr, "日志删除失败!")
			return
		}
		l.Report(l.Ctx.Input.IP(), "1", "DELETE", "0", "DelSysLog", models.GetLoginAdminUserName(), "日志删除成功", time.Now())
		l.OutPut(200, "日志删除成功")
	}
}

//DelSysLogs 批量删除日志
func (l *LogController) DelSysLogs() {
	//基础校验
	if ok := l.verificate(); !ok {
		return
	}
	//TODD:删除
	param := make(map[string]interface{})
	if json.Unmarshal(l.Ctx.Input.RequestBody, &param) == nil {
		if err := models.DelSysLogs(param); err != nil {
			l.Report(l.Ctx.Input.IP(), "1", "DELETE", "1", "DelSysLogs", models.GetLoginAdminUserName(), "日志批量删除失败", time.Now())
			l.OutPut(SysLogDelErr, "日志批量删除失败!")
			return
		}
		l.Report(l.Ctx.Input.IP(), "1", "DELETE", "0", "DelSysLogs", models.GetLoginAdminUserName(), "日志批量删除成功", time.Now())
		l.OutPut(200, "日志批量删除成功")
	}
}

//QuerySysLogs 分页查询日志列表
func (l *LogController) QuerySysLogs() {
	//基础校验
	if ok := l.verificate(); !ok {
		return
	}
	//TODD：查询
	var page common.Page
	if json.Unmarshal(l.Ctx.Input.RequestBody, &page) == nil {
		data, err := models.QuerySysLogs(page)
		if err != nil {
			l.Report(l.Ctx.Input.IP(), "1", "POST", "1", "QuerySysLogs", models.GetLoginAdminUserName(), "日志列表查询失败", time.Now())
			l.OutPut(SysLogQueryErr, "日志列表查询失败!")
			return
		}
		l.Report(l.Ctx.Input.IP(), "1", "POST", "0", "QuerySysLogs", models.GetLoginAdminUserName(), "日志列表查询成功", time.Now())
		l.OutPutList(200, "日志列表查询成功!", data)
	}
}
