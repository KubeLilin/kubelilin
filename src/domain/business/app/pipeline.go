package app

import (
	"encoding/json"
	"errors"
	"github.com/yoyofx/yoyogo/abstractions"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	pipelineV1 "sgr/pkg/pipeline"
	"time"
)

type PipelineService struct {
	db     *gorm.DB
	config abstractions.IConfiguration
}

func NewPipelineService(db *gorm.DB, config abstractions.IConfiguration) *PipelineService {
	return &PipelineService{db: db, config: config}
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
	now := time.Now()
	pipelineInfo.UpdateTime = &now

	dbRes = pipelineService.db.Model(&models.SgrTenantApplicationPipelines{}).Where("id=?", request.Id).Updates(pipelineInfo)
	if dbRes.Error != nil {
		return nil
	}
	return nil
}

func (pipelineService *PipelineService) UpdateDSL(request *req.EditPipelineReq) error {
	var pipelineStages []dto.StageInfo
	_ = json.Unmarshal([]byte(request.DSL), &pipelineStages)

	lilinHost := pipelineService.config.GetString("kubelilin.host")
	if lilinHost == "" {
		lilinHost = "localhost"
	}
	jenkinsUrl := pipelineService.config.GetString("pipeline.jenkins.url")
	jenkinsToken := pipelineService.config.GetString("pipeline.jenkins.token")
	jenkinsUser := pipelineService.config.GetString("pipeline.jenkins.username")
	jenkinsNamespace := pipelineService.config.GetString("pipeline.jenkins.k8s-namespace")
	buildImage := ""
	builder := pipelineV1.NewBuilder()
	builder.UseJenkins(jenkinsUrl, jenkinsUser, jenkinsToken).
		UseKubernetes(jenkinsNamespace).UseBuildImage(buildImage)

	//harborAddress := pipelineService.config.GetString("hub.harbor.url")
	//pipelineName := fmt.Sprintf("pipeline-%s-app-%s", request.Name, request.AppId)
	//imageName := fmt.Sprintf("app-%s-pipeline-%s", request.AppId, request.Id)
	//buildImage := ""
	//
	//// 转换DSL
	//env := []pipelineV1.EnvItem{
	//	// {Key: "SGR_DOCKER_FILE", Value: "./examples/simpleweb/Dockerfile"},
	//	{Key: "SGR_REPOSITORY_NAME", Value: fmt.Sprintf("%s/apps/%s", harborAddress, imageName)},
	//	{Key: "SGR_REGISTRY_ADDR", Value: fmt.Sprintf("https://%s/", harborAddress)},
	//	{Key: "SGR_REGISTRY_AUTH", Value: "YWRtaW46SGFyYm9yMTIzNDU="},
	//	{Key: "SGR_REGISTRY_CONFIG", Value: "/kaniko/.docker"},
	//}

	return nil
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
