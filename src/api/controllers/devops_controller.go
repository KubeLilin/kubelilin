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

func (controller *DevopsController) GetProjectList(ctx *context.HttpContext) mvc.ApiResult {
	return mvc.Success(nil)
}

func (controller *DevopsController) PostCreateProject(request *req.CreateNewProject) mvc.ApiResult {
	err := controller.devopsService.CreateProject(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("ok")
}
