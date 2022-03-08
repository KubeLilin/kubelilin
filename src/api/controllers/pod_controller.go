package controllers

import (
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/kubernetes"
)

type PodController struct {
	mvc.ApiController
	clusterService *kubernetes.ClusterService
}

func NewPodController(clusterService *kubernetes.ClusterService) *PodController {
	return &PodController{clusterService: clusterService}
}

func (controller PodController) GetTerminal(ctx *context.HttpContext) {
	var request req.PodTerminalExecRequest
	_ = ctx.BindWithUri(&request)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(request.TenantId, request.ClusterId)
	config, _ := controller.clusterService.GetClusterConfig(request.TenantId, request.ClusterId)

	web.Upgrade(ctx, func(conn *websocket.Conn) {
		terminal := kubernetes.NewWebTerminal(conn)
		err := kubernetes.Exec(client, config, terminal, request.Namespace, request.PodName, request.ContainerName)
		if err != nil {
			msg := fmt.Sprintf("Exec to pod error! err: %v", err)
			_, _ = terminal.Write([]byte(msg))
			terminal.Done()
		}
	})
}
