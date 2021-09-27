package controllers

import (
	"encoding/json"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
	"sgr/domain/database/models"
	"strconv"
)

type SysMenuController struct {
	mvc.ApiController
	service *tenant.SysMenuService
}

func NewSysMenuController(service *tenant.SysMenuService) *SysMenuController {
	return &SysMenuController{service: service}
}

func (c *SysMenuController) CreateMenu(ctx *context.HttpContext) mvc.ApiResult {
	var menu *models.SgrSysMenu
	err := ctx.Bind(&menu)
	if err != nil {
		return c.Fail(err.Error())
	}

	success, res := c.service.CreateMenu(menu)
	msg := "添加成功"
	if !success {
		msg = "添加失败"
	}
	return mvc.ApiResult{
		Success: success,
		Data:    res,
		Message: msg,
	}
}

func (c *SysMenuController) UpdateMenu(ctx *context.HttpContext) mvc.ApiResult {
	var menu *models.SgrSysMenu
	err := ctx.Bind(&menu)
	if err != nil {
		return c.Fail(err.Error())
	}
	success, res := c.service.UpdateMenu(menu)
	return mvc.ApiResult{
		Success: success,
		Data:    res,
		Message: "修改成功",
	}
}

func (c *SysMenuController) DeleteMenu(ctx *context.HttpContext) mvc.ApiResult {
	strId := ctx.Input.Query("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	res := c.service.DelMenu(id)
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

func (c *SysMenuController) GetUserMenuTree(ctx *context.HttpContext) mvc.ApiResult {
	var userId string
	userId = ctx.Input.Param("userId")
	menuTree := c.service.MenuTree(userId)
	return mvc.ApiResult{
		Success: true,
		Message: "查询成功",
		Data:    menuTree,
	}
}

func (c *SysMenuController) GetRoleMenuList(ctx *context.HttpContext) mvc.ApiResult {
	var strRoleId string
	strRoleId = ctx.Input.QueryDefault("roleId", "")
	if strRoleId == "" {
		return c.Fail("role id is null")
	}

	roleId, err := strconv.ParseInt(strRoleId, 10, 64)
	if err != nil {
		return c.Fail("role id format is error")
	}
	return c.OK(c.service.GetRoleMenuIdList(roleId))
}

// GetQueryList
// 目前返回值是字符串，最终 以对象形式输出，c.OK(menuList: []dto.SysMenuRoutes)
// 前端在 app.tsx/layout.menu 中修改对应菜单显示。
func (c *SysMenuController) GetQueryList(ctx *context.HttpContext) mvc.ApiResult {
	var userId string
	userId = ctx.Input.Param("userId")
	menuTree := c.service.MenuTree(userId)
	bytes, err := json.Marshal(menuTree)
	if err != nil {
		return mvc.ApiResult{}
	}
	return c.OK(string(bytes))
}

const menuList = `[
    {
        "icon": "user",
        "name": "account",
        "path": "/account",
        "routes": [
            {
                "component": "./account/manage",
                "name": "manage",
                "path": "/account/manage"
            },
			{
			  "component": "./account/role",
			  "name": "role",
			  "path": "/account/role"
			},
			{
			  "component": "./account/tenant",
			  "name": "tenant",
			  "path": "/account/tenant"
			}
        ]
    }
]`
