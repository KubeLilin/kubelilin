package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/pkg/page"
)

type UserRoleReq struct {
	ID     string `json:"id" uri:"id"`
	UserID string `json:"userId" uri:"userId"`
	RoleID string `json:"roleId" uri:"roleId"`
	page.PageRequest
}

type UserRoleListReq struct {
	mvc.RequestBody
	UserRoleList []UserRoleReq `json:"userRoleList"`
}
