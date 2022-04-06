package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/app"
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

func (controller *ServiceConnectionController) GetQueryServiceConnections(ctx *context.HttpContext, pageReq *req.ServiceConnectionPageReq) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	pageReq.TenantID = userInfo.TenantID
	res, err := controller.svc.QueryServiceConnections(pageReq)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}
