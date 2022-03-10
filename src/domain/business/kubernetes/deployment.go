package kubernetes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/cli-runtime/pkg/printers"
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	appsapplymetav1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubelilin/api/req"
	"kubelilin/api/res"
	"kubelilin/domain/database/models"
	"kubelilin/pkg/page"
	"strconv"
	"strings"
	"time"
)

type K8sApiVersion string

const (
	UNKNOW             K8sApiVersion = "UNKNOW"
	EXTENSION_V1_BETA1               = "extensions/v1beta1"
	APPS_V1_BETA1                    = "apps/v1beta1"
	APPS_V1_BETA2                    = "apps/v1beta2"
	APPS_V1                          = "apps/v1"
)

const (
	CLUSTER_IP   string = "ClusterIP"
	NODE_PORT    string = "NodePort"
	LOAD_BALANCE string = "LoadBalancer"
)

type DeploymentSupervisor struct {
	db             *gorm.DB
	clusterService *ClusterService
	k8sService     *ServiceSupervisor
}

func NewDeploymentSupervisor(db *gorm.DB, clusterService *ClusterService, k8sService *ServiceSupervisor) *DeploymentSupervisor {
	return &DeploymentSupervisor{
		db:             db,
		clusterService: clusterService,
		k8sService:     k8sService,
	}
}

func (ds *DeploymentSupervisor) ExecuteDeployment(execReq *req.ExecDeploymentRequest) (interface{}, error) {
	//region 参数校验
	if execReq.DpId == 0 {
		return nil, errors.New("未接收到部署ID")
	}
	if execReq.IsDiv && execReq.WholeImage == "" {
		return nil, errors.New("请填写正确的镜像地址")
	}
	if !execReq.IsDiv && (execReq.Image == "" || execReq.ImageTag == "") {
		return nil, errors.New("请填写正确的镜像地址")
	}
	if !execReq.IsDiv {
		execReq.WholeImage = execReq.Image + ":" + execReq.ImageTag
	}
	dpDatum := models.SgrTenantDeployments{}
	dpcDatum := models.SgrTenantDeploymentsContainers{}
	dbErr := ds.db.Model(&models.SgrTenantDeployments{}).Where("id=?", execReq.DpId).First(&dpDatum)
	if dbErr.Error != nil {
		return nil, errors.New("未找到相应的部署")
	}
	dbErr = ds.db.Model(&models.SgrTenantDeploymentsContainers{}).Where("deploy_id=?", execReq.DpId).Update("image", execReq.WholeImage)
	if dbErr.Error != nil {
		return nil, errors.New("设置镜像失败，请查看服务日志")
	}
	dbErr = ds.db.Model(&models.SgrTenantDeploymentsContainers{}).Where("deploy_id=?", execReq.DpId).First(&dpcDatum)
	if dbErr.Error != nil {
		return nil, errors.New("部署资源限制条件尚未维护，请添加资源限制条件")
	}
	if dpcDatum.Image == "" {
		return nil, errors.New("请维护部署镜像信息")
	}
	//endregion
	//记录发版记录
	exeRes, err := ds.InitDeploymentByApply(execReq.TenantId, &dpDatum, &dpcDatum)
	record := models.SgrTenantDeploymentRecord{AppID: dpDatum.AppID,
		DeploymentID: execReq.DpId,
		ApplyImage:   execReq.WholeImage,
		OpsType:      execReq.OpsType,
		Operator:     &execReq.Operator,
		CreationTime: time.Now(),
	}
	if execReq.OpsType == "" || execReq.OpsType == "manual" {
		execReq.OpsType = "githook"
	}
	if err != nil {
		record.State = "失败"
		record.Remark = err.Error()
	} else {
		record.State = "成功"
	}
	_ = ds.ReleaseRecord(record)
	return exeRes, err
}

// InitDeploymentByApply 使用apply的方式创建deployment/**
func (ds *DeploymentSupervisor) InitDeploymentByApply(tenantId uint64, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (interface{}, error) {
	clusterInfo := &models.SgrTenantCluster{}
	//and tenant_id=?   ,tenantId
	dbErr := ds.db.Model(&models.SgrTenantCluster{}).Where("id=?", dp.ClusterID).First(clusterInfo)
	if dbErr.Error != nil {
		return nil, errors.New("未找到集群信息")
	}
	clientSet, clientSetErr := ds.clusterService.GetClusterClientByTenantAndId(tenantId, clusterInfo.ID)
	if clientSetErr != nil {
		return nil, clientSetErr
	}
	var res []interface{}
	//创建pod
	dpRes, dpErr := ds.ApplyDeployment(clientSet, dp, dpc)
	if dpErr != nil {
		return nil, dpErr
	}
	res = append(res, dpRes)
	//创建svc
	svcRes, svcErr := ds.k8sService.ApplyService(clientSet.CoreV1(), dp)
	if svcErr != nil {
		return nil, svcErr
	}
	res = append(res, svcRes)
	return res, nil
}

func (ds *DeploymentSupervisor) ApplyDeployment(clientSet *kubernetes.Clientset, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	client := clientSet.AppsV1()
	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息"), nil
	}
	k8sDeployment := client.Deployments(namespace.Namespace)
	configuration := &appsapplyv1.DeploymentApplyConfiguration{}
	deploymentDatum := configuration.WithNamespace(namespace.Namespace)
	//metadata
	var name, kind, apiVersion string
	name = dp.Name
	kind = "Deployment"
	apiVersion = APPS_V1
	deploymentDatum.Name = &name
	deploymentDatum.APIVersion = &apiVersion
	deploymentDatum.Kind = &kind

	metaLabels := map[string]string{
		"kubelilin-default": "true",
		"appId":             strconv.FormatUint(dp.AppID, 10),
		"tenantId":          strconv.FormatUint(dp.TenantID, 10),
		"clusterId":         strconv.FormatUint(dp.ClusterID, 10),
		"namespaceId":       strconv.FormatUint(dp.NamespaceID, 10),
		"namespace":         namespace.Namespace,
		"k8s-app":           dp.Name,
	}

	//metalabel := make(map[string]string)
	//metalabel["k8s-app"] = dp.Name
	deploymentDatum.Labels = metaLabels
	//spec
	spec := appsapplyv1.DeploymentSpecApplyConfiguration{}
	replicas := int32(dp.Replicas)
	spec.Replicas = &replicas
	//strategy
	spec.Strategy = &appsapplyv1.DeploymentStrategyApplyConfiguration{
		RollingUpdate: &appsapplyv1.RollingUpdateDeploymentApplyConfiguration{
			MaxUnavailable: &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},
			MaxSurge:       &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},
		},
	}
	//selector
	//selectorMap := make(map[string]string)
	//selectorMap["k8s-app"] = dp.Name
	spec.Selector = &appsapplymetav1.LabelSelectorApplyConfiguration{
		MatchLabels: metaLabels,
	}
	//region  template
	specTemplate := corev1.PodTemplateSpecApplyConfiguration{}
	specTemplate.WithNamespace(namespace.Namespace)
	specTemplate.Labels = metaLabels
	//PodSpec
	podSpec := corev1.PodSpecApplyConfiguration{}
	containers, containerErr := ds.AssemblingContainerForApply(dp, dpc)
	if containerErr != nil {
		return nil, containerErr
	}
	podSpec.Containers = containers
	specTemplate.Spec = &podSpec
	spec.Template = &specTemplate
	//endregion
	deploymentDatum.Spec = &spec
	res, err := k8sDeployment.Apply(context.TODO(), deploymentDatum, metav1.ApplyOptions{Force: true, FieldManager: "deployment-apply-fields"})
	return res, err
}

// AssemblingContainerForApply 通过apply进行创建部署的容器信息/**
func (ds *DeploymentSupervisor) AssemblingContainerForApply(dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) ([]corev1.ContainerApplyConfiguration, error) {
	var containerArr []corev1.ContainerApplyConfiguration
	imagePullPolicy := v1.PullIfNotPresent
	cn := "app"
	container := corev1.ContainerApplyConfiguration{
		Name:            &cn,
		Image:           &dpc.Image,
		ImagePullPolicy: &imagePullPolicy,
	}
	//resources
	requestMap := v1.ResourceList{}
	limitMap := v1.ResourceList{}
	if dpc.RequestCPU > 0 {
		reqCpu, parseErr := resource.ParseQuantity(strconv.FormatFloat(dpc.RequestCPU, 'f', 2, 64))
		if parseErr != nil {
			return nil, parseErr
		}
		requestMap[v1.ResourceCPU] = reqCpu
	}
	if dpc.RequestMemory > 0 {
		reqMem, parseErr := resource.ParseQuantity(strconv.FormatFloat(dpc.RequestMemory, 'f', 2, 64) + "Mi")
		if parseErr != nil {
			return nil, parseErr
		}
		requestMap[v1.ResourceMemory] = reqMem
	}
	if dpc.LimitCPU > 0 {
		limitCpu, parseErr := resource.ParseQuantity(strconv.FormatFloat(dpc.LimitCPU, 'f', 2, 64))
		if parseErr != nil {
			return nil, parseErr
		}
		limitMap[v1.ResourceCPU] = limitCpu
	}
	if dpc.LimitMemory > 0 {
		limitMem, parseErr := resource.ParseQuantity(strconv.FormatFloat(dpc.LimitMemory, 'f', 2, 64) + "Mi")
		if parseErr != nil {
			return nil, parseErr
		}
		limitMap[v1.ResourceMemory] = limitMem
	}
	container.Resources = &corev1.ResourceRequirementsApplyConfiguration{
		Limits:   &limitMap,
		Requests: &requestMap,
	}
	//ports
	var ports []corev1.ContainerPortApplyConfiguration
	portName := "http"
	containerPort := int32(dp.ServicePort)
	protoctl := v1.ProtocolTCP
	ports = append(ports, corev1.ContainerPortApplyConfiguration{
		Name:          &portName,
		ContainerPort: &containerPort,
		Protocol:      &protoctl,
	})
	container.Ports = ports

	container.Env = injectionContainerEnv(dpc.Environments)

	containerArr = append(containerArr, container)
	return containerArr, nil
}

func injectionContainerEnv(envJson string) []corev1.EnvVarApplyConfiguration {
	var envs []corev1.EnvVarApplyConfiguration
	envs = append(envs,
		*corev1.EnvVar().WithName("MY_NODE_NAME").WithValueFrom(
			corev1.EnvVarSource().WithFieldRef(corev1.ObjectFieldSelector().WithFieldPath("spec.nodeName")),
		))

	envs = append(envs,
		*corev1.EnvVar().WithName("MY_POD_NAME").WithValueFrom(
			corev1.EnvVarSource().WithFieldRef(corev1.ObjectFieldSelector().WithFieldPath("metadata.name")),
		))
	envs = append(envs,
		*corev1.EnvVar().WithName("MY_POD_NAMESPACE").WithValueFrom(
			corev1.EnvVarSource().WithFieldRef(corev1.ObjectFieldSelector().WithFieldPath("metadata.namespace")),
		))
	envs = append(envs,
		*corev1.EnvVar().WithName("MY_POD_IP").WithValueFrom(
			corev1.EnvVarSource().WithFieldRef(corev1.ObjectFieldSelector().WithFieldPath("status.podIP")),
		))

	//region 添加录入对系统环境变量
	if envJson != "" {
		var envArr []req.DeploymentEnv
		envJsonErr := json.Unmarshal([]byte(envJson), &envArr)
		if envJsonErr == nil {
			for _, x := range envArr {
				envs = append(envs,
					*corev1.EnvVar().WithName(x.Key).WithValue(x.Value),
				)
			}
		}
	}

	//endregion
	return envs
}

func (ds *DeploymentSupervisor) GetDeploymentYaml(tenantId, dpId uint64) (string, error) {
	dpDatum := models.SgrTenantDeployments{}
	dbErr := ds.db.Model(&models.SgrTenantDeployments{}).Where("id=?", dpId).First(&dpDatum)
	if dbErr.Error != nil {
		return "", errors.New("未找到相应的部署")
	}
	namespace, err := ds.GetNameSpaceByDpId(dpId)
	if err != nil {
		return "", err
	}
	clusterInfo := &models.SgrTenantCluster{}
	dbErr = ds.db.Model(&models.SgrTenantCluster{}).Where("id=? ", dpDatum.ClusterID).First(clusterInfo)
	if dbErr.Error != nil {
		return "", errors.New("未找到集群信息")
	}
	clientSet, clientSetErr := ds.clusterService.GetClusterClientByTenantAndId(tenantId, clusterInfo.ID)
	if clientSetErr != nil {
		return "", clientSetErr
	}
	k8sDeployment, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), dpDatum.Name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	//yamlBytes, yamlErr := yaml.Marshal(k8sDeployment)
	yamlPrinter := printers.YAMLPrinter{}
	buffers := bytes.NewBufferString("")
	k8sDeployment.Kind = "Deployment"
	k8sDeployment.APIVersion = "apps/v1"
	yamlErr := yamlPrinter.PrintObj(k8sDeployment, buffers)

	return buffers.String(), yamlErr
}

func (ds *DeploymentSupervisor) GetNameSpaceByDpId(dpId uint64) (string, error) {

	var namespace string
	err := ds.db.Model(&models.SgrTenantNamespace{}).Raw(`select  t1.namespace from sgr_tenant_namespace as t1 
inner join  sgr_tenant_deployments std on t1.id=std.namespace_id and std.id=? `, dpId).Scan(&namespace)
	if err.Error != nil {
		return "", err.Error
	}
	if namespace == "" {
		return namespace, errors.New("没有找到部署的命名空间")
	}

	return namespace, nil
}

func (ds *DeploymentSupervisor) ReleaseRecord(record models.SgrTenantDeploymentRecord) error {
	recordRes := ds.db.Model(models.SgrTenantDeploymentRecord{}).Create(&record)
	if recordRes.Error != nil {
		return recordRes.Error
	}
	return nil
}

func (ds *DeploymentSupervisor) QueryReleaseRecord(appId, dpId uint64, req *page.PageRequest) (error, *page.Page) {
	var res []res.DeploymentReleaseRecordRes
	condition := ds.db.Model(models.SgrTenantDeploymentRecord{})
	var params []interface{}
	params = append(params, appId)
	sql := strings.Builder{}
	sql.WriteString("select stdr.app_id,stdr.deployment_id,std.name as deployment_name,stdr.apply_image,stdr.ops_type,stu.user_name as operator_name,stdr.creation_time ")
	sql.WriteString("from sgr_tenant_deployment_record as stdr ")
	sql.WriteString("inner join sgr_tenant_deployments std on stdr.deployment_id = std.id ")
	sql.WriteString("left join sgr_tenant_user as stu on stdr.operator=stu.id ")
	sql.WriteString("where stdr.app_id=? ")
	if dpId != 0 {
		sql.WriteString(" and stdr.deployment_id=?  ")
		params = append(params, dpId)
	}

	sql.WriteString("order by stdr.creation_time desc ")

	err, page := page.StartPage(condition, req.PageIndex, req.PageSize).DoScan(&res, sql.String(), params...)
	return err, page
}

//region 暂时弃用的代码，最开始的时候考虑为每个不通版本的k8s指定不通的api-version,最后发现可以统一用apps/v1

//func (ds *DeploymentSupervisor) InitExtensionV1Beta1deployment(client extensionsclient.ExtensionsV1beta1Interface, dp models.SgrTenantDeployments, dpc models.SgrTenantDeploymentsContainers) (interface{}, error) {
//	namespace := &models.SgrTenantNamespace{}
//	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
//	if dbErr.Error != nil {
//		return errors.New("未找到命名空间信息"), nil
//	}
//	k8sDeployment := client.Deployments(namespace.Namespace)
//	deploymentDatum := exv1beta1.Deployment{}
//	//metadata
//	deploymentDatum.Name = dp.Name
//	deploymentDatum.Kind = "Deployment"
//	deploymentDatum.APIVersion = EXTENSION_V1_BETA1
//	//spec
//	spec := exv1beta1.DeploymentSpec{}
//	spec.Replicas = &dp.Replicas
//	//strategy
//	spec.Strategy = exv1beta1.DeploymentStrategy{
//		RollingUpdate: &exv1beta1.RollingUpdateDeployment{
//			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//		},
//	}
//	//selector
//	selectorMap := make(map[string]string)
//	selectorMap["app"] = dp.Name
//	spec.Selector = &metav1.LabelSelector{
//		MatchLabels: selectorMap,
//	}
//	//region  template
//	specTemplate := &v1.PodTemplateSpec{}
//	specTemplate.Labels = selectorMap
//	//PodSpec
//	podSpec := &v1.PodSpec{}
//	containers, containerErr := ds.AssemblingContainer(dp, dpc)
//	if containerErr != nil {
//		return nil, containerErr
//	}
//	podSpec.Containers = containers
//	deploymentDatum.Spec = spec
//	//endregion
//	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
//	return res, err
//}
//
//func (ds *DeploymentSupervisor) InitAppsV1Beta1Deployment(client appsv1beta1client.AppsV1beta1Interface, dp models.SgrTenantDeployments, dpc models.SgrTenantDeploymentsContainers) (interface{}, error) {
//	namespace := &models.SgrTenantNamespace{}
//	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
//	if dbErr.Error != nil {
//		return errors.New("未找到命名空间信息"), nil
//	}
//	k8sDeployment := client.Deployments(namespace.Namespace)
//	deploymentDatum := appsv1beta1.Deployment{}
//	//metadata
//	deploymentDatum.Name = dp.Name
//	deploymentDatum.APIVersion = APPS_V1_BETA1
//	deploymentDatum.Kind = "Deployment"
//	//spec
//	spec := appsv1beta1.DeploymentSpec{}
//	spec.Replicas = &dp.Replicas
//	//strategy
//	spec.Strategy = appsv1beta1.DeploymentStrategy{
//		RollingUpdate: &appsv1beta1.RollingUpdateDeployment{
//			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//		},
//	}
//	//selector
//	selectorMap := make(map[string]string)
//	selectorMap["app"] = dp.Name
//	spec.Selector = &metav1.LabelSelector{
//		MatchLabels: selectorMap,
//	}
//	//region  template
//	specTemplate := &v1.PodTemplateSpec{}
//	specTemplate.Labels = selectorMap
//	//PodSpec
//	podSpec := &v1.PodSpec{}
//	containers, containerErr := ds.AssemblingContainer(dp, dpc)
//	if containerErr != nil {
//		return nil, containerErr
//	}
//	podSpec.Containers = containers
//	deploymentDatum.Spec = spec
//	//endregion
//	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
//	return res, err
//}
//
//func (ds *DeploymentSupervisor) InitAppsV1Beta2Deployment(client appsv1beta2client.AppsV1beta2Interface, dp models.SgrTenantDeployments, dpc models.SgrTenantDeploymentsContainers) (interface{}, error) {
//	namespace := &models.SgrTenantNamespace{}
//	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
//	if dbErr.Error != nil {
//		return errors.New("未找到命名空间信息"), nil
//	}
//	k8sDeployment := client.Deployments(namespace.Namespace)
//	deploymentDatum := appsv1beta2.Deployment{}
//	//metadata
//	deploymentDatum.Name = dp.Name
//	deploymentDatum.APIVersion = APPS_V1_BETA2
//	deploymentDatum.Kind = "Deployment"
//	deploymentDatum.Namespace = namespace.Namespace
//	//spec
//	spec := appsv1beta2.DeploymentSpec{}
//	spec.Replicas = &dp.Replicas
//	//strategy
//	spec.Strategy = appsv1beta2.DeploymentStrategy{
//		RollingUpdate: &appsv1beta2.RollingUpdateDeployment{
//			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//		},
//	}
//	//selector
//	selectorMap := make(map[string]string)
//	selectorMap["app"] = dp.Name
//	spec.Selector = &metav1.LabelSelector{
//		MatchLabels: selectorMap,
//	}
//	//region  template
//	specTemplate := v1.PodTemplateSpec{}
//	specTemplate.Labels = selectorMap
//	//PodSpec
//	podSpec := v1.PodSpec{}
//	containers, containerErr := ds.AssemblingContainer(dp, dpc)
//	if containerErr != nil {
//		return nil, containerErr
//	}
//	podSpec.Containers = containers
//	specTemplate.Spec = podSpec
//	spec.Template = specTemplate
//	//endregion
//	deploymentDatum.Spec = spec
//
//	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
//	return res, err
//}
//
//func (ds *DeploymentSupervisor) InitAppsV1Deployment(client appsv1client.AppsV1Interface, dp models.SgrTenantDeployments, dpc models.SgrTenantDeploymentsContainers) (interface{}, error) {
//	namespace := &models.SgrTenantNamespace{}
//	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
//	if dbErr.Error != nil {
//		return errors.New("未找到命名空间信息"), nil
//	}
//	k8sDeployment := client.Deployments(namespace.Namespace)
//	deploymentDatum := appsv1.Deployment{}
//	//metadata
//	deploymentDatum.Name = dp.Name
//	deploymentDatum.APIVersion = APPS_V1
//	deploymentDatum.Kind = "Deployment"
//	//spec
//	spec := appsv1.DeploymentSpec{}
//	spec.Replicas = &dp.Replicas
//	//strategy
//	spec.Strategy = appsv1.DeploymentStrategy{
//		RollingUpdate: &appsv1.RollingUpdateDeployment{
//			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
//		},
//	}
//	//selector
//	selectorMap := make(map[string]string)
//	selectorMap["app"] = dp.Name
//	spec.Selector = &metav1.LabelSelector{
//		MatchLabels: selectorMap,
//	}
//	//region  template
//	specTemplate := v1.PodTemplateSpec{}
//	specTemplate.Labels = selectorMap
//	//PodSpec
//	podSpec := v1.PodSpec{}
//	containers, containerErr := ds.AssemblingContainer(dp, dpc)
//	if containerErr != nil {
//		return nil, containerErr
//	}
//	podSpec.Containers = containers
//	specTemplate.Spec = podSpec
//	spec.Template = specTemplate
//	//endregion
//	deploymentDatum.Spec = spec
//	jsonStr, _ := json.Marshal(deploymentDatum)
//	fmt.Println(string(jsonStr))
//	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
//	return res, err
//}
//
//// SwitchApiVersion 根据k8sVersion选择对应的API版本，1.6之前EXTENSION_V1_BETA1  1.6-1.7 APPS_V1_BETA1 1.8 APPS_V1_BETA2 1.9以后 APPS_V1/**
//func (ds *DeploymentSupervisor) SwitchApiVersion(clusterVersion string) (error, K8sApiVersion) {
//	firstVersion := string(clusterVersion[3])
//	secondVersion, _ := strconv.ParseInt(string(clusterVersion[4]), 10, 32)
//	fmt.Println(firstVersion)
//	fmt.Println(secondVersion)
//	if firstVersion == "1" {
//		if secondVersion <= 6 {
//			return nil, EXTENSION_V1_BETA1
//		}
//		if secondVersion >= 6 && secondVersion < 8 {
//			return nil, APPS_V1_BETA1
//		}
//		if secondVersion == 8 {
//			return nil, APPS_V1_BETA2
//		}
//		return nil, APPS_V1
//	} else if firstVersion == "2" {
//		return nil, APPS_V1
//	}
//	return errors.New("未找到对应的API版本，请求检查集群版本号是否正确,当前版本号:" + clusterVersion), UNKNOW
//}

//endregion
