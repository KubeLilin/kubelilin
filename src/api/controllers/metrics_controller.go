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

func (controller MetricsController) GetNodes(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	list := controller.metricsService.GetNodeMetrics(clusterId)
	return mvc.Success(list)
}

func (controller MetricsController) GetStatistics(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	return mvc.Success(controller.metricsService.GetStatistics(clusterId))
}

func (controller MetricsController) GetWorkloads(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	return mvc.Success(controller.metricsService.GetResourceMetrics(clusterId))
}

func (controller MetricsController) GetProjects(ctx *context.HttpContext) mvc.ApiResult {
	return mvc.Success(controller.metricsService.GetProjectsMetrics())
}

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

// QueryPodMemoryRss
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

//QueryPodMemorySwap
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

// QueryPodNetworkReceiveBytes
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

// QueryPodNetworkTransmitBytes
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
