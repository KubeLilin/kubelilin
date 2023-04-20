package dto

import "time"

type ProjectPipelines struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"` // Pipeline ID
	AppId        uint64     `gorm:"column:appid;type:bigint unsigned;not null" json:"appid"`      // 应用ID
	AppName      string     `gorm:"column:appName;type:varchar(50);not null" json:"appName"`
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                        // 流水线名称, appid 下唯一
	Dsl          string     `gorm:"column:dsl;type:text;not null" json:"dsl"`                                 // 流水线DSL
	TaskStatus   *uint      `gorm:"column:taskStatus;type:int unsigned" json:"taskStatus"`                    // 流水线任务状态( ready=0 , running=1, success=2, fail=3,  )
	LastTaskID   string     `gorm:"column:lastTaskId;type:varchar(15);not null;default:''" json:"lastTaskId"` // 最后一次任务执行ID
	Status       uint8      `gorm:"column:status;type:tinyint unsigned;not null" json:"status"`
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
	LastCommit   string     `gorm:"column:last_commit;type:varchar(50);not null;default:''" json:"lastCommit"` // 最后一次提交的commit
}
