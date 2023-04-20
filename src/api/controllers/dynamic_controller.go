package controllers

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/kubernetes"
)

type DynamicController struct {
	mvc.ApiController
	service *kubernetes.DynamicResourceSupervisor
}

func NewDynamicController(service *kubernetes.DynamicResourceSupervisor) *DynamicController {
	return &DynamicController{service: service}
}

func (controller DynamicController) PostApply(request *requests.ApplyYAML) mvc.ApiResult {
	err := controller.service.CreateFromYAML(request.ClusterId, request.YAML)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.SuccessVoid()
}
