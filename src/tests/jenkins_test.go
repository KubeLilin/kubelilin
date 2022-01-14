package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	pipelineV1 "sgr/pkg/pipeline"
	"testing"
	"time"
)

func TestGetJenkinsJob(t *testing.T) {
	builder := pipelineV1.NewBuilder()
	builder.UseJenkins("http://152.136.141.235:32001", "jenkins", "11e681bb454a36a9ce0e0a6fd030d059a9").
		UseKubernetes("sgr-ci", "golang:1.16.5")

	pipeline, _ := builder.Build()

	ping, err := pipeline.Ping()
	if err != nil {
		return
	}
	fmt.Printf("jenkins version  %s", ping)

	job1, _ := pipeline.GetJobInfo("sample-pipeline-test", 25)

	assert.Equal(t, job1 != nil, true)
}

func TestJenkinsJob(t *testing.T) {

	builder := pipelineV1.NewBuilder()
	builder.UseJenkins("http://152.136.141.235:32001", "jenkins", "11d32a54cd6150bd626d8ed73c3bfa02d6").
		UseKubernetes("sgr-ci", "golang:1.16.5")

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
