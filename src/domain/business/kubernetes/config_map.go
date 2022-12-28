package kubernetes

import (
	"context"
	"errors"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	configCorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/utils"
)

type ConfigMapSupervisor struct {
	db             *gorm.DB
	clusterService *ClusterService
}

func NewConfigMapSupervisor(db *gorm.DB, clusterService *ClusterService) *ConfigMapSupervisor {
	return &ConfigMapSupervisor{
		db:             db,
		clusterService: clusterService,
	}
}

func (cs *ConfigMapSupervisor) QueryList(req *requests.ConfigMapPageReq) ([]requests.ConfigMap, error) {
	var configMapList []requests.ConfigMap
	namespace := models.SgrTenantNamespace{}
	dbErr := cs.db.Model(models.SgrTenantNamespace{}).Where("id=?", req.NamespaceID).First(&namespace)
	if dbErr.Error != nil {
		return configMapList, errors.New("未找到命名空间信息")
	}
	clientSet, clientSetErr := cs.clusterService.GetClusterClientByTenantAndId(0, req.ClusterID)
	if clientSetErr != nil {
		return configMapList, clientSetErr
	}
	listOptions := metav1.ListOptions{}
	labelSelectorMap := map[string]string{
		"kubelilin-default": "true",
	}
	if req.AppId > 0 {
		labelSelectorMap["appId"] = utils.ToString(req.AppId)
	}
	listOptions.LabelSelector = labels.SelectorFromSet(labelSelectorMap).String()
	configMaps, err := clientSet.CoreV1().ConfigMaps(namespace.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		return nil, err
	}

	for _, configmap := range configMaps.Items {
		item := requests.ConfigMap{Name: configmap.Name, AppId: req.AppId, NamespaceId: req.NamespaceID, ClusterId: req.ClusterID}
		for k, v := range configmap.Data {
			item.Items = append(item.Items, requests.ConfigMapDataItem{Key: k, Value: v})
		}
		configMapList = append(configMapList, item)
	}

	return configMapList, nil
}

func (cs *ConfigMapSupervisor) GetListByAppId(appId uint64) ([]dto.AppConfigmapInfo, error) {
	var list []dto.AppConfigmapInfo
	sql := `SELECT cm.id,cm.name,cluster.name as 'cluster',deploy.name as 'deployment', ns.namespace,
cm.cluster_id as 'clusterId',cm.namespace_id as 'namespaceId',cm.deployment_id as 'deploymentId',cm.app_id as 'appId', cm.creation_time as 'creationTime'
From sgr_tenant_config_map cm
INNER JOIN sgr_tenant_deployments deploy on deploy.id = cm.deployment_id
INNER JOIN sgr_tenant_cluster cluster on cluster.id = cm.cluster_id
INNER JOIN sgr_tenant_namespace ns on ns.id =cm.namespace_id
WHERE cm.app_id = ?`
	err := cs.db.Raw(sql, appId).Find(&list).Error
	//err := cs.db.Model(&models.SgrTenantConfigMap{}).Where("app_id = ?", appId).Find(&list).Error
	return list, err
}

func (cs *ConfigMapSupervisor) GetConfigMap(req *requests.ConfigMapPageReq) (requests.ConfigMap, error) {
	var configMapRet requests.ConfigMap
	namespace := models.SgrTenantNamespace{}
	dbErr := cs.db.Model(models.SgrTenantNamespace{}).Where("id=?", req.NamespaceID).First(&namespace)
	if dbErr.Error != nil {
		return configMapRet, errors.New("未找到命名空间信息")
	}
	clientSet, clientSetErr := cs.clusterService.GetClusterClientByTenantAndId(0, req.ClusterID)
	if clientSetErr != nil {
		return configMapRet, clientSetErr
	}
	configMap, err := clientSet.CoreV1().ConfigMaps(namespace.Namespace).Get(context.TODO(), req.Name, metav1.GetOptions{})
	if err != nil {
		return configMapRet, err
	}
	item := requests.ConfigMap{Name: configMap.Name, NamespaceId: req.NamespaceID, ClusterId: req.ClusterID}
	for k, v := range configMap.Data {
		item.Items = append(item.Items, requests.ConfigMapDataItem{Key: k, Value: v})
	}
	return item, nil
}

func (cs *ConfigMapSupervisor) Apply(configmap *requests.ConfigMap) error {
	namespace := &models.SgrTenantNamespace{}
	dbErr := cs.db.Model(models.SgrTenantNamespace{}).Where("id=?", configmap.NamespaceId).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息")
	}
	clientSet, clientSetErr := cs.clusterService.GetClusterClientByTenantAndId(0, configmap.ClusterId)
	if clientSetErr != nil {
		return clientSetErr
	}
	configMapSet := configCorev1.ConfigMap(configmap.Name, namespace.Namespace)
	metaLabels := map[string]string{
		"kubelilin-default": "true",
		"deployId":          utils.ToString(configmap.DeployId),
		"appId":             utils.ToString(configmap.AppId),
		"clusterId":         utils.ToString(configmap.ClusterId),
		"namespaceId":       utils.ToString(configmap.NamespaceId),
		"namespace":         namespace.Namespace,
	}
	configMapSet.Labels = metaLabels
	configMapSet.Data = make(map[string]string)
	for _, item := range configmap.Items {
		configMapSet.Data[item.Key] = item.Value
	}
	_, configApplyErr := clientSet.CoreV1().ConfigMaps(namespace.Namespace).
		Apply(context.TODO(), configMapSet, metav1.ApplyOptions{Force: true, FieldManager: "config-apply-fields"})
	if configApplyErr != nil {
		return configApplyErr
	}
	var configMapEntity models.SgrTenantConfigMap
	cs.db.Where("cluster_id=? AND namespace_id=? AND name=?", configmap.ClusterId, configmap.NamespaceId, configmap.Name).First(&configMapEntity)
	var updateErr error
	if configMapEntity.ID <= 0 { // new
		configMapEntity.Name = configmap.Name
		configMapEntity.NamespaceID = configmap.NamespaceId
		configMapEntity.ClusterID = configmap.ClusterId
		configMapEntity.AppID = configmap.AppId
		configMapEntity.DeploymentID = configmap.DeployId
		configMapEntity.Data = utils.ObjectToString(configmap.Items)
		updateErr = cs.db.Create(&configMapEntity).Error
	} else { // existing
		configMapEntity.Data = utils.ObjectToString(configmap.Items)
		updateErr = cs.db.Updates(configMapEntity).Error
	}
	return updateErr
}

func (cs *ConfigMapSupervisor) Delete(configmap *requests.ConfigMap) error {
	namespace := &models.SgrTenantNamespace{}
	dbErr := cs.db.Model(models.SgrTenantNamespace{}).Where("id=?", configmap.NamespaceId).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息")
	}
	clientSet, clientSetErr := cs.clusterService.GetClusterClientByTenantAndId(0, configmap.ClusterId)
	if clientSetErr != nil {
		return clientSetErr
	}
	configApplyErr := clientSet.CoreV1().ConfigMaps(namespace.Namespace).Delete(context.TODO(), configmap.Name, metav1.DeleteOptions{})
	if configApplyErr != nil {
		return configApplyErr
	}
	return cs.db.Model(&models.SgrTenantConfigMap{}).
		Delete("cluster_id=? AND namespace_id=? AND name=?", configmap.ClusterId, configmap.NamespaceId, configmap.Name).Error
}
