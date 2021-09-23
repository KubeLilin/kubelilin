package controllers

import (
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

// GetMenuByUser 返回路由信息接口
func (c *SysMenuController) GetQueryList(ctx *context.HttpContext) mvc.ApiResult {
	return c.OK(menuList)
}

const menuList = `[
    {
        "layout": false,
        "path": "/user",
        "routes": [
            {
                "component": "./user/Login",
                "layout": false,
                "name": "login",
                "path": "/user/login"
            },
            {
                "path": "/user",
                "redirect": "/user/login"
            },
            {
                "component": "./user/register-result",
                "icon": "smile",
                "name": "register-result",
                "path": "/user/register-result"
            },
            {
                "component": "./user/register",
                "layout": false,
                "icon": "smile",
                "name": "register",
                "path": "/user/register"
            },
            {
                "component": "404"
            }
        ]
    },
    {
        "icon": "dashboard",
        "name": "dashboard",
        "path": "/dashboard",
        "routes": [
            {
                "path": "/dashboard",
                "redirect": "/dashboard/analysis"
            },
            {
                "component": "./dashboard/analysis",
                "icon": "smile",
                "name": "analysis",
                "path": "/dashboard/analysis"
            },
            {
                "component": "./dashboard/monitor",
                "icon": "smile",
                "name": "monitor",
                "path": "/dashboard/monitor"
            },
            {
                "component": "./dashboard/workplace",
                "icon": "smile",
                "name": "workplace",
                "path": "/dashboard/workplace"
            }
        ]
    },
    {
        "icon": "form",
        "name": "form",
        "path": "/form",
        "routes": [
            {
                "path": "/form",
                "redirect": "/form/basic-form"
            },
            {
                "component": "./form/basic-form",
                "icon": "smile",
                "name": "basic-form",
                "path": "/form/basic-form"
            },
            {
                "component": "./form/step-form",
                "icon": "smile",
                "name": "step-form",
                "path": "/form/step-form"
            },
            {
                "component": "./form/advanced-form",
                "icon": "smile",
                "name": "advanced-form",
                "path": "/form/advanced-form"
            }
        ]
    },
    {
        "icon": "table",
        "name": "list",
        "path": "/list",
        "routes": [
            {
                "component": "./list/search",
                "name": "search-list",
                "path": "/list/search",
                "routes": [
                    {
                        "path": "/list/search",
                        "redirect": "/list/search/articles"
                    },
                    {
                        "component": "./list/search/articles",
                        "icon": "smile",
                        "name": "articles",
                        "path": "/list/search/articles"
                    },
                    {
                        "component": "./list/search/projects",
                        "icon": "smile",
                        "name": "projects",
                        "path": "/list/search/projects"
                    },
                    {
                        "component": "./list/search/applications",
                        "icon": "smile",
                        "name": "applications",
                        "path": "/list/search/applications"
                    }
                ]
            },
            {
                "path": "/list",
                "redirect": "/list/table-list"
            },
            {
                "component": "./list/table-list",
                "icon": "smile",
                "name": "table-list",
                "path": "/list/table-list"
            },
            {
                "component": "./list/basic-list",
                "icon": "smile",
                "name": "basic-list",
                "path": "/list/basic-list"
            },
            {
                "component": "./list/card-list",
                "icon": "smile",
                "name": "card-list",
                "path": "/list/card-list"
            }
        ]
    },
    {
        "icon": "profile",
        "name": "profile",
        "path": "/profile",
        "routes": [
            {
                "path": "/profile",
                "redirect": "/profile/basic"
            },
            {
                "component": "./profile/basic",
                "icon": "smile",
                "name": "basic",
                "path": "/profile/basic"
            },
            {
                "component": "./profile/advanced",
                "icon": "smile",
                "name": "advanced",
                "path": "/profile/advanced"
            }
        ]
    },
    {
        "icon": "CheckCircleOutlined",
        "name": "result",
        "path": "/result",
        "routes": [
            {
                "path": "/result",
                "redirect": "/result/success"
            },
            {
                "component": "./result/success",
                "icon": "smile",
                "name": "success",
                "path": "/result/success"
            },
            {
                "component": "./result/fail",
                "icon": "smile",
                "name": "fail",
                "path": "/result/fail"
            }
        ]
    },
    {
        "icon": "warning",
        "name": "exception",
        "path": "/exception",
        "routes": [
            {
                "path": "/exception",
                "redirect": "/exception/403"
            },
            {
                "component": "./exception/403",
                "icon": "smile",
                "name": "403",
                "path": "/exception/403"
            },
            {
                "component": "./exception/404",
                "icon": "smile",
                "name": "404",
                "path": "/exception/404"
            },
            {
                "component": "./exception/500",
                "icon": "smile",
                "name": "500",
                "path": "/exception/500"
            }
        ]
    },
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
    },

    {
        "hideInMenu": true,
        "icon": "highlight",
        "name": "editor",
        "path": "/editor",
        "routes": [
            {
                "path": "/editor",
                "redirect": "/editor/flow"
            },
            {
                "component": "./editor/flow",
                "icon": "smile",
                "name": "flow",
                "path": "/editor/flow"
            },
            {
                "component": "./editor/mind",
                "icon": "smile",
                "name": "mind",
                "path": "/editor/mind"
            },
            {
                "component": "./editor/koni",
                "icon": "smile",
                "name": "koni",
                "path": "/editor/koni"
            }
        ]
    },
    {
        "path": "/",
        "redirect": "/dashboard/analysis"
    },
    {
        "component": "404"
    }
]`
