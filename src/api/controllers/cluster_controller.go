package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/kubernetes"
	"strconv"
)

type ClusterController struct {
	mvc.ApiController
	clusterService *kubernetes.ClusterService
}

func NewClusterController(clusterService *kubernetes.ClusterService) *ClusterController {
	return &ClusterController{clusterService: clusterService}
}

func (controller ClusterController) GetPods(ctx *context.HttpContext) mvc.ApiResult {
	namespace := ctx.Input.QueryDefault("namespace", "")
	k8sapp := ctx.Input.QueryDefault("app", "")

	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.Atoi(strCid)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	podList := kubernetes.GetPodList(client, namespace, k8sapp)

	return controller.OK(podList)
}

func (controller ClusterController) GetNamespaces(ctx *context.HttpContext) mvc.ApiResult {
	//tenantId := ctx.Input.QueryDefault("tid","")
	// get k8s cluster client by tenant id
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.Atoi(strCid)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	namespaces := kubernetes.GetAllNamespaces(client)
	return controller.OK(namespaces)
}

func (controller ClusterController) GetDeployments(ctx *context.HttpContext) mvc.ApiResult {
	namespace := ctx.Input.QueryDefault("namespace", "")
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.Atoi(strCid)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	list := kubernetes.GetDeploymentList(client, namespace)
	return controller.OK(list)
}

func (controller ClusterController) GetNodes(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.Atoi(strCid)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	nodeList := kubernetes.GetNodeList(client)
	return controller.OK(nodeList)
}

func (controller ClusterController) GetList(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	tenantList, _ := controller.clusterService.GetClustersByTenant(userInfo.TenantID)
	return controller.OK(tenantList)
}
