package kubernetes

import (
	"context"
	"errors"
	"gorm.io/gorm"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
	exv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1client "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	extensionsclient "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"sgr/domain/database/models"
	"strconv"
)

type K8sApiVersion uint8

const (
	UNKNOW K8sApiVersion = iota
	EXTENSION_V1_BETA1
	APPS_V1_BETA1
	APPS_V1_BETA2
	APPS_V1
)

type DeploymentSupervisor struct {
	db             gorm.DB
	clusterService ClusterService
}

func (ds *DeploymentSupervisor) ExecuteDeployment(dpId, tenantId uint64) (interface{}, error) {

	//region 参数校验
	dpDatum := &models.SgrTenantDeployments{}
	dpcDatum := &models.SgrTenantDeploymentsContainers{}
	dbErr := ds.db.Model(&models.SgrTenantDeployments{}).Where("id=?", dpId).First(dpDatum)
	if dbErr.Error != nil {
		return nil, errors.New("未找到相应的部署")
	}
	dbErr = ds.db.Model(&models.SgrTenantDeploymentsContainers{}).Where("deploy_id=?", dpId).First(&dpcDatum)
	if dbErr.Error != nil {
		return nil, errors.New("部署资源限制条件尚未维护，请添加资源限制条件")
	}
	if dpcDatum.Image == "" {
		return nil, errors.New("请维护部署镜像信息")
	}
	//endregion
	return ds.InitDeploymentTemplate(tenantId, dpDatum, dpcDatum)
}

func (ds *DeploymentSupervisor) InitDeploymentTemplate(tenantId uint64, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (interface{}, error) {
	clusterInfo := &models.SgrTenantCluster{}
	dbErr := ds.db.Model(&models.SgrTenantCluster{}).Where("id=? and tenant_id=?", dp.ClusterID, tenantId).First(clusterInfo)
	if dbErr.Error != nil {
		return errors.New("未找到集群信息"), nil
	}
	apiVersionErr, apiVersion := ds.SwitchApiVersion(clusterInfo.Version)
	if apiVersionErr != nil {
		return apiVersionErr, nil
	}
	clientSet, clientSetErr := ds.clusterService.GetClusterClientByTenantAndId(tenantId, clusterInfo.ID)
	if clientSetErr != nil {
		return clientSetErr, nil
	}
	switch apiVersion {
	case EXTENSION_V1_BETA1:
		return ds.InitExtensionV1Beta1deployment(clientSet.ExtensionsV1beta1(), dp, dpc)
	case APPS_V1_BETA1:
		return ds.InitAppsV1Beta1Deployment(clientSet.AppsV1beta1(), dp, dpc)
	case APPS_V1_BETA2:
		return ds.InitAppsV1Beta2Deployment(clientSet.AppsV1beta2(), dp, dpc)
	case APPS_V1:
		return ds.InitAppsV1Deployment(clientSet.AppsV1(), dp, dpc)
	}
	return nil, errors.New("未找到当前集群版本的API")
}

func (ds *DeploymentSupervisor) InitExtensionV1Beta1deployment(client extensionsclient.ExtensionsV1beta1Interface, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息"), nil
	}
	k8sDeployment := client.Deployments(namespace.Namespace)
	deploymentDatum := exv1beta1.Deployment{}
	//metadata
	deploymentDatum.Name = dp.Name
	deploymentDatum.APIVersion = ""
	deploymentDatum.Kind = "Deployment"
	//spec
	spec := exv1beta1.DeploymentSpec{}
	spec.Replicas = &dp.Replicas
	//strategy
	spec.Strategy = exv1beta1.DeploymentStrategy{
		RollingUpdate: &exv1beta1.RollingUpdateDeployment{
			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
		},
	}
	//selector
	selectorMap := make(map[string]string)
	selectorMap["app"] = dp.Name
	spec.Selector = &metav1.LabelSelector{
		MatchLabels: selectorMap,
	}
	//region  template
	specTemplate := &v1.PodTemplateSpec{}
	specTemplate.Labels = selectorMap
	//PodSpec
	podSpec := &v1.PodSpec{}
	containers, containerErr := ds.AssemblingContainer(dp, dpc)
	if containerErr != nil {
		return nil, containerErr
	}
	podSpec.Containers = containers
	deploymentDatum.Spec = spec
	//endregion
	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
	return res, err
}

func (ds *DeploymentSupervisor) InitAppsV1Beta1Deployment(client appsv1beta1client.AppsV1beta1Interface, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息"), nil
	}
	k8sDeployment := client.Deployments(namespace.Namespace)
	deploymentDatum := appsv1beta1.Deployment{}
	//metadata
	deploymentDatum.Name = dp.Name
	deploymentDatum.APIVersion = ""
	deploymentDatum.Kind = "Deployment"
	//spec
	spec := appsv1beta1.DeploymentSpec{}
	spec.Replicas = &dp.Replicas
	//strategy
	spec.Strategy = appsv1beta1.DeploymentStrategy{
		RollingUpdate: &appsv1beta1.RollingUpdateDeployment{
			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
		},
	}
	//selector
	selectorMap := make(map[string]string)
	selectorMap["app"] = dp.Name
	spec.Selector = &metav1.LabelSelector{
		MatchLabels: selectorMap,
	}
	//region  template
	specTemplate := &v1.PodTemplateSpec{}
	specTemplate.Labels = selectorMap
	//PodSpec
	podSpec := &v1.PodSpec{}
	containers, containerErr := ds.AssemblingContainer(dp, dpc)
	if containerErr != nil {
		return nil, containerErr
	}
	podSpec.Containers = containers
	deploymentDatum.Spec = spec
	//endregion
	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
	return res, err
}

func (ds *DeploymentSupervisor) InitAppsV1Beta2Deployment(client appsv1beta2client.AppsV1beta2Interface, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息"), nil
	}
	k8sDeployment := client.Deployments(namespace.Namespace)
	deploymentDatum := appsv1beta2.Deployment{}
	//metadata
	deploymentDatum.Name = dp.Name
	deploymentDatum.APIVersion = ""
	deploymentDatum.Kind = "Deployment"
	//spec
	spec := appsv1beta2.DeploymentSpec{}
	spec.Replicas = &dp.Replicas
	//strategy
	spec.Strategy = appsv1beta2.DeploymentStrategy{
		RollingUpdate: &appsv1beta2.RollingUpdateDeployment{
			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
		},
	}
	//selector
	selectorMap := make(map[string]string)
	selectorMap["app"] = dp.Name
	spec.Selector = &metav1.LabelSelector{
		MatchLabels: selectorMap,
	}
	//region  template
	specTemplate := &v1.PodTemplateSpec{}
	specTemplate.Labels = selectorMap
	//PodSpec
	podSpec := &v1.PodSpec{}
	containers, containerErr := ds.AssemblingContainer(dp, dpc)
	if containerErr != nil {
		return nil, containerErr
	}
	podSpec.Containers = containers
	deploymentDatum.Spec = spec
	//endregion
	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
	return res, err
}
func (ds *DeploymentSupervisor) InitAppsV1Deployment(client appsv1client.AppsV1Interface, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息"), nil
	}
	k8sDeployment := client.Deployments(namespace.Namespace)
	deploymentDatum := appsv1.Deployment{}
	//metadata
	deploymentDatum.Name = dp.Name
	deploymentDatum.APIVersion = ""
	deploymentDatum.Kind = "Deployment"
	//spec
	spec := appsv1.DeploymentSpec{}
	spec.Replicas = &dp.Replicas
	//strategy
	spec.Strategy = appsv1.DeploymentStrategy{
		RollingUpdate: &appsv1.RollingUpdateDeployment{
			MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
			MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
		},
	}
	//selector
	selectorMap := make(map[string]string)
	selectorMap["app"] = dp.Name
	spec.Selector = &metav1.LabelSelector{
		MatchLabels: selectorMap,
	}
	//region  template
	specTemplate := &v1.PodTemplateSpec{}
	specTemplate.Labels = selectorMap
	//PodSpec
	podSpec := &v1.PodSpec{}
	containers, containerErr := ds.AssemblingContainer(dp, dpc)
	if containerErr != nil {
		return nil, containerErr
	}
	podSpec.Containers = containers
	deploymentDatum.Spec = spec
	//endregion
	res, err := k8sDeployment.Create(context.TODO(), &deploymentDatum, metav1.CreateOptions{})
	return res, err
}

// SwitchApiVersion 根据k8sVersion选择对应的API版本，1.6之前EXTENSION_V1_BETA1  1.6-1.7 APPS_V1_BETA1 1.8 APPS_V1_BETA2 1.9以后 APPS_V1/**
func (ds *DeploymentSupervisor) SwitchApiVersion(clusterVersion string) (error, K8sApiVersion) {
	firstVersion := string(clusterVersion[4])
	secondVersion, _ := strconv.ParseInt(string(clusterVersion[5]), 10, 32)
	if firstVersion == "1" {
		if secondVersion <= 6 {
			return nil, EXTENSION_V1_BETA1
		}
		if secondVersion >= 6 && secondVersion < 8 {
			return nil, APPS_V1_BETA1
		}
		if secondVersion == 8 {
			return nil, APPS_V1_BETA2
		}
		return nil, APPS_V1
	} else if firstVersion == "2" {
		return nil, APPS_V1
	}
	return errors.New("未找到对应的API版本，请求检查集群版本号是否正确"), UNKNOW
}

func (ds *DeploymentSupervisor) AssemblingContainer(dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) ([]v1.Container, error) {
	var containerArr []v1.Container
	container := v1.Container{
		Name:            dp.Name,
		Image:           dpc.Image,
		Command:         []string{dpc.RunCmd},
		ImagePullPolicy: v1.PullIfNotPresent,
	}
	//resources
	requestMap := map[v1.ResourceName]resource.Quantity{}
	limitMap := map[v1.ResourceName]resource.Quantity{}
	if dpc.RequestCPU > 0 {
		reqCpu, parseErr := resource.ParseQuantity(strconv.FormatFloat(dpc.RequestCPU, 'f', 2, 64))
		if parseErr != nil {
			return nil, parseErr
		}
		requestMap[v1.ResourceCPU] = reqCpu
	}
	if dpc.RequestMemory > 0 {
		reqMem, parseErr := resource.ParseQuantity(strconv.FormatFloat(dpc.RequestMemory, 'f', 2, 64))
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
		limitMem, parseErr := resource.ParseQuantity(strconv.FormatFloat(dpc.LimitMemory, 'f', 2, 64))
		if parseErr != nil {
			return nil, parseErr
		}
		limitMap[v1.ResourceMemory] = limitMem
	}
	container.Resources = v1.ResourceRequirements{
		Limits:   limitMap,
		Requests: requestMap,
	}
	//ports
	/*var ports []v1.ContainerPort
	append(ports,v1.ContainerPort{
		Name: "http",
	})*/
	containerArr = append(containerArr, container)
	return containerArr, nil
}
