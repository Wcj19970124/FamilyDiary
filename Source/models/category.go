package models

import (
	"strconv"
	"time"

	"../common"
	"github.com/astaxie/beego/logs"
)

//Category 后台帖子类别结构体
type Category struct {
	Id            int
	CategoryName  string
	ParentId      int
	Type          string
	Description   string
	CreateUser    string
	UpdateUser    string
	Status        string
	CreateTime    time.Time
	UpdateTime    time.Time
	ChildCategory []*Category
}

//GetCategoryByCategoryName 根据类别名判断类名是否已存在
func GetCategoryByCategoryName(categoryname string) bool {

	sql := "select id from fd_category where category_name = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return false
	}

	var category Category
	err = dbProxy.Raw(sql, categoryname).QueryRow(&category)
	if err == nil && strconv.Itoa(category.Id) != "" {
		return false
	}

	return true
}

//AddCategory 添加类别
func AddCategory(category Category) error {

	sql := "insert into fd_category(category_name,parent_id,type,description,create_user,create_time,update_user,update_time,status) values(?,?,?,?,?,?,?,?,?)"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, category.CategoryName, category.ParentId, category.Type, category.Description, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0).Exec()
	if err != nil {
		logs.Error("---- insert category failed,err:" + err.Error())
		return err
	}

	return nil
}

//DelCategory 删除类别
func DelCategory(id int) error {

	sql := "update fd_category set status = 1 where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, id).Exec()
	if err != nil {
		logs.Error("---- delete category failed,err:" + err.Error())
		return err
	}

	return nil
}

//UpdateCategory 更新类别
func UpdateCategory(category Category) error {

	sql := "update fd_category set category_name = ?,parent_id = ?,type=?,description=?,create_user=?,create_time=?,update_user=?,update_time=?,status=? where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return err
	}

	_, err = dbProxy.Raw(sql, category.CategoryName, category.ParentId, category.Type, category.Description, GetLoginAdminUserName(), time.Now(), GetLoginAdminUserName(), time.Now(), 0, category.Id).Exec()
	if err != nil {
		logs.Error("---- update category failed,err:" + err.Error())
		return err
	}

	return nil
}

//QueryCategory 查询单条类别
func QueryCategory(id int) (map[string]interface{}, error) {

	sql := "select category_name,parent_id,type,description from fd_category where id = ?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var category Category
	err = dbProxy.Raw(sql, id).QueryRow(&category)
	if err != nil {
		logs.Error("---- query category failed,err:" + err.Error())
		return nil, err
	}

	m := make(map[string]interface{})
	m["category"] = category

	return m, nil
}

//QueryCatrgories 类别列表分页查询
func QueryCatrgories(p common.Page) (map[string]interface{}, error) {

	sql1 := "select count(id) from fd_category"
	sql2 := "select id,category_name,parent_id,type,description,create_user,create_time,update_user,update_time,status from fd_category limit ?,?"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var categories []Category
	p.SetStartNo() //设置搜索的起始数据索引
	_, err = dbProxy.Raw(sql2, p.StartNo, p.PageSize).QueryRows(&categories)
	if err != nil {
		logs.Error("---- query category(page) failed,err:" + err.Error())
		return nil, err
	}

	err = dbProxy.Raw(sql1).QueryRow(&p.TotalCount)
	if err != nil {
		logs.Error("---- query category totalCount failed,err:" + err.Error())
		return nil, err
	}

	p.SetTotalPage() //设置总页数
	p.List = categories
	m := make(map[string]interface{})
	m["page"] = p

	return m, nil
}

//QueryCategoriesTree 查询类别树
func QueryCategoriesTree() (map[string]interface{}, error) {

	sql := "select id,category_name,parent_id,type,description,status from fd_category"
	dbProxy, err := store.GetDBProxy()
	if err != nil {
		logs.Error("---- get db proxy failed,err:" + err.Error() + " ----")
		return nil, err
	}

	var categories []*Category
	_, err = dbProxy.Raw(sql).QueryRows(&categories)
	if err != nil {
		logs.Error("---- query categoriesTree failed,err:" + err.Error())
		return nil, err
	}

	m := convertCategoriesToCategoriesTree(categories)
	return m, nil
}

//将切片形式数据转换为树状形式
func convertCategoriesToCategoriesTree(categories []*Category) map[string]interface{} {

	root := make(map[string]interface{})
	parentCategories := []*Category{}

	//获取根级类别
	for _, category := range categories {
		if category.ParentId == 0 {
			parentCategories = append(parentCategories, category)
		}
	}

	convertChildCategoriesToCategoriesTree(parentCategories, categories)
	root["categoriesTree"] = parentCategories

	return root
}

//convertChildCategoriesToCategoriesTree 将子类别转换为类别树
func convertChildCategoriesToCategoriesTree(parentCategories []*Category, categories []*Category) {

	if len(parentCategories) == 0 {
		return
	}

	//添加子类别进入类别树
	for _, parentCategory := range parentCategories {
		for _, category := range categories {
			if category.ParentId == parentCategory.Id {
				parentCategory.ChildCategory = append(parentCategory.ChildCategory, category)
			}
		}
	}

	//递归添加
	for _, childCategory := range parentCategories {
		convertChildCategoriesToCategoriesTree(childCategory.ChildCategory, categories)
	}
}
