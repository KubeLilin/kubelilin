package controllers

import (
	contextV1 "context"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
	requests2 "kubelilin/api/dto/requests"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/domain/dto"
	"kubelilin/utils"
	"strconv"
	"strings"
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
	workload := ctx.Input.QueryDefault("workload", "deployment")
	userInfo := requests2.GetUserInfo(ctx)
	cid, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	client, clientErr := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)
	if clientErr != nil {
		return mvc.FailWithMsg(nil, "Can't create cluster client")
	}
	podList := kubernetes.GetPodList(client, workload, namespace, k8snode, k8sapp)

	config, err1 := controller.clusterService.GetClusterConfig(0, cid)
	if err1 == nil {
		metricsClient, _ := metricsv.NewForConfig(config)
		emptyOptions := metav1.ListOptions{}
		if k8sapp != "" {
			emptyOptions.LabelSelector = "k8s-app=" + k8sapp
		}
		//if k8snode != "" {
		//	emptyOptions.FieldSelector = "spec.nodeName=" + k8snode
		//}
		if namespace == "" {
			emptyOptions.Limit = 500
		}
		podsMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(contextV1.TODO(), emptyOptions)
		if err == nil {
			for _, podsMetricsItem := range podsMetricsList.Items {
				for podindex, podItem := range podList {
					if podsMetricsItem.Name == podItem.PodName {
						podList[podindex].Usage = dto.NodeStatus{}
						for _, cmst := range podsMetricsItem.Containers {
							podList[podindex].Usage.CPU += cmst.Usage.Cpu().AsApproximateFloat64()
							podList[podindex].Usage.Memory += cmst.Usage.Memory().AsApproximateFloat64()
						}
					}
				}
			}
		}
	}

	return controller.OK(podList)
}

func (controller ClusterController) GetNamespaces(ctx *context.HttpContext) mvc.ApiResult {
	//tenantId := ctx.Input.QueryDefault("tid","")
	// get k8s cluster client by tenant id
	userInfo := requests2.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	namespaces := kubernetes.GetAllNamespaces(client)
	return controller.OK(namespaces)
}

func (controller ClusterController) GetNamespacesFromDB(ctx *context.HttpContext) mvc.ApiResult {
	//tenantId := ctx.Input.QueryDefault("tid","")
	// get k8s cluster client by tenant id
	userInfo := requests2.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.Atoi(strCid)
	res := controller.clusterService.GetNameSpacesFromDB(userInfo.TenantID, cid)
	return controller.OK(res)
}

func (controller ClusterController) GetNamespacesByTenantId(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	cid := utils.GetNumberOfParam[uint64](ctx, "cid")
	pageIndex := utils.GetNumberOfParam[int](ctx, "current")
	pageSize := utils.GetNumberOfParam[int](ctx, "pageSize")
	_, res := controller.clusterService.GetNameSpacesListForTenantId(cid, userInfo.TenantID, pageIndex, pageSize)
	return controller.OK(res)
}

func (controller ClusterController) GetNamespaceList(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	cid, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	tenantName := ctx.Input.QueryDefault("tenant", "")
	pageIndex, _ := utils.StringToInt(ctx.Input.QueryDefault("current", "0"))
	pageSize, _ := utils.StringToInt(ctx.Input.QueryDefault("pageSize", "0"))
	err, res := controller.clusterService.GetNameSpacesListForDB(cid, tenantName, pageIndex, pageSize)
	if err != nil {
		return controller.Fail(err.Error())
	}
	if cid > 0 {
		client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)
		resourceQuotasMaps := kubernetes.GetAllNamespaceResourceQuotas(client)
		nsList := res.Data.(*[]dto.NamespaceInfo)
		for i, nsItem := range *nsList {
			(*nsList)[i].QuotasSpec = resourceQuotasMaps["ns-"+nsItem.Namespace]
		}
	}
	return controller.OK(res)
}

func (controller ClusterController) GetDeployments(ctx *context.HttpContext) mvc.ApiResult {
	namespace := ctx.Input.QueryDefault("namespace", "")
	userInfo := requests2.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	list := kubernetes.GetDeploymentList(client, namespace)
	return controller.OK(list)
}

func (controller ClusterController) GetWorkloads(ctx *context.HttpContext) mvc.ApiResult {
	namespace := ctx.Input.QueryDefault("namespace", "")
	workload := ctx.Input.QueryDefault("workload", "")

	userInfo := requests2.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)
	var wordloads []dto.Workload
	switch strings.ToLower(workload) {
	case "deployment":
		wordloads = kubernetes.GetDeploymentList(client, namespace)
	case "statefulset":
		wordloads = kubernetes.GetStatefulSetList(client, namespace)
	case "daemonset":
		wordloads = kubernetes.GetDaemonSetList(client, namespace)
	}
	return controller.OK(wordloads)
}

func (controller ClusterController) GetNodes(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	nodeList := kubernetes.GetNodeList(client)
	return controller.OK(nodeList)
}

func (controller ClusterController) GetList(ctx *context.HttpContext) mvc.ApiResult {
	//userInfo := requests2.GetUserInfo(ctx)
	clusterName := ctx.Input.Query("name")
	tenantClusterList, _ := controller.clusterService.GetClustersByTenant(clusterName)
	return controller.OK(tenantClusterList)
}

func (controller ClusterController) PostClusterByConfig(ctx *context.HttpContext, request *requests2.ImportClusterReq) mvc.ApiResult {
	_, k8sFile, err := ctx.Input.FormFile("file1")
	if err != nil {
		return controller.Fail(err.Error())
	}
	configFile, _ := k8sFile.Open()
	//userInfo := requests.GetUserInfo(ctx)
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
	//userInfo := requests.GetUserInfo(ctx)
	cid := ctx.Input.QueryDefault("cid", "0")
	clusterId, _ := strconv.ParseUint(cid, 10, 64)
	namespace := ctx.Input.QueryDefault("namespace", "default")
	tenantId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("tentantId", "0"))
	//只能导入到 平台租户中，再进行分配
	created, err := controller.clusterService.CreateNamespace(tenantId, clusterId, namespace)
	if created {
		labels := map[string]string{
			"kubelilin-default": "true",
			"tenantId":          strconv.FormatUint(tenantId, 10),
			"clusterId":         strconv.FormatUint(clusterId, 10),
			"namespace":         namespace}

		clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(0, clusterId)
		err = kubernetes.CreateNamespace(clientSet, namespace, labels)
		if err != nil {
			return controller.Fail(err.Error())
		}
		return controller.OK(err == nil)
	}
	return controller.Fail(err.Error())
}

func (controller ClusterController) PutNewK8sNamespace(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	clusterId, _ := utils.StringToUInt64(ctx.Input.QueryDefault("cid", "0"))
	namespace := ctx.Input.QueryDefault("namespace", "")

	labels := map[string]string{
		"kubelilin-default": "true",
		"tenantId":          strconv.FormatUint(userInfo.TenantID, 10),
		"clusterId":         strconv.FormatUint(clusterId, 10),
		"namespace":         namespace}

	clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(0, clusterId)
	err := kubernetes.CreateNamespace(clientSet, namespace, labels)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(err == nil)
}

func (controller ClusterController) PutUpdateRuntime(ctx *context.HttpContext) mvc.ApiResult {
	namespaceId := utils.GetNumberOfParam[uint64](ctx, "namespaceId")
	enableRuntime, _ := utils.StringToBool(ctx.Input.QueryDefault("enableRuntime", "false"))
	runtimeName := ctx.Input.QueryDefault("runtimeName", "")
	err := controller.clusterService.UpdateRuntimeForNamespace(namespaceId, enableRuntime, runtimeName)
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
	var quotas dto.QuotasSpec
	err := ctx.Bind(&quotas)
	if err != nil {
		return controller.Fail(err.Error())
	}

	clientSet, _ := controller.clusterService.GetClusterClientByTenantAndId(0, quotas.ClusterId)
	err = kubernetes.CreateResourceQuotasByNamespace(clientSet, quotas)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(true)
}

func (controller ClusterController) GetIsInstalledDapr(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := requests2.GetUserInfo(ctx)
	strCid := ctx.Input.QueryDefault("cid", "0")
	cid, _ := strconv.ParseUint(strCid, 10, 64)
	client, _ := controller.clusterService.GetClusterClientByTenantAndId(userInfo.TenantID, cid)

	installed := kubernetes.IsInstallDAPRRuntime(client)
	return controller.OK(installed)
}
