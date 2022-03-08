package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type UserRoleReq struct {
	ID     string `json:"id" uri:"id"`
	UserID int64  `json:"userId" uri:"userId"`
	RoleID int64  `json:"roleId" uri:"roleId"`
	page.PageRequest
}

type UserRoleListReq struct {
	mvc.RequestBody
	UserRoleList []UserRoleReq `json:"userRoleList"`
}
