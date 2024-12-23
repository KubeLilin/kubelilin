package controllers

import (
	"encoding/json"
	"github.com/yoyofx/yoyogo/web/binding"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	requests2 "kubelilin/api/dto/requests"
	"kubelilin/domain/business/app"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/utils"
)

// ApplicationController 应用构造
type ApplicationController struct {
	mvc.ApiController
	service           *app.ApplicationService
	deploymentService *app.DeploymentService
	pipelineService   *app.PipelineService
}

// NewApplicationController 应用管理器构造函数
func NewApplicationController(service *app.ApplicationService, deploymentService *app.DeploymentService, pipelineService *app.PipelineService) *ApplicationController {
	return &ApplicationController{service: service, deploymentService: deploymentService, pipelineService: pipelineService}
}

// PostCreateApp 根据当前登录用户所述租户创建应用
func (c *ApplicationController) PostCreateApp(ctx *context.HttpContext, request *requests2.AppReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := c.service.CreateApp(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

// PostImportApp 当前登录用户根据 yaml 文件直接导入一个应用
func (c *ApplicationController) PostImportApp(ctx *context.HttpContext, request *requests2.ImportAppReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	// 解析传入的参数
	appreq := requests2.AppReq{
		TenantID:   request.TenantID,
		Name:       request.Name,
		Labels:     request.Name,
		Remarks:    request.Name,
		Git:        request.Git,
		Level:      request.Level,
		Language:   request.Language,
		Status:     1,
		SourceType: request.SourceType,
		SCID:       request.SCID,
		ProjectID:  request.ProjectID,
	}
	// create application
	err, app := c.service.CreateApp(&appreq)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	//便利导入的部署，顺道激发流水线
	for _, deployItem := range request.DeployList {
		deployModel := &requests2.DeploymentStepRequest{
			Name:            deployItem.DeployName,
			Nickname:        deployItem.DeployName,
			TenantID:        request.TenantID,
			ClusterID:       request.ClusterID,
			NamespaceID:     request.NamespaceId,
			AppID:           app.ID,
			AppName:         app.Name,
			Level:           "dev",
			ImageHub:        "",
			Status:          1,
			Replicas:        1,
			ServiceEnable:   true,
			ServiceAway:     "ClusterPort",
			ServicePortType: "TCP",
			ServicePort:     "8080",
			RequestCPU:      0.1,
			RequestMemory:   128,
			LimitCPU:        1,
			LimitMemory:     768,
			Environments:    nil,
			EnvJson:         "",
			Runtime:         "",
		}
		_ = deployModel
		// create deployment
		err, dinfo := c.deploymentService.CreateDeploymentStep1(deployModel)
		if err == nil {
			deployModel.ID = dinfo.ID
			// create container
			err, _ = c.deploymentService.CreateDeploymentStep2(deployModel)
		}
		// create pipeline by dsl
		pipelineReq := &requests2.AppNewPipelineReq{
			AppId: app.ID,
			Name:  "default-pipeline-" + deployItem.DeployName,
		}
		perr, pipelineInfo := c.pipelineService.NewPipeline(pipelineReq)
		if perr == nil {
			dsl := []dto.StageInfo{
				{
					Name: "代码",
					Steps: []dto.StepInfo{
						{
							Name: "拉取代码", Key: "git_pull",
							Content: map[string]interface{}{"git": request.Git, "branch": request.Ref},
						},
					},
				},
				{
					Name: "编译构建",
					Steps: []dto.StepInfo{
						{
							Name: "编译命令", Key: "code_build",
							Content: map[string]interface{}{"buildEnv": "golang", "buildImage": "golang:1.16.15", "buildScript": "# 编译命令，注：当前已在代码根路径下\n",
								"buildFile": "./" + deployItem.Dockerfile},
						},
					},
				},
				{
					Name: "部署",
					Steps: []dto.StepInfo{
						{
							Name: "应用部署", Key: "app_deploy",
							Content: map[string]interface{}{"depolyment": dinfo.ID},
						},
					},
				},
			}
			dslStr := utils.ObjectToString(dsl)
			editPipeline := &requests2.EditPipelineReq{
				Id:    pipelineInfo.ID,
				AppId: app.ID,
				Name:  "default-pipeline-" + deployItem.DeployName,
				DSL:   dslStr,
			}
			perr = c.pipelineService.UpdatePipeline(editPipeline)
			if perr == nil {
				_ = c.pipelineService.UpdateDSL(editPipeline)
			}
		}
	}

	return mvc.Success(true)
}

// PutEditApp 更新当前 app的基础信息
func (c *ApplicationController) PutEditApp(ctx *context.HttpContext, request *requests2.AppReq) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := c.service.UpdateApp(request)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

// DeleteApp 删除当前所操作的 APP
func (c *ApplicationController) DeleteApp(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appId", "0"))
	err := c.service.DeleteApp(appId)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(true)
}

// GetAppList 根据当前登录用户的所属租户，获取 APP列表
func (c *ApplicationController) GetAppList(ctx *context.HttpContext) mvc.ApiResult {
	request := requests2.AppReq{}
	_ = ctx.BindWithUri(&request)

	userInfo := requests2.GetUserInfo(ctx)
	request.TenantID = userInfo.TenantID
	err, res := c.service.QueryAppList(&request)
	//fmt.Println(res.Data)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(res)
}

// GetAppLanguage 获取 APP所支持的开发语言列表
func (c *ApplicationController) GetAppLanguage() mvc.ApiResult {
	res := c.service.QueryAppCodeLanguage()
	return mvc.Success(res)
}

// GetAppLevel 获取 APP所属等级
func (c *ApplicationController) GetAppLevel() mvc.ApiResult {
	res := c.service.QueryAppLevel()
	return mvc.Success(res)
}

// GetDeployLevel 获取部署等级，用于下拉列表
func (c *ApplicationController) GetDeployLevel() mvc.ApiResult {
	res := c.service.QueryDeployLevel()
	return mvc.Success(res)
}

// GetDeployLevelCounts 统计不同等级的Deploy数量
func (c *ApplicationController) GetDeployLevelCounts(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appid", "0"))
	res, _ := c.service.GetAppCountByDeployLevel(appId)
	return mvc.Success(res)
}

// GetProjectDeployLevelCounts 统计各个项目瞎不同等级的Deploy数量进行显示，用于填充表格顶部状态栏
func (c *ApplicationController) GetProjectDeployLevelCounts(ctx *context.HttpContext) mvc.ApiResult {
	projectId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("projectId", "0"))
	res, _ := c.service.GetProjectCountByDeployLevel(projectId)
	return mvc.Success(res)
}

// GetTeamDeployLevelCounts 统计当前租户下不同等级的Deploy用于进行抬头状态 Tab 的显示
func (c *ApplicationController) GetTeamDeployLevelCounts(ctx *context.HttpContext) mvc.ApiResult {
	tenantId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("tenantId", "0"))
	projectCounts, _ := c.service.GetTenantProjectCountByDeployLevel(tenantId)
	projectCount, _ := c.service.GetProjectCountByTenantId(tenantId)
	namespaceCount, _ := c.service.GetNamespaceCountByTenantId(tenantId)
	appCount, _ := c.service.GetAppCountByTenantId(tenantId)
	return mvc.Success(context.H{"insCounts": projectCounts, "proCount": projectCount, "namespaceCount": namespaceCount, "appCount": appCount})
}

// GetGitRepo get git address for application 获取当前所选应用的 GIT仓库地址
func (c *ApplicationController) GetGitRepo(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	appName := ctx.Input.Query("appName")
	cvsRes, err := c.service.InitGitRepository(userInfo.TenantID, appName)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(cvsRes)
}

// GetInfo get application information 获取当前所选应用的基础信息
func (c *ApplicationController) GetInfo(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appid", "0"))
	info, err := c.service.GetAppInfo(appId)
	if err != nil {
		return mvc.FailWithMsg(nil, err.Error())
	}
	return mvc.Success(info)
}

// GetGitBranches get git addresses & branches for pipeline 获取当前 git 仓库下的所有分支和所对应的流水线
func (c *ApplicationController) GetGitBranches(ctx *context.HttpContext) mvc.ApiResult {
	appId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("appid", "0"))
	gitAddr := ctx.Input.QueryDefault("git", "")
	gitType := ctx.Input.QueryDefault("gitType", "")

	appInfo, err := c.service.GetAppInfo(appId)
	scId := appInfo.SCID
	if err != nil {
		if appInfo.SCID <= 0 {
			scId = utils.GetNumberOfParam[uint64](ctx, "scid")
		}
	}
	token := ""
	if scId > 0 {
		scInfo, _ := c.service.GetServiceConnectionById(scId)
		var detail dto.ServiceConnectionDetails
		_ = json.Unmarshal([]byte(scInfo.Detail), &detail)
		token = detail.Token
	}

	if appInfo.SourceType != "" {
		gitType = appInfo.SourceType
	}
	if appInfo.Git != "" {
		gitAddr = appInfo.Git
	}

	names, _ := app.GetGitBranches(gitAddr, gitType, token)
	return mvc.Success(context.H{
		"git":      gitAddr,
		"branches": names,
	})
}

// GetSearchDockerfile 根据所选择的 SVC地址，在地址内搜索目录下的 Dockerfile文件列表
func (c *ApplicationController) GetSearchDockerfile(ctx *context.HttpContext) mvc.ApiResult {
	scid, _ := utils.StringToUInt64(ctx.Input.QueryDefault("scid", "0"))
	gitAddr := ctx.Input.QueryDefault("git", "")
	ref := ctx.Input.QueryDefault("ref", "")
	if gitAddr == "" {
		return mvc.Fail("git address is empty")
	}
	gitType := ctx.Input.QueryDefault("gitType", "")
	token := ""
	if scid > 0 {
		scInfo, _ := c.service.GetServiceConnectionById(scid)
		var detail dto.ServiceConnectionDetails
		_ = json.Unmarshal([]byte(scInfo.Detail), &detail)
		token = detail.Token
	}
	dockerPaths, err := app.FindFiles(gitAddr, gitType, token, ref, "Dockerfile")
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(dockerPaths)
}

// GetBuildImageByLanguageId 根据页面所选择的开发语言出发流水线进行镜像的打包操作
func (c *ApplicationController) GetBuildImageByLanguageId(ctx *context.HttpContext) mvc.ApiResult {
	languageId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("languageId", "0"))
	list, err := c.pipelineService.GetBuildImageByLanguageId(languageId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(list)
}

// GetBuildImageByLanguages 根据多种开发开发进行镜像的编译
func (c *ApplicationController) GetBuildImageByLanguages(ctx *context.HttpContext) mvc.ApiResult {
	languageId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("languageId", "0"))
	aliasName := ctx.Input.QueryDefault("aliasName", "")
	list, err := c.pipelineService.GetBuildImageBy(aliasName, languageId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(list)
}

// PostBuildImage 触发流水线进行镜像的编译
func (c *ApplicationController) PostBuildImage(ctx *context.HttpContext) mvc.ApiResult {
	var request models.ApplicationLanguageCompile
	_ = ctx.BindWith(&request, binding.JSON)
	request.Status = 1
	err := c.pipelineService.AddOrEditBuildImage(request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}

// DeleteBuildImage 删除已经编译好的镜像
func (c *ApplicationController) DeleteBuildImage(ctx *context.HttpContext) mvc.ApiResult {
	id, _ := utils.StringToUInt64(ctx.Input.QueryDefault("id", "0"))
	err := c.pipelineService.DeleteBuildImage(id)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}

// PostNewPipeline new pipeline only by name & id 根据所填写的名称创建一条新的流水线
func (c *ApplicationController) PostNewPipeline(req *requests2.AppNewPipelineReq) mvc.ApiResult {
	err, pipeline := c.pipelineService.NewPipeline(req)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(pipeline.ID)
}

// GetPipelines get pipeline list by application id. 根据当前所选的应用，获取应用下的流水线
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

// PostEditPipeline Save pipeline information and DSL. 保存流水线的信息和流程设计
func (c *ApplicationController) PostEditPipeline(request *requests2.EditPipelineReq) mvc.ApiResult {
	err := c.pipelineService.UpdatePipeline(request)
	if err == nil {
		err = c.pipelineService.UpdateDSL(request)
	}
	if err != nil {
		return mvc.FailWithMsg(false, err.Error())
	}

	return mvc.Success(true)
}

// GetPipeline get pipeline frontend json by id. 根据流水线的 id 获取流水线的 JSON数据结构描述信息
func (c *ApplicationController) GetPipeline(ctx *context.HttpContext) mvc.ApiResult {
	pipelineId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("id", "0"))
	pipeline, err := c.pipelineService.GetPipelineById(pipelineId)
	if err != nil {
		return mvc.Fail("not found pipeline!")
	}
	return mvc.Success(pipeline)
}

// PostAbortPipeline 终止当前所选的正在运行的流水线
func (c *ApplicationController) PostAbortPipeline(request *requests2.AbortPipelineReq) mvc.ApiResult {
	err := c.pipelineService.AbortPipeline(request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}

// PostRunPipeline 执行当前所选的流水线的内容
func (c *ApplicationController) PostRunPipeline(request *requests2.RunPipelineReq) mvc.ApiResult {
	taskId, err := c.pipelineService.RunPipeline(request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(taskId)
}

// PostRunPipelineWithBranch 使用指定的 SVC分支来运行流水线
func (c *ApplicationController) PostRunPipelineWithBranch(request *requests2.RunPipelineReq) mvc.ApiResult {
	taskId, err := c.pipelineService.RunPipelineWithParameters(request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(taskId)
}

// PostPipelineStatus 变更流水线的状态
func (c *ApplicationController) PostPipelineStatus(request *requests2.PipelineStatusReq) mvc.ApiResult {
	err := c.pipelineService.UpdatePipelineStatus(request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}

// DeletePipeline 删除当前所选流水线
func (c *ApplicationController) DeletePipeline(ctx *context.HttpContext) mvc.ApiResult {
	pipelineId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("id", "0"))
	err := c.pipelineService.DeletePipeline(pipelineId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}

// GetPipelineDetails 获取当前流水线的明细信息
func (c *ApplicationController) GetPipelineDetails(httpContext *context.HttpContext) mvc.ApiResult {
	var request requests2.PipelineDetailsReq
	_ = httpContext.BindWithUri(&request)
	job, err := c.pipelineService.GetDetails(&request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(job)
}

// GetPipelineLogs 获取当前流水线的执行日志
func (c *ApplicationController) GetPipelineLogs(httpContext *context.HttpContext) mvc.ApiResult {
	var request requests2.PipelineDetailsReq
	_ = httpContext.BindWithUri(&request)
	logs, err := c.pipelineService.GetLogs(&request)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(logs)
}
