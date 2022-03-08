package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/req"
	"kubelilin/domain/business/tenant"
)

type RoleMenuController struct {
	mvc.ApiController
	service *tenant.RoleMenuService
}

func NewRoleMenuController(service *tenant.RoleMenuService) *RoleMenuController {
	return &RoleMenuController{
		service: service,
	}
}

func (c *RoleMenuController) PostRoleMenuMap(req *req.RoleMenuListReq) mvc.ApiResult {
	fmt.Println(req)
	res := c.service.CreateRoleMenuMap(req)
	msg := "操作成功"
	if res != nil {
		msg = res.Error()
	}
	return mvc.ApiResult{
		Success: res == nil,
		Message: msg}
}
