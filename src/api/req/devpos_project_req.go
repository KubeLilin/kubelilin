package req

import "github.com/yoyofx/yoyogo/web/mvc"

type CreateNewProject struct {
	mvc.RequestBody
	Name      string   `json:"name" `
	TenantID  uint64   `json:"tenant_id"`
	AppIdList []uint64 `json:"appIdList" `
}
