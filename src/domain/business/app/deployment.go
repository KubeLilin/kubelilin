package app

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
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

func (deployment *DeploymentService) NewOrUpdateDeployment(deployModel *req.DeploymentStepRequest) (error, *req.DeploymentStepRequest) {
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

func (deployment *DeploymentService) CreateDeploymentStep1(deployModel *req.DeploymentStepRequest) (error, *models.SgrTenantDeployments) {
	dpModel := &models.SgrTenantDeployments{}
	err := copier.Copy(dpModel, deployModel)
	if err != nil {
		return err, nil
	}
	dbRes := deployment.db.Model(deployModel).Create(dpModel)
	return dbRes.Error, dpModel
}

func (deployment *DeploymentService) CreateDeploymentStep2(deployModel *req.DeploymentStepRequest) (error, *models.SgrTenantDeployments) {
	dpModel := models.SgrTenantDeployments{}
	dpcModel := models.SgrTenantDeploymentsContainers{}
	deployment.db.Model(deployModel).Where("app_id = ? and id = ?", deployModel.AppID, deployModel.ID).First(&dpModel)
	if dpModel.AppID == nil {
		return errors.New("未找到相应的部署数据"), nil
	}
	deployment.db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&models.SgrTenantDeployments{}).Updates(dpModel)
		tx.Model(&models.SgrTenantDeploymentsContainers{}).Create(dpcModel)
		return nil
	})
	return nil, nil
}

func (deployment *DeploymentService) GetDeployments(appId uint64, tenantId uint64, deployName string) ([]dto.DeploymentItemDto, error) {
	var deploymentList []dto.DeploymentItemDto
	dataSql := strings.Builder{}
	dataSql.WriteString(`SELECT d.id, d.nickname ,d.name, c.name  as 'clusterName' ,
  d.cluster_id as 'clusterId' , n.namespace ,d.last_image as 'lastImage', 0 'running' , 
  d.replicas 'expected', '0.0.0.0' as 'serviceIP', d.service_name as 'serviceName'
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
