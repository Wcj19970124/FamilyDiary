package controllers

import (
	"encoding/json"
	"time"

	"../common"
	"../models"
)

//CommentController 评论数据体
type CommentController struct {
	BaseController
}

//评论操作错误状态码
const (
	CommentDelErr   = 30002
	CommentQueryErr = 30004
)

//登录态和权限校验
func (c *CommentController) verificate() bool {
	return c.Verificate()
}

//DelComment 删除评论
func (c *CommentController) DelComment() {
	//基础校验
	if ok := c.verificate(); !ok {
		return
	}
	//TODD：删除
	if id, err := c.GetInt("id"); err == nil {
		if err := models.DelComment(id); err != nil {
			c.Report(c.Ctx.Input.IP(), "1", "DELETE", "1", "DelComment", models.GetLoginAdminUserName(), "评论删除失败", time.Now())
			c.OutPut(CommentDelErr, "评论删除失败!")
			return
		}
		c.Report(c.Ctx.Input.IP(), "1", "DELETE", "0", "DelComment", models.GetLoginAdminUserName(), "评论删除成功", time.Now())
		c.OutPut(200, "评论删除成功!")
	}
}

//QueryComments 分页查询评论
func (c *CommentController) QueryComments() {
	//基础校验
	if ok := c.verificate(); !ok {
		return
	}
	//查询
	var page common.Page
	if json.Unmarshal(c.Ctx.Input.RequestBody, &page) == nil {
		data, err := models.QueryComments(page)
		if err != nil {
			c.Report(c.Ctx.Input.IP(), "1", "POST", "1", "QueryComments", models.GetLoginAdminUserName(), "评论列表查询失败", time.Now())
			c.OutPut(CommentQueryErr, "评论列表查询失败!")
			return
		}
		c.Report(c.Ctx.Input.IP(), "1", "POST", "0", "QueryComments", models.GetLoginAdminUserName(), "评论列表查询成功", time.Now())
		c.OutPutList(200, "评论列表查询成功!", data)
	}
}
