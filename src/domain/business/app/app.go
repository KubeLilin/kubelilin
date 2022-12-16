package app

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yoyofx/yoyogo/abstractions"
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/pkg/page"
	"strings"
)

type ApplicationService struct {
	db         *gorm.DB
	vcsService VcsService
	config     abstractions.IConfiguration
}

func NewApplicationService(db *gorm.DB, config abstractions.IConfiguration) *ApplicationService {
	return &ApplicationService{db: db, config: config}
}

func (s *ApplicationService) CreateApp(req *requests.AppReq) (error, *models.SgrTenantApplication) {
	var exitCount int64
	s.db.Model(&models.SgrTenantApplication{}).Where("tenant_id=? and name=?", req.TenantID, req.Name).Count(&exitCount)
	if exitCount > 0 {
		return errors.New("already have the same name application"), nil
	}
	appModel := &models.SgrTenantApplication{}
	err := copier.Copy(appModel, req)
	appModel.ScID = &req.SCID
	appModel.GitType = req.SourceType
	if err != nil {
		return err, nil
	}
	dbErr := s.db.Transaction(func(tx *gorm.DB) error {
		dbRes := tx.Model(models.SgrTenantApplication{}).Create(appModel)
		if dbRes.Error != nil {
			return nil
		}
		if req.ProjectID > 0 {
			return tx.Create(&models.DevopsProjectsApps{ProjectID: req.ProjectID, ApplicationID: appModel.ID}).Error
		}
		return nil
	})
	//创建git仓库
	if dbErr != nil {
		return dbErr, nil
	}
	return nil, appModel
}

func (s *ApplicationService) DeleteApp(appId uint64) error {
	dbErr := s.db.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&models.SgrTenantDeployments{}).Where("app_id=?", appId).Count(&count)
		if count > 0 {
			return errors.New("application of deployments are not empty")
		}
		dbRes := tx.Delete(&models.SgrTenantApplication{}, "id=?", appId)
		if dbRes.Error != nil {
			return nil
		}
		return nil
	})
	return dbErr
}

func (s *ApplicationService) UpdateApp(req *requests.AppReq) (error, int64) {
	appModel := models.SgrTenantApplication{}
	appModel.Level = req.Level
	appModel.Remarks = req.Remarks
	appModel.Language = req.Language
	appModel.Status = req.Status
	appModel.GitType = req.SourceType
	appModel.ScID = &req.SCID
	appModel.Git = req.Git
	appModel.Labels = req.Labels
	dbRes := s.db.Model(&models.SgrTenantApplication{}).Where("id=?", req.ID).Updates(appModel)
	if dbRes.Error != nil {
		return nil, 0
	}
	return nil, dbRes.RowsAffected
}

func (s *ApplicationService) QueryAppList(req *requests.AppReq) (error, *page.Page) {
	res := &[]dto.ApplicationInfoDTO{}
	var sqlParams []interface{}
	sb := strings.Builder{}
	sb.WriteString(`SELECT t1.*,t1.git_type as SourceType,t1.sc_id as SCID,t2.name as language_name,t3.name as level_name ,
(SELECT count(1) FROM sgr_tenant_deployments where app_id = t1.id) as depCount
FROM sgr_tenant_application AS t1 
INNER JOIN sgr_code_application_language AS t2
ON t1.language = t2.id INNER JOIN sgr_code_application_level AS t3 ON t1.LEVEL = t3.id WHERE t1.status = 1`)
	if req.Name != "" {
		sb.WriteString(" AND t1.name like ?")
		sqlParams = append(sqlParams, "%"+req.Name+"%")
	}
	if req.Labels != "" {
		sb.WriteString(" AND t1.labels like ?")
		sqlParams = append(sqlParams, "%"+req.Labels+"%")
	}
	if req.Level != 0 {
		sb.WriteString(" AND t1.level=?")
		sqlParams = append(sqlParams, req.Level)
	}
	if req.Language != 0 {
		sb.WriteString(" AND t1.language=?")
		sqlParams = append(sqlParams, req.Language)
	}

	if req.TenantID > 0 {
		sb.WriteString(" AND t1.tenant_Id=?")
		sqlParams = append(sqlParams, req.TenantID)
	}

	return page.StartPage(s.db, req.PageIndex, req.PageSize).DoScan(res, sb.String(), sqlParams...)
}

func (s *ApplicationService) QueryAppCodeLanguage() []models.SgrCodeApplicationLanguage {
	var languageList []models.SgrCodeApplicationLanguage
	s.db.Model(&models.SgrCodeApplicationLanguage{}).Find(&languageList)
	return languageList
}

func (s *ApplicationService) QueryAppLevel() []models.SgrCodeApplicationLevel {
	var levelList []models.SgrCodeApplicationLevel
	s.db.Model(&models.SgrCodeApplicationLevel{}).Find(&levelList)
	return levelList
}

func (s *ApplicationService) QueryDeployLevel() []models.SgrCodeDeploymentLevel {
	var levelList []models.SgrCodeDeploymentLevel
	s.db.Model(&models.SgrCodeDeploymentLevel{}).Find(&levelList)
	return levelList
}

func (s *ApplicationService) InitGitRepository(tenantId uint64, appName string) (string, error) {
	tenant := models.SgrTenant{}
	dberr := s.db.Model(models.SgrTenant{}).Where("id=?", tenantId).First(&tenant)
	if dberr.Error != nil {
		return "", dberr.Error
	}
	sb := strings.Builder{}
	sb.WriteString(s.config.GetString("GIT_URL"))
	sb.WriteString("/")
	sb.WriteString(tenant.TCode)
	sb.WriteString("/")
	sb.WriteString(appName)
	sb.WriteString(".git")
	gitUrl := sb.String()
	return gitUrl, nil
}

func (s *ApplicationService) GetAppInfo(appId uint64) (dto.ApplicationDisplayDTO, error) {
	sql := `
SELECT t.t_name tenantName,app.name appName,app.labels,app.git,app.imagehub hub,lev.name level ,app.git_type ,app.sc_id, lang.name language ,app.status 
from sgr_tenant_application app 
INNER JOIN sgr_code_application_level lev on lev.id = app.level
INNER JOIN sgr_code_application_language lang on lang.id = app.language
INNER JOIN sgr_tenant t on t.id = app.tenant_Id 
WHERE app.id = ?
`
	var appInfo dto.ApplicationDisplayDTO
	err := s.db.Raw(sql, appId).First(&appInfo).Error
	return appInfo, err
}

func (s *ApplicationService) GetServiceConnectionById(Id uint64) (models.ServiceConnectionDetails, error) {
	var model models.ServiceConnectionDetails
	res := s.db.Model(&models.ServiceConnectionDetails{}).Where("main_id=?", Id).First(&model)
	return model, res.Error
}

func (s *ApplicationService) GetAppCountByDeployLevel(appId uint64) ([]dto.DeployLeveLCountInfo, error) {
	sql := `SELECT lev.name label,lev.code  value,IFNULL(dep.count,0) count FROM sgr_code_deployment_level lev
LEFT JOIN (
   SELECT  level,COUNT(level) count FROM sgr_tenant_deployments WHERE app_id = ? 
	 GROUP BY level
) dep on dep.level = lev.code`
	var list []dto.DeployLeveLCountInfo
	err := s.db.Raw(sql, appId).Find(&list).Error
	return list, err
}

func (s *ApplicationService) GetProjectCountByDeployLevel(projectId uint64) ([]dto.DeployLeveLCountInfo, error) {
	sql := `SELECT lev.name label,lev.code  value,IFNULL(dep.count,0) count FROM sgr_code_deployment_level lev
LEFT JOIN (
   SELECT  level,COUNT(level) count FROM sgr_tenant_deployments WHERE app_id in (select application_id from devops_projects_apps WHERE project_id =?)
	 GROUP BY level
) dep on dep.level = lev.code`
	var list []dto.DeployLeveLCountInfo
	err := s.db.Raw(sql, projectId).Find(&list).Error
	return list, err
}
