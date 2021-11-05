package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
)

type ImportClusterReq struct {
	mvc.RequestBody
	NickName string `form:"nickName"`
	TenantId uint64 `form:"tenantId"`
}
