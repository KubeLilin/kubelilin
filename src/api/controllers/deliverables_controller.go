package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests2 "kubelilin/api/dto/requests"
	"kubelilin/domain/business/deliverables"
)

func NewDeliverablesController(projectService *deliverables.TenantDeliverablesProjectService, deliverablesTreeService *deliverables.TenantDeliverablesTreeService) *DeliverablesController {
	return &DeliverablesController{
		projectService:          projectService,
		deliverablesTreeService: deliverablesTreeService,
	}
}

type DeliverablesController struct {
	mvc.ApiController
	projectService          *deliverables.TenantDeliverablesProjectService
	deliverablesTreeService *deliverables.TenantDeliverablesTreeService
}

func (c DeliverablesController) PostTenantDeliverablesProject(ctx *context.HttpContext, reqData *requests2.CreateTenantDeliverablesProjectReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err := c.projectService.CreateTenantDeliverablesProject(reqData)
	if err != nil {
		mvc.Fail(err)
	}
	return mvc.Success(reqData.Id)
}

func (c DeliverablesController) GetTenantDeliverablesProject(ctx *context.HttpContext, reqData *requests2.QueryTenantDeliverablesProjectReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err, res := c.projectService.QueryTenantDeliverablesProject(reqData)
	if err != nil {
		mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (c DeliverablesController) EditDeliverableTree(ctx *context.HttpContext, reqData *requests2.EditTenantDeliverablesTreeReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err := c.deliverablesTreeService.EditTree(reqData)
	if err != nil {
		mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("")
}
