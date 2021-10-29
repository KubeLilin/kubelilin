package req

import (
	"fmt"
	"github.com/yoyofx/yoyogo/utils/jwt"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/pkg/page"
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
	TenantID int64
}

func GetUserInfo(ctx *context.HttpContext) *UserInfo {
	mappings := ctx.GetItem("userinfo")
	maps := mappings.(map[string]interface{})

	if maps != nil {
		uid, _ := strconv.ParseInt(fmt.Sprintf("%v", maps["uid"]), 10, 64)
		tid, _ := strconv.ParseInt(fmt.Sprintf("%v", maps["tenantId"]), 10, 64)
		return &UserInfo{
			UserId:   uid,
			TenantID: tid,
		}
	}
	return nil
}
