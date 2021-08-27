package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/pkg/page"
)

type TenantRequest struct {
	mvc.RequestBody
	//properties
	ID     uint64 `json:"id"`
	TName  string `json:"tName"`
	TCode  string `json:"tCode"`  // 租户编码
	Status int8   `json:"status"` // 状态

	page.PageRequest
}
