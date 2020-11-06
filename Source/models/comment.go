package models

import (
	"time"

	"../common"
	"github.com/astaxie/beego/logs"
)

//Comment 后台评论数据体
type Comment struct {
	Id         int
	TouristId  int
	DiaryId    int
	Content    string
	ParentId   int
	Goods      int
	Elite      string
	Top        string
	CreateUser string
	UpdateUser string
	Status     string
	CreateTime time.Time
	UpdateTime time.Time
}

//DelComment 删除评论
func DelComment(id int) error {

	sql := "update fd_diary_comment set status = 1 where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, id).Exec()
	if err != nil {
		logs.Error("---- delete comment failed,err:" + err.Error())
		return err
	}

	return nil
}

//QueryComments 分页查询评论列表
func QueryComments(p common.Page) (map[string]interface{}, error) {

	sql1 := "select count(id) from fd_diary_comment"
	sql2 := "select id,tourist_id,diary_id,content,parent_id,goods,elite,top,status from fd_diary_comment limit ?,?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var comments []Comment
	p.SetStartNo() //设置搜索的起始索引
	_, err = dbProxy.Raw(sql2, p.StartNo, p.PageSize).QueryRows(&comments)
	if err != nil {
		logs.Error("---- get comment(page) failed,err:" + err.Error())
		return nil, err
	}

	err = dbProxy.Raw(sql1).QueryRow(&p.TotalCount)
	if err != nil {
		logs.Error("---- get comment totalCount failed,err:" + err.Error())
		return nil, err
	}

	p.SetTotalPage() //设置总页数
	p.List = comments
	m := make(map[string]interface{})
	m["page"] = p

	return m, nil
}
