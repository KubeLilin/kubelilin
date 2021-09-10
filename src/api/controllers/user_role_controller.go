package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
)

type UserRoleController struct {
	mvc.ApiController
	Service *tenant.UserRoleService
}

func NewUserRoleController(sc *tenant.UserRoleService) *UserRoleController {
	return &UserRoleController{
		Service: sc,
	}
}

func (c *UserRoleController) PostUserRole(req *req.UserRoleListReq) mvc.ApiResult {
	success, res := c.Service.CreateUserRole(req)
	return mvc.ApiResult{
		Success: success,
		Message: "操作成功",
		Data:    res,
	}
}

func (c *UserRoleController) DeleteUserRole(ctx context.HttpContext) mvc.ApiResult {
	id := ctx.Input.Query("id")
	res := c.Service.DeleteUserRole(id)
	return mvc.ApiResult{
		Success: res,
		Message: "操作完成",
	}
}

func (c *UserRoleController) QueryUserRole(req req.UserRoleReq) mvc.ApiResult {
	res := c.Service.QueryUserRole(req)
	return mvc.ApiResult{
		Success: res != nil,
		Data:    res,
		Message: "查询成功",
	}
}
