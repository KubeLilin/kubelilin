package pipeline

import (
	"fmt"
	"log"
	"sgr/pkg/pipeline/templates"
	"strings"
)

type Options struct {
	jenkinsUrl       string // jenkins 地址
	jenkinsUser      string // jenkins 用户名
	jenkinsUserToken string // jenkins 用户token

	k8sNamespace     string // jenkins slave node running in k8s namespace
	dockerBuildImage string // jenkins slave node running docker image
}

type Builder struct {
	Options *Options
}

func NewBuilder() *Builder {
	return &Builder{&Options{}}
}

func (builder *Builder) UseKubernetes(namespace string, image string) *Builder {
	builder.Options.k8sNamespace = namespace
	builder.Options.dockerBuildImage = image
	return builder
}

func (builder *Builder) UseJenkins(url string, user string, token string) *Builder {
	builder.Options.jenkinsUrl = url
	builder.Options.jenkinsUser = user
	builder.Options.jenkinsUserToken = token
	return builder
}

func (builder *Builder) Build() (Pipeline, error) {
	return newWorkFlow(DriverJenkins.String(), builder.Options.jenkinsUrl, builder.Options.jenkinsUser, builder.Options.jenkinsUserToken,
		"", nil)
}

// WorkFlowProcessor 生成流水线代码 /**
func (builder *Builder) WorkFlowProcessor(inputParams []EnvItem, callback DeployRequest,
	checkoutSteps []StepItem, buildSteps []StepItem, imageSteps []StepItem) FlowProcessor {
	envVars := []EnvItem{
		{Key: "JENKINS_SLAVE_WORKSPACE", Value: "/home/jenkins/agent"},
		{Key: "ACCESS_TOKEN", Value: builder.Options.jenkinsUserToken},
	}
	envVars = append(envVars, inputParams...)

	containerTemplates := []ContainerEnv{
		{
			Name:       "jnlp",
			Image:      "jenkins/inbound-agent:4.10-3",
			WorkingDir: "/home/jenkins/agent",
		},
		{
			Name:       "build",
			Image:      builder.Options.dockerBuildImage,
			CommandArr: []string{"sleep"},
			ArgsArr:    []string{"99d"},
		},
	}

	checkoutItems := map[string]interface{}{"CheckoutItems": checkoutSteps}
	buildItems := map[string]interface{}{"BuildItems": buildSteps}
	imageItems := map[string]interface{}{"ImageItems": imageSteps}

	var taskPipelineXMLStrArr []string
	checkoutTasks, _ := GeneratePipelineXMLStr(templates.Checkout, checkoutItems)
	compileTasks, _ := GeneratePipelineXMLStr(templates.Compile, buildItems)
	buildTasks, _ := GeneratePipelineXMLStr(templates.BuildImage, imageItems)

	taskPipelineXMLStrArr = append(taskPipelineXMLStrArr, checkoutTasks)
	taskPipelineXMLStrArr = append(taskPipelineXMLStrArr, compileTasks)
	taskPipelineXMLStrArr = append(taskPipelineXMLStrArr, buildTasks)
	pipelineJson := strings.Join(taskPipelineXMLStrArr, " ")
	flowProcessor := &CIContext{
		EnvVars:            envVars,
		ContainerTemplates: containerTemplates,
		Stages:             pipelineJson,
		CommonContext: CommonContext{
			Namespace: builder.Options.k8sNamespace,
		},
		CallBack: callback,
	}
	return flowProcessor
}

// NewWorkFlow new workflow factory
func newWorkFlow(driver, addr, user, token, jobName string, flowProcessor FlowProcessor) (Pipeline, error) {
	var err error
	var workFlowProvider Pipeline
	switch {
	case driver == DriverJenkins.String():
		workFlowProvider, err = NewJenkinsClient(
			URL(addr),
			JenkinsUser(user),
			JenkinsToken(token),
			JenkinsJob(jobName),
			Processor(flowProcessor))

		if err != nil {
			log.Print(err)
			return nil, err
		}
		return workFlowProvider, nil
	}
	log.Print("work flow system not configured")
	return nil, fmt.Errorf("work flow system not configured")
}
