package tenant

import (
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/api/res"
	"sgr/domain/database/models"
	"sgr/pkg/page"
)

type UserRoleService struct {
	db *gorm.DB
}

func NewUserRoleService(db *gorm.DB) *UserRoleService {
	return &UserRoleService{
		db: db,
	}
}

func (urs *UserRoleService) CreateUserRole(req *req.UserRoleListReq) (bool, *[]models.SgrTenantUserRole) {
	fmt.Println(req)
	var userRoleData = make([]models.SgrTenantUserRole, 0)
	if req == nil {
		return false, nil
	}
	for _, x := range req.UserRoleList {
		var userRoleEle = &models.SgrTenantUserRole{}
		copier.Copy(userRoleEle, &x)
		userRoleData = append(userRoleData, *userRoleEle)
	}
	urs.db.Transaction(func(db *gorm.DB) error {
		if err := db.Model(&models.SgrTenantUserRole{}).Delete(models.SgrTenantUserRole{}, "user_id=?", req.UserRoleList[0].UserID).Error; err != nil {
			return err
		}
		if err := db.Model(&models.SgrTenantUserRole{}).Create(userRoleData).Error; err != nil {
			return err
		}
		return nil
	})
	return true, &userRoleData
}

func (urs *UserRoleService) DeleteUserRole(id string) bool {
	res := urs.db.Model(&models.SgrTenantUserRole{}).Delete(&models.SgrTenantUserRole{}, id)
	return res.RowsAffected > 0
}

func (urs *UserRoleService) QueryUserRole(req req.UserRoleReq) *page.Page {
	var resData = []res.UserRoleRes{}
	condition := urs.db.Raw("select t1.id,t1.role_id,t1.user_id,t2.role_name from sgr_tenant_user_role as t1 inner join"+
		"sgr_tenant_role as t2 on t1.role_id=t2.role_id where t1.user_id=? ", req.UserID)
	return page.StartPage(condition, req.PageIndex, req.PageSize).DoScan(resData)
}
