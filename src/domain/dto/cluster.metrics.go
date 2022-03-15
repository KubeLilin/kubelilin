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

type WorkloadsMetrics struct {
	PodsRunning  int `json:"podsRunning"`
	PodsCount    int `json:"podsCount"`
	Deployment   int `json:"deployment"`
	StatefulSets int `json:"statefulSets"`
	DaemonSets   int `json:"daemonSets"`
	ReplicaSets  int `json:"replicaSets"`
	CronJobs     int `json:"cronJobs"`
	Jobs         int `json:"jobs"`
}

type ProjectCountMetrics struct {
	Publish      int64 `json:"publish"`
	Applications int64 `json:"applications"`
	Deploys      int64 `json:"deploys"`
}
