package controllers

import (
	"github.com/yoyofx/yoyogo/web/binding"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/dto"
	"sgr/domain/business/tenant"
	"sgr/domain/database/models"
	"sgr/pkg/page"
)

type TenantController struct {
	mvc.ApiController
	Service *tenant.TenantService
}

func NewTenantController(service *tenant.TenantService) *TenantController {
	return &TenantController{Service: service}
}

/*
	PostCreate Create tenant
	URL: http://localhost:8080/v1/tenant/create
	BODY:
	{
		"id":0,
		"name":"曹操爱吃",
		"code":"666",
		"status":0
	}
*/
func (controller TenantController) PostCreate(tenant *dto.TenantRequest) mvc.ApiResult {
	res := controller.Service.CreateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: tenant.Status,
	})
	return mvc.ApiResult{Data: res}
}

func (controller TenantController) PostUpdate(tenant *dto.TenantRequest) mvc.ApiResult {
	res := controller.Service.UpdateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: tenant.Status,
	})
	return mvc.ApiResult{Data: res}
}

func (controller TenantController) PostStatus(tenant *dto.TenantRequest) mvc.ApiResult {
	res := controller.Service.ChangeStatus(tenant.ID, tenant.Status)
	return mvc.ApiResult{Data: res}
}

func (controller TenantController) GetTenantList(ctx *context.HttpContext) *page.Page {
	var tenantRequest = &dto.TenantRequest{}
	err := ctx.BindWith(tenantRequest, binding.Form)
	if err != nil {
		panic(err)
	}
	res := controller.Service.QueryTenantList(tenantRequest)
	return res
}
