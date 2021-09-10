package tenant

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/domain/dto"
	"sgr/pkg/page"
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

func (sms *SysMenuService) DelMenu(menu *models.SgrSysMenu) bool {
	res := sms.db.Delete(menu)
	return res.RowsAffected > 0
}

func (sms *SysMenuService) QueryMenuList(menuReq *req.SysMenuReq) *page.Page {
	data := &[]models.SgrSysMenu{}
	params := &models.SgrSysMenu{}
	err := copier.Copy(menuReq, params)
	if err != nil {
		panic(err)
	}
	condition := sms.db.Model(&models.SgrSysMenu{}).Where(params)
	return page.StartPage(condition, menuReq.PageIndex, menuReq.PageSize).DoFind(data)
}

func (sms *SysMenuService) MenuTree() *[]dto.SysMenuTreeDTO {
	/*var userMenuList []string
	sms.db.Model(&models.SgrTenantRole{}).Select("Mnu").Where()*/
	var dataMenuList = &[]models.SgrSysMenu{}
	//sms.db.Model(&models.SgrSysMenu{}).Where(" menu_code IN ",userMenuList).Find(dataMenuList)
	sms.db.Model(&models.SgrSysMenu{}).Find(dataMenuList)
	var menuList = []dto.SysMenuTreeDTO{}
	for _, ele := range *dataMenuList {
		if ele.IsRoot == 1 {
			rootMenu := dto.SysMenuTreeDTO{
				ID:       ele.ID,
				TenantID: ele.TenantID,
				MenuCode: ele.MenuCode,
				MenuName: ele.MenuName,
				IsRoot:   ele.IsRoot,
				Sort:     ele.Sort,
				ParentID: ele.ParentID,
				Status:   ele.Status,
			}
			rootMenu.ChildrenMenu = Recursion(rootMenu, dataMenuList)
			menuList = append(menuList, rootMenu)
		}
	}
	return &menuList
}

func Recursion(parentMenu dto.SysMenuTreeDTO, sourceData *[]models.SgrSysMenu) *[]dto.SysMenuTreeDTO {
	var targetData []dto.SysMenuTreeDTO
	for _, ele := range *sourceData {
		if ele.ParentID == parentMenu.ID {
			childMenu := dto.SysMenuTreeDTO{
				ID:       ele.ID,
				TenantID: ele.TenantID,
				MenuCode: ele.MenuCode,
				MenuName: ele.MenuName,
				IsRoot:   ele.IsRoot,
				Sort:     ele.Sort,
				ParentID: ele.ParentID,
				Status:   ele.Status,
			}
			for _, y := range *sourceData {
				if y.ParentID == childMenu.ID {
					childMenu.ChildrenMenu = Recursion(childMenu, sourceData)
				}
			}
			targetData = append(targetData, childMenu)
		}
	}
	return &targetData
}
