package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type ConfigMapPageReq struct {
	mvc.RequestBody
	page.PageRequest

	ConfigName string `json:"configName" uri:"configName"`
}
