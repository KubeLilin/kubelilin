package tenant

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/pkg/page"
)

type TenantRoleService struct {
	db *gorm.DB
}

func NewTenantRoleService(db *gorm.DB) *TenantRoleService {
	return &TenantRoleService{
		db: db,
	}
}

func (trs *TenantRoleService) CreateTenantRole(req *req.TenantRoleReq) (bool, *models.SgrTenantRole) {
	var tenantRoelModel = &models.SgrTenantRole{}
	copier.Copy(tenantRoelModel, req)
	trs.db.Create(tenantRoelModel)
	return tenantRoelModel.ID != 0, tenantRoelModel
}

func (trs *TenantRoleService) UpdateTenantRole(req *req.TenantRoleReq) (bool, *models.SgrTenantRole) {
	var tenantRoelModel = &models.SgrTenantRole{}
	copier.Copy(tenantRoelModel, req)
	res := trs.db.Save(tenantRoelModel)
	return res.RowsAffected > 0, tenantRoelModel
}

func (trs *TenantRoleService) DeleteTenantRole(id string) bool {
	res := trs.db.Delete(&models.SgrTenantRole{}, id)
	return res.RowsAffected > 0
}

func (trs *TenantRoleService) QueryTenantRoleList(req *req.TenantRoleReq) *page.Page {
	var data = &[]models.SgrTenantRole{}
	var params = models.SgrTenantRole{}
	copier.Copy(params, req)
	condition := trs.db.Model(&models.SgrTenantRole{}).Where(params)
	return page.StartPage(condition, req.PageIndex, req.PageSize).DoFind(data)
}