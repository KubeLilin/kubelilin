package requests

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type AppReq struct {
	mvc.RequestBody
	page.PageRequest
	ID         uint64 `json:"id" uri:"id"`
	TenantID   uint64 `json:"tenantId" uri:"tenantId"`
	Name       string `json:"name" uri:"name"`
	Nickname   string `json:"nickname" uri:"nickname"`
	Labels     string `json:"labels" uri:"labels"`
	Remarks    string `json:"remarks" uri:"remarks"`
	Git        string `json:"git" uri:"git"`
	Level      uint16 `json:"level" uri:"level"`
	Language   uint16 `json:"language" uri:"language"`
	Status     int8   `json:"status" uri:"status"`
	SourceType string `json:"sourceType" uri:"sourceType"`
	SCID       uint64 `json:"sources" uri:"sources"`
	ProjectID  uint64 `json:"pid" uri:"pid"`
}

type ImportAppReq struct {
	mvc.RequestBody
	TenantID    uint64 `json:"tenantId" uri:"tenantId"`
	Name        string `json:"name" uri:"name"`
	Git         string `json:"git" uri:"git"`
	Ref         string `json:"ref" uri:"ref"`
	Level       uint16 `json:"level" uri:"level"`
	Language    uint16 `json:"language" uri:"language"`
	SourceType  string `json:"sourceType" uri:"sourceType"`
	SCID        uint64 `json:"sources" uri:"sources"`
	ProjectID   uint64 `json:"projectId" uri:"projectId"`
	ClusterID   uint64 `json:"clusterId" uri:"clusterId"`
	NamespaceId uint64 `json:"namespaceId" uri:"namespaceId"`

	DeployList []struct {
		DeployName string `json:"deployName"`
		Dockerfile string `json:"dockerfile"`
	} `json:"deployList"`
}

type AppNewPipelineReq struct {
	mvc.RequestBody
	AppId uint64 `json:"appid" uri:"appid"`
	Name  string `json:"name" uri:"name"`
}

type EditPipelineReq struct {
	mvc.RequestBody
	Id    uint64 `json:"id"`
	AppId uint64 `json:"appid"`
	Name  string `gorm:"json:"name"`
	DSL   string `json:"dsl"`
}

type RunPipelineReq struct {
	mvc.RequestBody
	Id     uint64 `json:"id"`
	AppId  uint64 `json:"appid"`
	Branch string `json:"branch"`
}

type AbortPipelineReq struct {
	mvc.RequestBody
	Id     uint64 `json:"id"`
	AppId  uint64 `json:"appid"`
	TaskId int64  `json:"taskId"`
}

type PipelineStatusReq struct {
	mvc.RequestBody
	Id     uint64 `json:"id"`
	Status int8   `json:"status" uri:"status"`
}

type PipelineDetailsReq struct {
	mvc.RequestBody
	Id     uint64 `json:"id" uri:"id"`
	AppId  uint64 `json:"appId" uri:"appId"`
	TaskId int64  `json:"taskId" uri:"taskId"`
}
