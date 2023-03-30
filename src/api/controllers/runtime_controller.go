package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/utils"
)

type RuntimeController struct {
	mvc.ApiController
	clusterService *kubernetes.ClusterService
}

func NewRuntimeController(clusterService *kubernetes.ClusterService) *RuntimeController {
	return &RuntimeController{clusterService: clusterService}
}

func (controller RuntimeController) GetIsInstalledRuntime(ctx *context.HttpContext) mvc.ApiResult {
	namespaceId := utils.GetNumberOfParam[uint64](ctx, "namespaceId")
	ns := controller.clusterService.GetNameSpacesById(namespaceId)
	if ns != nil && ns.EnableRuntime > 0 {
		return mvc.Success(ns.RuntimeName)
	}
	return mvc.Fail("namespace not found")
}
