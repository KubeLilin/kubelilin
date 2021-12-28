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
	deploymentService    *app.DeploymentService
	deploymentSupervisor *kubernetes.DeploymentSupervisor
	clusterService       *kubernetes.ClusterService
}

func NewDeploymentController(deploymentService *app.DeploymentService, clusterService *kubernetes.ClusterService, deploymentSupervisor *kubernetes.DeploymentSupervisor) *DeploymentController {
	return &DeploymentController{deploymentService: deploymentService, clusterService: clusterService, deploymentSupervisor: deploymentSupervisor}
}

func (controller DeploymentController) PostExecuteDeployment(ctx *context.HttpContext, execReq *req.ExecDeploymentRequest) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	execReq.TenantId = userInfo.TenantID
	res, err := controller.deploymentSupervisor.ExecuteDeployment(execReq)
	if err == nil {
		return mvc.Success(res)
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
	appName := ctx.Input.QueryDefault("appName", "")

	userInfo := req.GetUserInfo(ctx)
	var tenantID uint64 = 0
	if userInfo != nil {
		tenantID = userInfo.TenantID
	}
	appid, _ := strconv.ParseUint(strAppId, 10, 64)
	depolymentList, err := controller.deploymentService.GetDeployments(appid, tenantID, deployName, appName)
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

func (controller DeploymentController) PostReplicas(request *req.ScaleRequest, ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, request.ClusterId)
	ret, err := kubernetes.SetReplicasNumber(client, request.Namespace, request.DeploymentName, request.Number)
	if err != nil {
		panic(err)
	}
	return mvc.Success(ret)
}

func (controller DeploymentController) PostReplicasById(request *req.ScaleV1Request, ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	deployment, _ := controller.deploymentService.GetDeploymentByID(request.DeploymentId)
	if deployment.ID == 0 {
		return mvc.Fail("未找到部署信息")
	}
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, deployment.ClusterId)

	ret, _ := kubernetes.SetReplicasNumber(client, deployment.NameSpace, deployment.Name, request.Number)
	if ret {
		_ = controller.deploymentService.SetReplicas(request.DeploymentId, request.Number)
	} else {
		return mvc.Fail("操作失败")
	}
	return mvc.Success(ret)
}

func (controller DeploymentController) PostDestroyPod(request *req.DestroyPodRequest, ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, request.ClusterId)
	err := kubernetes.DestroyPod(client, request.Namespace, request.PodName)
	if err != nil {
		return mvc.FailWithMsg("操作失败", err.Error())
	}
	return mvc.Success(true)
}

func (controller DeploymentController) GetPodLogs(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	var request *req.PodLogsRequest
	_ = ctx.BindWithUri(&request)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, request.ClusterId)
	logs, err := kubernetes.GetLogs(client, request.Namespace, request.PodName, request.ContainerName, request.Lines)
	if err != nil {
		return mvc.FailWithMsg("获取Pod日志失败！", err.Error())
	}
	return mvc.Success(logs)
}

func (controller DeploymentController) GetEvents(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	var request *req.EventsRequest
	_ = ctx.BindWithUri(&request)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, request.ClusterId)
	events := kubernetes.GetEvents(client, request.Namespace, request.Deployment)
	return mvc.Success(events)
}

func (controller DeploymentController) GetYaml(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	dpIdStr := ctx.Input.Query("dpId")
	dpId, _ := strconv.ParseUint(dpIdStr, 10, 64)
	yamlStr, err := controller.deploymentSupervisor.GetDeploymentYaml(userInfo.TenantID, dpId)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(yamlStr)
}
