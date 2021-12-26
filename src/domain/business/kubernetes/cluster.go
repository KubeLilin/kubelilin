package kubernetes

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"mime/multipart"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	"strings"
	"sync"
)

var k8sClientMemoryCache = map[string]*kubernetes.Clientset{}
var mutex sync.RWMutex

type ClusterService struct {
	db *gorm.DB
}

func NewClusterService(db *gorm.DB) *ClusterService {
	return &ClusterService{db: db}
}

func (cluster *ClusterService) GetClustersByTenant(tenantId uint64, clusterName string) ([]dto.ClusterInfo, error) {
	var data []models.SgrTenantCluster
	var clusterList []dto.ClusterInfo
	sb := strings.Builder{}
	sb.WriteString("tenant_id = ?")
	if clusterName != "" {
		sb.WriteString(" and name=? ")
	}
	cluster.db.Model(&models.SgrTenantCluster{}).Where(sb.String(), tenantId, clusterName).Find(&data)
	for _, item := range data {
		t := dto.ClusterInfo{}
		copier.Copy(&t, item)
		clusterList = append(clusterList, t)
	}
	return clusterList, nil
}

func (cluster *ClusterService) GetNameSpacesFromDB(tenantId uint64, clusterId int) []models.SgrTenantNamespace {
	var res []models.SgrTenantNamespace
	cluster.db.Model(&models.SgrTenantNamespace{}).Where("tenant_id=? and cluster_id=?", tenantId, clusterId).Find(&res)
	return res
}

func (cluster *ClusterService) CreateNamespace(clusterId uint64, namespace string) (bool, error) {
	var exitsCount int64
	cluster.db.Model(models.SgrTenantNamespace{}).Where("cluster_id=? and namespace=?", clusterId, namespace).Count(&exitsCount)
	if exitsCount > 0 {
		return false, errors.New("already have the same namespace")
	}

	return true, nil
}

func (cluster *ClusterService) GetClusterClientByTenantAndId(tenantId uint64, clusterId uint64) (*kubernetes.Clientset, error) {
	//判断缓存是否存在
	key := "t" + string(tenantId) + "c" + string(clusterId)
	clientValue, ok := k8sClientMemoryCache[key]
	if ok {
		healthy, healthyErr := cluster.ClientHealthCheck(clientValue)
		if ok && healthy {
			return clientValue, nil
		}
		return nil, healthyErr
	} else {
		mutex.Lock()
		var data models.SgrTenantCluster
		cluster.db.Model(&models.SgrTenantCluster{}).Where("tenant_id = ? AND id = ?", tenantId, clusterId).First(&data)
		client, err := NewClientSetWithFileContent(data.Config)
		if err == nil {
			healthy, healthyErr := cluster.ClientHealthCheck(client)
			if healthy {
				k8sClientMemoryCache[key] = client
			} else {
				return nil, healthyErr
			}
		}
		mutex.Unlock()
		return client, err
	}
}

func (cluster *ClusterService) GetClusterConfig(tenantId uint64, clusterId uint64) (*rest.Config, error) {
	var data models.SgrTenantCluster
	cluster.db.Model(&models.SgrTenantCluster{}).Where("tenant_id = ? AND id = ?", tenantId, clusterId).First(&data)
	return clientcmd.RESTConfigFromKubeConfig([]byte(data.Config))
}

func (cluster *ClusterService) ClientHealthCheck(client *kubernetes.Clientset) (bool, error) {
	_, err := client.ServerVersion()
	return err == nil, err
}
func (cluster *ClusterService) ImportK8sConfig(configFile multipart.File, clusterName string, tenantId uint64) (res *models.SgrTenantCluster, err error) {
	//判断集群是否已经存在
	var exitsCount int64
	cluster.db.Model(models.SgrTenantCluster{}).Where("tenant_id=? and name=?", tenantId, clusterName).Count(&exitsCount)
	if exitsCount > 0 {
		return nil, errors.New("already have the same cluster")
	}
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

func (cluster *ClusterService) DeleteCluster(clusterId int64) (err error) {
	res := cluster.db.Model(&models.SgrTenantCluster{}).Delete(&models.SgrTenantCluster{}, clusterId)
	return res.Error
}
