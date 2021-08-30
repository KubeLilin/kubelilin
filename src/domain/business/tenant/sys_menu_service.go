package tenant

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/pkg/page"
)

type SysMenuService struct {
	db *gorm.DB
}

func NewSysMenuService(db *gorm.DB) *SysMenuService {
	return &SysMenuService{db: db}
}

func (sms *SysMenuService) CreateMenu(menu *models.SgrSysMenu) (bool, *models.SgrSysMenu) {
	res := sms.db.Create(menu)
	return res.RowsAffected > 0, menu
}

func (sms *SysMenuService) UpdateMenu(menu *models.SgrSysMenu) (bool, *models.SgrSysMenu) {
	res := sms.db.Save(menu)
	return res.RowsAffected > 0, menu
}

func (sms *SysMenuService) DelMenu(menu *models.SgrSysMenu) bool {
	res := sms.db.Delete(menu)
	return res.RowsAffected > 0
}

func (sms *SysMenuService) QueryMenuList(menuReq *req.SysMenuReq) *page.Page {
	data := &[]models.SgrSysMenu{}
	params := &models.SgrSysMenu{}
	err := copier.Copy(menuReq, params)
	if err != nil {
		panic(err)
	}
	condition := sms.db.Model(&models.SgrSysMenu{}).Where(params)
	return page.StartPage(condition, menuReq.PageIndex, menuReq.PageSize).DoSelect(data)

}
