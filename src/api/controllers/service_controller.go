package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/kubernetes"
)

type ServiceController struct {
	mvc.ApiController
	clusterService *kubernetes.ClusterService
	svcSupervisor  *kubernetes.ServiceSupervisor
}

func NewServiceController(clusterService *kubernetes.ClusterService, svcSupervisor *kubernetes.ServiceSupervisor) *ServiceController {
	return &ServiceController{
		clusterService: clusterService,
		svcSupervisor:  svcSupervisor,
	}
}

func (c *ServiceController) GetServiceList(ctx *context.HttpContext) mvc.ApiResult {
	reqParam := req.ServiceRequest{}
	_ = ctx.BindWithUri(&reqParam)
	reqParam.TenantId = 1
	reqParam.ClusterId = 3
	reqParam.Namespace = "yoyogo"
	fmt.Println("666")
	fmt.Println(reqParam)
	//userInfo := req.GetUserInfo(ctx)
	list, err := c.svcSupervisor.QueryServiceList(reqParam)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(list)
}
