package dto

type ClusterMetrics struct {
	Usage       NodeStatus `json:"usage"`       // 使用率
	Allocatable NodeStatus `json:"allocatable"` //可用率
	Capacity    NodeStatus `json:"capacity"`    //总量

	Nodes ClusterNodesMetrics `json:"nodes"`
}

type ClusterNodesMetrics struct {
	Available int64 `json:"available"` //可用数
	Count     int64 `json:"count"`     //总数
}
