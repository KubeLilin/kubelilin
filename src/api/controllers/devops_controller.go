package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/app"
)

type DevopsController struct {
	mvc.ApiController
	devopsService *app.DevopsService
}

func NewDevopsController(devops *app.DevopsService) *DevopsController {
	return &DevopsController{devopsService: devops}
}

func (controller DevopsController) PostCreateProject(ctx *context.HttpContext, request *req.NewProject) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err := controller.devopsService.CreateProject(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("ok")
}

func (controller DevopsController) PostEditProject(ctx *context.HttpContext, request *req.NewProject) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err := controller.devopsService.EditProject(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("ok")
}

func (controller DevopsController) GetProjectList(ctx *context.HttpContext, request *req.DevopsProjectReq) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := controller.devopsService.GetProjectList(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (controller DevopsController) GetAppList(ctx *context.HttpContext, request *req.AppReq) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := controller.devopsService.GetAppList(request)
	fmt.Println(res.Data)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}
