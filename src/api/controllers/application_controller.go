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
	service *app.ApplicationService
}

func NewApplicationController(service *app.ApplicationService) *ApplicationController {
	return &ApplicationController{service: service}
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
		return mvc.Success(names)
	}
	// appInfo.Git
	return mvc.Fail("no data")
}

func (c *ApplicationController) GetBuildScripts(ctx *context.HttpContext) mvc.ApiResult {
	return mvc.Success(
		context.H{
			"golang": `# 编译命令，注：当前已在代码根路径下
go env -w GOPROXY=https://goproxy.cn,direct
go build -ldflags="-s -w" -o app .
`,
			"java": `# 编译命令，注：当前已在代码根路径下
mvn clean package                         
`,
			"nodejs": `# 编译命令，注：当前已在代码根路径下
npm config set registry https://registry.npm.taobao.org --global
npm install
npm run build
`,
		})
}
