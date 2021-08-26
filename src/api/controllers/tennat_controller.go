package controllers

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/domain/business/tenant"
	"sgr/domain/database/models"
	"time"
)

type TenantController struct {
	mvc.ApiController
	Service *tenant.TenantService
}

type SgrTentDTO struct {
	ID           uint64
	TName        string     `json:"tName"`
	TCode        string     `json:"tCode"`  // 租户编码
	Status       int8       `json:"status"` // 状态
	CreationTime *time.Time // 创建时间
	UpdateTime   *time.Time
}

func NewTenantController(service *tenant.TenantService) *TenantController {
	return &TenantController{Service: service}
}

func (controller TenantController) PostCreateTenant(tenant SgrTentDTO) mvc.ApiResult {
	res := controller.Service.CreateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: tenant.Status,
	})
	return mvc.ApiResult{Data: res}
}

func (controller TenantController) PostUpdateTenant(tenant SgrTentDTO) mvc.ApiResult {
	res := controller.Service.UpdateTenant(&models.SgrTenant{
		TName:  tenant.TName,
		TCode:  tenant.TCode,
		Status: tenant.Status,
	})
	return mvc.ApiResult{Data: res}
}
