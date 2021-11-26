package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/app"
)

type DeploymentController struct {
	mvc.ApiController
	deploymentService *app.DeploymentService
}

func NewDeploymentController(deploymentService *app.DeploymentService) *DeploymentController {
	return &DeploymentController{deploymentService: deploymentService}
}

func (controller DeploymentController) PostStepV1(ctx *context.HttpContext, request *req.DeploymentRequest) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	//default values
	//request.Replicas = 1
	//request.WorkloadType = app.Deployment
	err, m := controller.deploymentService.NewOrUpdateDeployment(request.SgrTenantDeployments)
	if err == nil {
		return mvc.Success(m)
	}

	return mvc.Fail(err.Error())
}
