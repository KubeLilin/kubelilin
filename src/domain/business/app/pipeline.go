package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drone/go-scm/scm"
	"github.com/yoyofx/yoyogo/abstractions"
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	pipelineV1 "kubelilin/pkg/pipeline"
	"kubelilin/utils"
	"strconv"
	"time"
)

type PipelineService struct {
	db                *gorm.DB
	jenkinsBuilder    *pipelineV1.Builder
	serviceConnection *ServiceConnectionService
	appservice        *ApplicationService
	config            abstractions.IConfiguration
	onReady           bool
}

func NewPipelineService(db *gorm.DB, jenkins *pipelineV1.Builder, sc *ServiceConnectionService, appService *ApplicationService, config abstractions.IConfiguration) *PipelineService {
	// set pipeline config
	peSC, err := sc.GetPipelineEngine()
	onReady := err == nil
	if onReady {
		jenkins.UseJenkins(peSC.Repo, peSC.UserName, peSC.Token).UseKubernetes(peSC.Password)
	} else {
		panic("not found pipeline settings. Pipeline be can't run.")
	}
	return &PipelineService{db: db, jenkinsBuilder: jenkins, serviceConnection: sc, appservice: appService, onReady: onReady, config: config}
}

func (pipelineService *PipelineService) GetBuildImageByLanguageId(languageId uint64) ([]models.ApplicationLanguageCompile, error) {
	var languageCompileList []models.ApplicationLanguageCompile
	dbRes := pipelineService.db.Model(&models.ApplicationLanguageCompile{}).Where("language_id=? AND status = 1", languageId).Order("compile_image, sort DESC").Find(&languageCompileList)
	return languageCompileList, dbRes.Error
}

func (pipelineService *PipelineService) GetBuildImageBy(alias string, languageId uint64) ([]models.ApplicationLanguageCompile, error) {
	var languageCompileList []models.ApplicationLanguageCompile
	dbRes := pipelineService.db.Model(&models.ApplicationLanguageCompile{}).Where("status = 1")
	if languageId > 0 {
		dbRes.Where("language_id=?", languageId)
	}
	if alias != "" {
		dbRes.Where("compile_image like ?", "%"+alias+"%")
	}
	dbRes.Order("compile_image, sort DESC").Find(&languageCompileList)
	return languageCompileList, dbRes.Error
}

func (pipelineService *PipelineService) AddOrEditBuildImage(request models.ApplicationLanguageCompile) error {
	if request.ID > 0 {
		return pipelineService.db.Model(&models.ApplicationLanguageCompile{}).Where("id=?", request.ID).Updates(&request).Error
	} else {
		return pipelineService.db.Model(&models.ApplicationLanguageCompile{}).Create(&request).Error
	}
}

func (pipelineService *PipelineService) DeleteBuildImage(id uint64) error {
	dbRes := pipelineService.db.Model(&models.ApplicationLanguageCompile{}).Where("id=?", id).Update("status", 0)
	return dbRes.Error
}

/*
NewPipeline 新建流水线(仅名称)
*/
func (pipelineService *PipelineService) NewPipeline(req *requests.AppNewPipelineReq) (error, *models.SgrTenantApplicationPipelines) {
	var exitCount int64
	pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("appid=? and name=?", req.AppId, req.Name).Count(&exitCount)
	if exitCount > 0 {
		return errors.New("already have the same name pipeline"), nil
	}

	now := time.Now()
	appModel := &models.SgrTenantApplicationPipelines{
		Appid:        req.AppId,
		Name:         req.Name,
		Dsl:          "",
		LastTaskID:   "",
		Status:       uint8(1),
		CreationTime: &now,
		UpdateTime:   &now,
	}
	dbRes := pipelineService.db.Model(models.SgrTenantApplicationPipelines{}).Create(appModel)
	if dbRes.Error != nil {
		return nil, appModel
	}

	return nil, appModel
}

/*
GetAppPipelines 获取流水线列表
*/
func (pipelineService *PipelineService) GetAppPipelines(appId uint64) ([]dto.PipelineInfo, error) {
	sql := `SELECT id,appid,name,dsl,taskStatus,lastTaskId,last_commit lastCommit FROM sgr_tenant_application_pipelines WHERE status=1 AND appid=?`
	var pipelineInfoList []dto.PipelineInfo
	err := pipelineService.db.Raw(sql, appId).Find(&pipelineInfoList).Error
	return pipelineInfoList, err
}

/*
GetPipelineById 按ID获取流水线
*/
func (pipelineService *PipelineService) GetPipelineById(id uint64) (dto.PipelineInfo, error) {
	sql := `SELECT pipe.id,pipe.appid,pipe.name,pipe.dsl,pipe.taskStatus,pipe.lastTaskId,lang.name language FROM sgr_tenant_application_pipelines pipe
INNER JOIN sgr_tenant_application app on app.id = pipe.appid
INNER JOIN sgr_code_application_language lang on lang.id = app.language
WHERE pipe.id = ?`
	var pipelineInfo dto.PipelineInfo
	err := pipelineService.db.Raw(sql, id).First(&pipelineInfo).Error
	return pipelineInfo, err
}

/*
UpdatePipeline 更新流水线基本信息
*/
func (pipelineService *PipelineService) UpdatePipeline(request *requests.EditPipelineReq) error {
	var pipelineInfo models.SgrTenantApplicationPipelines
	dbRes := pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).First(&pipelineInfo)
	if dbRes.Error != nil {
		return dbRes.Error
	}
	pipelineInfo.Name = request.Name
	pipelineInfo.Dsl = request.DSL
	taskStatus := uint(0)
	pipelineInfo.TaskStatus = &taskStatus
	now := time.Now()
	pipelineInfo.UpdateTime = &now

	dbRes = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).Updates(pipelineInfo)
	if dbRes.Error != nil {
		return nil
	}
	return nil
}

func (pipelineService *PipelineService) UpdateDSL(request *requests.EditPipelineReq) error {
	// Generate pipeline name and docker image name.
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	// pipeline json from frontend
	var pipelineStages []dto.StageInfo
	_ = json.Unmarshal([]byte(request.DSL), &pipelineStages)

	// global context
	context := make(map[string]string)
	context["deployUrl"] = "localhost"
	context["pipelineName"] = pipelineName
	// get config by db
	systemCallbackSC, err := pipelineService.serviceConnection.GetSystemCallback()
	if err != nil {
		return errors.New("not found system callback , pipeline will be not deploy application")
	}
	context["deployUrl"] = systemCallbackSC.Repo
	context["deployToken"] = systemCallbackSC.Token
	//deployUrl := pipelineService.config.GetString("kubelilin.deploy.url")

	//	 harbor config
	//harborAddress := pipelineService.config.GetString("hub.harbor.url")
	//harborToken := pipelineService.config.GetString("hub.harbor.token")
	imageHubSC, err := pipelineService.serviceConnection.GetImageHub()
	if err != nil {
		return errors.New("not found image hub settings , pipeline will be not deploy application")
	}
	context["imageHubAddress"] = imageHubSC.Repo
	context["imageHubToken"] = imageHubSC.Token

	// Conversion JSON to DSL
	// Set SGR_DOCKER_FILE value, that apply to code_build step .
	env := []pipelineV1.EnvItem{
		// {Key: "SGR_DOCKER_FILE", Value: "./examples/simpleweb/Dockerfile"},
		{Key: "SGR_REPOSITORY_NAME", Value: fmt.Sprintf("%s/apps/%s", context["imageHubAddress"], context["pipelineName"])},
		{Key: "SGR_REGISTRY_ADDR", Value: fmt.Sprintf("https://%s/", context["imageHubAddress"])},
		{Key: "SGR_REGISTRY_AUTH", Value: context["imageHubToken"]},
		{Key: "SGR_REGISTRY_CONFIG", Value: "/kaniko/.docker"},
		{Key: "PID", Value: request.Id},
		{Key: "APPID", Value: request.AppId},
	}
	//var buildImage string
	//var branch string
	//var deployId uint64
	var dslStageList []pipelineV1.StageItem
	for _, stage := range pipelineStages {
		dslStageItem := pipelineV1.StageItem{Name: stage.Name}
		for _, step := range stage.Steps {
			switch step.Key { // git_pull  , code_build  ,  cmd_shell ,  app_deploy
			case "git_pull":
				dslStageItem.Steps = append(dslStageItem.Steps, pipelineV1.StepItem{Name: step.Name,
					Command: fmt.Sprintf(`
					checkout([
                    	$class: 'GitSCM', branches: [[name: "${params.BRANCH_NAME}"]],
                    	doGenerateSubmoduleConfigurations: false,extensions: [[$class:'CheckoutOption',timeout:30],[$class:'CloneOption',depth:0,noTags:false,reference:'',shallow:false,timeout:30]], submoduleCfg: [],
                    	userRemoteConfigs: [[ url: "%s"]]
                	])`, step.Content["git"])})
				context["branch"] = step.Content["branch"].(string)
				break
			case "cmd_shell":
				dslStageItem.Steps = append(dslStageItem.Steps, pipelineV1.StepItem{Name: step.Name,
					Command: fmt.Sprintf(`
					sh '''
					 %s
					'''`, step.Content["shell"].(string))})
				break
			case "app_deploy":
				dslStageItem.Steps = append(dslStageItem.Steps, pipelineV1.StepItem{Name: step.Name,
					Command: fmt.Sprintf(`
				   script{
					   def rbody = "{\"wholeImage\": \"${env.SGR_REPOSITORY_NAME}:v${env.BUILD_NUMBER}\", \"IsDiv\":true , \"dpId\": %v, \"tenantId\": 0 }"
					   httpRequest acceptType: 'APPLICATION_JSON', contentType: 'APPLICATION_JSON', httpMode: 'POST', requestBody:rbody , responseHandle: 'NONE', timeout: 30, url: '%s/v1/deployment/executedeployment'
				   }
				`, step.Content["depolyment"], context["deployUrl"])})
				context["deployId"] = utils.ToString(uint64(step.Content["depolyment"].(float64)))
				break
			case "code_build":
				// 添加编译环境,Dockerfile 文件位置
				env = append(env, pipelineV1.EnvItem{Key: "SGR_DOCKER_FILE", Value: step.Content["buildFile"].(string)})
				// 添加编译环境 编译镜像：版本
				buildEnv := step.Content["buildEnv"].(string)
				buildImage, hasBuildImage := step.Content["buildImage"].(string)
				if hasBuildImage {
					context["buildImage"] = buildImage
				} else { // 没有镜像设置,则使用默认语言镜像
					context["buildImage"] = getDefaultBuildImageByLanguage(buildEnv)

				}

				dslStageItem.Steps = append(dslStageItem.Steps, pipelineV1.StepItem{Name: "code build",
					Command: fmt.Sprintf(`
					container('build') {
						sh '''
						%s
						'''
					}`, step.Content["buildScript"])})

				dslStageItem.Steps = append(dslStageItem.Steps, pipelineV1.StepItem{Name: "docker build",
					Command: `
					container('docker') {
						sh "[ -d $SGR_REGISTRY_CONFIG ] || mkdir -pv $SGR_REGISTRY_CONFIG"
   						sh """#!/busybox/sh -e
                    		echo '{"auths": {"'$SGR_REGISTRY_ADDR'": {"auth": "'$SGR_REGISTRY_AUTH'"}}}' > $SGR_REGISTRY_CONFIG/config.json
						"""
						sh '''#!/busybox/sh
							/kaniko/executor -f $SGR_DOCKER_FILE -c . --destination=$SGR_REPOSITORY_NAME:v$BUILD_NUMBER  --insecure --ignore-path=/product_uuid --skip-tls-verify -v=debug
						''' 
					}`})
				break
			case "publish_notify":
				dslStageItem.Steps = append(dslStageItem.Steps, pipelineV1.StepItem{Name: step.Name,
					Command: fmt.Sprintf(`
				   script{
					   def rbody = "{\"version\": \"v${env.BUILD_NUMBER}\",  \"dpId\": %v, \"branch\": \"%s\" , \"notifyType\": \"%s\" , \"notifyKey\": \"%s\" }"
					   httpRequest acceptType: 'APPLICATION_JSON', contentType: 'APPLICATION_JSON', httpMode: 'POST', requestBody:rbody , responseHandle: 'NONE', timeout: 30, url: '%s/v1/deployment/notify'
				   }
				`, context["deployId"], context["branch"], step.Content["notifyType"], step.Content["notifyKey"], context["deployUrl"])})
				break
			}
		}

		dslStageList = append(dslStageList, dslStageItem)
	}

	stageItems := map[string]interface{}{"pipelineStages": dslStageList}
	parameters := []pipelineV1.ParamItem{
		{Name: "BRANCH_NAME", DefaultValue: context["branch"], Description: "分支名称"},
		{Name: "VERSION", DefaultValue: "1.0", Description: "构建序号"},
	}

	builder := pipelineService.jenkinsBuilder.UseBuildImage(context["buildImage"])

	processor := builder.CICDProcessor(parameters, env, stageItems, &pipelineV1.DeployRequest{URL: context["deployUrl"] + "/v1/pipeline/webhook"})
	pipeline, _ := builder.Build()

	return pipeline.SaveJob(pipelineName, processor)
}

func getDefaultBuildImageByLanguage(languageName string) string {
	buildImage := ""
	switch languageName {
	case "java":
		buildImage = "maven:3.8.4-jdk-8"
		break
	case "golang":
		buildImage = "golang:1.18.1"
		break
	case "nodejs":
		buildImage = "node:14-alpine"
	case "dotnet":
		buildImage = "mcr.microsoft.com/dotnet/sdk:5.0"
	}
	return buildImage
}

func (pipelineService *PipelineService) AbortPipeline(request *requests.AbortPipelineReq) error {
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	builder := pipelineService.jenkinsBuilder
	pipeline, _ := builder.Build()
	return pipeline.Abort(pipelineName, request.TaskId)
}

func (pipelineService *PipelineService) RunPipeline(request *requests.RunPipelineReq) (int64, error) {
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	builder := pipelineService.jenkinsBuilder
	pipeline, _ := builder.Build()
	taskId, err := pipeline.RunJob(pipelineName)

	//get git last commit
	appInfo, _ := pipelineService.appservice.GetAppInfo(request.AppId)
	token := ""
	if appInfo.SCID > 0 {
		scInfo, _ := pipelineService.appservice.GetServiceConnectionById(appInfo.SCID)
		var detail dto.ServiceConnectionDetails
		_ = json.Unmarshal([]byte(scInfo.Detail), &detail)
		token = detail.Token
	}
	var commit *scm.Commit
	if appInfo.Git != "" {
		commit, _ = GetLastCommit(appInfo.Git, appInfo.SourceType, token)
	}
	// update databse
	var pipelineInfo models.SgrTenantApplicationPipelines
	_ = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).First(&pipelineInfo)
	now := time.Now()
	pipelineInfo.UpdateTime = &now
	pipelineInfo.LastTaskID = strconv.FormatInt(taskId, 10)
	taskStatus := uint(1)
	pipelineInfo.TaskStatus = &taskStatus
	if commit != nil {
		commitMessage, _ := json.Marshal(commit)
		pipelineInfo.LastCommit = string(commitMessage)
	}
	_ = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).Updates(pipelineInfo)

	return taskId, err
}

func (pipelineService *PipelineService) RunPipelineWithParameters(request *requests.RunPipelineReq) (int64, error) {
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	builder := pipelineService.jenkinsBuilder
	pipeline, _ := builder.Build()
	taskId, err := pipeline.RunJobWithParameters(pipelineName, request.Branch)

	//get git last commit
	appInfo, _ := pipelineService.appservice.GetAppInfo(request.AppId)
	token := ""
	if appInfo.SCID > 0 {
		scInfo, _ := pipelineService.appservice.GetServiceConnectionById(appInfo.SCID)
		var detail dto.ServiceConnectionDetails
		_ = json.Unmarshal([]byte(scInfo.Detail), &detail)
		token = detail.Token
	}
	var commit *scm.Commit
	if appInfo.Git != "" {
		commit, _ = GetLastCommit(appInfo.Git, appInfo.SourceType, token)
	}
	// update databse
	var pipelineInfo models.SgrTenantApplicationPipelines
	_ = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).First(&pipelineInfo)
	now := time.Now()
	pipelineInfo.UpdateTime = &now
	pipelineInfo.LastTaskID = strconv.FormatInt(taskId, 10)
	taskStatus := uint(1)
	pipelineInfo.TaskStatus = &taskStatus
	if commit != nil {
		commitMessage, _ := json.Marshal(commit)
		pipelineInfo.LastCommit = string(commitMessage)
	}
	_ = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).Updates(pipelineInfo)

	return taskId, err
}

func (pipelineService *PipelineService) UpdatePipelineStatus(request *requests.PipelineStatusReq) error {
	// update databse
	var pipelineInfo models.SgrTenantApplicationPipelines
	_ = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).First(&pipelineInfo)
	now := time.Now()
	pipelineInfo.UpdateTime = &now
	taskStatus := uint(request.Status)
	pipelineInfo.TaskStatus = &taskStatus
	dbRes := pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).Updates(pipelineInfo)
	return dbRes.Error
}

func (pipelineService *PipelineService) DeletePipeline(pipelineId uint64) error {
	sql := `update sgr_tenant_application_pipelines SET status=0 where id=?`
	return pipelineService.db.Exec(sql, pipelineId).Error
}

func (pipelineService *PipelineService) GetDetails(request *requests.PipelineDetailsReq) (*pipelineV1.JobInfo, error) {
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	builder := pipelineService.jenkinsBuilder
	pipeline, _ := builder.Build()
	return pipeline.GetJobInfo(pipelineName, request.TaskId)

	//job1.Result (IN_PROGRESS, SUCCESS , FAILED , ABORTED)
}
func (pipelineService *PipelineService) GetLogs(request *requests.PipelineDetailsReq) (string, error) {
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	builder := pipelineService.jenkinsBuilder
	pipeline, _ := builder.Build()
	return pipeline.GetJobLogs(pipelineName, request.TaskId)
}
