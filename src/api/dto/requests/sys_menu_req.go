package requests

import "kubelilin/pkg/page"

type SysMenuReq struct {
	ID       uint64 `json:"id" uri:"id"`
	TenantID int64  `json:"tenantId" uri:"tenantId"`
	MenuCode string `json:"menuCode" uri:"menuCode"`
	MenuName string `json:"menuName" uri:"menuName"`
	IsRoot   int8   `json:"isRoot" uri:"isRoot"`
	ParentID int64  `json:"parentId" uri:"parentId"`
	Sort     int    `json:"sort" uri:"sort"`
	Status   int8   `json:"status" uri:"status"`
	*page.PageRequest
}
