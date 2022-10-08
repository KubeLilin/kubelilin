package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests2 "kubelilin/api/dto/requests"
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
	reqParam := requests2.ServiceRequest{}
	_ = ctx.BindWithUri(&reqParam)
	userInfo := requests2.GetUserInfo(ctx)
	reqParam.TenantId = userInfo.TenantID
	//userInfo := requests.GetUserInfo(ctx)
	list, err := c.svcSupervisor.QueryServiceList(reqParam)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(list)
}

func (c *ServiceController) GetInfo(ctx *context.HttpContext) mvc.ApiResult {
	reqParam := requests2.ServiceRequest{}
	_ = ctx.BindWithUri(&reqParam)
	userInfo := requests2.GetUserInfo(ctx)
	reqParam.TenantId = userInfo.TenantID
	//userInfo := requests.GetUserInfo(ctx)
	list, err := c.svcSupervisor.QueryServiceInfo(reqParam)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(list)
}

func (c *ServiceController) GetNamespaceByTenant(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	list := c.svcSupervisor.QueryNameSpaceByTenant(userInfo.TenantID)
	return mvc.Success(list)
}

func (c *ServiceController) PostChangeService(ctx *context.HttpContext, svcReq *requests2.ServiceInfoReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	svcReq.TenantId = userInfo.TenantID
	marshal, _ := json.Marshal(svcReq)
	fmt.Println(string(marshal))
	err := c.svcSupervisor.ChangeService(svcReq)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(nil)
}
