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
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/pkg/page"
	"mime/multipart"
	"strings"
	"sync"
	"time"
)

var k8sClientMemoryCache sync.Map

//var k8sClientMemoryCache = map[string]*kubernetes.Clientset{}
var mutex sync.Mutex

type ClusterService struct {
	db *gorm.DB
}

func NewClusterService(db *gorm.DB) *ClusterService {
	return &ClusterService{db: db}
}

func (cluster *ClusterService) GetClustersByTenant(clusterName string) ([]dto.ClusterInfo, error) {
	var data []models.SgrTenantCluster
	var clusterList []dto.ClusterInfo

	query := cluster.db.Model(&models.SgrTenantCluster{})
	if len(clusterName) > 0 {
		query.Where(" name like ? ", "%"+clusterName+"%")
	}
	query.Find(&data)

	for _, item := range data {
		t := dto.ClusterInfo{}
		copier.Copy(&t, item)
		clusterList = append(clusterList, t)
	}
	return clusterList, nil
}
func (cluster *ClusterService) GetNameSpacesFromDB(tenantId uint64, clusterId int) []models.SgrTenantNamespace {
	var res []models.SgrTenantNamespace
	cluster.db.Model(&models.SgrTenantNamespace{}).Where(" cluster_id=? and tenant_id=?", clusterId, tenantId).Find(&res)
	return res
}

func (cluster *ClusterService) GetNameSpacesListForDB(clusterId uint64, tenantName string, PageIndex int, PageSize int) (error, *page.Page) {
	var res []dto.NamespaceInfo
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString(`select ns.id,ns.tenant_id tenantId,ns.cluster_id clusterId,ns.namespace , clr.name clusterName ,tt.t_code tenantCode, tt.t_name tenantName from sgr_tenant_namespace ns 
INNER JOIN sgr_tenant_cluster clr on clr.id = ns.cluster_id
INNER JOIN sgr_tenant tt on tt.id = ns.tenant_id WHERE 1=1 `)
	var params []interface{}
	if clusterId > 0 {
		sqlBuilder.WriteString(" AND ns.cluster_id=?")
		params = append(params, clusterId)
	}

	if tenantName != "" {
		sqlBuilder.WriteString(" AND tt.t_name like ?")
		params = append(params, "%"+tenantName+"%")
	}

	//err := cluster.db.Raw(sqlBuilder.String(), params...).Scan(&responses).Error

	return page.StartPage(cluster.db, PageIndex, PageSize).DoScan(&res, sqlBuilder.String(), params...)
}

func (cluster *ClusterService) CreateNamespace(tenantID uint64, clusterId uint64, namespace string) (bool, error) {
	var exitsCount int64
	cluster.db.Model(models.SgrTenantNamespace{}).Where("cluster_id=? and namespace=?", clusterId, namespace).Count(&exitsCount)
	if exitsCount > 0 {
		return false, errors.New("already have the same namespace")
	}

	now := time.Now()
	record := &models.SgrTenantNamespace{
		TenantID:   tenantID,
		ClusterID:  clusterId,
		Namespace:  namespace,
		CreateTime: &now,
		UpdateTime: &now,
		Status:     1,
	}
	err := cluster.db.Model(models.SgrTenantNamespace{}).Create(record).Error
	return err == nil, err
}

func (cluster *ClusterService) GetClusterClientByTenantAndId(tenantId uint64, clusterId uint64) (*kubernetes.Clientset, error) {
	//判断缓存是否存在
	//mutex.Lock()
	//defer mutex.Unlock()
	key := "c" + string(clusterId)
	clientValue, ok := k8sClientMemoryCache.Load(key)
	//clientValue, ok := k8sClientMemoryCache[key]
	if ok {
		healthy, healthyErr := cluster.ClientHealthCheck(clientValue.(*kubernetes.Clientset))
		if ok && healthy {
			return clientValue.(*kubernetes.Clientset), nil
		}
		return nil, healthyErr
	} else {
		var data models.SgrTenantCluster
		cluster.db.Model(&models.SgrTenantCluster{}).Where(" id = ?", clusterId).First(&data)
		client, err := NewClientSetWithFileContent(data.Config)
		if err == nil {
			healthy, healthyErr := cluster.ClientHealthCheck(client)
			if healthy {
				k8sClientMemoryCache.Store(key, client)
				//k8sClientMemoryCache[key] = client
			} else {
				return nil, healthyErr
			}
		}
		return client, err
	}

}

func (cluster *ClusterService) GetClusterConfig(tenantId uint64, clusterId uint64) (*rest.Config, error) {
	var data models.SgrTenantCluster
	cluster.db.Model(&models.SgrTenantCluster{}).Where(" id = ?", clusterId).First(&data)
	return clientcmd.RESTConfigFromKubeConfig([]byte(data.Config))
}

func (cluster *ClusterService) ClientHealthCheck(client *kubernetes.Clientset) (bool, error) {
	_, err := client.ServerVersion()
	return err == nil, err
}
func (cluster *ClusterService) ImportK8sConfig(configFile multipart.File, clusterName string, tenantId uint64) (res *models.SgrTenantCluster, err error) {
	//判断集群是否已经存在
	var exitsCount int64
	//tenant_id=? and  tenantId
	cluster.db.Model(models.SgrTenantCluster{}).Where(" name=?", clusterName).Count(&exitsCount)
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
	//clusterConfig := clusters[0].(map[string]interface{})
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
		Name:     clusterName, //clusterConfig["name"].(string),
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
