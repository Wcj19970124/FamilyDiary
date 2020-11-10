package models

import (
	"time"

	"../common"
	"github.com/astaxie/beego/logs"
)

//Diary 后台帖子结构体
type Diary struct {
	Id         int
	Title      string
	Content    string
	CategoryId int
	Comments   int
	Goods      int
	Glance     int
	CreatUser  string
	UpdateUser string
	Status     string
	CreateTime time.Time
	UpdateTime time.Time
}

//AddDiary 添加帖子
func AddDiary(diary Diary) error {

	sql := "insert into fd_diary(title,content,category_id,comments,goods,glance,create_user,create_time,update_user,update_time,status) values(?,?,?,?,?,?,?,?,?,?,?)"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, diary.Title, diary.Content, diary.CategoryId, 0, 0, 0, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0).Exec()
	if err != nil {
		logs.Error("---- insert diary failed,err:" + err.Error())
		return err
	}

	return nil
}

//DelDiary 删除帖子
func DelDiary(id int) error {

	sql := "update fd_diary set status = 1 where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, id).Exec()
	if err != nil {
		logs.Error("---- delete diary failed,err:" + err.Error())
		return err
	}

	return nil
}

//UpdateDiary 更新帖子
func UpdateDiary(diary Diary) error {

	sql := "update fd_diary set title = ?,content=?,category_id=?,update_user = ?,update_time=? where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, diary.Title, diary.Content, diary.CategoryId, GetLoginAdminUserName(), time.Now(), diary.Id).Exec()
	if err != nil {
		logs.Error("---- update diary failed,err:" + err.Error())
		return err
	}

	return nil
}

//QueryDiary 根据id查询帖子
func QueryDiary(id int) (map[string]interface{}, error) {

	sql := "select id,title,content,category_id,comments,goods,glance,status from fd_diary where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var diary Diary
	err = dbProxy.Raw(sql, id).QueryRow(&diary)
	if err != nil {
		logs.Error("---- get diary by id failed,err:" + err.Error())
		return nil, err
	}

	m := make(map[string]interface{})
	m["diary"] = diary

	return m, nil
}

//QueryDiaries 分页查询帖子列表
func QueryDiaries(p common.Page) (map[string]interface{}, error) {

	sql1 := "select count(id) from fd_diary"
	sql2 := "select id,title,content,category_id,comments,goods,glance,create_user,create_time,update_user,update_time,status from fd_diary limit ?,?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var diaries []Diary
	p.SetStartNo() //设置起始所搜索引
	_, err = dbProxy.Raw(sql2, p.StartNo, p.PageSize).QueryRows(&diaries)
	if err != nil {
		logs.Error("---- get diary(page) failed,err:" + err.Error())
		return nil, err
	}

	err = dbProxy.Raw(sql1).QueryRow(&p.TotalCount)
	if err != nil {
		logs.Error("---- get diary totalCount failed,err:" + err.Error())
		return nil, err
	}

	p.SetTotalPage() //设置总页数
	p.List = diaries
	m := make(map[string]interface{})
	m["page"] = p

	return m, nil
}
