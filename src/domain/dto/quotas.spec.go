package dto

type QuotasSpec struct {
	TenantID    uint64 `json:"tenantId"`
	ClusterId   uint64 `json:"clusterId" uri:"clusterId"`
	Namespace   string `json:"namespace"`
	LimitCpu    int    `json:"cpu"`
	LimitMemory int    `json:"memory"`
	LimitPods   int    `json:"pods"`
}
