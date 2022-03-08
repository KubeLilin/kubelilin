package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type TenantRequest struct {
	*mvc.RequestBody
	//properties
	ID     uint64 `json:"id"`
	TName  string `json:"tName" uri:"tName"`
	TCode  string `json:"tCode" uri:"tCode"`   // 租户编码
	Status *int8  `json:"status" uri:"status"` // 状态

	page.PageRequest
}
