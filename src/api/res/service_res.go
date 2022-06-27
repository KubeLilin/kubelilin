package res

import "kubelilin/pkg/page"

type ServiceListRes struct {
	page.Page
}

type NamespaceList struct {
	Namespace string `gorm:"column:namespace;type:varchar(50);not null" json:"namespace"`
}
