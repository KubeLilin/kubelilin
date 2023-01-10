package requests

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type ConfigMapPageReq struct {
	mvc.RequestBody
	page.PageRequest

	ClusterID   uint64 `json:"clusterId" uri:"clusterId"` // 集群id
	NamespaceID uint64 `json:"namespaceId" uri:"namespaceId"`
	AppId       uint64 `json:"appId" uri:"appId"`
	Name        string `json:"name" uri:"name"`
}
