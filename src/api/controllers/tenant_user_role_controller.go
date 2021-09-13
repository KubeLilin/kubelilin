package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
)

type TenantUserRoleController struct {
	mvc.ApiController
	Service *tenant.TenantUserRoleService
}

func NewTenantUserRoleController(sc *tenant.TenantUserRoleService) *TenantUserRoleController {
	return &TenantUserRoleController{
		Service: sc,
	}
}

func (c *TenantUserRoleController) PostUserRole(req *req.UserRoleListReq) mvc.ApiResult {
	success, res := c.Service.CreateUserRole(req)
	return mvc.ApiResult{
		Success: success,
		Message: "操作成功",
		Data:    res,
	}
}

func (c *TenantUserRoleController) DeleteUserRole(ctx *context.HttpContext) mvc.ApiResult {
	id := ctx.Input.Query("id")
	res := c.Service.DeleteUserRole(id)
	return mvc.ApiResult{
		Success: res,
		Message: "操作完成",
	}
}

func (c *TenantUserRoleController) GetUserRole(ctx *context.HttpContext) mvc.ApiResult {
	req := &req.UserRoleReq{}
	err := ctx.BindWithUri(req)
	if err != nil {
		panic(err)
	}
	err, res := c.Service.QueryUserRole(*req)
	return mvc.ApiResult{
		Success: err == nil,
		Data:    res,
		Message: "666",
	}
}
