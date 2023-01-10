package dto

import "time"

type AppConfigmapInfo struct {
	Id             string    `json:"id" gorm:"column:id"`
	Name           string    `json:"name" gorm:"column:name"`
	ClusterName    string    `json:"cluster" gorm:"column:cluster"`
	NamespaceName  string    `json:"namespace" gorm:"column:namespace"`
	DeploymentName string    `json:"deployment" gorm:"column:deployment"`
	AppId          uint64    `json:"appId" gorm:"column:appId"`
	ClusterId      uint64    `json:"clusterId" gorm:"column:clusterId"`
	NamespaceId    uint64    `json:"namespaceId" gorm:"column:namespaceId"`
	DeploymentId   uint64    `json:"deploymentId" gorm:"column:deploymentId"`
	CreationTime   time.Time `json:"creationTime" gorm:"column:creationTime"`
}
