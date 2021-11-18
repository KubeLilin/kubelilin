package kubernetes

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"mime/multipart"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	"strings"
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

func (cluster *ClusterService) ImportK8sConfig(configFile multipart.File, clusterName string, tenantId uint64) (res *models.SgrTenantCluster, err error) {
	//读取配置文件
	defer configFile.Close()
	content, _ := ioutil.ReadAll(configFile)
	var configStr = string(content)
	configV := viper.New()
	configV.SetConfigType("yaml")
	configV.ReadConfig(strings.NewReader(configStr))
	clusters := configV.Get("clusters").([]interface{})
	if len(clusters) == 0 {
		return nil, errors.New("config file must contain more than one cluster configuration points")
	}
	clusterConfig := clusters[0].(map[interface{}]interface{})
	if len(configStr) == 0 {
		return nil, errors.New("config can not be empty")
	}
	//调用k8s获取集群描述信息
	client, err := NewClientSetWithFileContent(configStr)
	if err != nil {
		return nil, err
	}
	versionInfo, err := client.Discovery().ServerVersion()
	if err != nil {
		return nil, err
	}
	clusterData := &models.SgrTenantCluster{
		TenantID: tenantId,
		Name:     clusterConfig["name"].(string),
		Nickname: clusterName,
		Version:  versionInfo.GitVersion,
		Config:   configStr,
		Status:   1,
	}
	err = cluster.db.Model(&models.SgrTenantCluster{}).Create(clusterData).Error
	return clusterData, err
}

func (cs *ClusterService) DeleteCluster(clusterId int64) (err error) {
	res := cs.db.Model(&models.SgrTenantCluster{}).Delete(&models.SgrTenantCluster{}, clusterId)
	return res.Error
}
