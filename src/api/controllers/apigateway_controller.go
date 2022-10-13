package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests "kubelilin/api/dto/requests"
	"kubelilin/domain/business/networks"
	"kubelilin/domain/database/models"
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

func (controller *ApiGatewayController) PostCreateOrEditTeam(ctx *context.HttpContext, request *requests.GatewayTeamRequest) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	tenantID := userInfo.TenantID
	err := controller.service.CreateOrUpdateTeam(models.ApplicationAPIGatewayTeams{
		ID:        request.Id,
		Name:      request.TeamName,
		GatewayID: request.GatewayId,
		TenantID:  tenantID,
		Level:     request.Level,
		Status:    1,
	})
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(true)
}

func (controller *ApiGatewayController) GetRouterList(request *requests.GatewayRouterRequest) mvc.ApiResult {
	list, err := controller.service.GetRouterList(request)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(list)
}

func (controller *ApiGatewayController) GetAppList(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	tenantID := userInfo.TenantID
	list, err := controller.service.GetAppList(tenantID)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(list)
}

func (controller *ApiGatewayController) GetDeploymentList(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	tenantID := userInfo.TenantID
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("clusterId", "0"))
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appId", "0"))

	list, err := controller.service.GetDeploymentList(tenantID, clusterId, appId)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(list)
}
