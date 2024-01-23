package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/app"
	"kubelilin/utils"
)

type PipelineController struct {
	mvc.ApiController

	pipelineService *app.PipelineService
}

func NewPipelineController(pipelineService *app.PipelineService) *PipelineController {
	return &PipelineController{pipelineService: pipelineService}
}

func (controller PipelineController) PostWebHook(request *requests.PipelineStatusRequest, ctx *context.HttpContext) mvc.ApiResult {
	//map[appid:2 branch:dev buildNumber:7 image:harbor.yoyogo.run/apps/pipeline-7-app-2:v7 pid:7 status:SUCCESS timestamp:Thu Sep 21 03:27:23 UTC 2023]
	pid, _ := utils.StringToUInt64(request.Pid)
	// write database for commit deliverables

	status := 2
	if request.Status != "SUCCESS" {
		status = 3
	}
	_ = controller.pipelineService.UpdatePipelineStatus(&requests.PipelineStatusReq{Id: pid, Status: int8(status)})

	return controller.OK("ok")
}
