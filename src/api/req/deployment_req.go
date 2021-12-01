package req

import "github.com/yoyofx/yoyogo/web/mvc"

type DeploymentStepRequest struct {
	mvc.RequestBody
	ID              uint64  `json:"id"`          // 部署ID
	Name            string  `json:"name"`        // 部署名称(英文唯一)
	Nickname        string  `json:"nickname"`    // 部署中文名称#
	TenantID        uint64  `json:"tenantId"`    // 租户ID
	ClusterID       uint64  `json:"clusterId"`   // 集群ID
	NamespaceID     uint64  `json:"namespaceId"` // 命名空间ID
	AppID           *uint64 `json:"appId"`       // 应用ID
	AppName         string  `json:"appName"`     // 应用名称(英文唯一)
	LastImage       string  `json:"lastImage"`
	Level           string  `json:"level"`           // 环境级别 ( Prod , Test , Dev )
	ImageHub        string  `json:"imageHub"`        // 自动生成的镜像仓库地址( hub域名/apps/{应用名-部署名} , 如 http://hub.yoyogo.run/apps/demo-prod )
	Status          uint8   `json:"status"`          // 状态
	WorkloadType    string  `json:"workloadType"`    // 部署类型(Deployment、DaemonSet、StatefulSet、CronJob)
	Replicas        uint32  `json:"replicas"`        // 部署副本数#
	ServiceEnable   bool    `json:"serviceEnable"`   // 是否开启 Service
	ServiceAway     string  `json:"serviceAway"`     // Service访问方式(NodePort、ClusterPort)
	ServicePortType string  `json:"servicePortType"` // Service端口映射类型(TCP/UDP)
	ServicePort     *uint16 `json:"servicePort"`     // Service端口映射(容器端口->服务端口) 	// 更新时间
}
