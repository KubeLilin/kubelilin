package req

import "kubelilin/pkg/page"

type ServiceRequest struct {
	page.PageRequest
	TenantId  uint64
	ClusterId uint64
	Namespace string
}
