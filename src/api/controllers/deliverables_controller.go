package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/deliverables"
)

type ArtifactController struct {
	mvc.ApiController
	projectService deliverables.TenantDeliverablesProjectService
}

func (c *ArtifactController) CreateTenantDeliverablesProject(ctx *context.HttpContext, reqData requests.CreateTenantDeliverablesProjectReq) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	c.projectService.CreateTenantDeliverablesProject(reqData)
	return mvc.Success(reqData.Id)
}

func (c *ArtifactController) QueryTenantDeliverablesProject(ctx *context.HttpContext, reqData requests.QueryTenantDeliverablesProjectReq) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err, res := c.projectService.QueryTenantDeliverablesProject(reqData)
	if err != nil {
		panic(err)
	}
	return mvc.Success(res)
}
