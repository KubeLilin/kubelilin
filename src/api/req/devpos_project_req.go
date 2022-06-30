package req

import "github.com/yoyofx/yoyogo/web/mvc"

type CreateNewProject struct {
	mvc.RequestBody
	Name      string  `json:"name" `
	TenantID  int64   `json:"tenant_id"`
	AppIdList []int64 `json:"appIdList" `
}
