package controllers

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/app"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/domain/business/notice"
	"kubelilin/pkg/page"
	"kubelilin/utils"
	"strconv"
	"time"
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
	execReq.Operator = uint64(userInfo.UserId)
	res, err := controller.deploymentSupervisor.ExecuteDeployment(execReq)
	if err == nil {
		return mvc.Success(res)
	}
	return controller.ApiResult().StatusCode(500).Build()
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
	fmt.Println(deployModel)
	err, res := controller.deploymentService.CreateDeploymentStep2(deployModel)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller DeploymentController) GetList(ctx *context.HttpContext) mvc.ApiResult {
	var request req.DeploymentGetListRequest
	_ = ctx.BindWithUri(&request)
	userInfo := req.GetUserInfo(ctx)
	var tenantID uint64 = 0
	if userInfo != nil {
		tenantID = userInfo.TenantID
	}
	err, deploymentList := controller.deploymentService.GetDeployments(request.Profile, request.AppID, tenantID,
		request.DeployName, request.AppName, request.ClusterId, request.CurrentPage, request.PageSize)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(deploymentList)
}

func (controller DeploymentController) GetDeploymentFormInfo(ctx *context.HttpContext) mvc.ApiResult {
	strDpId := ctx.Input.Query("dpId")
	fmt.Println(strDpId)
	dpId, err := strconv.ParseUint(strDpId, 10, 64)
	if err != nil {
		return mvc.FailWithMsg(nil, "??????id??????????????????????????????id")
	}
	resErr, res := controller.deploymentService.GetDeploymentForm(dpId)
	if resErr != nil {
		return mvc.FailWithMsg(nil, resErr.Error())
	}
	return mvc.Success(res)
}

func (controller DeploymentController) DeleteDeployment(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	deploymentId, err := utils.StringToUInt64(ctx.Input.QueryDefault("dpId", "0"))
	if err != nil {
		return mvc.FailWithMsg(nil, "??????id??????????????????????????????id")
	}
	err = controller.deploymentSupervisor.DeleteDeployment(userInfo.TenantID, deploymentId)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(true)

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
		return mvc.Fail("?????????????????????")
	}
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, deployment.ClusterId)

	ret, _ := kubernetes.SetReplicasNumber(client, deployment.NameSpace, deployment.Name, request.Number)
	if ret {
		_ = controller.deploymentService.SetReplicas(request.DeploymentId, request.Number)
	} else {
		return mvc.Fail("????????????")
	}
	return mvc.Success(ret)
}

func (controller DeploymentController) PostDestroyPod(request *req.DestroyPodRequest, ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, request.ClusterId)
	err := kubernetes.DestroyPod(client, request.Namespace, request.PodName)
	if err != nil {
		return mvc.FailWithMsg("????????????", err.Error())
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
		return mvc.FailWithMsg("??????Pod???????????????", err.Error())
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

func (controller DeploymentController) GetReleaseRecord(ctx *context.HttpContext) mvc.ApiResult {
	dpId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("dpId", "0"))
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appId", "0"))
	level := ctx.Input.QueryDefault("dpLevel", "")
	pageReq := page.InitPageByCtx(ctx)
	err, res := controller.deploymentSupervisor.QueryReleaseRecord(appId, dpId, level, pageReq)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller DeploymentController) PostNotify(notifyReq *req.DeployNotifyRequest) mvc.ApiResult {
	notifyPlugin := linq.From(notice.Plugins).WhereT(func(item notice.Plugin) bool {
		return item.Value == notifyReq.NotifyType
	}).First().(notice.Plugin)

	notifier := notifyPlugin.New(notifyReq.NotifyKey)

	_, deployment := controller.deploymentService.GetDeploymentForm(notifyReq.DeployId)
	message := notice.Message{
		App:         deployment.Nickname,
		Level:       deployment.Level,
		Environment: deployment.Name,
		Version:     notifyReq.Version,
		Branch:      notifyReq.Branch,
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
		Success:     "????????????",
	}

	err := notifier.PostMessage(message)
	if err != nil {
		return mvc.ApiResult{Message: err.Error(), Status: 500}
	}

	return mvc.Success(true)

}

func (c DeploymentController) PostRollBackByReleaseRecord(ctx *context.HttpContext, execReq *req.ExecDeploymentRequest) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	execReq.TenantId = userInfo.TenantID
	execReq.Operator = uint64(userInfo.UserId)
	res, err := c.deploymentSupervisor.ExecuteDeployment(execReq)
	if err == nil {
		return mvc.Success(res)
	}
	return mvc.Fail(err.Error())
}

func (controller DeploymentController) GetNotifications() mvc.ApiResult {
	return mvc.Success(notice.Plugins)
}
