package requests

import (
	"fmt"
	"github.com/yoyofx/yoyogo/utils/jwt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/global"
	"kubelilin/pkg/page"
	"strconv"
)

type QueryUserRequest struct {
	mvc.RequestBody

	TenantID int64  `json:"tenantId" uri:"tenantId"`
	UserName string `json:"userName" uri:"userName"`
	Mobile   string `json:"mobile" uri:"mobile"`
	Email    string `json:"email" uri:"email"`
	Status   *int8  `json:"status" uri:"status"` // 状态
	*page.PageRequest
}

// LoginRequest 用户登录请求
type (
	LoginRequest struct {
		mvc.RequestBody

		UserName  string `json:"username"`
		Password  string `json:"password"`
		AutoLogin bool   `json:"autoLogin"`
		LoginType string `json:"type"`
	}

	LoginResult struct {
		UserId    uint64 `json:"userId"`
		Status    string `json:"status"`
		LoginType string `json:"type"`
		Authority string `json:"currentAuthority"`

		Token   string `json:"token"`
		Expires int64  `json:"expires"`
	}
)

type (
	// UserInfoResponse 用户信息响应
	UserInfoResponse struct {
		Name        string        `json:"name"`
		Avatar      string        `json:"avatar"`
		Userid      string        `json:"userid"`
		Email       string        `json:"email"`
		Signature   string        `json:"signature"`
		Title       string        `json:"title"`
		Group       string        `json:"group"`
		Tags        []interface{} `json:"tags"`
		NotifyCount int           `json:"notifyCount"`
		UnreadCount int           `json:"unreadCount"`
		Country     string        `json:"country"`
		Access      string        `json:"access"`
		Address     string        `json:"address"`
		Phone       string        `json:"phone"`
	}
)

type JwtCustomClaims struct {
	jwt.StandardClaims

	// addition
	Uid      uint  `json:"uid"`
	TenantId int64 `json:"tenantId"`
	Admin    bool  `json:"admin"`
}

type UserInfo struct {
	UserId   int64
	TenantID uint64
}

func GetUserInfo(ctx *context.HttpContext) *UserInfo {
	defaultUserInfo := &UserInfo{
		UserId:   0,
		TenantID: 0,
	}
	mappings := ctx.GetItem("userinfo")

	if mappings != nil {
		maps := mappings.(map[string]interface{})
		defaultUserInfo.UserId, _ = strconv.ParseInt(fmt.Sprintf("%v", maps["uid"]), 10, 64)
		defaultUserInfo.TenantID, _ = strconv.ParseUint(fmt.Sprintf("%v", maps["tenantId"]), 10, 64)
		return defaultUserInfo
	} else {
		global.GlobalLogger.Error("Not found user info by Jwt claims !")
		global.GlobalLogger.Error("Can't get user info! Please turn on JWT Authentication At yoyogo.application.server.jwt.enabled=true with config file, and restart. So that is need to login again ! ")
	}

	return defaultUserInfo
}
