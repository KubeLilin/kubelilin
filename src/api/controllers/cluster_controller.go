package controllers

import (
	contextv1 "context"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"gorm.io/gorm"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (controller ClusterController) GetDeployments(ctx *context.HttpContext) mvc.ApiResult {
	emptyOptions := v1.ListOptions{}
	client, _ := kubernetes.NewClientSet("")
	list, _ := client.AppsV1().Deployments("yoyogo").List(contextv1.TODO(), emptyOptions)
	return controller.OK(list.Items)
}
