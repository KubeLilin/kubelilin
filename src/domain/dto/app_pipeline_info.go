package dto

type PipelineInfo struct {
	Id         uint64 `gorm:"column:id;" json:"id"`
	AppId      uint64 `gorm:"column:appid;" json:"appid"`
	Name       string `gorm:"column:name;" json:"name"`
	DSL        string `gorm:"column:dsl;" json:"dsl"`
	Status     uint32 `gorm:"column:taskStatus;" json:"status"`
	LastTaskID string `gorm:"column:lastTaskId;" json:"taskid"`
}
