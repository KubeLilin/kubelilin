package kubernetes

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	"time"
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

func (cluster *ClusterService) ImportK8sConfig(configStr, clusterName string, tenantId uint64) (res *models.SgrTenantCluster, err error) {
	//调用k8s获取集群描述信息
	if len(configStr) == 0 {
		return nil, errors.New("config can not be empty")
	}
	client, err := NewClientSetWithFileContent(configStr)
	if err != nil {
		return nil, err
	}
	versionInfo, err := client.Discovery().ServerVersion()
	if err != nil {
		return nil, err
	}
	fmt.Println(versionInfo)
	t := time.Now()
	clusterData := &models.SgrTenantCluster{
		TenantID:   tenantId,
		Name:       clusterName,
		Nickname:   clusterName,
		Version:    versionInfo.GitVersion,
		Config:     configStr,
		Status:     1,
		CreateTime: &t,
		UpdateTime: &t,
	}
	err = cluster.db.Model(&models.SgrTenantCluster{}).Create(clusterData).Error
	return clusterData, err
}
