package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests2 "kubelilin/api/dto/requests"
	"kubelilin/domain/business/app"
	"kubelilin/domain/business/devops"
	"kubelilin/utils"
)

type DevopsController struct {
	mvc.ApiController
	devopsService  *app.DevopsService
	projectService *devops.ProjectService
}

func NewDevopsController(devops *app.DevopsService, project *devops.ProjectService) *DevopsController {
	return &DevopsController{devopsService: devops, projectService: project}
}

func (controller DevopsController) PostCreateProject(ctx *context.HttpContext, request *requests2.NewProject) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err := controller.devopsService.CreateProject(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("ok")
}

func (controller DevopsController) PostEditProject(ctx *context.HttpContext, request *requests2.NewProject) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err := controller.devopsService.EditProject(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("ok")
}

func (controller DevopsController) DeleteProject(ctx *context.HttpContext) mvc.ApiResult {
	projectId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("id", "0"))
	err := controller.devopsService.DeleteProject(projectId)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("ok")
}

func (controller DevopsController) GetProjectList(ctx *context.HttpContext, request *requests2.DevopsProjectReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := controller.devopsService.GetProjectList(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller DevopsController) GetAppList(ctx *context.HttpContext, request *requests2.AppReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := controller.devopsService.GetAppList(request)
	fmt.Println(res.Data)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller DevopsController) GetPipelineList(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	projectId := utils.GetNumberOfParam[uint64](ctx, "projectId")
	res, err := controller.devopsService.GetPipelineListByProjectId(projectId, userInfo.TenantID)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller DevopsController) GetResourceMetrics(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	projectId := utils.GetNumberOfParam[uint64](ctx, "projectId")
	res, err := controller.projectService.GetResourceMetrics(userInfo.TenantID, projectId)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}
