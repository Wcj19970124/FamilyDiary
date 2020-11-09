package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

//GetByKey 读redis
func GetByKey(key string) (string, int, error) {
	return getRedis(key)
}

//SetByKey 存redis
func SetByKey(key string, val string, ttl int) (int, error) {
	return setRedis(key, val, ttl)
}

//Expire 设置缓存时间
func Expire(key string, ttl int) (int, error) {
	return expireRedis(key, ttl)
}

//DelKey 删除缓存的key
func DelKey(key string) (int, error) {
	return deleteKey(key)
}

func getRedis(key string) (string, int, error) {
	logs.Debug("---- get redis starting ----")
	conn, err := store.GetRedis()
	if err != nil {
		logs.Error("get redis conn failed,err:" + err.Error())
		return "", 0, err
	}

	defer conn.Close() //最终需要关闭连接

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		logs.Debug("redis exec get failed,err:" + err.Error())
		return "", 0, err
	}

	//截取打印字符串---val
	tmp := val
	r := []rune(tmp)
	if len(r) > 1000 {
		tmp = string(r[0:1000])
	}
	logs.Debug("---- get redis val:" + tmp + " ----")

	return val, 0, nil
}

func setRedis(key string, val string, ttl int) (int, error) {
	logs.Debug("---- set redis starting ----")
	conn, err := store.GetRedis()
	if err != nil {
		logs.Error("---- get redis conn failed,err:" + err.Error() + " ----")
		return 0, err
	}

	defer conn.Close() //最终关闭连接

	_, err = conn.Do("SETEX", key, ttl, val)
	if err != nil {
		logs.Error("redis set key:" + key + " failed,err:" + err.Error())
		return 0, err
	}

	//截取打印字符串---val
	tmp := val
	r := []rune(tmp)
	if len(r) > 1000 {
		tmp = string(r[0:1000])
	}
	logs.Debug("---- get redis val:" + tmp + " ----")

	return 0, nil
}

func expireRedis(key string, ttl int) (int, error) {
	logs.Debug("---- expire redis starting ----")
	conn, err := store.GetRedis()
	if err != nil {
		logs.Error("---- get redis conn failed,err:" + err.Error() + " ----")
		return 0, err
	}

	defer conn.Close()

	_, err = conn.Do("expire", key, ttl)
	if err != nil {
		logs.Error("redis expire key:" + key + " failed,err:" + err.Error())
		return 0, err
	}

	return 0, nil
}

func deleteKey(key string) (int, error) {
	logs.Debug("---- expire redis starting ----")
	conn, err := store.GetRedis()
	if err != nil {
		logs.Error("---- get redis conn failed,err:" + err.Error() + " ----")
		return 0, err
	}

	defer conn.Close()

	_, err = conn.Do("del", key)
	if err != nil {
		logs.Error("del key:" + key + " falied,err:" + err.Error())
		return 0, err
	}

	return 0, nil
}
