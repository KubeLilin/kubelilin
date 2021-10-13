package tenant

import (
	"fmt"
	"github.com/jinzhu/copier"
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

func (ts *TenantService) CreateTenant(tenant *models.SgrTenant) bool {
	t1 := models.SgrTenant{}
	res := ts.db.Model(models.SgrTenant{}).Where("t_code = ?", tenant.TCode).First(&t1)
	if t1.ID > 0 {
		// create 重复 Code
		return false
	}
	res = ts.db.Create(tenant)
	return res.RowsAffected > 0
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
	params := &models.SgrTenant{}
	err := copier.Copy(params, request)
	if err != nil {
		panic(err)
	}
	condition := ts.db.Model(&models.SgrTenant{}).Where(params)
	return page.StartPage(condition, request.PageIndex, request.PageSize).DoFind(&[]models.SgrTenant{})
}
