package app

import (
	"errors"
	"gorm.io/gorm"
	"kubelilin/api/req"
	"kubelilin/domain/database/models"
)

type DevopsService struct {
	db *gorm.DB
}

func NewDevopsService(db *gorm.DB) *DevopsService {
	return &DevopsService{db: db}
}

func (service *DevopsService) CreateProject(project *req.CreateNewProject) error {
	var exitCount int64
	service.db.Model(&models.SgrTenantApplication{}).Where("tenant_id=? and name=?", project.TenantID, project.Name).Count(&exitCount)
	if exitCount > 0 {
		return errors.New("already have the same name project")
	}
	return nil
}
