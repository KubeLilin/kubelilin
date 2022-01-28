package dto

type QuotasSpec struct {
	Namespace   string `json:"namespace"`
	LimitCpu    int    `json:"cpu"`
	LimitMemory int    `json:"memory"`
	LimitPods   int    `json:"pods"`
}
