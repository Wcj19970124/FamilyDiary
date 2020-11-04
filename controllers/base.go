package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

//BaseController 基础数据体
type BaseController struct {
	beego.Controller
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
