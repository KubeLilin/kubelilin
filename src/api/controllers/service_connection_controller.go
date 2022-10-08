package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests2 "kubelilin/api/dto/requests"
	"kubelilin/domain/business/app"
	"kubelilin/utils"
	"strconv"
)

type ServiceConnectionController struct {
	mvc.ApiController
	svc *app.ServiceConnectionService
}

func NewServiceConnectionController(svc *app.ServiceConnectionService) *ServiceConnectionController {
	return &ServiceConnectionController{
		svc: svc,
	}
}

func (controller *ServiceConnectionController) PostCreateServiceConnection(ctx *context.HttpContext, request *requests2.ServiceConnectionReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	res, err := controller.svc.CreateServiceConnection(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller *ServiceConnectionController) PostUpdateServiceConnection(req *requests2.ServiceConnectionReq) mvc.ApiResult {
	res, err := controller.svc.UpdateServiceConnection(req)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller *ServiceConnectionController) GetQueryServiceConnections(ctx *context.HttpContext) mvc.ApiResult {
	var pageReq requests2.ServiceConnectionPageReq
	userInfo := requests2.GetUserInfo(ctx)
	ctx.BindWithUri(&pageReq)
	pageReq.TenantID = userInfo.TenantID
	res, err := controller.svc.QueryServiceConnections(&pageReq)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller *ServiceConnectionController) GetServiceConnectionInfo(ctx *context.HttpContext) mvc.ApiResult {
	id, _ := strconv.ParseInt(ctx.Input.Query("id"), 10, 64)
	res, err := controller.svc.QueryServiceConnectionInfo(id)
	if err == nil {
		return mvc.Success(res)
	}
	return mvc.FailWithMsg(nil, err.Error())
}

func (controller *ServiceConnectionController) DeleteServiceConnectionInfo(ctx *context.HttpContext) mvc.ApiResult {
	id, _ := utils.StringToUInt64(ctx.Input.QueryDefault("id", "0"))
	err := controller.svc.DeleteServiceConnectionInfo(id)
	if err == nil {
		return mvc.Success(true)
	}
	return mvc.FailWithMsg(nil, err.Error())
}

func (controller *ServiceConnectionController) GetRepoListByType(ctx *context.HttpContext) mvc.ApiResult {
	repoType := ctx.Input.Query("repoType")
	userInfo := requests2.GetUserInfo(ctx)
	res, err := controller.svc.QueryRepoListByType(userInfo.TenantID, repoType)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}
