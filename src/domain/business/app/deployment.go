package app

import (
	"gorm.io/gorm"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	"strings"
)

const (
	Deployment  = "Deployment"
	DaemonSet   = "DaemonSet"
	StatefulSet = "StatefulSet"
	CronJob     = "CronJob"
)

type DeploymentService struct {
	db *gorm.DB
}

func NewDeploymentService(db *gorm.DB) *DeploymentService {
	return &DeploymentService{db: db}
}

func (deployment *DeploymentService) NewOrUpdateDeployment(deployModel *models.SgrTenantDeployments) (error, *models.SgrTenantDeployments) {
	var deploy *models.SgrTenantDeployments

	deployment.db.Model(deployModel).Where("app_id = ? and name = ?", deployModel.AppID, deployModel.Name).First(&deploy)
	if deploy != nil { // new
		dbRes := deployment.db.Model(deployModel).Create(deployModel)
		return dbRes.Error, deployModel
	} else { // update
		dbRes := deployment.db.Model(deployModel).Updates(deployModel)
		return dbRes.Error, deployModel
	}
}

func (deployment *DeploymentService) GetDeployments(appId uint64, tenantId uint64, deployName string) ([]dto.DeploymentItemDto, error) {
	var deploymentList []dto.DeploymentItemDto
	dataSql := strings.Builder{}
	dataSql.WriteString(`SELECT d.id, d.nickname , c.name  as 'clusterName' , n.namespace ,d.last_image as 'lastImage', 0 'running' , d.replicas 'expected', '0.0.0.0' as 'serviceIP', '' as 'serviceName'
  FROM sgr_tenant_deployments d
  INNER JOIN sgr_tenant_cluster c on c.id = d.cluster_id
  INNER JOIN sgr_tenant_namespace n on n.id = d.namespace_id
  WHERE  d.app_id = ? AND d.tenant_id =? `)
	if deployName != "" {
		dataSql.WriteString("AND d.nickname like '%" + deployName + "%'")
	}
	dataRes := deployment.db.Raw(dataSql.String(), appId, tenantId).Scan(&deploymentList)
	return deploymentList, dataRes.Error
}
