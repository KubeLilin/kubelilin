package tenant

import (
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	dbmodels "kubelilin/domain/database/models"
	"kubelilin/pkg/page"
)

type UserService struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (user *UserService) GetById(id int64) *dbmodels.SgrTenantUser {
	var tenantUser *dbmodels.SgrTenantUser
	res := user.db.First(&tenantUser, "id = ? AND status = ?", id, 1)
	if res.Error != nil {
		return nil
	}
	return tenantUser
}

func (user *UserService) GetUserByName(name string) *dbmodels.SgrTenantUser {
	var tenantUser *dbmodels.SgrTenantUser
	res := user.db.First(&tenantUser, "user_name = ?", name)
	if res.Error != nil {
		return nil
	}
	return tenantUser
}

func (user *UserService) GetTenantById(tenantId uint64) *dbmodels.SgrTenant {
	var tenant *dbmodels.SgrTenant
	res := user.db.First(&tenant, "id = ?", tenantId)
	if res.Error != nil {
		return nil
	}
	return tenant
}

func (user *UserService) GetUserByNameAndPassword(name string, password string) *dbmodels.SgrTenantUser {
	var tenantUser *dbmodels.SgrTenantUser
	res := user.db.First(&tenantUser, "user_name = ? AND  password = ? AND  status = ?", name, password, 1)
	if res.Error != nil {
		return nil
	}
	return tenantUser
}

func (user *UserService) Register(newUser *dbmodels.SgrTenantUser) bool {
	res := user.db.Create(newUser)
	return res.RowsAffected > 0
}

func (user *UserService) Update(modifyUser *dbmodels.SgrTenantUser) bool {
	res := user.db.Save(modifyUser)
	return res.RowsAffected > 0
}

func (user *UserService) SetStatus(id int64, status int) bool {
	res := user.db.Model(&dbmodels.SgrTenantUser{}).Where("id=?", id).Update(dbmodels.SgrTenantUserColumns.Status, status)
	return res.RowsAffected > 0
}

func (user *UserService) Delete(id int64) bool {
	res := user.db.Model(&dbmodels.SgrTenantUser{}).Where("id=?", id).Update(dbmodels.SgrTenantUserColumns.Status, 0)
	return res.RowsAffected > 0
}

func (user *UserService) QueryUserList(request *requests.QueryUserRequest) *page.Page {
	condition := user.db.Model(&dbmodels.SgrTenantUser{})

	if request.TenantID > 0 {
		condition.Where("tenant_id = ?", request.TenantID)
	}
	if request.UserName != "" {
		condition.Where("user_name LIKE ?", request.UserName+"%")
	}

	if request.Mobile != "" {
		condition.Where("mobile LIKE ?", "%"+request.Mobile+"%")
	}

	if request.Email != "" {
		condition.Where("email LIKE ?", request.Email+"%")
	}
	if request.Status != nil {
		condition.Where("status = ?", request.Status)
	}
	condition.Order("status DESC")
	return page.StartPage(condition, request.PageIndex, request.PageSize).DoFind(&[]dbmodels.SgrTenantUser{})
}
