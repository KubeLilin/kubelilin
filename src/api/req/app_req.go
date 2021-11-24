package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/pkg/page"
)

type AppReq struct {
	mvc.RequestBody
	page.PageRequest
	ID       uint64 `json:"id" uri:"id"`
	TenantId uint64
	Name     string `json:"name" uri:"name"`
	Labels   string `json:"labels" uri:"labels"`
	Remarks  string `json:"remarks" uri:"remarks"`
	Git      string `json:"git" uri:"git"`
	Level    uint16 `json:"level" uri:"level"`
	Language uint16 `json:"language" uri:"language"`
	Status   int8   `json:"status" uri:"status"`
}