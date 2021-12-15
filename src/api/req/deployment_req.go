package req

import "github.com/yoyofx/yoyogo/web/mvc"

type DeploymentStepRequest struct {
	mvc.RequestBody
	ID              uint64  `json:"id" gorm:"primaryKey;column:id;"`                                            // 部署ID
	DPCID           uint64  `json:"dpcId" gorm:"column:dpc_id;"`                                                //限制ID
	Name            string  `json:"name" gorm:"column:name;type:varchar(30);not null"`                          // 部署名称(英文唯一)
	Nickname        string  `json:"nickname" gorm:"column:nickname;type:varchar(50);not null" `                 // 部署中文名称#
	TenantID        uint64  `json:"tenantId" gorm:"column:tenant_id;type:bigint(20) unsigned;not null"`         // 租户ID
	ClusterID       uint64  `json:"clusterId" gorm:"column:cluster_id;type:bigint(20) unsigned;not null"`       // 集群ID
	NamespaceID     uint64  `json:"namespaceId" gorm:"column:namespace_id;type:bigint(20) unsigned;not null" `  // 命名空间ID
	AppID           uint64  `json:"appId" gorm:"column:app_id;type:bigint(20) unsigned" `                       // 应用ID
	AppName         string  `json:"appName" gorm:"column:app_name;type:varchar(50);not null" `                  // 应用名称(英文唯一)
	Level           string  `json:"level" gorm:"column:level;type:varchar(8);not null" `                        // 环境级别 ( Prod , Test , Dev )
	ImageHub        string  `json:"imageHub" gorm:"column:image_hub;type:varchar(200)"`                         // 自动生成的镜像仓库地址( hub域名/apps/{应用名-部署名} , 如 http://hub.yoyogo.run/apps/demo-prod )
	Status          uint8   `json:"status"`                                                                     // 状态
	Replicas        uint32  `json:"replicas" gorm:"column:replicas;type:int(10) unsigned;not null;default:1"`   // 部署副本数#
	ServiceEnable   bool    `json:"serviceEnable" gorm:"column:service_enable;type:tinyint(1);not null"`        // 是否开启 Service
	ServiceAway     string  `json:"serviceAway" gorm:"column:service_away;type:varchar(10)"`                    // Service访问方式(NodePort、ClusterPort)
	ServicePortType string  `json:"servicePortType" gorm:"column:service_port_type;type:varchar(8)"`            // Service端口映射类型(TCP/UDP)
	ServicePort     string  `json:"servicePort" gorm:"column:service_port;type:smallint(5) unsigned" `          // Service端口映射(容器端口->服务端口) 	// 更新时间
	RequestCPU      float64 `json:"requestCpu" gorm:"column:request_cpu;type:decimal(4,2) unsigned;not null"`   // CPU限制Core - request
	RequestMemory   float64 `json:"requestMemory" gorm:"column:request_memory;type:decimal(5,0);not null"`      // 内存限制MiB - request
	LimitCPU        float64 `json:"limitCpu" gorm:"column:limit_cpu;type:decimal(4,2) unsigned;not null"`       // CPU限制Core - limit
	LimitMemory     float64 `json:"limitMemory" gorm:"column:limit_memory;type:decimal(5,0) unsigned;not null"` // 内存限制MiB
}

type ScaleRequest struct {
	mvc.RequestBody
	Namespace      string `json:"namespace" uri:"namespace"`
	DeploymentName string `json:"deploymentName" uri:"deploymentName"`
	Number         int32  `json:"number" uri:"number"`
	ClusterId      uint64 `json:"clusterId" uri:"clusterId"`
}

type ScaleV1Request struct {
	mvc.RequestBody
	DeploymentId uint64 `json:"deployId" uri:"deployId"`
	Number       int32  `json:"number" uri:"number"`
}

type DestroyPodRequest struct {
	mvc.RequestBody
	ClusterId uint64 `json:"clusterId" uri:"clusterId"`
	Namespace string `json:"namespace" uri:"namespace"`
	PodName   string `json:"podName"`
}

type PodLogsRequest struct {
	mvc.RequestBody
	ClusterId     uint64 `json:"clusterId" uri:"clusterId"`
	Namespace     string `json:"namespace" uri:"namespace"`
	PodName       string `json:"podName" uri:"podName"`
	ContainerName string `json:"containerName" uri:"containerName"`
	Lines         int64  `json:"lines" uri:"lines"`
}
