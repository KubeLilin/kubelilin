package tenant

import (
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
)

type RoleMenuService struct {
	db *gorm.DB
}

func NewRoleMenuService(db *gorm.DB) *RoleMenuService {
	return &RoleMenuService{db: db}
}

func (rms *RoleMenuService) CreateRoleMenuMap(req *req.RoleMenuListReq) error {
	var menuRoleData = make([]models.SgrRoleMenuMap, 0)
	for _, x := range req.RoleMenuList {
		var menuRoleEle = models.SgrRoleMenuMap{}
		copier.Copy(&menuRoleEle, &x)
		fmt.Println(x)
		fmt.Println(menuRoleEle)
		menuRoleData = append(menuRoleData, menuRoleEle)
	}
	return rms.db.Transaction(func(db *gorm.DB) error {
		if err := db.Model(&models.SgrRoleMenuMap{}).Delete(models.SgrRoleMenuMap{}, "role_id=?", req.RoleMenuList[0].RoleID).Error; err != nil {
			return err
		}
		if err := db.Model(&models.SgrRoleMenuMap{}).Create(menuRoleData).Error; err != nil {
			return err
		}
		return nil
	})
}

func (rms *RoleMenuService) QueryRleMenuMap(roleIdArr []string) *[]models.SgrRoleMenuMap {
	var data = &[]models.SgrRoleMenuMap{}
	rms.db.Model(&models.SgrRoleMenuMap{}).Where("roleId IN ？", roleIdArr).Find(data)
	return data
}
