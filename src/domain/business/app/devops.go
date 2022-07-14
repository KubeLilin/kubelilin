package app

import (
	"errors"
	"gorm.io/gorm"
	"kubelilin/api/req"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/pkg/page"
	"kubelilin/utils"
)

type DevopsService struct {
	db *gorm.DB
}

func NewDevopsService(db *gorm.DB) *DevopsService {
	return &DevopsService{db: db}
}

func (service *DevopsService) CreateProject(requestProject *req.CreateNewProject) error {
	var exitCount int64
	service.db.Model(&models.DevopsProjects{}).Where("tenant_id=? and name=?", requestProject.TenantID, requestProject.Name).Count(&exitCount)
	if exitCount > 0 {
		return errors.New("already have the same name project")
	}

	devProject := models.DevopsProjects{
		Name:         requestProject.Name,
		TenantID:     requestProject.TenantID,
		CreationTime: utils.TimeNowPtr(),
	}
	dpcRes := service.db.Model(&models.DevopsProjects{}).Create(&devProject)
	if dpcRes.Error != nil {
		return dpcRes.Error
	}
	for _, appId := range requestProject.AppIdList {
		devProjectApps := models.DevopsProjectsApps{
			ProjectID:     devProject.ID,
			ApplicationID: appId,
		}
		service.db.Model(&models.DevopsProjectsApps{}).Create(&devProjectApps)
	}

	return nil
}

func (service *DevopsService) GetProjectList(request *req.DevopsProjectReq) (error, *page.Page) {
	sql := `SELECT p.id,p.name 'project_name', p.creation_time, 
(SELECT count(1) FROM  devops_projects_apps pas WHERE p.id = pas.project_id) 'app_count',
(SELECT GROUP_CONCAT(pas.application_id) FROM  devops_projects_apps pas WHERE p.id = pas.project_id) 'app_ids'
FROM devops_projects p`
	var sqlParams []interface{}
	if request.Name != "" {
		sql += " WHERE p.name like ?"
		sqlParams = append(sqlParams, "%"+request.Name+"%")
	}
	res := &[]dto.DevOpsProjectsDTO{}
	return page.StartPage(service.db, request.CurrentPage, request.PageSize).DoScan(res, sql, sqlParams...)
}
