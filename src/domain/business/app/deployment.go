package app

import (
	"gorm.io/gorm"
	"sgr/domain/database/models"
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
