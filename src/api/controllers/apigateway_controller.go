package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests "kubelilin/api/dto/requests"
	"kubelilin/domain/business/networks"
	"kubelilin/utils"
)

type ApiGatewayController struct {
	mvc.ApiController
	service *networks.ApiGatewayService
}

func NewApiGatewayController(service *networks.ApiGatewayService) *ApiGatewayController {
	return &ApiGatewayController{service: service}
}

func (controller *ApiGatewayController) GetList(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("adminUri", "0"))
	list, err := controller.service.GetAllGatewayList(clusterId)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(list)
}

func (controller *ApiGatewayController) GetTeamList(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	tenantID := userInfo.TenantID
	gatewayId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("gatewayId", "0"))
	list, err := controller.service.GetAllGatewayTeamList(gatewayId, tenantID)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(list)
}
