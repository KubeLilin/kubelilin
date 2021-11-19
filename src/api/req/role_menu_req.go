package req

import "github.com/yoyofx/yoyogo/web/mvc"

type RoleMenuReq struct {
	RoleID int64 `json:"roleId" uri:"roleId"`
	MenuID int64 `json:"menuId" uri:"menuId"`
}

type RoleMenuListReq struct {
	mvc.RequestBody
	RoleMenuList []RoleMenuReq `json:"roleMenuList" uri:"roleMenuList"`
}
