package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/app"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/utils"
)

type RuntimeController struct {
	mvc.ApiController
	clusterService *kubernetes.ClusterService
	runtimeService *app.RuntimeService
}

func NewRuntimeController(clusterService *kubernetes.ClusterService, runtimeService *app.RuntimeService) *RuntimeController {
	return &RuntimeController{clusterService: clusterService, runtimeService: runtimeService}
}

func (controller RuntimeController) GetIsInstalledRuntime(ctx *context.HttpContext) mvc.ApiResult {
	namespaceId := utils.GetNumberOfParam[uint64](ctx, "namespaceId")
	ns := controller.clusterService.GetNameSpacesById(namespaceId)
	if ns != nil && ns.EnableRuntime > 0 {
		return mvc.Success(ns.RuntimeName)
	}
	return mvc.Fail("namespace not found")
}

func (controller RuntimeController) PostSaveDaprComponent(request *requests.RuntimeReq) mvc.ApiResult {
	model, err := controller.runtimeService.SaveDaprComponent(request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(model)
}

func (controller RuntimeController) GetDaprComponentList() mvc.ApiResult {
	list, err := controller.runtimeService.GetDaprComponentList()
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(list)
}

func (controller RuntimeController) DeleteDaprComponent(ctx *context.HttpContext) mvc.ApiResult {
	id := utils.GetNumberOfParam[uint64](ctx, "id")
	err := controller.runtimeService.DeleteDaprComponent(id)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(nil)
}

func (controller RuntimeController) GetDaprResourceList(ctx *context.HttpContext) mvc.ApiResult {
	clusterId := utils.GetNumberOfParam[uint64](ctx, "clusterId")
	namespace := ctx.Input.QueryDefault("namespace", "")
	restConfig, err := controller.clusterService.GetClusterConfig(0, clusterId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	list, err := kubernetes.GetDaprComponentResource(restConfig, namespace)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(list)
}
