package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
	"strconv"
)

type TenantRoleController struct {
	mvc.ApiController
	service *tenant.TenantRoleService
}

func NewTenantRoleController(service *tenant.TenantRoleService) *TenantRoleController {
	return &TenantRoleController{service: service}
}

func (c *TenantRoleController) PostTenantRole(req *req.TenantRoleReq) mvc.ApiResult {
	success, res := c.service.CreateTenantRole(req)
	return mvc.ApiResult{
		Success: success,
		Data:    res,
		Message: "添加成功",
	}
}

func (c *TenantRoleController) PutTenantRole(req *req.TenantRoleReq) mvc.ApiResult {
	success, res := c.service.UpdateTenantRole(req)
	return mvc.ApiResult{
		Success: success,
		Data:    res,
		Message: "修改成功",
	}
}

func (c *TenantRoleController) DeleteTenantRole(ctx *context.HttpContext) mvc.ApiResult {
	id := ctx.Input.Param("id")
	success := c.service.DeleteTenantRole(id)
	return mvc.ApiResult{
		Success: success,
		Data:    id,
		Message: "删除成功",
	}
}

func (c *TenantRoleController) GetTenantRoleList(ctx *context.HttpContext) mvc.ApiResult {
	keyword := ctx.Input.QueryDefault("keyword", "")
	strPageIndex := ctx.Input.QueryDefault("pageIndex", "")
	strPageSize := ctx.Input.QueryDefault("pageSize", "")
	pageIndex, _ := strconv.Atoi(strPageIndex)
	pageSize, _ := strconv.Atoi(strPageSize)

	res := c.service.QueryTenantRoleList(keyword, pageIndex, pageSize)
	return mvc.ApiResult{
		Success: res != nil,
		Data:    res,
	}
}
