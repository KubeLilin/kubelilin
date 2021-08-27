package req

import "sgr/pkg/page"

type SysMenuReq struct {
	ID       uint64
	TenantID int64
	MenuCode string
	MenuName string
	IsRoot   int8  // 是否是根目录
	ParentID int64 // 父层级id
	Sort     int   // 权重，正序排序
	Status   int8  // 状态
	page.PageRequest
}
