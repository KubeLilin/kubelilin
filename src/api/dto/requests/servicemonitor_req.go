package requests

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"time"
)

type ServiceMonitorRequest struct {
	mvc.RequestBody

	ID             uint64     `json:"id"`
	Name           string     `json:"name"`           // ServiceMonitor名称
	AppID          uint64     `json:"appId"`          // 应用ID
	ClusterID      uint64     `json:"clusterId"`      // 集群ID
	Namespace      string     `json:"namespace"`      // 目标命名空间
	DeploymentID   uint64     `json:"deploymentId"`   // 部署ID
	DeploymentName string     `json:"deploymentName"` // 部署名称
	Interval       uint       `json:"interval"`       // 采集间隔时间
	Port           string     `json:"port"`           // 采集服务端口名称,
	Path           string     `json:"path"`           // 采集指标端点
	CreateTime     *time.Time `json:"createTime"`     // 创建时间
	UpdateTime     *time.Time `json:"updateTime"`     // 更新时间
	Status         uint8      `json:"status"`         // 状态
}
