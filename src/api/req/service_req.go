package req

import "kubelilin/pkg/page"

type ServiceRequest struct {
	page.PageRequest
	TenantId    uint64 `json:"tenantId" uri:"tenantId"`
	ClusterId   uint64 `json:"clusterId" uri:"clusterId"`
	Namespace   string `json:"namespace" uri:"namespace"`
	ContinueStr string `json:"continueStr" uri:"continueStr"`
}
