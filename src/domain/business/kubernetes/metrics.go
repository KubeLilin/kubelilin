package kubernetes

import (
	"context"
	"gorm.io/gorm"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
	"kubelilin/domain/database/models"
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
	for index, node := range nodeList {
		if err == nil {
			for _, nodeMetrics := range nodeMetricsList.Items {
				if node.Name == nodeMetrics.Name {
					nodeList[index].Usage = dto.NodeStatus{
						CPU:     nodeMetrics.Usage.Cpu().AsApproximateFloat64(),
						Memory:  nodeMetrics.Usage.Memory().AsApproximateFloat64(),
						Pods:    nodeMetrics.Usage.Pods().Value(),
						Storage: nodeMetrics.Usage.StorageEphemeral().AsApproximateFloat64(),
					}
				}
			}
		} else {
			nodeList[index].Usage = dto.NodeStatus{
				CPU:     0,
				Memory:  0,
				Pods:    0,
				Storage: 0,
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

func (metrics *MetricsServer) GetResourceMetrics(clusterId uint64) dto.WorkloadsMetrics {
	clientSet, _ := metrics.clusterService.GetClusterClientByTenantAndId(0, clusterId)
	deploymentList, _ := clientSet.AppsV1().Deployments("").List(context.TODO(), v1.ListOptions{})
	stateful, _ := clientSet.AppsV1().StatefulSets("").List(context.TODO(), v1.ListOptions{})
	replicaSets, _ := clientSet.AppsV1().ReplicaSets("").List(context.TODO(), v1.ListOptions{})
	daemonSets, _ := clientSet.AppsV1().DaemonSets("").List(context.TODO(), v1.ListOptions{})
	cronJobs, _ := clientSet.BatchV1().CronJobs("").List(context.TODO(), v1.ListOptions{})
	jobs, _ := clientSet.BatchV1().Jobs("").List(context.TODO(), v1.ListOptions{})
	pods, _ := clientSet.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})

	podsRunning := 0
	for _, item := range pods.Items {
		if item.Status.Phase == "Running" {
			podsRunning++
		}
	}

	workloadMetrics := dto.WorkloadsMetrics{
		PodsRunning:  podsRunning,
		PodsCount:    len(pods.Items),
		Deployment:   len(deploymentList.Items),
		StatefulSets: len(stateful.Items),
		DaemonSets:   len(daemonSets.Items),
		ReplicaSets:  len(replicaSets.Items),
		CronJobs:     len(cronJobs.Items),
		Jobs:         len(jobs.Items),
	}
	return workloadMetrics
}

func (metrics *MetricsServer) GetProjectsMetrics() dto.ProjectCountMetrics {
	var publishCount, appCount, depolyCount int64
	metrics.db.Model(models.SgrTenantDeploymentRecord{}).Count(&publishCount)
	metrics.db.Model(models.SgrTenantApplication{}).Count(&appCount)
	metrics.db.Model(models.SgrTenantDeployments{}).Count(&depolyCount)
	return dto.ProjectCountMetrics{
		Publish:      publishCount,
		Applications: appCount,
		Deploys:      depolyCount,
	}
}
