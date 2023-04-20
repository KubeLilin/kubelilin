package controllers

import (
	"github.com/yoyofx/glinq"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"kubelilin/api/dto/requests"
	"kubelilin/api/dto/responses"
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

func (controller RuntimeController) GetDaprComponentTypeList(ctx *context.HttpContext) mvc.ApiResult {
	list, err := controller.runtimeService.GetDaprComponentTypes()
	if err != nil {
		return mvc.Fail(err.Error())
	}

	res := glinq.Map(glinq.From(list), func(val string) responses.LabelValues {
		return responses.LabelValues{
			Label: val,
			Value: val,
		}
	})
	return mvc.Success(res.ToSlice())
}

// GetDaprComponentTemplateByType 根据组件类型,获取dapr组件模板
func (controller RuntimeController) GetDaprComponentTemplateByType(ctx *context.HttpContext) mvc.ApiResult {
	componentType := ctx.Input.QueryDefault("componentType", "")
	templateYaml, err := controller.runtimeService.GetDaprComponentTemplateByType(componentType)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	var data map[string]interface{}
	_ = utils.YamlToObject(templateYaml, &data)
	return mvc.Success(data)
}

// PostCreateOrUpdateDaprComponent 创建或更新dapr组件
func (controller RuntimeController) PostCreateOrUpdateDaprComponent(request *requests.DaprComponentRequest) mvc.ApiResult {
	restConfig, err := controller.clusterService.GetClusterConfig(0, request.ClusterId)
	if err != nil {
		return mvc.Fail(err.Error())
	}

	var metaData []map[string]any
	for k, v := range request.Metadata {
		metaData = append(metaData, map[string]any{
			"name":  k,
			"value": v,
		})
	}

	component := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "dapr.io/v1alpha1",
			"kind":       "Component",
			"metadata": map[string]interface{}{
				"name":      request.Name,
				"namespace": request.Namespace,
			},
			"spec": map[string]interface{}{
				"type":     request.Type,
				"version":  request.Version,
				"metadata": metaData,
			},
		},
	}

	err = kubernetes.CreateOrUpdateDaprComponentResource(restConfig, request.Namespace, component)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(nil)
}

// delete dapr component
func (controller RuntimeController) DeleteDaprComponentResource(ctx *context.HttpContext) mvc.ApiResult {
	clusterId := utils.GetNumberOfParam[uint64](ctx, "clusterId")
	namespace := ctx.Input.QueryDefault("namespace", "")
	name := ctx.Input.QueryDefault("name", "")
	restConfig, err := controller.clusterService.GetClusterConfig(0, clusterId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	err = kubernetes.DeleteDaprComponentResource(restConfig, namespace, name)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}
