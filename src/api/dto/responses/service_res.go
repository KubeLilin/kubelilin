package responses

import (
	v1 "k8s.io/api/core/v1"
	"kubelilin/pkg/page"
	"time"
)

type ServiceListRes struct {
	page.Page
}

type NamespaceList struct {
	Namespace string `gorm:"column:namespace;type:varchar(50);not null" json:"namespace"`
}

type ServiceInfo struct {
	Namespace  string
	Name       string
	Type       string
	Labels     string
	Selector   string
	CreateTime time.Time
	Port       []v1.ServicePort
}
