package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/pkg/page"
)

type QueryUserRequest struct {
	mvc.RequestBody

	TenantID int64  `json:"tenantId" uri:"tenantId"`
	UserName string `json:"userName" uri:"userName"`
	Mobile   string `json:"mobile" uri:"mobile"`
	Email    string `json:"email" uri:"email"`

	*page.PageRequest
}
