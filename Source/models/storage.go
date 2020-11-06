package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql" //数据库驱动
)

//Storage 提供redis和db连接
type Storage struct {
	holder map[string]interface{}
}

var store Storage

//单例模式保证只有一个holder
func (s *Storage) getHolder() map[string]interface{} {
	if len(s.holder) < 1 {
		s.holder = make(map[string]interface{})
	}
	return s.holder
}

//GetRedis 获取redis连接
//第一步：先获取连接池
//第二步：再冲连接池中取出一个连接
func (s *Storage) GetRedis() (redis.Conn, error) {
	logs.Debug("---- 开始获取redis连接 ----")
	pool, err := s.getRedisPool()
	if err != nil {
		return nil, err
	}
	return pool.Get(), nil
}

//获取redis连接池
//第一步：检查store中是否已经有连接池,如果有,直接返回
//第二步：如果没有,则获取配置文件的信息，创建一个连接池返回
func (s *Storage) getRedisPool() (*redis.Pool, error) {
	logs.Debug("---- 开始获取redis连接池 ----")
	if _, ok := s.holder["redis"].(*redis.Pool); ok {
		return s.holder["redis"].(*redis.Pool), nil
	}

	maxIdle, maxActive, host, port, pwd, err := s.getRedisConfig()
	if err != nil {
		return nil, err
	}

	//创建一个redis连接池
	s.getHolder()["redis"] = &redis.Pool{
		MaxIdle:   maxIdle,
		MaxActive: maxActive,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				if _, err := conn.Do("AUTH", pwd); err != nil {
					return nil, err
				}
			}
			return conn, nil
		},
	}

	return s.holder["redis"].(*redis.Pool), nil
}

//获取配置文件中的redis配置信息
func (s *Storage) getRedisConfig() (maxIdle int, maxActive int, host string, port string, pwd string, err error) {
	logs.Debug("---- 开始获取配置文件信息redis信息 ----")
	maxIdle, err = beego.AppConfig.Int("redis::maxIdle")
	if err != nil {
		maxIdle, err = 10, nil
	}

	maxActive, err = beego.AppConfig.Int("redis::maxActive")
	if err != nil {
		maxActive, err = 100, nil
	}

	txt := ""
	host = beego.AppConfig.String("redis::host")
	if host == "" {
		txt += "redis host is not exist... "
	}

	port = beego.AppConfig.String("redis::port")
	if port == "" {
		txt += "redis port is not exists... "
	}

	pwd = beego.AppConfig.String("redis::auth")
	if pwd == "" {
		txt += "redis password is not exist... "
	}

	if len(txt) > 0 {
		err = errors.New(txt)
	}

	logs.Debug("---- host:" + host + ",port:" + port + ",pwd:" + pwd + ",maxIdle:" + strconv.Itoa(maxIdle) + ",maxActive:" + strconv.Itoa(maxActive) + " ----")
	return maxIdle, maxActive, host, port, pwd, err
}

//GetDBProxy 获取数据库连接
//第一步：如果已经存在，直接返回
//第二步：如果不存在，获取配置信息，连接数据库，返回连接
func (s *Storage) GetDBProxy() (orm.Ormer, error) {
	logs.Debug("---- 开始获取数据库连接 ----")
	if _, ok := s.holder["db"].(orm.Ormer); ok {
		return s.holder["db"].(orm.Ormer), nil
	}

	maxIdle, maxConn, user, pwd, dbName, err := s.getDBConfig()
	if err != nil {
		return nil, err
	}

	//注册数据库驱动/数据库
	dataSource := user + ":" + pwd + "@/" + dbName + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)

	//创建连接
	o := orm.NewOrm()
	o.Using("default")
	s.getHolder()["db"] = o

	return s.holder["db"].(orm.Ormer), nil
}

//获取配置文件中数据库的配置信息
func (s *Storage) getDBConfig() (maxIdle int, maxConn int, user string, pwd string, dbName string, err error) {
	logs.Debug("---- 开始获取配置文件中DB信息 ----")
	maxIdle, err = beego.AppConfig.Int("db::maxIdle")
	if err != nil {
		maxIdle, err = 10, nil
	}

	maxConn, err = beego.AppConfig.Int("db::maxConn")
	if err != nil {
		maxConn, err = 30, nil
	}

	txt := ""
	user = beego.AppConfig.String("db::user")
	if user == "" {
		txt += "db user is not exist.. "
	}

	pwd = beego.AppConfig.String("db::password")
	if pwd == "" {
		txt += "db password is not exist.. "
	}

	dbName = beego.AppConfig.String("db::dbname")
	if dbName == "" {
		txt += "db dbName is not exist.. "
	}

	if len(txt) > 0 {
		err = errors.New(txt)
	}

	logs.Debug("user:" + user + ",password:" + pwd + ",dbname:" + dbName + ",maxIdle:" + strconv.Itoa(maxIdle) + ",maxConn:" + strconv.Itoa(maxConn))
	return maxIdle, maxConn, user, pwd, dbName, err
}
