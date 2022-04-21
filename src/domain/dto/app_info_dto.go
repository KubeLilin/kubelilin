package dto

import "time"

type ApplicationInfoDTO struct {
	Id           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	Name         string     `gorm:"column:name;type:varchar(50);not null" `           // 集群应用名称(英文唯一)
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint(11);not null" `       // 租户
	Labels       string     `gorm:"column:labels;type:varchar(50);not null" `         // 应用中文名称
	Remarks      string     `gorm:"column:remarks;type:varchar(200);not null" `       // 集群应用备注
	Git          string     `gorm:"column:git;type:varchar(500);not null" `           // 集群应用绑定的git地址
	Level        uint16     `gorm:"column:level;type:smallint(6) unsigned;not null" ` // 应用级别
	LevelName    string     `gorm:"column:level_name;type:smallint(6) unsigned;not null" `
	Language     uint16     `gorm:"column:language;type:smallint(5) unsigned;not null" ` // 开发语言
	LanguageName string     `gorm:"column:language_name;type:smallint(5) unsigned;not null" `
	Status       int8       `gorm:"column:status;type:tinyint(4);not null;default:0" ` // 状态
	CreateTime   *time.Time `gorm:"column:create_time;type:datetime" `                 // 创建时间
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" `                 // 更新时间
	SourceType   string     `gorm:"column:git_type" `
	SCID         uint64     `gorm:"column:sc_id"`
}

type ApplicationDisplayDTO struct {
	Name       string `gorm:"column:appName;" `                         // 集群应用名称(英文唯一)
	TenantName string `gorm:"column:tenantName;" `                      // 租户
	Labels     string `gorm:"column:labels;type:varchar(50);not null" ` // 应用中文名称
	Git        string `gorm:"column:git;type:varchar(500);not null" `   // 集群应用绑定的git地址
	Hub        string `gorm:"column:hub;"`
	Level      string `gorm:"column:level;" `
	Language   string `gorm:"column:language;" `                                 // 开发语言
	Status     int8   `gorm:"column:status;type:tinyint(4);not null;default:0" ` // 状态
	SourceType string `gorm:"column:git_type" `
	SCID       uint64 `gorm:"column:sc_id"`
}
