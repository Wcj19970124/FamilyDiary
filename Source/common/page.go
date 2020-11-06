package common

import "math"

//Page 分页结构体
type Page struct {
	PageNo     int         //当前页码
	PageSize   int         //每页的数据条数
	TotalCount int         //数据总条数
	TotalPage  int         //总页数
	StartNo    int         //DB查询时limit开始的数据
	List       interface{} //数据
}

//SetTotalPage 设置总页数
func (p *Page) SetTotalPage() {
	totalPage := math.Ceil(float64(p.TotalCount) / float64(p.PageSize))
	p.TotalPage = int(totalPage)
}

//SetStartNo 设置数据库查询时的起始数据索引
func (p *Page) SetStartNo() {
	startNo := p.PageNo * p.PageSize
	p.StartNo = startNo
}
