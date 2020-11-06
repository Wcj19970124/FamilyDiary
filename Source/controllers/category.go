package controllers

import (
	"encoding/json"
	"time"

	"../common"
	"../models"
)

//CategoryController 类别数据体
type CategoryController struct {
	BaseController
}

//类别操作错误状态码
const (
	CategoryAddErr        = 30001 //类别添加失败
	CategoryDelErr        = 30002 //类别删除失败
	CategoryUpdateErr     = 30003 //类别更新失败
	CategoryQueryErr      = 30004 //类别查询失败
	CategoryAlreadyExists = 30006 //类别已存在
)

//登录态和权限校验
func (t *CategoryController) verificate() bool {
	return t.Verificate()
}

//AddCategory 添加类别
func (t *CategoryController) AddCategory() {
	//基础校验
	if ok := t.verificate(); !ok {
		return
	}
	//TODD：添加类别
	var category models.Category
	if json.Unmarshal(t.Ctx.Input.RequestBody, &category) == nil {
		if ok := models.GetCategoryByCategoryName(category.CategoryName); !ok {
			t.Report(t.Ctx.Input.IP(), "1", "POST", "1", "AddCategory", models.GetLoginAdminUserName(), "类别已存在", time.Now())
			t.OutPut(CategoryAlreadyExists, "类别已存在!")
			return
		}
		if err := models.AddCategory(category); err != nil {
			t.Report(t.Ctx.Input.IP(), "1", "POST", "1", "AddCategory", models.GetLoginAdminUserName(), "类别添加失败", time.Now())
			t.OutPut(CategoryAddErr, "类别添加失败!")
			return
		}
		t.Report(t.Ctx.Input.IP(), "1", "POST", "0", "AddCategory", models.GetLoginAdminUserName(), "类别添加成功", time.Now())
		t.OutPut(200, "类别添加成功!")
	}
}

//DelCategory 删除类别
func (t *CategoryController) DelCategory() {
	//基础校验
	if ok := t.verificate(); !ok {
		return
	}
	//TODD:删除
	if id, err := t.GetInt("id"); err == nil {
		if err := models.DelCategory(id); err != nil {
			t.Report(t.Ctx.Input.IP(), "1", "DELETE", "1", "DelCategory", models.GetLoginAdminUserName(), "类别删除失败", time.Now())
			t.OutPut(CategoryDelErr, "类别删除失败!")
			return
		}
		t.Report(t.Ctx.Input.IP(), "1", "DELETE", "0", "DelCategory", models.GetLoginAdminUserName(), "类别删除成功", time.Now())
		t.OutPut(200, "类别删除成功!")
	}
}

//UpdateCategory 更新类别
func (t *CategoryController) UpdateCategory() {
	//基础校验
	if ok := t.verificate(); !ok {
		return
	}
	//TODD:更新
	var category models.Category
	if json.Unmarshal(t.Ctx.Input.RequestBody, &category) == nil {
		if ok := models.GetCategoryByCategoryName(category.CategoryName); !ok {
			t.Report(t.Ctx.Input.IP(), "1", "PUT", "1", "UpdateCategory", models.GetLoginAdminUserName(), "类别已存在", time.Now())
			t.OutPut(CategoryAlreadyExists, "类别已存在!")
			return
		}
		if err := models.UpdateCategory(category); err != nil {
			t.Report(t.Ctx.Input.IP(), "1", "PUT", "1", "UpdateCategory", models.GetLoginAdminUserName(), "类别更新失败", time.Now())
			t.OutPut(CategoryUpdateErr, "类别更新失败!")
			return
		}
		t.Report(t.Ctx.Input.IP(), "1", "PUT", "0", "UpdateCategory", models.GetLoginAdminUserName(), "类别更新成功", time.Now())
		t.OutPut(200, "类别更新成功!")
	}
}

//QueryCategory 单条类别查询
func (t *CategoryController) QueryCategory() {
	//基础校验
	if ok := t.verificate(); !ok {
		return
	}
	//TODD:查询
	if id, err := t.GetInt("id"); err == nil {
		data, err := models.QueryCategory(id)
		if err != nil {
			t.Report(t.Ctx.Input.IP(), "1", "GET", "1", "QueryCategory", models.GetLoginAdminUserName(), "类别查询失败", time.Now())
			t.OutPut(CategoryQueryErr, "类别查询失败!")
		}
		t.Report(t.Ctx.Input.IP(), "1", "GET", "0", "QueryCategory", models.GetLoginAdminUserName(), "类别查询成功", time.Now())
		t.OutPutList(200, "类别查询成功!", data)
	}
}

//QueryCategories 类别列表查询
func (t *CategoryController) QueryCategories() {
	//基础校验
	if ok := t.verificate(); !ok {
		return
	}
	//TODD:查询
	var page common.Page
	if json.Unmarshal(t.Ctx.Input.RequestBody, &page) == nil {
		data, err := models.QueryCatrgories(page)
		if err != nil {
			t.Report(t.Ctx.Input.IP(), "1", "POST", "1", "QueryCategories", models.GetLoginAdminUserName(), "类别列表查询失败", time.Now())
			t.OutPut(CategoryQueryErr, "类别列表查询失败!")
			return
		}
		t.Report(t.Ctx.Input.IP(), "1", "POST", "0", "QueryCategories", models.GetLoginAdminUserName(), "类别列表查询成功", time.Now())
		t.OutPutList(200, "类别列表查询成功!", data)
	}
}

//QueryCategoriesTree 类别树查询
func (t *CategoryController) QueryCategoriesTree() {
	//基础校验
	if ok := t.verificate(); !ok {
		return
	}
	//TODD:类别树查询
	data, err := models.QueryCategoriesTree()
	if err != nil {
		t.Report(t.Ctx.Input.IP(), "1", "GET", "1", "QueryCategoriesTree", models.GetLoginAdminUserName(), "类别树查询失败", time.Now())
		t.OutPut(CategoryQueryErr, "类别树查询失败!")
		return
	}
	t.Report(t.Ctx.Input.IP(), "1", "GET", "0", "QueryCategoriesTree", models.GetLoginAdminUserName(), "类别树查询成功", time.Now())
	t.OutPutList(200, "类别树查询成功!", data)
}
