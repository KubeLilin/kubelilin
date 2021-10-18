package dto

type NodeStatus struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
	Pods   int     `json:"pods"`
}

type NodeAddress struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}
