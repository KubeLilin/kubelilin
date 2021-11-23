package dto

import "time"

type ClusterInfo struct {
	ID           uint64 `json:"id"`           // ID
	TenantID     *int64 `json:"tenantId"`     // 租户ID
	Name         string `json:"name"`         // 集群名称
	Version      string `json:"version"`      // k8s 版本号
	Distribution string `json:"distribution"` // 来源
	Sort         *int   `json:"sort"`         // 排序
	Status       int8   `json:"status"`       // 状态
}

type Pod struct {
	Namespace   string        `json:"namespace"`
	PodName     string        `json:"name"`
	PodIP       string        `json:"ip"`
	HostIP      string        `json:"hostIP"`
	ClusterName string        `json:"clusterName"`
	Count       int           `json:"podCount"`
	Ready       int           `json:"podReadyCount"`
	StartTime   string        `json:"startTime"`
	Age         time.Duration `json:"age"`
	Status      string        `json:"status"`
	Restarts    int           `json:"restarts"`
}

type Namespace struct {
	Name   string        `json:"name"`
	Status string        `json:"status"`
	Age    time.Duration `json:"age"`
}

type Node struct {
	Uid                     string        `json:"uid"`
	Name                    string        `json:"name"`
	PodCIDR                 string        `json:"podCIDR"`
	Addresses               []NodeAddress `json:"addresses"`
	Capacity                NodeStatus    `json:"capacity"`
	Allocatable             NodeStatus    `json:"allocatable"`
	OSImage                 string        `json:"osImage"`
	ContainerRuntimeVersion string        `json:"containerRuntimeVersion"`
	KubeletVersion          string        `json:"kubeletVersion"`
	OperatingSystem         string        `json:"operatingSystem"`
	Architecture            string        `json:"architecture"`
	Status                  string        `json:"status"`
}
