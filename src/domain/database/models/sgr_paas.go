package models

import "time"

// SgrRoleMenuMap 角色菜单权限影射
type SgrRoleMenuMap struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	RoleID       uint64     `gorm:"column:role_id;type:bigint(20) unsigned;not null" json:"roleId"`
	MenuID       uint64     `gorm:"column:menu_id;type:bigint(20) unsigned;not null" json:"menuId"`
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrRoleMenuMap) TableName() string {
	return "sgr_role_menu_map"
}

// SgrRoleMenuMapColumns get sql column name.获取数据库列名
var SgrRoleMenuMapColumns = struct {
	ID           string
	RoleID       string
	MenuID       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	RoleID:       "role_id",
	MenuID:       "menu_id",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrSysMenu 菜单
type SgrSysMenu struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	TenantID     int64      `gorm:"column:tenant_id;type:bigint(11);not null" json:"tenantId"`                    // 租户
	MenuCode     string     `gorm:"column:menu_code;type:varchar(100);not null" json:"menuCode"`                  // 编码
	MenuName     string     `gorm:"column:menu_name;type:varchar(50);not null" json:"menuName"`                   // 目录名称
	Icon         string     `gorm:"column:icon;type:varchar(50)" json:"icon"`                                     // 图标
	Path         string     `gorm:"column:path;type:varchar(100);not null" json:"path"`                           // 路由路径
	Component    string     `gorm:"column:component;type:varchar(100)" json:"component"`                          // react组件路径
	IsRoot       int8       `gorm:"column:is_root;type:tinyint(3);not null;default:0" json:"isRoot"`              // 是否是根目录
	ParentID     uint64     `gorm:"column:parent_id;type:bigint(20) unsigned;not null;default:0" json:"parentId"` // 父层级id
	Sort         int        `gorm:"column:sort;type:int(11);not null;default:0" json:"sort"`                      // 权重，正序排序
	Status       int8       `gorm:"column:status;type:tinyint(4);not null;default:0" json:"status"`               // 状态
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
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
	Icon         string
	Path         string
	Component    string
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
	Icon:         "icon",
	Path:         "path",
	Component:    "component",
	IsRoot:       "is_root",
	ParentID:     "parent_id",
	Sort:         "sort",
	Status:       "status",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrTenant 租户
type SgrTenant struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	TName        string     `gorm:"column:t_name;type:varchar(50);not null" json:"tName"`           // 租户名称
	TCode        string     `gorm:"unique;column:t_code;type:varchar(16);not null" json:"tCode"`    // 租户编码
	Status       int8       `gorm:"column:status;type:tinyint(3);not null;default:0" json:"status"` // 状态
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime" json:"creationTime"`         // 创建时间
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`             // 修改时间
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

// SgrTenantCluster [...]
type SgrTenantCluster struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"` // ID
	TenantID     *int64     `gorm:"column:tenant_id;type:bigint(20)" json:"tenantId"`                 // 租户ID
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                // 集群名称
	Version      string     `gorm:"column:version;type:varchar(50)" json:"version"`                   // k8s 版本号
	Distribution string     `gorm:"column:distribution;type:varchar(30)" json:"distribution"`         // 来源
	Config       string     `gorm:"column:config;type:text;not null" json:"-"`                        // k8s config text
	Sort         *int       `gorm:"column:sort;type:int(11)" json:"sort"`                             // 排序
	Status       int8       `gorm:"column:status;type:tinyint(4);not null" json:"status"`             // 状态
	CreateTime   *time.Time `gorm:"column:create_time;type:datetime;not null" json:"createTime"`      // 创建时间
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;not null" json:"updateTime"`      // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantCluster) TableName() string {
	return "sgr_tenant_cluster"
}

// SgrTenantClusterColumns get sql column name.获取数据库列名
var SgrTenantClusterColumns = struct {
	ID           string
	TenantID     string
	Name         string
	Version      string
	Distribution string
	Config       string
	Sort         string
	Status       string
	CreateTime   string
	UpdateTime   string
}{
	ID:           "id",
	TenantID:     "tenant_id",
	Name:         "name",
	Version:      "version",
	Distribution: "distribution",
	Config:       "config",
	Sort:         "sort",
	Status:       "status",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// SgrTenantRole 租户角色
type SgrTenantRole struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	RoleCode     string     `gorm:"uniqueIndex:un_role_code_name;column:role_code;type:varchar(30);not null" json:"roleCode"` // 角色编码
	RoleName     string     `gorm:"uniqueIndex:un_role_code_name;column:role_name;type:varchar(50);not null" json:"roleName"` // 角色名称
	Description  string     `gorm:"column:description;type:varchar(50)" json:"description"`                                   // 角色描述
	Status       int8       `gorm:"column:status;type:tinyint(3);not null;default:0" json:"status"`                           // 状态
	TenantID     int64      `gorm:"column:tenant_id;type:bigint(11);not null" json:"tenantId"`                                // 租户
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
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
	Description  string
	Status       string
	TenantID     string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	RoleCode:     "role_code",
	RoleName:     "role_name",
	Description:  "description",
	Status:       "status",
	TenantID:     "tenant_id",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrTenantUser 用户信息
type SgrTenantUser struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	TenantID     uint64      `gorm:"column:tenant_id;type:bigint(11);not null" json:"tenantId"`      // 租户
	UserName     string     `gorm:"column:user_name;type:varchar(50)" json:"userName"`              // 用户名
	Account      string     `gorm:"column:account;type:varchar(50);not null" json:"account"`        // 账号
	Password     string     `gorm:"column:password;type:varchar(255);not null" json:"password"`     // 密码
	Mobile       string     `gorm:"column:mobile;type:varchar(20)" json:"mobile"`                   // 手机
	Email        string     `gorm:"column:email;type:varchar(50)" json:"email"`                     // 邮箱
	Status       int8       `gorm:"column:status;type:tinyint(3);not null;default:0" json:"status"` // 状态
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	UserID       uint64      `gorm:"column:user_id;type:bigint(20);not null" json:"userId"` // 用户id
	RoleID       int64      `gorm:"column:role_id;type:bigint(20);not null" json:"roleId"` // 角色id
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
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
