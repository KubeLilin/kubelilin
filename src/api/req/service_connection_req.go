package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type ServiceConnectionReq struct {
	ID          uint64 `json:"id"`
	TenantID    uint64 `json:"tenantId"`                                               // 租户id
	Name        string `json:"name"`                                                   // 连接名称
	ServiceType int    `json:"serviceType"`                                            // 连接类型 1凭证 2连接
	Type        int    `gorm:"column:type;type:int;not null" json:"type"`              // 凭证类型 1.github 2..gitlab 3.gogos 4.gitee
	Detail      string `gorm:"column:detail;type:varchar(500);not null" json:"detail"` // 凭证信息/连接信息
}

type ServiceConnectionPageReq struct {
	mvc.RequestBody
	Name     string `json:"name" uri:"name"`
	TenantID uint64 `json:"tenantId" uri:"tenantId"`
	page.Page
}
