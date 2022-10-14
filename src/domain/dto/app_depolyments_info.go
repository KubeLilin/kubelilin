package dto

import "time"

type DeploymentItemDto struct {
	ID          uint64 `gorm:"column:id;" json:"id,omitempty"`
	Name        string `gorm:"column:name" json:"name"`
	NickName    string `gorm:"column:nickname" json:"nickname"`
	NameSpace   string `gorm:"column:namespace" json:"namespace"`
	ClusterName string `gorm:"column:clusterName" json:"clusterName"`
	ClusterId   uint64 `gorm:"column:clusterId" json:"clusterId"`
	Status      string `gorm:"column:status" json:"status"`
	LastImage   string `gorm:"column:lastImage" json:"lastImage"`
	Running     uint64 `gorm:"column:running" json:"running"`
	Expected    uint64 `gorm:"column:expected" json:"expected"`
	ServiceIP   string `gorm:"column:serviceIP"  json:"serviceIP"`
	ServiceName string `gorm:"column:serviceName" json:"serviceName"`
	ServicePort int    `gorm:"column:servicePort" json:"servicePort"`
	AppName     string `gorm:"column:appName" json:"appName"`
	Level       string `gorm:"column:level" json:"level"`
}

type EventItemDto struct {
	FirstTime   time.Time `json:"firstTime"`
	LastTime    time.Time `json:"lastTime"`
	Name        string    `json:"name"`
	Level       string    `json:"level"`
	Reason      string    `json:"reason"`
	Information string    `json:"infomation"`
	Kind        string    `json:"kind"`
}
