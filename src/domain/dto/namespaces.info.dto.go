package dto

type NamespaceInfo struct {
	Id          uint64 `json:"id" gorm:"column:id"`
	TenantId    uint64 `json:"tenantId" gorm:"column:tenantId;"`
	ClusterId   uint64 `json:"clusterId" gorm:"column:clusterId;"`
	TenantCode  string `json:"tenantCode" gorm:"column:tenantCode;"`
	ClusterName string `json:"clusterName" gorm:"column:clusterName;"`
	Namespace   string `json:"namespace" gorm:"column:namespace"`
	TenantName  string `json:"tenantName" gorm:"column:tenantName"`
}
