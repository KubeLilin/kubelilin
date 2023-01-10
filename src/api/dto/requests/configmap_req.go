package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type ConfigMap struct {
	mvc.RequestBody
	DeployId    uint64              `json:"deployId" uri:"deployId"`
	AppId       uint64              `json:"appId" uri:"appId"`
	ClusterId   uint64              `json:"clusterId" uri:"clusterId"`
	NamespaceId uint64              `json:"namespaceId" uri:"namespaceId" `
	Name        string              `json:"name" uri:"name"`
	Items       []ConfigMapDataItem `json:"items" uri:"items"`
}

type ConfigMapDataItem struct {
	Key   string `json:"key" uri:"key"`
	Value string `json:"value" uri:"value"`
}
