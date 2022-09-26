package kubernetes

import (
	"errors"
	"gorm.io/gorm"
	configCorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"kubelilin/api/req"
	"kubelilin/domain/database/models"
	"kubelilin/pkg/page"
	"strconv"
	"strings"
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

func (cs *ConfigMapSupervisor) QueryConfigList(req req.ConfigMapPageReq) []models.SgrTenantConfigMap {
	var configMapList []models.SgrTenantConfigMap
	sb := strings.Builder{}
	sb.WriteString("select * from sgr_tenant_config_map ")
	if len(req.ConfigName) > 0 {
		sb.WriteString(" where name like '%")
		sb.WriteString(req.ConfigName)
		sb.WriteString("%'")
	}
	page.StartPage(cs.db, req.PageIndex, req.PageSize).DoScan(configMapList, sb.String())
	return configMapList
}

func (cs *ConfigMapSupervisor) ApplyConfigMap(client corev1.CoreV1Interface, cp *models.SgrTenantConfigMap) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := cs.db.Model(models.SgrTenantNamespace{}).Where("id=?", cp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return nil, errors.New("未找到命名空间信息")
	}
	//k8sConfigMaps := client.ConfigMaps(namespace.Namespace)
	configuration := configCorev1.ConfigMapApplyConfiguration{}
	cpDatum := configuration.WithNamespace(namespace.Namespace)
	var name, kind, apiVersion string
	name = cp.Name
	kind = "Deployment"
	apiVersion = APPS_V1
	cpDatum.Name = &name
	cpDatum.APIVersion = &apiVersion
	cpDatum.Kind = &kind
	metaLabels := map[string]string{
		"kubelilin-default": "true",
		"tenantId":          strconv.FormatUint(cp.TenantID, 10),
		"clusterId":         strconv.FormatUint(cp.ClusterID, 10),
		"namespaceId":       strconv.FormatUint(cp.NamespaceID, 10),
		"namespace":         namespace.Namespace,
	}
	cpDatum.Labels = metaLabels
	return nil, nil
	//k8sConfigMaps.Apply(context.TODO())
}
