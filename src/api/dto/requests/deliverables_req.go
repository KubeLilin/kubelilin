package requests

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type CreateTenantDeliverablesProjectReq struct {
	mvc.RequestBody
	Id                  uint64 `json:"id"`
	TenantId            uint64 `json:"tenantId"`
	ProjectName         string `json:"projectName"`
	ServiceConnectionId uint64 `json:"serviceConnectionId"`
}

type QueryTenantDeliverablesProjectReq struct {
	mvc.RequestBody
	page.PageRequest
	ProjectName string `json:"projectName" uri:"projectName" `
	TenantId    uint64 `json:"tenantId"`
}

type EditTenantDeliverablesTreeReq struct {
	mvc.RequestBody
	Id        uint64 `json:"id"`
	TenantId  uint64 `json:"tenantId"`
	Name      string `json:"name"`
	ProjectId uint64 `json:"projectId"`
	ParentID  uint64 `json:"parentId"`
}
