package tenant

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/pkg/page"
	"sort"
)

type SysMenuService struct {
	db *gorm.DB
}

func NewSysMenuService(db *gorm.DB) *SysMenuService {
	return &SysMenuService{db: db}
}

func (sms *SysMenuService) CreateMenu(menu *models.SgrSysMenu) (bool, *models.SgrSysMenu) {
	res := sms.db.Create(menu)
	return res.RowsAffected > 0, menu
}

func (sms *SysMenuService) UpdateMenu(menu *models.SgrSysMenu) (bool, *models.SgrSysMenu) {
	res := sms.db.Save(menu)
	return res.RowsAffected > 0, menu
}

func (sms *SysMenuService) GetById(id uint64) *models.SgrSysMenu {
	var menu *models.SgrSysMenu
	res := sms.db.First(&menu, "id = ?", id)
	if res.Error != nil {
		return nil
	}
	return menu
}

func (sms *SysMenuService) GetByPath(path string) *models.SgrSysMenu {
	var menu *models.SgrSysMenu
	res := sms.db.First(&menu, "path = ?", path)
	if res.Error != nil {
		return nil
	}
	return menu
}

func (sms *SysMenuService) DelMenu(id int64) bool {
	res := sms.db.Delete(&models.SgrSysMenu{}, id)
	return res.RowsAffected > 0
}

func (sms *SysMenuService) QueryMenuList(menuReq *requests.SysMenuReq) *page.Page {
	data := &[]models.SgrSysMenu{}
	params := &models.SgrSysMenu{}
	err := copier.Copy(menuReq, params)
	if err != nil {
		panic(err)
	}
	condition := sms.db.Model(&models.SgrSysMenu{}).Where(params)
	return page.StartPage(condition, menuReq.PageIndex, menuReq.PageSize).DoFind(data)
}

func (sms *SysMenuService) MenuTree(userId string) *[]dto.SysMenuRoutes {
	var menuList []dto.SysMenuRoutes
	var userRoleList []models.SgrTenantUserRole
	//查询用户角色
	queryUser := sms.db.Model(&models.SgrTenantUserRole{})
	if userId != "" {
		queryUser.Where("user_id=?", userId)
	}
	queryUser.Find(&userRoleList)
	if userRoleList == nil || len(userRoleList) == 0 {
		return &menuList
	}

	//查询角色对应的菜单
	var roleIdArr []int64
	for _, x := range userRoleList {
		roleIdArr = append(roleIdArr, x.RoleID)
	}
	var roleMenuMap []models.SgrRoleMenuMap
	if userId != "" {
		sms.db.Model(&models.SgrRoleMenuMap{}).Where("role_id in ?", roleIdArr).Find(&roleMenuMap)
		if len(roleMenuMap) == 0 {
			return &menuList
		}
	}
	//查询菜单列表进行匹配
	var userMenuArr []uint64
	for _, x := range roleMenuMap {
		userMenuArr = append(userMenuArr, x.MenuID)
	}
	var dataMenuList []models.SgrSysMenu
	db := sms.db.Model(&models.SgrSysMenu{})
	if userId != "" {
		db.Where(" id IN ?", userMenuArr)
	}
	db.Find(&dataMenuList)

	for _, ele := range dataMenuList {
		if ele.IsRoot == 1 {
			rootMenu := dto.SysMenuRoutes{
				ID:        ele.ID,
				Component: ele.Component,
				Icon:      ele.Icon,
				Path:      ele.Path,
				Layout:    false,
				Sort:      ele.Sort,
				Name:      ele.MenuName,
				ParentID:  ele.ParentID,
			}
			rootMenu.Routes = Recursion(rootMenu, &dataMenuList)
			menuList = append(menuList, rootMenu)
		}
	}
	return &menuList
}

func Recursion(parentMenu dto.SysMenuRoutes, sourceData *[]models.SgrSysMenu) *[]dto.SysMenuRoutes {
	var targetData []dto.SysMenuRoutes
	for _, ele := range *sourceData {
		if ele.ParentID == parentMenu.ID {
			childMenu := dto.SysMenuRoutes{
				ID:        ele.ID,
				Component: ele.Component,
				Icon:      ele.Icon,
				Path:      ele.Path,
				Layout:    false,
				Sort:      ele.Sort,
				Name:      ele.MenuName,
				ParentID:  ele.ParentID,
			}
			for _, y := range *sourceData {
				if y.ParentID == childMenu.ID {
					childMenu.Routes = Recursion(childMenu, sourceData)
				}
			}
			targetData = append(targetData, childMenu)
		}
	}
	sort.SliceStable(targetData, func(i, j int) bool {
		return targetData[i].Sort < targetData[j].Sort
	})
	return &targetData
}

func (sms *SysMenuService) GetRoleMenuIdList(roleId int64) []int64 {
	sql := "SELECT sm.id FROM  sgr_role_menu_map rmm INNER JOIN `sgr_sys_menu` sm on rmm.`menu_id` = sm.`id` WHERE rmm.role_id = ?"
	var menuIdList []int64
	sms.db.Raw(sql, roleId).Scan(&menuIdList)
	return menuIdList
}
