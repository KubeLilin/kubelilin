package dto

import "time"

type Pod struct {
	Namespace   string        `json:"namespace"`
	PodName     string        `json:"name"`
	PodIP       string        `json:"ip"`
	HostIP      string        `json:"hostIP"`
	ClusterName string        `json:"clusterName"`
	Count       int           `json:"podCount"`
	Ready       int           `json:"podReadyCount"`
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
	Uid                     string     `json:"uid"`
	Name                    string     `json:"name"`
	PodCIDR                 string     `json:"podCIDR"`
	InternalIP              string     `json:"internalIP"`
	ExternalIP              string     `json:"externalIP"`
	HostName                string     `json:"hostName"`
	Capacity                NodeStatus `json:"capacity"`
	Allocatable             NodeStatus `json:"allocatable"`
	OSImage                 string     `json:"osImage"`
	ContainerRuntimeVersion string     `json:"containerRuntimeVersion"`
	KubeletVersion          string     `json:"kubeletVersion"`
	OperatingSystem         string     `json:"operatingSystem"`
	Architecture            string     `json:"architecture"`
}
