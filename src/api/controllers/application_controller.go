package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/app"
)

type ApplicationController struct {
	mvc.ApiController
	service    *app.ApplicationService
	vscService *app.VcsService
}

func NewApplicationController(service *app.ApplicationService, vscService *app.VcsService) *ApplicationController {
	return &ApplicationController{service: service, vscService: vscService}
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
	cvsRes, err := c.vscService.InitGitRepository(userInfo.TenantID, appName)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(cvsRes)
}
