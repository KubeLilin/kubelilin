package kubernetes

import (
	"errors"
	"gorm.io/gorm"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
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

func (ds *DeploymentSupervisor) ExecuteDeployment(dpId, tenantId uint64) error {

	//region 参数校验
	dpDatum := &models.SgrTenantDeployments{}
	dpcDatum := &models.SgrTenantDeploymentsContainers{}
	dbErr := ds.db.Model(&models.SgrTenantDeployments{}).Where("id=?", dpId).First(dpDatum)
	if dbErr.Error != nil {
		return errors.New("未找到相应的部署")
	}
	dbErr = ds.db.Model(&models.SgrTenantDeploymentsContainers{}).Where("deploy_id=?", dpId).First(&dpcDatum)
	if dbErr.Error != nil {
		return errors.New("部署资源限制条件尚未维护，请添加资源限制条件")
	}
	if dpcDatum.Image == "" {
		return errors.New("请维护部署镜像信息")
	}
	//endregion
	return nil
}

func (ds *DeploymentSupervisor) InitDeploymentTemplate(tenantId uint64, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (error, interface{}) {
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
		ds.InitExtensionV1Beta1deployment(clientSet.ExtensionsV1beta1(), nil, nil)
	}

	return nil, nil
}

func (ds *DeploymentSupervisor) InitExtensionV1Beta1deployment(client extensionsv1beta1.ExtensionsV1beta1Interface, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) (error, interface{}) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := ds.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息"), nil
	}
	_ = client.Deployments(namespace.Namespace)

	//k8sDeployment.Create(context.TODO())
	return nil, nil
}

func (ds *DeploymentSupervisor) InitAppsV1Beta1Deployment(client extensionsv1beta1.ExtensionsV1beta1Interface) {

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
