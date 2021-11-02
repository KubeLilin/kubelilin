package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"mime/multipart"
)

type ImportClusterReq struct {
	mvc.RequestBody
	File     *multipart.FileHeader `form:"file1"`
	NickName string                `form:"nickName"`
	TenantId uint64                `form:"tenantId"`
}
