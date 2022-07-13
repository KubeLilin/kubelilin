package dto

import "time"

type DevOpsProjectsDTO struct {
	ID         uint64     `gorm:"column:id;" json:"id"`
	Name       string     `gorm:"column:project_name;" json:"name" `                     // 项目名称
	CreateTime *time.Time `gorm:"column:creation_time;type:datetime" json:"createTime" ` // 创建时间
	AppCount   uint64     `gorm:"column:app_count;" json:"appCount" `
	AppIdList  string     `gorm:"column:app_ids;" json:"appList"`
}
