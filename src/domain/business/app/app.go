package app

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yoyofx/yoyogo/abstractions"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	"sgr/pkg/page"
	"strings"
)

type ApplicationService struct {
	db         *gorm.DB
	vcsService VcsService
	config     abstractions.IConfiguration
}

func NewApplicationService(db *gorm.DB, gogsVcsService *GogsVcsService, config abstractions.IConfiguration) *ApplicationService {
	return &ApplicationService{db: db, vcsService: gogsVcsService, config: config}
}

func (s *ApplicationService) CreateApp(req *req.AppReq) (error, *models.SgrTenantApplication) {
	var exitCount int64
	s.db.Model(&models.SgrTenantApplication{}).Where("tenant_id=? and name=?", req.TenantID, req.Name).Count(&exitCount)
	if exitCount > 0 {
		return errors.New("already have the same name application"), nil
	}
	appModel := &models.SgrTenantApplication{}
	err := copier.Copy(appModel, req)
	if err != nil {
		return err, nil
	}
	dbErr := s.db.Transaction(func(tx *gorm.DB) error {
		dbRes := s.db.Model(models.SgrTenantApplication{}).Create(appModel)
		if dbRes.Error != nil {
			return nil
		}
		_, repoErr := s.vcsService.CreateTenantRepository(req.TenantID, req.Name)
		if repoErr != nil {
			return repoErr
		}
		return nil
	})
	//创建git仓库
	if dbErr != nil {
		return dbErr, nil
	}
	return nil, appModel
}

func (s *ApplicationService) UpdateApp(req *req.AppReq) (error, int64) {
	appModel := models.SgrTenantApplication{}
	appModel.Level = req.Level
	appModel.Remarks = req.Remarks
	appModel.Language = req.Language
	appModel.Status = req.Status
	//appModel.Git = req.Git
	appModel.Labels = req.Labels
	dbRes := s.db.Model(&models.SgrTenantApplication{}).Where("id=?", req.ID).Updates(appModel)
	if dbRes.Error != nil {
		return nil, 0
	}
	return nil, dbRes.RowsAffected
}

func (s *ApplicationService) QueryAppList(req *req.AppReq) (error, *page.Page) {
	res := &[]dto.ApplicationInfoDTO{}
	var sqlParams []interface{}
	sb := strings.Builder{}
	sb.WriteString("SELECT t1.*,t2.name as language_name,t3.name as level_name FROM sgr_tenant_application AS t1 INNER JOIN sgr_code_application_language AS t2  ")
	sb.WriteString(" ON t1.language = t2.id INNER JOIN sgr_code_application_level AS t3 ON t1.LEVEL = t3.id WHERE t1.status = 1 ")
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

func (s *ApplicationService) InitGitRepository(tenantId uint64, appName string) (string, error) {
	tenant := models.SgrTenant{}
	dberr := s.db.Model(models.SgrTenant{}).Where("id=?", tenantId).First(&tenant)
	if dberr.Error != nil {
		return "", dberr.Error
	}
	sb := strings.Builder{}
	sb.WriteString(s.config.GetString(GIT_URL))
	sb.WriteString("/")
	sb.WriteString(tenant.TCode)
	sb.WriteString("/")
	sb.WriteString(appName)
	sb.WriteString(".git")
	gitUrl := sb.String()
	return gitUrl, nil
}
