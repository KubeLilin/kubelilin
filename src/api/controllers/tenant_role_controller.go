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
	req.Status = 1
	success, res := c.service.CreateTenantRole(req)
	return mvc.ApiResult{
		Success: success,
		Data:    res,
		Message: "添加成功",
	}
}

func (c *TenantRoleController) PostUpdateTenantRole(req *req.TenantRoleReq) mvc.ApiResult {
	req.Status = 1
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
	roleId := ctx.Input.QueryDefault("keyword", "")
	strTenantId := ctx.Input.QueryDefault("tenantId", "")
	strPageIndex := ctx.Input.QueryDefault("pageIndex", "")
	strPageSize := ctx.Input.QueryDefault("pageSize", "")
	pageIndex, _ := strconv.Atoi(strPageIndex)
	pageSize, _ := strconv.Atoi(strPageSize)
	tenantId, _ := strconv.Atoi(strTenantId)

	res := c.service.QueryTenantRoleList(roleId, tenantId, pageIndex, pageSize)
	return mvc.ApiResult{
		Success: res != nil,
		Data:    res,
	}
}
