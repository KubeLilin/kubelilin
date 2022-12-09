package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests "kubelilin/api/dto/requests"
	"kubelilin/domain/business/app"
	"kubelilin/domain/business/networks"
	"kubelilin/domain/database/models"
	"kubelilin/utils"
)

type ApiGatewayController struct {
	mvc.ApiController
	service           *networks.ApiGatewayService
	deploymentService *app.DeploymentService
}

func NewApiGatewayController(service *networks.ApiGatewayService, deploymentService *app.DeploymentService) *ApiGatewayController {
	return &ApiGatewayController{service: service, deploymentService: deploymentService}
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

func (controller *ApiGatewayController) GetRouterList(request *requests.GatewayRouterListRequest) mvc.ApiResult {
	list, err := controller.service.GetRouterList(request)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(list)
}

func (controller *ApiGatewayController) GetRouterListBy(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appId", "0"))
	deployId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("deployId", "0"))

	list, err := controller.service.GetRouterListBy(appId, deployId)
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

func (controller *ApiGatewayController) PostCreateOrEditRouter(request *requests.GatewayRouterRequest) mvc.ApiResult {
	deployment, err := controller.deploymentService.GetDeploymentByID(uint64(request.DeploymentID))
	if err != nil {
		return mvc.FailWithMsg(false, "没有找到对应的部署")
	}
	model, rerr := controller.service.CreateOrEditRouter(request, deployment)
	if rerr != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	// add router for apisix api
	var gatewayEntity models.ApplicationAPIGateway
	if request.GatewayID > 0 {
		gatewayEntity, _ = controller.service.GetById(request.GatewayID)
	} else {
		gatewayEntity, _ = controller.service.GetByClusterId(deployment.ClusterId)
	}

	apisixProxy := networks.NewAPISIXProxy(gatewayEntity.AdminURI, gatewayEntity.AccessToken)
	err = apisixProxy.CreateOrUpdateRoute(utils.ToString(model.ID), *model)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(true)
}

func (controller *ApiGatewayController) DeleteRoute(ctx *context.HttpContext) mvc.ApiResult {
	id, _ := utils.StringToUInt64(ctx.Input.QueryDefault("id", "0"))
	gatewayId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("gatewayId", "0"))
	err := controller.service.DeleteRouter(id)
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}

	gatewayEntity, _ := controller.service.GetById(gatewayId)
	apisixProxy := networks.NewAPISIXProxy(gatewayEntity.AdminURI, gatewayEntity.AccessToken)
	err = apisixProxy.DeleteRoute(utils.ToString(id))
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}
	return mvc.Success(true)
}
