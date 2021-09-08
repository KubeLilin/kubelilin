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
	service *tenant.SysMenuService
}

func NewSysMenuController(service *tenant.SysMenuService) *SysMenuController {
	return &SysMenuController{service: service}
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
                "path": "/account",
                "redirect": "/account/center"
            },
            {
                "component": "./account/center",
                "icon": "smile",
                "name": "center",
                "path": "/account/center"
            },
            {
                "component": "./account/settings",
                "icon": "smile",
                "name": "settings",
                "path": "/account/settings"
            }
        ]
    },
    {
        
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
