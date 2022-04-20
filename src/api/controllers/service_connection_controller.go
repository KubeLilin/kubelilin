package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/app"
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

func (controller *ServiceConnectionController) PostCreateServiceConnection(ctx *context.HttpContext, request *req.ServiceConnectionReq) mvc.ApiResult {

	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	res, err := controller.svc.CreateServiceConnection(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller *ServiceConnectionController) PostUpdateServiceConnection(req *req.ServiceConnectionReq) mvc.ApiResult {
	res, err := controller.svc.UpdateServiceConnection(req)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller *ServiceConnectionController) GetQueryServiceConnections(ctx *context.HttpContext) mvc.ApiResult {
	var pageReq req.ServiceConnectionPageReq
	userInfo := req.GetUserInfo(ctx)
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

func (controller *ServiceConnectionController) GetRepoListByType(ctx *context.HttpContext) mvc.ApiResult {
	repoType := ctx.Input.Query("repoType")
	userInfo := req.GetUserInfo(ctx)
	res, err := controller.svc.QueryRepoListByType(userInfo.TenantID, repoType)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}
