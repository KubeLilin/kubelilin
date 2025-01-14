package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/domain/business/metrics"
	"kubelilin/utils"
	"time"
)

type MetricsController struct {
	mvc.ApiController
	metricsService *kubernetes.MetricsServer
	chartService   *metrics.Chart
}

func NewMetricsController(metricsServer *kubernetes.MetricsServer, chartService *metrics.Chart) *MetricsController {
	return &MetricsController{metricsService: metricsServer, chartService: chartService}
}

// GetNodes 获取所有的节点信息
func (controller MetricsController) GetNodes(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	list := controller.metricsService.GetNodeMetrics(clusterId)
	return mvc.Success(list)
}

// GetStatistics 获取集群的各项指标统计信息
func (controller MetricsController) GetStatistics(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	return mvc.Success(controller.metricsService.GetStatistics(clusterId))
}

// GetWorkloads 获取工作看板
func (controller MetricsController) GetWorkloads(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	return mvc.Success(controller.metricsService.GetResourceMetrics(clusterId))
}

// GetProjects 获取项目
func (controller MetricsController) GetProjects(ctx *context.HttpContext) mvc.ApiResult {
	return mvc.Success(controller.metricsService.GetProjectsMetrics())
}

// GetCustomMetrics 获取用户的各项指标
func (controller MetricsController) GetCustomMetrics(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	pql := ctx.Input.QueryDefault("pql", "")
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryMetrics(pql, start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}

// GetNodeCpuUtilisation 获取 CPU的各项指标信息
func (controller MetricsController) GetNodeCpuUtilisation(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	// uint64 to unix time
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryNodeCpuUtilisation(start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}

// GetPodCPUUsage 获取当前集群所选 POD 的 CPU使用率
func (controller MetricsController) GetPodCPUUsage(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	namespace := ctx.Input.QueryDefault("namespace", "")
	workload := ctx.Input.QueryDefault("workload", "")
	// uint64 to unix time
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryPodCPUUsage(namespace, workload, start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}

// GetPodMemoryUsage 获取当前集群所选 POD 的内存使用使用率
func (controller MetricsController) GetPodMemoryUsage(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	namespace := ctx.Input.QueryDefault("namespace", "")
	workload := ctx.Input.QueryDefault("workload", "")
	// uint64 to unix time
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryPodMemoryUsage(namespace, workload, start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}

// QueryPodMemoryRss 获取当前 POD的内存交换
func (controller MetricsController) GetPodMemoryRss(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	namespace := ctx.Input.QueryDefault("namespace", "")
	workload := ctx.Input.QueryDefault("workload", "")
	// uint64 to unix time
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryPodMemoryRss(namespace, workload, start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}

// QueryPodMemorySwap 货权当前 POD 的内存 SWAP
func (controller MetricsController) GetPodMemorySwap(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	namespace := ctx.Input.QueryDefault("namespace", "")
	workload := ctx.Input.QueryDefault("workload", "")
	// uint64 to unix time
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryPodMemorySwap(namespace, workload, start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}

// QueryPodNetworkReceiveBytes 获取当前 POD的网络流量流入
func (controller MetricsController) GetPodNetworkReceiveBytes(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	namespace := ctx.Input.QueryDefault("namespace", "")
	workload := ctx.Input.QueryDefault("workload", "")
	// uint64 to unix time
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryPodNetworkReceiveBytes(namespace, workload, start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}

// QueryPodNetworkTransmitBytes 获取当前 POD的的网络交货流量
func (controller MetricsController) GetPodNetworkTransmitBytes(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	startTime := utils.GetNumberOfParam[uint64](ctx, "startTime")
	endTime := utils.GetNumberOfParam[uint64](ctx, "endTime")
	namespace := ctx.Input.QueryDefault("namespace", "")
	workload := ctx.Input.QueryDefault("workload", "")
	// uint64 to unix time
	start := time.Unix(int64(startTime), 0)
	end := time.Unix(int64(endTime), 0)
	chartData, err := controller.chartService.Get(clusterId).QueryPodNetworkTransmitBytes(namespace, workload, start, end)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(chartData)
}
