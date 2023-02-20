package dal

import (
	"gorm.io/gorm"
	"time"
)

/*
 @Author: 71made
 @Date: 2023/02/17 23:20
 @ProductName: page.go
 @Description:
*/

type Page struct {
	Limit     int       // 查询记录数
	Offset    int       // 查询起始位置
	StartTime time.Time // 查询记录最晚创建/更新时间
}

func (p *Page) Exec(db *gorm.DB) *gorm.DB {
	tmp := db.Offset(p.Offset)
	if p.Limit != 0 {
		tmp = tmp.Limit(p.Limit)
	}
	if p.StartTime != defaultTime {
		tmp = tmp.Where("created_at >= ?", p.StartTime)
	}
	return tmp
}

var defaultTime time.Time

func DefaultPage() *Page {
	return &Page{
		Limit:     0,
		Offset:    0,
		StartTime: defaultTime,
	}
}

type PageOption func(page *Page)

func PageLimit(limit int) PageOption {
	return func(page *Page) {
		page.Limit = limit
	}
}

func PageOffset(offset int) PageOption {
	return func(page *Page) {
		page.Offset = offset
	}
}

func PageAfter(startTime interface{}) PageOption {
	return func(page *Page) {
		switch startTime.(type) {
		case time.Time:
			page.StartTime = startTime.(time.Time)
		case int64:
			page.StartTime = time.Unix(startTime.(int64), 0)
		default:
			page.StartTime = defaultTime
		}
	}
}
