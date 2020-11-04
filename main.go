package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//初始化操作：日志输出等
func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"file/log/family_diary.log"}`)
}

func main() {
	beego.Run()
}
