package req

import "github.com/yoyofx/yoyogo/web/mvc"

type RoleMenuReq struct {
	mvc.RequestBody
	RoleID uint64 `json:"roleId" uri:"roleId"`
	MenuID uint64 `json:"menuId" uri:"menuId"`
}

type RoleMenuListReq struct {
	mvc.RequestBody
	RoleMenuList []RoleMenuReq `json:"roleMenuList" uri:"roleMenuList"`
}
