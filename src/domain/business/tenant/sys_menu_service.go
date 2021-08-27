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

func (sms *SysMenuService) QueryMenuList(menuReq req.SysMenuReq) *page.Page {
	data := &[]models.SgrSysMenu{}
	var count int64
	condition := &models.SgrSysMenu{}
	copier.Copy(&menuReq, condition)
	sms.db.Model(&models.SgrSysMenu{}).Where(condition).Offset(menuReq.OffSet()).Limit(menuReq.PageSize)
	sms.db.Model(&models.SgrSysMenu{}).Where(condition).Count(&count)
	return &page.Page{
		PageIndex: menuReq.PageIndex,
		PageSize:  menuReq.PageSize,
		Total:     count,
		Data:      data,
	}

}
