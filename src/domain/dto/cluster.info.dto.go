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
