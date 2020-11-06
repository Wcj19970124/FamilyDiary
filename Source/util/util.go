package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

//MD5 加密
func MD5(val string) string {
	h := md5.New()
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))
}

//ConvertStructToString 结构体转换为字符串形式
func ConvertStructToString(v interface{}) string {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		logs.Error("convert struct to string failed,err:" + err.Error())
	}
	return string(jsonStr)
}
