package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"gorm.io/gorm"
	"sgr/domain/business/kubernetes"
)

type ClusterController struct {
	mvc.ApiController

	db *gorm.DB
}

func NewClusterController(database *gorm.DB) *ClusterController {
	return &ClusterController{db: database}
}

func (controller ClusterController) GetPods(ctx *context.HttpContext) mvc.ApiResult {
	namespace := ctx.Input.QueryDefault("namespace", "")
	//tenantId := ctx.Input.QueryDefault("tid","")
	// get k8s cluster client by tenant id
	client, err := kubernetes.NewClientSet("")
	if err != nil {
		panic(err)
	}

	podList := kubernetes.GetPodList(client, namespace)

	return controller.OK(podList)
}

func (controller ClusterController) GetNamespaces(ctx *context.HttpContext) mvc.ApiResult {
	//tenantId := ctx.Input.QueryDefault("tid","")
	// get k8s cluster client by tenant id
	client, err := kubernetes.NewClientSet("")
	if err != nil {
		panic(err)
	}
	namespaces := kubernetes.GetAllNamespaces(client)
	return controller.OK(namespaces)
}
