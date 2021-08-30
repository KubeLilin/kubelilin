package page

import "gorm.io/gorm"

type Page struct {
	PageIndex int
	PageSize  int
	Total     int64
	Data      interface{}
}

type PageRequest struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

func (pg PageRequest) OffSet() int {
	return (pg.PageIndex - 1) * pg.PageSize
}

type PageHelper struct {
	pageInfo *PageRequest
	db       *gorm.DB
}

/**
传入设置好where条件的db开始分页
*/
func StartPage(db *gorm.DB, pageIndex, pageSize int) *PageHelper {
	ph := &PageHelper{
		pageInfo: &PageRequest{PageIndex: pageIndex, PageSize: pageSize},
		db:       db,
	}
	return ph
}

/**
执行分页进行赋值
*/
func (ph *PageHelper) DoSelect(data interface{}) *Page {
	var count int64
	ph.db.Offset(ph.pageInfo.OffSet()).Limit(ph.pageInfo.PageSize).Find(data)
	ph.db.Count(&count)
	return &Page{
		Data:      data,
		Total:     count,
		PageIndex: ph.pageInfo.PageIndex,
		PageSize:  ph.pageInfo.PageSize,
	}
}
