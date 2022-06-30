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
	userInfo := req.GetUserInfo(ctx)
	reqParam.TenantId = userInfo.TenantID
	//userInfo := req.GetUserInfo(ctx)
	list, err := c.svcSupervisor.QueryServiceList(reqParam)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(list)
}

func (c *ServiceController) GetInfo(ctx *context.HttpContext) mvc.ApiResult {
	reqParam := req.ServiceRequest{}
	_ = ctx.BindWithUri(&reqParam)
	userInfo := req.GetUserInfo(ctx)
	reqParam.TenantId = userInfo.TenantID
	//userInfo := req.GetUserInfo(ctx)
	list, err := c.svcSupervisor.QueryServiceInfo(reqParam)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(list)
}

func (c *ServiceController) GetNamespaceByTenant(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	list := c.svcSupervisor.QueryNameSpaceByTenant(userInfo.TenantID)
	return mvc.Success(list)
}

func (c *ServiceController) PostChangeService(ctx *context.HttpContext, svcReq *req.ServiceInfoReq) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	svcReq.TenantId = userInfo.TenantID
	fmt.Println(svcReq)
	err := c.svcSupervisor.ChangeService(svcReq)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(nil)
}
