package tenant

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"kubelilin/api/req"
	"kubelilin/api/res"
	"kubelilin/domain/database/models"
	"kubelilin/pkg/page"
)

type TenantUserRoleService struct {
	db *gorm.DB
}

func NewTenantUserRoleService(db *gorm.DB) *TenantUserRoleService {
	return &TenantUserRoleService{
		db: db,
	}
}

func (urs *TenantUserRoleService) CreateUserRole(req *req.UserRoleListReq) (bool, *[]models.SgrTenantUserRole) {
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

func (urs *TenantUserRoleService) DeleteUserRole(id string) bool {
	res := urs.db.Model(&models.SgrTenantUserRole{}).Delete(&models.SgrTenantUserRole{}, id)
	return res.RowsAffected > 0
}

func (urs *TenantUserRoleService) QueryUserRole(req req.UserRoleReq) (error, *page.Page) {
	var resData = &[]res.UserRoleRes{}
	sql := "select t1.id,t1.role_id,t1.user_id,t2.role_name from sgr_tenant_user_role as t1 inner join " +
		"sgr_tenant_role as t2 on t1.role_id=t2.id where t1.user_id=? "
	return page.StartPage(urs.db, req.PageIndex, req.PageSize).DoScan(resData, sql, req.UserID)
}

func (urs *TenantUserRoleService) AllRoles() ([]res.UserRoleRes, error) {
	var resData []res.UserRoleRes
	sql := "select t2.id as role_id,t2.role_name from sgr_tenant_role as t2"
	err := urs.db.Raw(sql).Scan(&resData).Error
	return resData, err
}
