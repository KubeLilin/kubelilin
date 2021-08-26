package controllers

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/domain/business/tenant"
	"sgr/domain/database/models"
)

type TenantController struct {
	mvc.ApiController
	Service *tenant.TenantService
}

type TenantRequest struct {
	mvc.RequestBody
	//properties
	ID     uint64 `json:"id"`
	TName  string `json:"name"`
	TCode  string `json:"code"`   // 租户编码
	Status int8   `json:"status"` // 状态
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
func (controller TenantController) PostCreate(tenant *TenantRequest) mvc.ApiResult {
	res := controller.Service.CreateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: tenant.Status,
	})
	return mvc.ApiResult{Data: res}
}

func (controller TenantController) PostUpdate(tenant *TenantRequest) mvc.ApiResult {
	res := controller.Service.UpdateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: tenant.Status,
	})
	return mvc.ApiResult{Data: res}
}
