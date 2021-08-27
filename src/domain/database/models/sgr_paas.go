package models

import "time"

// SgrSysMenu 菜单
type SgrSysMenu struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	TenantID     int64      `gorm:"column:tenant_id;type:bigint(11);not null"`           // 租户
	MenuCode     string     `gorm:"unique;column:menu_code;type:varchar(100);not null"`  // 编码
	MenuName     string     `gorm:"column:menu_name;type:varchar(50);not null"`          // 目录名称
	IsRoot       int8       `gorm:"column:is_root;type:tinyint(3);not null;default:0"`   // 是否是根目录
	ParentID     int64      `gorm:"column:parent_id;type:bigint(20);not null;default:0"` // 父层级id
	Sort         int        `gorm:"column:sort;type:int(11);not null;default:0"`         // 权重，正序排序
	Status       int8       `gorm:"column:status;type:tinyint(4);not null;default:0"`    // 状态
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrSysMenu) TableName() string {
	return "sgr_sys_menu"
}

// SgrSysMenuColumns get sql column name.获取数据库列名
var SgrSysMenuColumns = struct {
	ID           string
	TenantID     string
	MenuCode     string
	MenuName     string
	IsRoot       string
	ParentID     string
	Sort         string
	Status       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	TenantID:     "tenant_id",
	MenuCode:     "menu_code",
	MenuName:     "menu_name",
	IsRoot:       "is_root",
	ParentID:     "parent_id",
	Sort:         "sort",
	Status:       "status",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrTenant 租户
type SgrTenant struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	TName        string     `gorm:"column:t_name;type:varchar(50);not null"`          // 租户名称
	TCode        string     `gorm:"unique;column:t_code;type:varchar(16);not null"`   // 租户编码
	Status       int8       `gorm:"column:status;type:tinyint(3);not null;default:0"` // 状态
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime"`               // 创建时间
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime"`                 // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenant) TableName() string {
	return "sgr_tenant"
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	RoleCode     string     `gorm:"uniqueIndex:un_role_code_name;column:role_code;type:varchar(30);not null"` // 角色编码
	RoleName     string     `gorm:"uniqueIndex:un_role_code_name;column:role_name;type:varchar(50);not null"` // 角色名称
	Status       int8       `gorm:"column:status;type:tinyint(3);not null;default:0"`                         // 状态
	TenantID     int64      `gorm:"column:tenant_id;type:bigint(11);not null"`                                // 租户
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantRole) TableName() string {
	return "sgr_tenant_role"
}

// SgrTenantRoleColumns get sql column name.获取数据库列名
var SgrTenantRoleColumns = struct {
	ID           string
	RoleCode     string
	RoleName     string
	Status       string
	TenantID     string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	RoleCode:     "role_code",
	RoleName:     "role_name",
	Status:       "status",
	TenantID:     "tenant_id",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrTenantUser 用户信息
type SgrTenantUser struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	TenantID     int64      `gorm:"column:tenant_id;type:bigint(11);not null"`        // 租户
	UserName     string     `gorm:"column:user_name;type:varchar(50)"`                // 用户名
	Account      string     `gorm:"column:account;type:varchar(50);not null"`         // 账号
	Password     string     `gorm:"column:password;type:varchar(255);not null"`       // 密码
	Mobile       string     `gorm:"column:mobile;type:varchar(10)"`                   // 手机
	Email        string     `gorm:"column:email;type:varchar(50)"`                    // 邮箱
	Status       int8       `gorm:"column:status;type:tinyint(3);not null;default:0"` // 状态
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantUser) TableName() string {
	return "sgr_tenant_user"
}

// SgrTenantUserColumns get sql column name.获取数据库列名
var SgrTenantUserColumns = struct {
	ID           string
	TenantID     string
	UserName     string
	Account      string
	Password     string
	Mobile       string
	Email        string
	Status       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	TenantID:     "tenant_id",
	UserName:     "user_name",
	Account:      "account",
	Password:     "password",
	Mobile:       "mobile",
	Email:        "email",
	Status:       "status",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrTenantUserRole 用户角色
type SgrTenantUserRole struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	UserID       int64      `gorm:"column:user_id;type:bigint(20);not null"` // 用户id
	RoleID       int64      `gorm:"column:role_id;type:bigint(20);not null"` // 角色id
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantUserRole) TableName() string {
	return "sgr_tenant_user_role"
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
