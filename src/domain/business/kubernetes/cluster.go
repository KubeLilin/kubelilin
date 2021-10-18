package kubernetes

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"sgr/domain/database/models"
	"sgr/domain/dto"
)

type ClusterService struct {
	db *gorm.DB
}

func NewClusterService(db *gorm.DB) *ClusterService {
	return &ClusterService{db: db}
}

func (cluster *ClusterService) GetClustersByTenant(tenantId int64) ([]dto.ClusterInfo, error) {
	var data []models.SgrTenantCluster
	var clusterList []dto.ClusterInfo
	cluster.db.Model(&models.SgrTenantCluster{}).Where("tenant_id = ?", tenantId).Find(&data)
	for _, item := range data {
		t := dto.ClusterInfo{}
		copier.Copy(&t, item)
		clusterList = append(clusterList, t)
	}
	return clusterList, nil
}

func (cluster *ClusterService) GetClusterClientByTenantAndId(tenantId int64, clusterId int) (*kubernetes.Clientset, error) {
	var data models.SgrTenantCluster
	cluster.db.Model(&models.SgrTenantCluster{}).Where("tenant_id = ? AND id = ?", tenantId, clusterId).First(&data)
	return NewClientSetWithFileContent(data.Config)
}
