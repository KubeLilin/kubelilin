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
	k8snode := ctx.Input.QueryDefault("node", "")

	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	podList := kubernetes.GetPodList(client, namespace, k8snode, k8sapp)

	return controller.OK(podList)
}

func (controller ClusterController) GetNamespaces(ctx *context.HttpContext) mvc.ApiResult {
	//tenantId := ctx.Input.QueryDefault("tid","")
	// get k8s cluster client by tenant id
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	namespaces := kubernetes.GetAllNamespaces(client)
	return controller.OK(namespaces)
}

func (controller ClusterController) GetNamespacesFromDB(ctx *context.HttpContext) mvc.ApiResult {
	//tenantId := ctx.Input.QueryDefault("tid","")
	// get k8s cluster client by tenant id
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.Atoi(strCid)
	res := controller.clusterService.GetNameSpacesFromDB(userInfo.TenantID, cid)
	return controller.OK(res)
}

func (controller ClusterController) GetDeployments(ctx *context.HttpContext) mvc.ApiResult {
	namespace := ctx.Input.QueryDefault("namespace", "")
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	list := kubernetes.GetDeploymentList(client, namespace)
	return controller.OK(list)
}

func (controller ClusterController) GetNodes(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	nodeList := kubernetes.GetNodeList(client)
	return controller.OK(nodeList)
}

func (controller ClusterController) GetList(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	clusterName := ctx.Input.Query("name")
	tenantClusterList, _ := controller.clusterService.GetClustersByTenant(userInfo.TenantID, clusterName)
	return controller.OK(tenantClusterList)
}

func (controller ClusterController) PostClusterByConfig(ctx *context.HttpContext, request *req.ImportClusterReq) mvc.ApiResult {
	_, k8sFile, err := ctx.Input.FormFile("file1")
	if err != nil {
		return controller.Fail(err.Error())
	}
	configFile, _ := k8sFile.Open()
	userInfo := req.GetUserInfo(ctx)
	config, err := controller.clusterService.ImportK8sConfig(configFile, request.NickName, uint64(userInfo.TenantID))
	if err == nil {
		return controller.OK(config)
	}

	return controller.Fail(err.Error())
}

func (controller ClusterController) DeleteDelClusterInfo(ctx *context.HttpContext) mvc.ApiResult {
	id := ctx.Input.Query("id")
	clusterId, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		controller.clusterService.DeleteCluster(clusterId)
		return controller.OK(err == nil)
	}
	return controller.Fail(err.Error())

}

func (controller ClusterController) PutNewNamespace(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := req.GetUserInfo(ctx)
	cid := ctx.Input.QueryDefault("cid", "0")
	clusterId, _ := strconv.ParseUint(cid, 10, 64)
	namespace := ctx.Input.QueryDefault("namespace", "default")

	clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, clusterId)
	err := kubernetes.CreateNamespace(clientSet, namespace)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(err == nil)
}
