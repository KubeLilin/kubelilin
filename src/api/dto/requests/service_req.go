package requests

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"k8s.io/apimachinery/pkg/util/intstr"
	"kubelilin/pkg/page"
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
	TenantId  uint64            `json:"tenantId" uri:"tenantId"`
	Namespace string            `json:"namespace" uri:"namespace"`
	Name      string            `json:"name" uri:"name"`
	Type      string            `json:"type" uri:"type"`
	Labels    string            `json:"labels" uri:"labels"`
	Selector  string            `json:"selector" uri:"selector"`
	Port      []ServicePortInfo `json:"port" uri:"port"`
}

type ServicePortInfo struct {
	Name       string             `json:"name"`
	Port       intstr.IntOrString `json:"port"`
	TargetPort intstr.IntOrString `json:"targetPort"`
	NodePort   intstr.IntOrString `json:"nodePort"`
	Protocol   string             `json:"protocol"`
}
