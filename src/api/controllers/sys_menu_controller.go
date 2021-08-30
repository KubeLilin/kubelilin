package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
	"sgr/domain/database/models"
)

type SysMenuController struct {
	mvc.ApiController
	service tenant.SysMenuService
}

func (c *SysMenuController) CreateMenu(menu *models.SgrSysMenu) mvc.ApiResult {
	success, res := c.service.CreateMenu(menu)
	return mvc.ApiResult{
		Success: success,
		Data:    res,
		Message: "添加成功",
	}
}

func (c *SysMenuController) UpdateMenu(menu *models.SgrSysMenu) mvc.ApiResult {
	success, res := c.service.UpdateMenu(menu)
	return mvc.ApiResult{
		Success: success,
		Data:    res,
		Message: "修改成功",
	}
}

func (c *SysMenuController) DeleteMenu(menu *models.SgrSysMenu) mvc.ApiResult {
	res := c.service.DelMenu(menu)
	return mvc.ApiResult{
		Success: res,
		Data:    res,
		Message: "删除成功",
	}
}

func (c *SysMenuController) GetMenuList(ctx *context.HttpContext) mvc.ApiResult {
	var sysReq = &req.SysMenuReq{}
	err := ctx.BindWithUri(sysReq)
	if err != nil {
		panic(err)
	}
	res := c.service.QueryMenuList(sysReq)
	return mvc.ApiResult{
		Success: res != nil,
		Data:    res,
		Message: "查询成功",
	}
}
