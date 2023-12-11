package requests

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type CreateTenantArtifactsProjectReq struct {
	mvc.RequestBody
	Id              uint64 `json:"id"`
	TenantId        uint64 `json:"tenantId"`
	ProjectName     string `json:"projectName"`
	HarborProjectId uint64 `json:"projectId"`
}

type QueryTenantArtifactsProjectReq struct {
	mvc.RequestBody
	page.PageRequest
	ProjectName string `json:"projectName" uri:"projectName" `
	TenantId    uint64 `json:"tenantId"`
}
