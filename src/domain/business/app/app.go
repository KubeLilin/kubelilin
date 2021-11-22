package app

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/pkg/page"
)

type ApplicationService struct {
	db *gorm.DB
}

func NewApplicationService(db *gorm.DB) *ApplicationService {
	return &ApplicationService{db: db}
}

func (s *ApplicationService) CreateApp(req *req.AppReq) (error, *models.SgrTenantApplication) {
	var exitCount int64
	s.db.Model(&models.SgrTenantApplication{}).Where("tenant_id=? and name=?", req.TenantId, req.Name).Count(&exitCount)
	if exitCount > 0 {
		return errors.New("already have the same name application"), nil
	}
	appModel := &models.SgrTenantApplication{}
	err := copier.Copy(appModel, req)
	if err != nil {
		return err, nil
	}
	dbRes := s.db.Model(models.SgrTenantApplication{}).Create(appModel)
	if dbRes.Error != nil {
		return dbRes.Error, nil
	}
	return nil, appModel
}

func (s *ApplicationService) UpdateApp(req *req.AppReq) (error, int64) {
	appModel := models.SgrTenantApplication{}
	appModel.Level = req.Level
	appModel.Remarks = req.Remarks
	appModel.Language = req.Language
	dbRes := s.db.Model(&models.SgrTenantApplication{}).Where("id=?", req.ID).Updates(appModel)
	if dbRes.Error != nil {
		return nil, 0
	}
	return nil, dbRes.RowsAffected
}

func (s *ApplicationService) QueryAppList(req *req.AppReq) *page.Page {
	appModel := &models.SgrTenantApplication{}
	err := copier.Copy(appModel, req)
	if err != nil {
		return nil
	}
	return page.StartPage(s.db, req.PageIndex, req.PageSize).DoFind(&[]models.SgrTenantApplication{})
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
