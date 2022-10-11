package controllers

import (
	"encoding/json"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/tenant"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"strconv"
	"time"
)

type SysMenuController struct {
	mvc.ApiController
	service *tenant.SysMenuService
}

func NewSysMenuController(service *tenant.SysMenuService) *SysMenuController {
	return &SysMenuController{service: service}
}

func (c *SysMenuController) PostCreateOrUpdateMenu(ctx *context.HttpContext) mvc.ApiResult {
	var menuDto *dto.SysMenuRoutes
	err := ctx.Bind(&menuDto)
	if err != nil {
		return c.Fail(err.Error())
	}
	exitsPathMenu := c.service.GetByPath(menuDto.Name)
	if exitsPathMenu != nil {
		return c.Fail("已存在路由信息!")
	}

	t := time.Now()
	menu := c.service.GetById(menuDto.ID)
	if menu == nil {
		// create
		menu = &models.SgrSysMenu{}
		menu.CreationTime = &t
		menu.UpdateTime = &t
		menu.Status = 1
		menu.ParentID = menuDto.ParentID
		menu.MenuName = menuDto.Name
		menu.Path = menuDto.Path
		menu.Component = "." + menuDto.Path
		menu.Icon = menuDto.Icon
		menu.Sort = menuDto.Sort
		menu.IsRoot = menuDto.IsRoot
		success, res := c.service.CreateMenu(menu)
		return mvc.ApiResult{
			Success: success,
			Data:    res,
		}

	} else {
		// update
		menu.UpdateTime = &t
		menu.MenuName = menuDto.Name
		menu.Path = menuDto.Path
		menu.Component = "." + menuDto.Path
		menu.Icon = menuDto.Icon
		menu.Sort = menuDto.Sort
		menu.Status = 1
		success, res := c.service.UpdateMenu(menu)
		return mvc.ApiResult{
			Success: success,
			Data:    res,
		}
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
	var sysReq = &requests.SysMenuReq{}
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
	userId = ctx.Input.QueryDefault("id", "")
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
