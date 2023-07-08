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
		pods, _ := clientSet.CoreV1().Pods("").List(context.Background(), v1.ListOptions{
			FieldSelector: "spec.nodeName=" + node.Name,
		})
		if err == nil {
			for _, nodeMetrics := range nodeMetricsList.Items {
				if node.Name == nodeMetrics.Name {
					nodeList[index].Usage = dto.NodeStatus{
						CPU:     nodeMetrics.Usage.Cpu().AsApproximateFloat64(),
						Memory:  nodeMetrics.Usage.Memory().AsApproximateFloat64(),
						Pods:    int64(len(pods.Items)),
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
		// Requests
		clusterMetrics.Requests.CPU += node.Requests.CPU
		clusterMetrics.Requests.Memory += node.Requests.Memory
		// Limits
		clusterMetrics.Limits.CPU += node.Limits.CPU
		clusterMetrics.Limits.Memory += node.Limits.Memory

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

// GetTeamSpacesByTenantId 所在团队的命名空间数量
//  label,value,count
//	cls-hbktlqm        3 	2
//	microk8s-cluster   4    2
//	kind-kind	       5	1
func (metrics *MetricsServer) GetTeamSpacesByTenantId(tenantId uint64) ([]dto.DeployLeveLCountInfo, error) {
	var sqlParams []interface{}
	where := ""
	if tenantId > 0 {
		where = "WHERE ns.tenant_id = ?"
		sqlParams = append(sqlParams, tenantId)
	}
	sql := `SELECT clu.name label,ns.cluster_id value ,COUNT(ns.cluster_id) count FROM sgr_tenant_namespace ns
INNER JOIN sgr_tenant_cluster clu on clu.id = ns.cluster_id` + where + `GROUP BY ns.cluster_id`
	var list []dto.DeployLeveLCountInfo
	err := metrics.db.Raw(sql, sqlParams).Find(&list).Error
	return list, err
}

// GetDeployLevelCountByTenantId 团队环境级别的部署数量
//	开发环境		dev	       4
//	测试环境		test	   3
//	预发布环境	release	   0
//	生产环境		prod	   4
func (metrics *MetricsServer) GetDeployLevelCountByTenantId(tenantId uint64) ([]dto.DeployLeveLCountInfo, error) {
	var sqlParams []interface{}
	where := ""
	if tenantId > 0 {
		where = "WHERE tenant_Id = ?"
		sqlParams = append(sqlParams, tenantId)
	}
	sql := `SELECT lev.name label,lev.code  value,IFNULL(dep.count,0) count FROM sgr_code_deployment_level lev
LEFT JOIN (
   SELECT  level,COUNT(level) count FROM sgr_tenant_deployments ` + where +
		`	 GROUP BY level
) dep on dep.level = lev.code`
	var list []dto.DeployLeveLCountInfo
	err := metrics.db.Raw(sql, sqlParams).Find(&list).Error
	return list, err
}
