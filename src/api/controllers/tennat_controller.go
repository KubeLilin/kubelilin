package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
	"sgr/domain/database/models"
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
func (controller TenantController) PostCreate(tenant *req.TenantRequest) mvc.ApiResult {
	fmt.Println(tenant)
	res := controller.Service.CreateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: *tenant.Status,
	})
	return mvc.ApiResult{Data: res, Success: res, Message: ""}
}

func (controller TenantController) PostUpdate(tenant *req.TenantRequest) mvc.ApiResult {
	res := controller.Service.UpdateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: *tenant.Status,
	})
	return mvc.ApiResult{Data: res}
}

func (controller TenantController) PostStatus(tenant *req.TenantRequest) mvc.ApiResult {
	res := controller.Service.ChangeStatus(tenant.ID, *tenant.Status)
	return mvc.ApiResult{Data: res, Success: res != 0}
}

func (controller TenantController) GetTenantList(ctx *context.HttpContext) mvc.ApiResult {
	var tenantRequest = &req.TenantRequest{}
	err := ctx.BindWithUri(tenantRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println(tenantRequest)
	res := controller.Service.QueryTenantList(tenantRequest)
	return controller.OK(res)
}
