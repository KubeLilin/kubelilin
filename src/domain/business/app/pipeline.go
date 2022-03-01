package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	pipelineV1 "sgr/pkg/pipeline"
	"strconv"
	"time"
)

type PipelineService struct {
	db             *gorm.DB
	jenkinsBuilder *pipelineV1.Builder
	config         abstractions.IConfiguration
}

func NewPipelineService(db *gorm.DB, jenkins *pipelineV1.Builder, config abstractions.IConfiguration) *PipelineService {
	return &PipelineService{db: db, jenkinsBuilder: jenkins, config: config}
}

/*
GetBuildScripts 获取各语言流水线构建编译命令
*/
func (pipelineService *PipelineService) GetBuildScripts() map[string]string {
	return map[string]string{
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
		"dotnet": `# 编译命令，注：当前已在代码根路径下
dotnet restore
dotnet publish -p:PublishSingleFile=true -r linux-musl-x64 --self-contained true -p:PublishTrimmed=True -p:TrimMode=Link -c Release -o /app/publish                       
`,
	}

}

/*
NewPipeline 新建流水线(仅名称)
*/
func (pipelineService *PipelineService) NewPipeline(req *req.AppNewPipelineReq) (error, *models.SgrTenantApplicationPipelines) {
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
	sql := `SELECT id,appid,name,dsl,taskStatus,lastTaskId FROM sgr_tenant_application_pipelines WHERE appid=?`
	var pipelineInfoList []dto.PipelineInfo
	err := pipelineService.db.Raw(sql, appId).Find(&pipelineInfoList).Error
	return pipelineInfoList, err
}

/*
GetPipelineById 按ID获取流水线
*/
func (pipelineService *PipelineService) GetPipelineById(id uint64) (dto.PipelineInfo, error) {
	sql := `SELECT id,appid,name,dsl,taskStatus,lastTaskId FROM sgr_tenant_application_pipelines WHERE id=?`
	var pipelineInfo dto.PipelineInfo
	err := pipelineService.db.Raw(sql, id).First(&pipelineInfo).Error
	return pipelineInfo, err
}

/*
UpdatePipeline 更新流水线基本信息
*/
func (pipelineService *PipelineService) UpdatePipeline(request *req.EditPipelineReq) error {
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

func (pipelineService *PipelineService) UpdateDSL(request *req.EditPipelineReq) error {
	// Generate pipeline name and docker image name.
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	imageName := pipelineName
	// pipeline json from frontend
	var pipelineStages []dto.StageInfo
	_ = json.Unmarshal([]byte(request.DSL), &pipelineStages)

	// get config by configuration
	lilinHost := pipelineService.config.GetString("kubelilin.host")
	if lilinHost == "" {
		lilinHost = "localhost"
	}
	//// 	 jenkins config
	//jenkinsUrl := pipelineService.config.GetString("pipeline.jenkins.url")
	//jenkinsToken := pipelineService.config.GetString("pipeline.jenkins.token")
	//jenkinsUser := pipelineService.config.GetString("pipeline.jenkins.username")
	//jenkinsNamespace := pipelineService.config.GetString("pipeline.jenkins.k8s-namespace")
	//	 harbor config
	harborAddress := pipelineService.config.GetString("hub.harbor.url")
	harborToken := pipelineService.config.GetString("hub.harbor.token")

	// Conversion JSON to DSL
	// Set SGR_DOCKER_FILE value, that apply to code_build step .
	env := []pipelineV1.EnvItem{
		// {Key: "SGR_DOCKER_FILE", Value: "./examples/simpleweb/Dockerfile"},
		{Key: "SGR_REPOSITORY_NAME", Value: fmt.Sprintf("%s/apps/%s", harborAddress, imageName)},
		{Key: "SGR_REGISTRY_ADDR", Value: fmt.Sprintf("https://%s/", harborAddress)},
		{Key: "SGR_REGISTRY_AUTH", Value: harborToken},
		{Key: "SGR_REGISTRY_CONFIG", Value: "/kaniko/.docker"},
	}
	var buildImage string

	var dslStageList []pipelineV1.StageItem
	for _, stage := range pipelineStages {
		dslStageItem := pipelineV1.StageItem{Name: stage.Name}
		for _, step := range stage.Steps {
			switch step.Key { // git_pull  , code_build  ,  cmd_shell ,  app_deploy
			case "git_pull":
				dslStageItem.Steps = append(dslStageItem.Steps, pipelineV1.StepItem{Name: step.Name,
					Command: fmt.Sprintf("	git url: '%s', branch: '%s'", step.Content["git"], step.Content["branch"])})
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
					sh '''
					curl -H "Accept: application/json" -H "Content-type: application/json" -X POST -d "{"wholeImage": "${SGR_REPOSITORY_NAME}:v${BUILD_NUMBER}", "IsDiv":true , "dpId": %v, "tenantId": 0 }" https://%s/v1/deployment/executedeployment
					# echo "{"wholeImage": "${SGR_REPOSITORY_NAME}:v${BUILD_NUMBER}", "IsDiv":true , "dpId": 1, "tenantId": 0 }"
					'''`, lilinHost, step.Content["depolyment"])})
				break
			case "code_build":
				// 添加编译环境,Dockerfile 文件位置
				env = append(env, pipelineV1.EnvItem{Key: "SGR_DOCKER_FILE", Value: step.Content["buildFile"].(string)})
				// 添加编译环境 编译镜像：版本
				buildEnv := step.Content["buildEnv"].(string)
				buildImage = getBuildImageByLanguage(buildEnv)

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
   						sh """
                    		echo '{"auths": {"'$SGR_REGISTRY_ADDR'": {"auth": "'$SGR_REGISTRY_AUTH'"}}}' > $SGR_REGISTRY_CONFIG/config.json
						"""
						sh '''#!/busybox/sh
							/kaniko/executor -f $SGR_DOCKER_FILE -c . --destination=$SGR_REPOSITORY_NAME:v$BUILD_NUMBER  --insecure --skip-tls-verify -v=debug
						''' 
					}`})
				break
			}
		}

		dslStageList = append(dslStageList, dslStageItem)
	}

	stageItems := map[string]interface{}{"pipelineStages": dslStageList}

	// connect jenkins and save the job
	//builder := pipelineV1.NewBuilder()
	//builder.UseJenkins(jenkinsUrl, jenkinsUser, jenkinsToken).
	//	UseKubernetes(jenkinsNamespace).UseBuildImage(buildImage)

	builder := pipelineService.jenkinsBuilder.UseBuildImage(buildImage)

	processor := builder.CICDProcessor(env, stageItems)
	pipeline, _ := builder.Build()

	return pipeline.SaveJob(pipelineName, processor)
}

func getBuildImageByLanguage(languageName string) string {
	buildImage := ""
	switch languageName {
	case "java":
		buildImage = "maven:3.8.4-jdk-8"
		break
	case "golang":
		buildImage = "golang:1.16.5"
		break
	case "nodejs":
		buildImage = "node:16-alpine"
	case "dotnet":
		buildImage = "mcr.microsoft.com/dotnet/sdk:5.0"
	}
	return buildImage
}

func (pipelineService *PipelineService) RunPipeline(request *req.RunPipelineReq) (int64, error) {
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	builder := pipelineService.jenkinsBuilder
	pipeline, _ := builder.Build()
	taskId, err := pipeline.RunJob(pipelineName)

	// update databse
	var pipelineInfo models.SgrTenantApplicationPipelines
	_ = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).First(&pipelineInfo)
	now := time.Now()
	pipelineInfo.UpdateTime = &now
	pipelineInfo.LastTaskID = strconv.FormatInt(taskId, 10)
	taskStatus := uint(1)
	pipelineInfo.TaskStatus = &taskStatus
	_ = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).Updates(pipelineInfo)

	return taskId, err
}

func (pipelineService *PipelineService) UpdatePipelineStatus(request *req.PipelineStatusReq) error {
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

func (pipelineService *PipelineService) GetDetails(request *req.PipelineDetailsReq) (*pipelineV1.JobInfo, error) {
	pipelineName := fmt.Sprintf("pipeline-%v-app-%v", request.Id, request.AppId)
	builder := pipelineService.jenkinsBuilder
	pipeline, _ := builder.Build()
	return pipeline.GetJobInfo(pipelineName, request.TaskId)

	//job1.Result (IN_PROGRESS, SUCCESS , FAILED , ABORTED)
}
