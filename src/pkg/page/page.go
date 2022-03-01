package page

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Page struct {
	PageIndex int
	PageSize  int
	Total     int64
	Data      interface{}
}

type PageRequest struct {
	PageIndex   int    `json:"pageIndex" uri:"pageIndex"`
	PageSize    int    `json:"pageSize" uri:"pageSize"`
	CurrentPage int    `json:"current" uri:"current"`
	Keyword     string `json:"keyword" uri:"keyword"`
}

func (pg PageRequest) OffSet() int {
	return (pg.PageIndex - 1) * pg.PageSize
}

func InitPageByCtx(ctx *context.HttpContext) *PageRequest {

	pageIndex, _ := strconv.ParseInt(ctx.Input.Query("current"), 10, 32)
	paagSize, _ := strconv.ParseInt(ctx.Input.Query("pageSize"), 10, 32)
	return &PageRequest{
		PageIndex: int(pageIndex),
		PageSize:  int(paagSize),
	}
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
func (ph *PageHelper) DoFind(data interface{}) *Page {
	var count int64
	ph.db.Count(&count)
	ph.db.Offset(ph.pageInfo.OffSet()).Limit(ph.pageInfo.PageSize).Find(data)
	return &Page{
		Data:      data,
		Total:     count,
		PageIndex: ph.pageInfo.PageIndex,
		PageSize:  ph.pageInfo.PageSize,
	}
}

func (ph *PageHelper) DoScan(data interface{}, sql string, values ...interface{}) (error, *Page) {
	var count int64
	var countSql string
	var dataSql string
	var err error
	countSql = "select count(0) from (" + sql + ") as c1"
	var sb strings.Builder
	sb.WriteString(sql)
	sb.WriteString(" limit ")
	sb.WriteString(strconv.Itoa(ph.pageInfo.OffSet()))
	sb.WriteString(",")
	sb.WriteString(strconv.Itoa(ph.pageInfo.PageSize))
	dataSql = sb.String()
	fmt.Println(dataSql)
	countRes := ph.db.Raw(countSql, values...).Scan(&count)
	dataRes := ph.db.Raw(dataSql, values...).Scan(data)
	if countRes.Error != nil {
		err = countRes.Error
	}
	if dataRes.Error != nil {
		err = dataRes.Error
	}
	return err, &Page{
		Data:      data,
		Total:     count,
		PageIndex: ph.pageInfo.PageIndex,
		PageSize:  ph.pageInfo.PageSize,
	}
}
