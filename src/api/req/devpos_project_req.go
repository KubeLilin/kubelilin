package req

import "github.com/yoyofx/yoyogo/web/mvc"

type NewProject struct {
	mvc.RequestBody
	ProjectId uint64   `json:"project_id"`
	Name      string   `json:"name" `
	TenantID  uint64   `json:"tenant_id"`
	AppIdList []uint64 `json:"appIdList" `
}
