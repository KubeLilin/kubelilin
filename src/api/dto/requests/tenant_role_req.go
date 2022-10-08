package requests

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type TenantRoleReq struct {
	mvc.RequestBody
	ID       uint64 `json:"id" uri:"id"`
	RoleCode string `json:"roleCode" uri:"roleCode"` // 角色编码
	RoleName string `json:"roleName" uri:"roleName"` // 角色名称
	Status   int8   `json:"status" uri:"status"`     // 状态
	TenantID int64  `json:"tenantId" uri:"tenantId"`
	page.PageRequest
}
