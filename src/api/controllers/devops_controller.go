package controllers

import (
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

func (controller *DevopsController) PostCreateProject(ctx *context.HttpContext, request *req.CreateNewProject) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err := controller.devopsService.CreateProject(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("ok")
}

func (controller *DevopsController) GetProjectList(ctx *context.HttpContext) mvc.ApiResult {
	var request req.DevopsProjectReq
	_ = ctx.BindWithUri(&request)
	err, res := controller.devopsService.GetProjectList(&request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}
