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

type CreateConfigMapReq struct {
	ID          int64  `json:"id" uri:"id"`
	Name        string `json:"name" uri:"name"`
	TenantID    uint64 `json:"tenantId" uri:"tenantId"`
	ClusterID   uint64 `json:"clusterId" uri:"clusterId"` // 集群id
	NamespaceID uint64 `json:"namespaceId" uri:"namespaceId"`
	Type        string `json:"type" uri:"type"` // 文件类型，yaml、properties
	Data        string `json:"data" uri:"data"` // 配置内容
}
