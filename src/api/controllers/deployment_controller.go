package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/app"
	"sgr/domain/business/kubernetes"
	"strconv"
)

type DeploymentController struct {
	mvc.ApiController
	deploymentService *app.DeploymentService
	clusterService    *kubernetes.ClusterService
}

func NewDeploymentController(deploymentService *app.DeploymentService, clusterService *kubernetes.ClusterService) *DeploymentController {
	return &DeploymentController{deploymentService: deploymentService, clusterService: clusterService}
}

func (controller DeploymentController) PostStepV1(ctx *context.HttpContext, request *req.DeploymentStepRequest) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	//default values
	//request.Replicas = 1
	//request.WorkloadType = app.Deployment
	err, m := controller.deploymentService.NewOrUpdateDeployment(request)
	if err == nil {
		return mvc.Success(m)
	}
	return mvc.Fail(err.Error())
}

func (controller *DeploymentController) PostCreateDeploymentStep1(ctx *context.HttpContext, deployModel *req.DeploymentStepRequest) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	var tenantID uint64 = 0
	if userInfo != nil {
		tenantID = userInfo.TenantID
	}
	deployModel.TenantID = tenantID
	err, res := controller.deploymentService.CreateDeploymentStep1(deployModel)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller *DeploymentController) PostCreateDeploymentStep2(deployModel *req.DeploymentStepRequest) mvc.ApiResult {
	err, res := controller.deploymentService.CreateDeploymentStep2(deployModel)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller DeploymentController) GetList(ctx *context.HttpContext) mvc.ApiResult {
	strAppId := ctx.Input.QueryDefault("appid", "0")
	deployName := ctx.Input.QueryDefault("nickname", "")
	userInfo := req.GetUserInfo(ctx)
	var tenantID uint64 = 0
	if userInfo != nil {
		tenantID = userInfo.TenantID
	}
	appid, _ := strconv.ParseUint(strAppId, 10, 64)
	depolymentList, err := controller.deploymentService.GetDeployments(appid, tenantID, deployName)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(depolymentList)
}

func (controller DeploymentController) GetDeploymentFormInfo(ctx *context.HttpContext) mvc.ApiResult {
	strDpId := ctx.Input.Query("dpId")
	fmt.Println(strDpId)
	dpId, err := strconv.ParseUint(strDpId, 10, 64)
	if err != nil {
		return mvc.FailWithMsg(nil, "部署id无效或者未接收到部署id")
	}
	resErr, res := controller.deploymentService.GetDeploymentForm(dpId)
	if resErr != nil {
		return mvc.FailWithMsg(nil, resErr.Error())
	}
	return mvc.Success(res)
}

func (controller DeploymentController) PostReplicas(request *req.SacRequest, ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.Atoi(strCid)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)
	ret, err := kubernetes.SetReplicasNumber(client, request.Namespace, request.DeploymentName, request.Number)
	if err != nil {
		panic(err)
	}
	return mvc.Success(ret)
}
