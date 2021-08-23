package models

import (
	"time"
)

// SgrSysMenu 菜单
type SgrSysMenu struct {
	ID             uint64
	TenantID       int64  // 租户
	MenuCode       string // 编码
	MenuName       string // 目录名称
	IsRoot         int8   // 是否是根目录
	ParentID       int64  // 父层级id
	Sort           int    // 权重，正序排序
	Status         int8   // 状态
	CreatetionTime time.Time
	UpdateTime     time.Time
}

// SgrSysMenuColumns get sql column name.获取数据库列名
var SgrSysMenuColumns = struct {
	ID             string
	TenantID       string
	MenuCode       string
	MenuName       string
	IsRoot         string
	ParentID       string
	Sort           string
	Status         string
	CreatetionTime string
	UpdateTime     string
}{
	ID:             "id",
	TenantID:       "tenant_id",
	MenuCode:       "menu_code",
	MenuName:       "menu_name",
	IsRoot:         "is_root",
	ParentID:       "parent_id",
	Sort:           "sort",
	Status:         "status",
	CreatetionTime: "createtion_time",
	UpdateTime:     "update_time",
}

// SgrTenant 租户
type SgrTenant struct {
	ID           uint64
	TName        string    // 租户名称
	TCode        string    // 租户编码
	Status       int8      // 状态
	CreationTime time.Time // 创建时间
	UpdateTime   time.Time // 修改时间
}

// SgrTenantColumns get sql column name.获取数据库列名
var SgrTenantColumns = struct {
	ID           string
	TName        string
	TCode        string
	Status       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	TName:        "t_name",
	TCode:        "t_code",
	Status:       "status",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrTenantRole 租户角色
type SgrTenantRole struct {
	ID             uint64
	RoleCode       string // 角色编码
	RoleName       string // 角色名称
	Status         int8   // 状态
	TenantID       int64  // 租户
	CreatetionTime time.Time
	UpdateTime     time.Time
}

// SgrTenantRoleColumns get sql column name.获取数据库列名
var SgrTenantRoleColumns = struct {
	ID             string
	RoleCode       string
	RoleName       string
	Status         string
	TenantID       string
	CreatetionTime string
	UpdateTime     string
}{
	ID:             "id",
	RoleCode:       "role_code",
	RoleName:       "role_name",
	Status:         "status",
	TenantID:       "tenant_id",
	CreatetionTime: "createtion_time",
	UpdateTime:     "update_time",
}

// SgrTenantUser 用户信息
type SgrTenantUser struct {
	ID             uint64
	TenantID       int64  // 租户
	UserName       string // 用户名
	Account        string // 账号
	Password       string // 密码
	Mobile         string // 手机
	Email          string // 邮箱
	Status         int8   // 状态
	CreatetionTime time.Time
	UpdateTime     time.Time
}

// SgrTenantUserColumns get sql column name.获取数据库列名
var SgrTenantUserColumns = struct {
	ID             string
	TenantID       string
	UserName       string
	Account        string
	Password       string
	Mobile         string
	Email          string
	Status         string
	CreatetionTime string
	UpdateTime     string
}{
	ID:             "id",
	TenantID:       "tenant_id",
	UserName:       "user_name",
	Account:        "account",
	Password:       "password",
	Mobile:         "mobile",
	Email:          "email",
	Status:         "status",
	CreatetionTime: "createtion_time",
	UpdateTime:     "update_time",
}

// SgrTenantUserRole 用户角色
type SgrTenantUserRole struct {
	ID           uint64
	UserID       int64 // 用户id
	RoleID       int64 // 角色id
	CreationTime time.Time
	UpdateTime   time.Time
}

// SgrTenantUserRoleColumns get sql column name.获取数据库列名
var SgrTenantUserRoleColumns = struct {
	ID           string
	UserID       string
	RoleID       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	UserID:       "user_id",
	RoleID:       "role_id",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}
