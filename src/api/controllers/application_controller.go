package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/app"
	"sgr/utils"
)

type ApplicationController struct {
	mvc.ApiController
	service         *app.ApplicationService
	pipelineService *app.PipelineService
}

func NewApplicationController(service *app.ApplicationService, pipelineService *app.PipelineService) *ApplicationController {
	return &ApplicationController{service: service, pipelineService: pipelineService}
}

func (c *ApplicationController) PostCreateApp(ctx *context.HttpContext, request *req.AppReq) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := c.service.CreateApp(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (c *ApplicationController) PutEditApp(ctx *context.HttpContext, request *req.AppReq) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := c.service.UpdateApp(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (c *ApplicationController) GetAppList(ctx *context.HttpContext) mvc.ApiResult {
	request := req.AppReq{}
	ctx.BindWithUri(&request)

	userInfo := req.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := c.service.QueryAppList(&request)
	fmt.Println(res.Data)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

func (c *ApplicationController) GetAppLanguage() mvc.ApiResult {
	res := c.service.QueryAppCodeLanguage()
	return mvc.Success(res)
}

func (c *ApplicationController) GetAppLevel() mvc.ApiResult {
	res := c.service.QueryAppLevel()
	return mvc.Success(res)
}

func (c *ApplicationController) GetGitRepo(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	appName := ctx.Input.Query("appName")
	cvsRes, err := c.service.InitGitRepository(userInfo.TenantID, appName)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(cvsRes)
}

func (c *ApplicationController) GetInfo(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appid", "0"))
	info, err := c.service.GetAppInfo(appId)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(info)
}

func (c *ApplicationController) GetGitBranches(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appid", "0"))
	appInfo, _ := c.service.GetAppInfo(appId)
	if appInfo.Git != "" {
		names, _ := c.service.VCSService.GetGitBranches(appInfo.Git)
		return mvc.Success(context.H{
			"git":      appInfo.Git,
			"branches": names,
		})
	}
	// appInfo.Git
	return mvc.Fail("no data")
}

func (c *ApplicationController) GetBuildScripts(ctx *context.HttpContext) mvc.ApiResult {
	return mvc.Success(c.pipelineService.GetBuildScripts())
}

func (c *ApplicationController) PostNewPipeline(ctx *context.HttpContext, req *req.AppNewPipelineReq) mvc.ApiResult {
	err, pipeline := c.pipelineService.NewPipeline(req)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(pipeline.ID)
}

func (c *ApplicationController) GetPipelines(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appid", "0"))
	if appId == 0 {
		return mvc.Fail("没有找到应用")
	}
	pipelines, err := c.pipelineService.GetAppPipelines(appId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(pipelines)
}

func (c *ApplicationController) PostEditPipeline(request *req.EditPipelineReq) mvc.ApiResult {
	err := c.pipelineService.UpdatePipeline(request)
	if err != nil {
		return mvc.Fail(false)
	} else {

	}
	return mvc.Success(true)
}

func (c *ApplicationController) GetPipeline(ctx *context.HttpContext) mvc.ApiResult {
	pipelineId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("id", "0"))
	pipeline, err := c.pipelineService.GetPipelineById(pipelineId)
	if err != nil {
		return mvc.Fail("not found pipeline!")
	}
	return mvc.Success(pipeline)
}
