package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/domain/dto"
	"kubelilin/utils"
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

func (controller ClusterController) GetNamespaceList(ctx *context.HttpContext) mvc.ApiResult {
	strCid := ctx.Input.QueryDefault("cid", "0")
	tenantName := ctx.Input.QueryDefault("tenant", "")
	pageIndex, _ := utils.StringToInt(ctx.Input.QueryDefault("current", "0"))
	pageSize, _ := utils.StringToInt(ctx.Input.QueryDefault("pageSize", "0"))
	cid, _ := strconv.Atoi(strCid)
	err, res := controller.clusterService.GetNameSpacesListForDB(cid, tenantName, pageIndex, pageSize)
	if err != nil {
		return controller.Fail(err.Error())
	}
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
	//userInfo := req.GetUserInfo(ctx)
	// 只能导入到 平台租户中，再进行分配
	config, err := controller.clusterService.ImportK8sConfig(configFile, request.NickName, uint64(1))
	if err == nil {
		return controller.OK(config)
	}

	return controller.Fail(err.Error())
}

func (controller ClusterController) DeleteDelClusterInfo(ctx *context.HttpContext) mvc.ApiResult {
	id := ctx.Input.Query("id")
	clusterId, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		_ = controller.clusterService.DeleteCluster(clusterId)
		return controller.OK(err == nil)
	}
	return controller.Fail(err.Error())

}

func (controller ClusterController) PutNewNamespace(ctx *context.HttpContext) mvc.ApiResult {
	//userInfo := req.GetUserInfo(ctx)
	cid := ctx.Input.QueryDefault("cid", "0")
	clusterId, _ := strconv.ParseUint(cid, 10, 64)
	namespace := ctx.Input.QueryDefault("namespace", "default")
	tenantId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("tentantId", "0"))
	//只能导入到 平台租户中，再进行分配
	created, err := controller.clusterService.CreateNamespace(tenantId, clusterId, namespace)
	if created {
		clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(0, clusterId)
		err = kubernetes.CreateNamespace(clientSet, namespace)
		if err != nil {
			return controller.Fail(err.Error())
		}
		return controller.OK(err == nil)
	}
	return controller.Fail(err.Error())
}

func (controller ClusterController) PutNewK8sNamespace(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	namespace := ctx.Input.QueryDefault("namespace", "")
	clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(0, clusterId)
	err := kubernetes.CreateNamespace(clientSet, namespace)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(err == nil)
}

func (controller ClusterController) GetResourceQuota(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	namespace := ctx.Input.QueryDefault("namespace", "")
	clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(0, clusterId)
	req, _ := kubernetes.GetResourceQuotasByNamespace(clientSet, namespace)
	return controller.OK(req)
}

func (controller ClusterController) PostResourceQuota(ctx *context.HttpContext) mvc.ApiResult {
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(0, clusterId)
	var quotas dto.QuotasSpec
	err := ctx.Bind(&quotas)
	if err != nil {
		return controller.Fail(err.Error())
	}

	err = kubernetes.CreateResourceQuotasByNamespace(clientSet, quotas)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(true)
}
