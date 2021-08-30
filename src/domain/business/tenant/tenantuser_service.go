package tenant

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	dbmodels "sgr/domain/database/models"
	"sgr/pkg/page"
)

type UserService struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (user *UserService) Register(newUser *dbmodels.SgrTenantUser) bool {
	res := user.db.Create(newUser)
	return res.RowsAffected > 0
}

func (user *UserService) Update(modifyUser *dbmodels.SgrTenantUser) bool {
	res := user.db.Save(modifyUser)
	return res.RowsAffected > 0
}

func (user *UserService) Delete(id int64) bool {
	res := user.db.Model(&dbmodels.SgrTenantUser{}).Where("id=?", id).Update(dbmodels.SgrTenantUserColumns.Status, 0)
	return res.RowsAffected > 0
}

func (user *UserService) QueryUserList(request *req.QueryUserRequest) *page.Page {
	var params dbmodels.SgrTenantUser
	err := copier.Copy(&params, request)
	if err != nil {
		panic(err)
	}
	condition := user.db.Model(&dbmodels.SgrTenantUser{}).Where("status = ?", 1)
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

	return page.StartPage(condition, request.PageIndex, request.PageSize).DoSelect(&[]dbmodels.SgrTenantUser{})
}
