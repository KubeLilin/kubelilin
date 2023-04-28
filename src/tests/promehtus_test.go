package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"kubelilin/domain/business/metrics"
	"testing"
	"time"
)

func TestQueryNodeCpuUtilisation(t *testing.T) {
	provider := metrics.NewPrometheusMetrics("http://49.232.111.253:39090")
	chart := metrics.NewChart(provider)
	chartDataset, _ := chart.QueryNodeCpuUtilisation(time.Now().Add(-time.Hour), time.Now())
	fmt.Println(chartDataset)
	assert.NotEmpty(t, chartDataset)
}

func TestQueryPodCPUUsage(t *testing.T) {
	provider := metrics.NewPrometheusMetrics("http://49.232.111.253:39090")
	namespace := "klns-administration"
	workload := "dev-yoyogodemo-kind-kind"
	chart := metrics.NewChart(provider)
	chartDataset, _ := chart.QueryPodCPUUsage(namespace, workload, time.Now().Add(-time.Hour), time.Now())
	fmt.Println(chartDataset)
	assert.NotEmpty(t, chartDataset)
	//	pql := `sum(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{cluster="", namespace="%s"}
	//  * on(namespace,pod) group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload="%s", workload_type="deployment"}
	//) by (pod)`
	//	query, _ := metrics.Query(fmt.Sprintf(pql, namespace, namespace, workload), time.Now().Add(-time.Hour), time.Now())
	//	fmt.Println(query)
}

func TestQueryPodMemoryUsage(t *testing.T) {
	metrics := metrics.NewPrometheusMetrics("http://49.232.111.253:39090")
	namespace := "klns-administration"
	workload := "dev-yoyogodemo-kind-kind"
	pql := `sum(
    container_memory_working_set_bytes{cluster="", namespace="%s", container!="", image!=""}
  * on(namespace,pod)
    group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster="", namespace="%s", workload="%s", workload_type="deployment"}
) by (pod)`
	query, _ := metrics.Query(fmt.Sprintf(pql, namespace, namespace, workload), time.Now().Add(-time.Hour), time.Now())
	fmt.Println(query)
}
