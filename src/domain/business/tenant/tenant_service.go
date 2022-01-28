package tenant

import (
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
	ts.db.Model(models.SgrTenant{}).Where("t_code = ?", tenant.TCode).First(&t1)
	if t1.ID > 0 {
		// create 重复 Code
		return false
	}
	err := ts.db.Transaction(func(tx *gorm.DB) error {
		//创建租户
		dbRes := ts.db.Create(tenant)
		if dbRes.Error != nil {
			return dbRes.Error
		}
		//创建第一个用户
		tenantAdmin := &models.SgrTenantUser{
			TenantID: tenant.ID,
			Account:  tenant.TName + "admin",
			UserName: tenant.TCode + "-admin",
			Password: "1234abcd",
		}
		if err := ts.db.Create(tenantAdmin).Error; err != nil {
			return err
		}
		//给用户分配权限 先鸽了哪天想写了再写
		userRole := &models.SgrTenantUserRole{
			UserID: tenantAdmin.ID,
			RoleID: 2, // 2: admin 租户管理员
		}
		if err := ts.db.Create(userRole).Error; err != nil {
			return err
		}
		return nil
	})

	return err == nil
}

func (ts *TenantService) UpdateTenant(tenant *models.SgrTenant) *models.SgrTenant {
	ts.db.Save(tenant)
	return tenant
}

func (ts *TenantService) ChangeStatus(id uint64, status int8) bool {
	if status == 1 {
		status = 2
	} else {
		status = 1
	}
	res := ts.db.Model(&models.SgrTenant{}).Where("id=?", id).Update(models.SgrTenantColumns.Status, status)
	return res.Error == nil
}

func (ts *TenantService) QueryTenantList(request *req.TenantRequest) *page.Page {
	params := &models.SgrTenant{}
	err := copier.Copy(params, request)
	params.TName = ""
	if err != nil {
		panic(err)
	}
	condition := ts.db.Model(&models.SgrTenant{}).Where(params)
	if request.TName != "" {
		condition.Where("t_name like ?", "%"+request.TName+"%")
	}
	return page.StartPage(condition, request.PageIndex, request.PageSize).DoFind(&[]models.SgrTenant{})
}
