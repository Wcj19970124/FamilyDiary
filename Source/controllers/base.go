package controllers

import (
	"encoding/json"
	"time"

	"../models"
	"../verificate"
	"github.com/astaxie/beego"
)

//BaseController 基础数据体
type BaseController struct {
	beego.Controller
}

//Verificate 校验用户登陆状态和权限
func (b *BaseController) Verificate() bool {
	//TODD:登陆校验
	if ok := verificate.AdminLogin(); !ok {
		b.OutPut(UserNotLogin, "登陆已失效,请重新登陆!")
		return false
	}
	//TODD:权限校验
	if ok := verificate.AdminPermission(); !ok {
		b.OutPut(PermissionDenied, "对不起,您没有此权限！")
		return false
	}
	return true
}

//OutPut 无data数据输出
//第一步：将数据添加进入map
//第二步：将数据编码成json
//第三步：设置响应头部，输出数据
func (b *BaseController) OutPut(ret int, msg string) {
	m := make(map[string]interface{})
	m["ret"] = ret
	m["msg"] = msg

	res, _ := json.Marshal(&m)

	b.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=utf-8")
	b.Ctx.WriteString(string(res))
}

//OutPutList 混合类型数据输出
func (b *BaseController) OutPutList(ret int, msg string, data map[string]interface{}) {
	m := make(map[string]interface{})
	m["ret"] = ret
	m["msg"] = msg
	m["data"] = data

	res, _ := json.Marshal(&m)

	b.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=utf-8")
	b.Ctx.WriteString(string(res))
}

//Report 数据上报 --- 写日志
func (b *BaseController) Report(ip string, t string, method string, fail string, function string, operateUser string, description string, operateTime time.Time) {
	models.AddSysLog(ip, t, method, fail, function, operateUser, description, operateTime)
}
