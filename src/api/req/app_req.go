package req

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
	Labels     string `json:"labels" uri:"labels"`
	Remarks    string `json:"remarks" uri:"remarks"`
	Git        string `json:"git" uri:"git"`
	Level      uint16 `json:"level" uri:"level"`
	Language   uint16 `json:"language" uri:"language"`
	Status     int8   `json:"status" uri:"status"`
	SourceType string `json:"sourceType" uri:"sourceType"`
	SCID       uint64 `json:"sources" uri:"sources"`
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
	Id    uint64 `json:"id"`
	AppId uint64 `json:"appid"`
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
