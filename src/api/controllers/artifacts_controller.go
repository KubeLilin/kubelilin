package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/artifacts"
)

type ArtifactController struct {
	mvc.ApiController
	projectService artifacts.TenantArtifactsProjectService
}

func (c *ArtifactController) CreateTenantArtifactsProject(ctx *context.HttpContext, reqData requests.CreateTenantArtifactsProjectReq) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	c.projectService.CreateTenantArtifactsProject(reqData)
	return mvc.Success(reqData.Id)
}

func (c *ArtifactController) QueryTenantArtifactsProject(ctx *context.HttpContext, reqData requests.QueryTenantArtifactsProjectReq) mvc.ApiResult {
	userInfo := requests.GetUserInfo(ctx)
	reqData.TenantId = userInfo.TenantID
	err, res := c.projectService.QueryTenantArtifactsProject(reqData)
	if err != nil {
		panic(err)
	}
	return mvc.Success(res)
}
