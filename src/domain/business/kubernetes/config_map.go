package kubernetes

import (
	"errors"
	"gorm.io/gorm"
	configCorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"kubelilin/domain/database/models"
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

func (cs *ConfigMapSupervisor) ApplyConfigMap(client corev1.CoreV1Interface, dp *models.SgrTenantDeployments) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := cs.db.Model(models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return nil, errors.New("未找到命名空间信息")
	}
	//k8sConfigMaps := client.ConfigMaps(namespace.Namespace)
	configuration := configCorev1.ConfigMapApplyConfiguration{}
	configuration.WithName("")
	return nil, nil
	//k8sConfigMaps.Apply(context.TODO())
}
