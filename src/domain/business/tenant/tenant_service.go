package tenant

import (
	"fmt"
	"gorm.io/gorm"
	"sgr/domain/database/models"
)

type TenantService struct {
	db *gorm.DB
}

func NewTenantService(db *gorm.DB) *TenantService {
	return &TenantService{db: db}
}

var dsn = "user:pass@tcp(49.232.153.51:3306)/sgr_platform?charset=utf8mb4&parseTime=True&loc=Local"

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

func (ts *TenantService) ChangeStatus(id int64, status bool) int64 {
	res := ts.db.Model(&models.SgrTenant{}).Where("id=?", id).Update(models.SgrTenantColumns.Status, status)
	return res.RowsAffected
}
