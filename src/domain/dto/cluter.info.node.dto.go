package dto

type NodeStatus struct {
	CPU     float64 `json:"cpu"`
	Memory  float64 `json:"memory"`
	Storage float64 `json:"storage"`
	Pods    int64   `json:"pods"`
}

type NodeAddress struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}
