package kubernetes

import (
	"context"
	"gorm.io/gorm"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
	"kubelilin/domain/dto"
)

type MetricsServer struct {
	db             *gorm.DB
	clusterService *ClusterService
}

func NewMetricsServer(db *gorm.DB, clusterService *ClusterService) *MetricsServer {
	return &MetricsServer{db: db, clusterService: clusterService}
}

func (metrics *MetricsServer) GetNodeMetrics(clusterId uint64) []dto.Node {
	config, _ := metrics.clusterService.GetClusterConfig(0, clusterId)
	clientSet, _ := metrics.clusterService.GetClusterClientByTenantAndId(0, clusterId)
	nodeList := GetNodeList(clientSet)
	metricsClient, _ := metricsv.NewForConfig(config)
	nodeMetricsList, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for index, node := range nodeList {
		for _, nodeMetrics := range nodeMetricsList.Items {
			if node.Name == nodeMetrics.Name {
				nodeList[index].Usage = dto.NodeStatus{
					CPU:     nodeMetrics.Usage.Cpu().AsApproximateFloat64(),
					Memory:  nodeMetrics.Usage.Memory().AsApproximateFloat64(),
					Pods:    nodeMetrics.Usage.Pods().Size(),
					Storage: nodeMetrics.Usage.StorageEphemeral().AsApproximateFloat64(),
				}
			}
		}
	}
	return nodeList
}

func (metrics *MetricsServer) GetStatistics(clusterId uint64) dto.ClusterMetrics {
	nodeList := metrics.GetNodeMetrics(clusterId)
	clusterMetrics := dto.ClusterMetrics{Usage: dto.NodeStatus{}, Allocatable: dto.NodeStatus{},
		Capacity: dto.NodeStatus{}, Nodes: dto.ClusterNodesMetrics{}}
	for _, node := range nodeList {
		// Usage
		clusterMetrics.Usage.CPU += node.Usage.CPU
		clusterMetrics.Usage.Memory += node.Usage.Memory
		clusterMetrics.Usage.Pods += node.Usage.Pods
		clusterMetrics.Usage.Storage += node.Usage.Storage
		// Allocatable
		clusterMetrics.Allocatable.CPU += node.Allocatable.CPU
		clusterMetrics.Allocatable.Memory += node.Allocatable.Memory
		clusterMetrics.Allocatable.Pods += node.Allocatable.Pods
		clusterMetrics.Allocatable.Storage += node.Allocatable.Storage
		// Capacity
		clusterMetrics.Capacity.CPU += node.Capacity.CPU
		clusterMetrics.Capacity.Memory += node.Capacity.Memory
		clusterMetrics.Capacity.Pods += node.Capacity.Pods
		clusterMetrics.Capacity.Storage += node.Capacity.Storage

		clusterMetrics.Nodes.Count++
		if node.Status == "ready" {
			clusterMetrics.Nodes.Available++
		}

	}
	return clusterMetrics
}
