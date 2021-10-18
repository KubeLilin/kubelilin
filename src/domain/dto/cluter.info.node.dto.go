package dto

type NodeStatus struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
	Pods   int    `json:"pods"`
}
