package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sgr/domain/dto"
	pipelineV1 "sgr/pkg/pipeline"
	"testing"
	"time"
)

func TestGetJenkinsJob(t *testing.T) {
	builder := pipelineV1.NewBuilder()
	builder.UseJenkins("http://152.136.141.235:32001", "jenkins", "11e681bb454a36a9ce0e0a6fd030d059a9").
		UseKubernetes("sgr-ci").UseBuildImage("golang:1.16.5")

	pipeline, _ := builder.Build()

	ping, err := pipeline.Ping()
	if err != nil {
		return
	}
	fmt.Printf("jenkins version  %s", ping)

	job1, _ := pipeline.GetJobInfo("sample-pipeline-test", 25)

	assert.Equal(t, job1 != nil, true)
}

// add /etc/hosts jenkins.sgr-ci IP
func TestJenkinsJob(t *testing.T) {

	builder := pipelineV1.NewBuilder()
	builder.UseJenkins("http://jenkins.sgr-ci:32001", "jenkins", "11d32a54cd6150bd626d8ed73c3bfa02d6").
		UseKubernetes("sgr-ci").UseBuildImage("golang:1.16.5")

	pipeline, _ := builder.Build()

	ping, err := pipeline.Ping()
	if err != nil {
		return
	}
	fmt.Printf("jenkins version  %s", ping)

	assert.Equal(t, ping != "", true)

	processor := builder.WorkFlowProcessor(
		[]pipelineV1.EnvItem{
			{Key: "SGR_REPOSITORY_NAME", Value: "yoyofx/yoyogo-demo"},
			{Key: "SGR_DOCKER_FILE", Value: "./examples/simpleweb/Dockerfile"},

			{Key: "SGR_REGISTRY_ADDR", Value: "https://index.docker.io/v1/"},
			{Key: "SGR_REGISTRY_AUTH", Value: "eW95b2Z4OnpsMTI1MzMwMw=="},
			{Key: "SGR_REGISTRY_CONFIG", Value: "/kaniko/.docker"},
		},
		//&pipelineV1.DeployRequest{
		//	Token: "xx1123123123df",
		//	URL:   "http://152.136.141.235:32001",
		//	Body:  "{\"publish_job_id\": 1}",
		//},
		nil,
		[]pipelineV1.StepItem{ // Checkout 阶段 ： 只能串行
			{
				Name:    "git",
				Command: "git url: 'https://gitee.com/yoyofx/yoyogo.git', branch: 'dev'",
			},
			{
				Name: "ls",
				Command: `
				sh '''
				ls
				'''
				`,
			},
		},
		[]pipelineV1.StepItem{ // Compile 阶段：  并行执行
			{
				Name:          "go build",
				ContainerName: "build",
				Command: `
			  sh '''
				echo "hello world"
				go version
				go env -w GOPROXY=https://goproxy.cn,direct
				go test -v ./tests/
			  '''
			`,
			},
		},
		[]pipelineV1.StepItem{ // Build 阶段 ：并行执行
			{
				Name:          "golang_image_build",
				ContainerName: "docker",
				Command: `
				sh "[ -d $SGR_REGISTRY_CONFIG ] || mkdir -pv $SGR_REGISTRY_CONFIG"
   				sh """
                    echo '{"auths": {"'$SGR_REGISTRY_ADDR'": {"auth": "'$SGR_REGISTRY_AUTH'"}}}' > $SGR_REGISTRY_CONFIG/config.json
				"""
				sh '''#!/busybox/sh
					/kaniko/executor -f $SGR_DOCKER_FILE -c . --destination=$SGR_REPOSITORY_NAME:$BUILD_NUMBER  --insecure --skip-tls-verify -v=debug
				''' `,
			},
		},
	)

	//docker login index.docker.io --username=yoyofx --password=zl1253303
	//docker build -t $SGR_REPOSITORY_NAME:$BUILD_NUMBER -f $SGR_DOCKER_FILE .
	//	docker push yoyofx/yoyogo-demo:$BUILD_NUMBER

	_ = pipeline.SaveJob("sample-pipeline-test", processor)

	runId, _ := pipeline.RunJob("sample-pipeline-test")

	assert.Equal(t, runId > 0, true)

	time.Sleep(1500)

	job, _ := pipeline.GetJobInfo("sample-pipeline-test", runId)

	assert.Equal(t, job != nil, true)

	logs, err := pipeline.GetJobLogs("sample-pipeline-test", runId)

	assert.Equal(t, logs != "", true)
}

func TestParsePipelineJob(t *testing.T) {
	pipelineJson := `[{"name":"代码","steps":[{"name":"拉取代码","key":"git_pull","save":true,"content":{"git":"https://gitee.com/yoyofx/yoyogo.git","branch":"dev"}}]},{"name":"编译构建","steps":[{"name":"编译命令","key":"code_build","save":true,"content":{"buildEnv":"golang","buildScript":"# 编译命令，注：当前已在代码根路径下\ngo env -w GOPROXY=https://goproxy.cn,direct\ngo build -ldflags=\"-s -w\" -o app .\n","buildFile":"./examples/simpleweb/Dockerfile"}},{"name":"命令执行","key":"cmd_shell","save":true,"content":{"shell":"# bash"}}]},{"name":"部署","steps":[{"name":"应用部署","key":"app_deploy","save":true,"content":{"depolyment":1}}]},{"name":"通知","steps":[{"name":"命令执行","key":"cmd_shell","save":true,"content":{"shell":"# bash"}}]}]`
	var pipelineStages []dto.StageInfo
	_ = json.Unmarshal([]byte(pipelineJson), &pipelineStages)

	lilinHost := "api.xx.com"
	imageName := "dev-yoyogo-demo-cls-hbktlqm5"
	harborAddress := "harbor.xiaocui.site"
	buildImage := ""

	// 转换DSL
	env := []pipelineV1.EnvItem{
		// {Key: "SGR_DOCKER_FILE", Value: "./examples/simpleweb/Dockerfile"},
		{Key: "SGR_REPOSITORY_NAME", Value: fmt.Sprintf("%s/apps/%s", harborAddress, imageName)},
		{Key: "SGR_REGISTRY_ADDR", Value: fmt.Sprintf("https://%s/", harborAddress)},
		{Key: "SGR_REGISTRY_AUTH", Value: "YWRtaW46SGFyYm9yMTIzNDU="},
		{Key: "SGR_REGISTRY_CONFIG", Value: "/kaniko/.docker"},
	}

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
					# curl -H "Accept: application/json" -H "Content-type: application/json" -X POST -d "{"wholeImage": "${SGR_REPOSITORY_NAME}:v${BUILD_NUMBER}", "IsDiv":true , "dpId": %v, "tenantId": 0 }" https://%s/v1/deployment/executedeployment
					echo "{"wholeImage": "${SGR_REPOSITORY_NAME}:v${BUILD_NUMBER}", "IsDiv":true , "dpId": 1, "tenantId": 0 }"
					'''`, lilinHost, step.Content["depolyment"])})
				break
			case "code_build":
				// 添加编译环境,Dockerfile 文件位置
				env = append(env, pipelineV1.EnvItem{Key: "SGR_DOCKER_FILE", Value: step.Content["buildFile"].(string)})
				// 添加编译环境 编译镜像：版本
				buildEnv := step.Content["buildEnv"].(string)
				switch buildEnv {
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

	//pipelineDSL, _ := pipelineV1.GeneratePipelineXMLStr(templates.CICD, stageItems)

	builder := pipelineV1.NewBuilder()
	builder.UseJenkins("http://152.136.141.235:32001", "jenkins", "11d32a54cd6150bd626d8ed73c3bfa02d6").
		UseKubernetes("sgr-ci").UseBuildImage(buildImage)

	processor := builder.CICDProcessor(env, stageItems)
	pipeline, _ := builder.Build()
	_ = pipeline.SaveJob("jenkins-unit-test", processor)

}
