package dto

type SysMenuTreeDTO struct {
	ID           uint64
	TenantID     int64  // 租户
	MenuCode     string // 编码
	MenuName     string // 目录名称
	IsRoot       int8   // 是否是根目录
	ParentID     uint64 // 父层级id
	Sort         int    // 权重，正序排序
	Status       int8   // 状态
	ChildrenMenu *[]SysMenuTreeDTO
}
