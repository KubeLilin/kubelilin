package tenant

import (
	"fmt"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/pkg/page"
)

type TenantService struct {
	db *gorm.DB
}

func NewTenantService(db *gorm.DB) *TenantService {
	return &TenantService{db: db}
}

func (ts *TenantService) CreateTenant(tenant *models.SgrTenant) *models.SgrTenant {
	res := ts.db.Create(tenant)
	fmt.Printf("插入条数：%d", res.RowsAffected)
	return tenant
}

func (ts *TenantService) UpdateTenant(tenant *models.SgrTenant) *models.SgrTenant {
	res := ts.db.Save(tenant)
	fmt.Printf("更新条数：%d", res.RowsAffected)
	return tenant
}

func (ts *TenantService) ChangeStatus(id uint64, status int8) int64 {
	res := ts.db.Model(&models.SgrTenant{}).Where("id=?", id).Update(models.SgrTenantColumns.Status, status)
	return res.RowsAffected
}

func (ts *TenantService) QueryTenantList(request *req.TenantRequest) *page.Page {
	data := &[]models.SgrTenant{}
	ts.db.Model(&models.SgrTenant{}).Where(&models.SgrTenant{
		TName:  request.TName,
		TCode:  request.TCode,
		Status: request.Status,
	}).Offset(request.OffSet()).Limit(request.PageSize).Find(data)
	var count int64
	ts.db.Model(&models.SgrTenant{}).Where(&models.SgrTenant{
		TName:  request.TName,
		TCode:  request.TCode,
		Status: request.Status,
	}).Count(&count)
	return &page.Page{
		Data:      data,
		Total:     count,
		PageIndex: request.PageIndex,
		PageSize:  request.PageSize,
	}
}
