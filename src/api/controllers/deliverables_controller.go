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

// DeliverablesController 租户制品管理
type DeliverablesController struct {
	mvc.ApiController
	projectService          *deliverables.TenantDeliverablesProjectService
	deliverablesTreeService *deliverables.TenantDeliverablesTreeService
}

// PostTenantDeliverablesProject 根据当前登录用户的所属租户，创建制品
func (c DeliverablesController) PostTenantDeliverablesProject(ctx *context.HttpContext, reqData *requests2.CreateTenantDeliverablesProjectReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err := c.projectService.CreateTenantDeliverablesProject(reqData)
	if err != nil {
		mvc.Fail(err)
	}
	return mvc.Success(reqData.Id)
}

// GetTenantDeliverablesProject 获取当前租户下的制品项目
func (c DeliverablesController) GetTenantDeliverablesProject(ctx *context.HttpContext, reqData *requests2.QueryTenantDeliverablesProjectReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err, res := c.projectService.QueryTenantDeliverablesProject(reqData)
	if err != nil {
		mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

// EditDeliverableTree 修改制品仓库的树形结构
func (c DeliverablesController) EditDeliverableTree(ctx *context.HttpContext, reqData *requests2.EditTenantDeliverablesTreeReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err := c.deliverablesTreeService.EditTree(reqData)
	if err != nil {
		mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success("")
}
