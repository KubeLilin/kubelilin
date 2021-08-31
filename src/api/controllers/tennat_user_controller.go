package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
	dbmodels "sgr/domain/database/models"
	"strconv"
)

type UserController struct {
	mvc.ApiController
	Service *tenant.UserService
}

func NewUserController(service *tenant.UserService) *UserController {
	return &UserController{Service: service}
}

func (user *UserController) GetInfo(ctx *context.HttpContext) mvc.ApiResult {
	strId := ctx.Input.QueryDefault("id", "")
	userId, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		user.Fail(err.Error())
	}
	userInfo := user.Service.GetById(userId)
	return mvc.ApiResult{
		Success: userInfo != nil,
		Message: "获取用户信息",
		Data:    userInfo,
	}
}

func (user *UserController) PostRegister(ctx *context.HttpContext) mvc.ApiResult {
	var registerUser *dbmodels.SgrTenantUser
	_ = ctx.Bind(&registerUser)

	registerUser.Status = 1
	ok := user.Service.Register(registerUser)
	return mvc.ApiResult{
		Success: ok,
		Message: "注册成功",
	}
}

func (user *UserController) PostUpdate(modifyUser *dbmodels.SgrTenantUser) mvc.ApiResult {
	ok := user.Service.Update(modifyUser)
	return mvc.ApiResult{
		Success: ok,
		Message: "修改成功",
	}
}

func (user *UserController) DeleteUnRegister(ctx *context.HttpContext) mvc.ApiResult {
	idStr := ctx.Input.QueryDefault("id", "")
	userId, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	ok := user.Service.Delete(userId)

	return mvc.ApiResult{
		Success: ok,
		Message: "删除成功",
	}
}

func (user *UserController) GetList(ctx *context.HttpContext) mvc.ApiResult {
	request := &req.QueryUserRequest{}
	err := ctx.BindWithUri(request)
	if err != nil {
		panic(err)
	}
	res := user.Service.QueryUserList(request)
	return mvc.ApiResult{
		Success: res != nil,
		Data:    res,
		Message: "查询成功",
	}
}
