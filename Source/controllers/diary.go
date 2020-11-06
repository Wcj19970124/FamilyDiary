package controllers

import (
	"encoding/json"
	"time"

	"../common"
	"../models"
)

//DiaryController 帖子数据体
type DiaryController struct {
	BaseController
}

//帖子操作错误状态码
const (
	DiaryAddErr    = 30001 //帖子添加失败
	DiaryDelErr    = 30002 //帖子删除失败
	DiaryUpdateErr = 30003 //帖子更新失败
	DiaryQueryErr  = 30004 //帖子查询失败
)

//登录态和权限校验
func (d *DiaryController) verificate() bool {
	return d.Verificate()
}

//AddDiary 添加帖子
func (d *DiaryController) AddDiary() {
	//基础校验
	if ok := d.verificate(); !ok {
		return
	}
	//TODD：添加
	var diary models.Diary
	if json.Unmarshal(d.Ctx.Input.RequestBody, &diary) == nil {
		if err := models.AddDiary(diary); err != nil {
			d.Report(d.Ctx.Input.IP(), "1", "POST", "1", "AddDiary", models.GetLoginAdminUserName(), "帖子添加失败", time.Now())
			d.OutPut(DiaryAddErr, "帖子添加失败!")
			return
		}
		d.Report(d.Ctx.Input.IP(), "1", "POST", "0", "AddDiary", models.GetLoginAdminUserName(), "帖子添加成功", time.Now())
		d.OutPut(200, "帖子添加成功!")
	}

}

//DelDiary 删除帖子
func (d *DiaryController) DelDiary() {
	//基础校验
	if ok := d.verificate(); !ok {
		return
	}
	//TODD:删除
	if id, err := d.GetInt("id"); err == nil {
		if err := models.DelDiary(id); err != nil {
			d.Report(d.Ctx.Input.IP(), "1", "DELETE", "1", "DelDiary", models.GetLoginAdminUserName(), "帖子删除失败", time.Now())
			d.OutPut(DiaryDelErr, "帖子删除失败!")
			return
		}
		d.Report(d.Ctx.Input.IP(), "1", "DELETE", "0", "DelDiary", models.GetLoginAdminUserName(), "帖子删除成功", time.Now())
		d.OutPut(200, "帖子删除成功!")
	}
}

//UpdateDiary 更新帖子
func (d *DiaryController) UpdateDiary() {
	//基础校验
	if ok := d.verificate(); !ok {
		return
	}
	//TODD:更新
	var diary models.Diary
	if json.Unmarshal(d.Ctx.Input.RequestBody, &diary) == nil {
		if err := models.UpdateDiary(diary); err != nil {
			d.Report(d.Ctx.Input.IP(), "1", "DELETE", "1", "UpdateDiary", models.GetLoginAdminUserName(), "帖子更新失败", time.Now())
			d.OutPut(DiaryUpdateErr, "帖子更新失败!")
			return
		}
		d.Report(d.Ctx.Input.IP(), "1", "DELETE", "0", "UpdateDiary", models.GetLoginAdminUserName(), "帖子更新成功", time.Now())
		d.OutPut(200, "帖子更新成功!")
	}
}

//QueryDiary 单条帖子查询
func (d *DiaryController) QueryDiary() {
	//基础校验
	if ok := d.verificate(); !ok {
		return
	}
	//TODD:查询
	if id, err := d.GetInt("id"); err == nil {
		data, err := models.QueryDiary(id)
		if err != nil {
			d.Report(d.Ctx.Input.IP(), "1", "GET", "1", "QueryDiary", models.GetLoginAdminUserName(), "帖子查询失败", time.Now())
			d.OutPut(DiaryQueryErr, "帖子查询失败!")
			return
		}
		d.Report(d.Ctx.Input.IP(), "1", "GET", "0", "QueryDiary", models.GetLoginAdminUserName(), "帖子查询成功", time.Now())
		d.OutPutList(200, "帖子查询成功!", data)
	}
}

//QueryDiaries 帖子列表查询
func (d *DiaryController) QueryDiaries() {
	//基础校验
	if ok := d.verificate(); !ok {
		return
	}
	//TODD:查询
	var page common.Page
	if json.Unmarshal(d.Ctx.Input.RequestBody, &page) == nil {
		data, err := models.QueryDiaries(page)
		if err != nil {
			d.Report(d.Ctx.Input.IP(), "1", "POST", "1", "QueryDiaries", models.GetLoginAdminUserName(), "帖子列表查询失败", time.Now())
			d.OutPut(DiaryQueryErr, "帖子列表查询失败!")
			return
		}
		d.Report(d.Ctx.Input.IP(), "1", "POST", "0", "QueryDiaries", models.GetLoginAdminUserName(), "帖子列表查询成功", time.Now())
		d.OutPutList(200, "帖子列表查询成功!", data)
	}
}
