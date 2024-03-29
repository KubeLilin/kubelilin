package dto

type NamespaceInfo struct {
	Id            uint64               `json:"id" gorm:"column:id"`
	TenantId      uint64               `json:"tenantId" gorm:"column:tenantId;"`
	ClusterId     uint64               `json:"clusterId" gorm:"column:clusterId;"`
	TenantCode    string               `json:"tenantCode" gorm:"column:tenantCode;"`
	ClusterName   string               `json:"clusterName" gorm:"column:clusterName;"`
	Namespace     string               `json:"namespace" gorm:"column:namespace"`
	TenantName    string               `json:"tenantName" gorm:"column:tenantName"`
	EnableRuntime bool                 `json:"enableRuntime" gorm:"column:enableRuntime"`
	RuntimeName   string               `json:"runtimeName" gorm:"column:runtimeName"`
	DeployCount   uint64               `json:"deployCount" gorm:"column:deployCount"`
	InsCount      uint64               `json:"insCount" gorm:"column:insCount"`
	QuotasSpec    []ResourceQuotasItem `json:"quotasSpec" gorm:"-"`
}
