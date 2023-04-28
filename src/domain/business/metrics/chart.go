package metrics

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Chart struct {
	metricsProvider *PrometheusMetrics
	db              *gorm.DB
}

func NewChartWithDb(db *gorm.DB) *Chart {
	return &Chart{
		db: db,
	}
}

func NewChart(metricsProvider *PrometheusMetrics) *Chart {
	return &Chart{
		metricsProvider: metricsProvider,
	}
}

func (chart *Chart) Get(clusterId uint64) *Chart {
	// get from db with clusterId,that is the dataSource of prometheus client api provider
	dataSource := "http://49.232.111.253:39090"
	// db.table("").where("cluster_id = ?", clusterId).first(&dataSource)

	chart.metricsProvider = NewPrometheusMetrics(dataSource)
	return chart
}

func (chart *Chart) QueryNodeCpuUtilisation(startTime time.Time, endTime time.Time) (string, error) {
	query, err := chart.metricsProvider.Query(`(( instance:node_cpu_utilisation:rate5m{job="node-exporter", cluster=""} * instance:node_num_cpu:sum{job="node-exporter", cluster=""}) != 0 )
			/ scalar(sum(instance:node_num_cpu:sum{job="node-exporter", cluster=""}))`, startTime, endTime)
	return query, err
}

func (chart *Chart) QueryPodCPUUsage(namespace string, workload string, startTime time.Time, endTime time.Time) (string, error) {
	pql := `sum(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{cluster="", namespace="%s"}
  * on(namespace,pod) group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload="%s", workload_type="deployment"}
) by (pod)`
	query, err := chart.metricsProvider.Query(fmt.Sprintf(pql, namespace, namespace, workload), startTime, endTime)
	return query, err
}

// pod memory usage
func (chart *Chart) QueryPodMemoryUsage(namespace string, workload string, startTime time.Time, endTime time.Time) (string, error) {
	pql := `sum(
    container_memory_working_set_bytes{cluster="", namespace="%s", container!="", image!=""}
  * on(namespace,pod)
    group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload="%s", workload_type="deployment"}
) by (pod)
`
	query, err := chart.metricsProvider.Query(fmt.Sprintf(pql, namespace, namespace, workload), startTime, endTime)
	return query, err
}

// container_memory_rss
func (chart *Chart) QueryPodMemoryRss(namespace string, workload string, startTime time.Time, endTime time.Time) (string, error) {
	pql := `sum(
    container_memory_rss{cluster="", namespace="%s", container!="", image!=""}
  * on(namespace,pod)
    group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload="%s", workload_type="deployment"}
) by (pod)
`
	query, err := chart.metricsProvider.Query(fmt.Sprintf(pql, namespace, namespace, workload), startTime, endTime)
	return query, err
}

// container_memory_swap
func (chart *Chart) QueryPodMemorySwap(namespace string, workload string, startTime time.Time, endTime time.Time) (string, error) {
	pql := `sum(
	container_memory_swap{cluster="", namespace="%s", container!="", image!=""}
  * on(namespace,pod)
	group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload="%s", workload_type="deployment"}
) by (pod)
`
	query, err := chart.metricsProvider.Query(fmt.Sprintf(pql, namespace, namespace, workload), startTime, endTime)
	return query, err
}

// container_network_receive_bytes_total
func (chart *Chart) QueryPodNetworkReceiveBytes(namespace string, workload string, startTime time.Time, endTime time.Time) (string, error) {
	pql := `(sum(irate(container_network_receive_bytes_total{job="kubelet", metrics_path="/metrics/cadvisor", cluster="", namespace="%s"}[5m])
* on (namespace,pod)
group_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload=~"%s", workload_type="deployment"}) by (pod))
`
	query, err := chart.metricsProvider.Query(fmt.Sprintf(pql, namespace, namespace, workload), startTime, endTime)
	return query, err
}

//container_network_transmit_bytes_total
func (chart *Chart) QueryPodNetworkTransmitBytes(namespace string, workload string, startTime time.Time, endTime time.Time) (string, error) {
	pql := `(sum(irate(container_network_transmit_bytes_total{job="kubelet", metrics_path="/metrics/cadvisor", cluster="", namespace="%s"}[5m])
* on (namespace,pod)
group_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload=~"%s", workload_type="deployment"}) by (pod))
`
	query, err := chart.metricsProvider.Query(fmt.Sprintf(pql, namespace, namespace, workload), startTime, endTime)
	return query, err
}
