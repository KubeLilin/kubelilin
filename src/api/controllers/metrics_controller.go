package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/utils"
)

type MetricsController struct {
	mvc.ApiController
	metricsService *kubernetes.MetricsServer
}

func NewMetricsController(metricsServer *kubernetes.MetricsServer) *MetricsController {
	return &MetricsController{metricsService: metricsServer}
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
