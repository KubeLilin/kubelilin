package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/api/req"
	"sgr/domain/business/tenant"
	dbmodels "sgr/domain/database/models"
	"strconv"
	"time"
)

type UserController struct {
	mvc.ApiController
	Service *tenant.UserService
}

func NewUserController(service *tenant.UserService) *UserController {
	return &UserController{Service: service}
}

func (user *UserController) PostLogin(ctx *context.HttpContext, loginRequest *req.LoginRequest) mvc.ApiResult {
	if loginRequest.UserName == "" || loginRequest.Password == "" {
		ctx.Output.SetStatus(401)
		return user.Fail("no username or password")
	}
	queryUser := user.Service.GetUserByNameAndPassword(loginRequest.UserName, loginRequest.Password)

	if queryUser == nil {
		return mvc.ApiResult{Success: true, Message: "can not find user be", Data: req.LoginResult{Status: "false"}}
	}

	return user.OK(req.LoginResult{Status: "ok", UserId: queryUser.ID, LoginType: loginRequest.LoginType, Authority: "admin"})
}

func (user *UserController) GetInfo(ctx *context.HttpContext) mvc.ApiResult {
	strId := ctx.Input.QueryDefault("id", "")
	userId, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return user.Fail(err.Error())
	}
	userInfo := user.Service.GetById(userId)
	if userInfo == nil {
		return user.Fail("fail")
	}

	return mvc.ApiResult{
		Success: userInfo != nil,
		Message: "获取用户信息",
		Data: req.UserInfoResponse{
			Name:        userInfo.UserName,
			Avatar:      "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
			Userid:      strconv.FormatUint(userInfo.ID, 10),
			Email:       userInfo.Email,
			Signature:   "",
			Title:       "",
			Group:       strconv.FormatInt(userInfo.TenantID, 10),
			Tags:        nil,
			NotifyCount: 0,
			UnreadCount: 0,
			Country:     "china",
			Access:      "-",
			Address:     "-",
			Phone:       userInfo.Mobile,
		},
	}
}

func (user *UserController) PostRegister(ctx *context.HttpContext) mvc.ApiResult {
	var registerUser *dbmodels.SgrTenantUser
	_ = ctx.Bind(&registerUser)

	ok := false
	retMessage := "注册成功"
	exitsUser := user.Service.GetUserByName(registerUser.UserName)
	if exitsUser == nil {
		t := time.Now()
		registerUser.Status = 1
		registerUser.CreationTime = &t
		registerUser.UpdateTime = &t
		ok = user.Service.Register(registerUser)
	} else {
		retMessage = "注册失败"
	}
	return mvc.ApiResult{
		Success: ok,
		Message: retMessage,
	}
}

func (user *UserController) PostUpdate(ctx *context.HttpContext) mvc.ApiResult {
	var modifyUser *dbmodels.SgrTenantUser
	_ = ctx.Bind(&modifyUser)
	t := time.Now()
	modifyUser.UpdateTime = &t
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

func (user *UserController) PutStatus(ctx *context.HttpContext) mvc.ApiResult {
	idStr := ctx.Input.QueryDefault("id", "")
	statusStr := ctx.Input.QueryDefault("status", "")
	userId, _ := strconv.ParseInt(idStr, 10, 32)
	status, _ := strconv.Atoi(statusStr)

	ok := user.Service.SetStatus(userId, status)
	return mvc.ApiResult{
		Success: ok,
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
