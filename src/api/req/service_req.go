package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	v1 "k8s.io/api/core/v1"
	"kubelilin/pkg/page"
	"time"
)

type ServiceRequest struct {
	page.PageRequest
	Name        string `json:"name" uri:"name"`
	TenantId    uint64 `json:"tenantId" uri:"tenantId"`
	ClusterId   uint64 `json:"clusterId" uri:"clusterId"`
	Namespace   string `json:"namespace" uri:"namespace"`
	ContinueStr string `json:"continueStr" uri:"continueStr"`
}

type ServiceInfoReq struct {
	mvc.RequestBody
	TenantId   uint64 `json:"tenantId" uri:"tenantId"`
	Namespace  string `json:"namespace" uri:"namespace"`
	Name       string `json:"name" uri:"name"`
	Type       string `json:"type" uri:"type"`
	Labels     string `json:"labels" uri:"labels"`
	Selector   string `json:"selector" uri:"selector"`
	CreateTime time.Time
	Port       []v1.ServicePort `json:"port" uri:"port"`
}
