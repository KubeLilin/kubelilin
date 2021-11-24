package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/kubernetes"
	"strconv"
)

type DeploymentController struct {
	mvc.ApiController
	clusterService *kubernetes.ClusterService
}

func NewDeploymentController(clusterService *kubernetes.ClusterService) *DeploymentController {
	return &DeploymentController{clusterService: clusterService}
}

func (controller DeploymentController) PostNew(ctx *context.HttpContext, request *req.DeploymentReq) mvc.ApiResult {

	return controller.OK(request)
}

func (controller DeploymentController) PostModify(ctx *context.HttpContext, request *req.DeploymentReq) mvc.ApiResult {
	pDeployId := ctx.Input.QueryDefault("deployid", "0")
	deployId, _ := strconv.ParseInt(pDeployId, 10, 64)
	if deployId > 0 {
		return controller.OK(deployId)
	}

	return controller.OK("ok")
}
