package models

import "time"

// SgrCodeApplicationLanguage 字典-应用开发语言
type SgrCodeApplicationLanguage struct {
	ID   uint16 `gorm:"primaryKey;column:id;type:smallint(5) unsigned;not null" json:"id"`
	Code string `gorm:"column:code;type:varchar(8)" json:"code"`
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Sort uint16 `gorm:"column:sort;type:smallint(5) unsigned;not null;default:0" json:"sort"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrCodeApplicationLanguage) TableName() string {
	return "sgr_code_application_language"
}

// SgrCodeApplicationLanguageColumns get sql column name.获取数据库列名
var SgrCodeApplicationLanguageColumns = struct {
	ID   string
	Code string
	Name string
	Sort string
}{
	ID:   "id",
	Code: "code",
	Name: "name",
	Sort: "sort",
}

// SgrCodeApplicationLevel 字典-应用级别
type SgrCodeApplicationLevel struct {
	ID   uint16 `gorm:"primaryKey;column:id;type:smallint(10) unsigned;not null" json:"id"`
	Code string `gorm:"column:code;type:varchar(8)" json:"code"`
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Sort uint16 `gorm:"column:sort;type:smallint(5) unsigned;not null;default:0" json:"sort"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrCodeApplicationLevel) TableName() string {
	return "sgr_code_application_level"
}

// SgrCodeApplicationLevelColumns get sql column name.获取数据库列名
var SgrCodeApplicationLevelColumns = struct {
	ID   string
	Code string
	Name string
	Sort string
}{
	ID:   "id",
	Code: "code",
	Name: "name",
	Sort: "sort",
}

// SgrRoleMenuMap 角色菜单权限影射
type SgrRoleMenuMap struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	RoleID       uint64     `gorm:"column:role_id;type:bigint(20) unsigned;not null" json:"roleId"` // 角色ID
	MenuID       uint64     `gorm:"column:menu_id;type:bigint(20) unsigned;not null" json:"menuId"` // 菜单ID
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

// SgrTenantApplication 集群应用
type SgrTenantApplication struct {
	ID         uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	TenantID   *uint64    `gorm:"column:tenant_Id;type:bigint(20) unsigned" json:"tenantId"`
	Name       string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                  // 集群应用名称(英文唯一)
	Labels     string     `gorm:"column:labels;type:varchar(50);not null" json:"labels"`              // 应用中文名称
	Remarks    string     `gorm:"column:remarks;type:varchar(200);not null" json:"remarks"`           // 集群应用备注
	Git        string     `gorm:"column:git;type:varchar(500);not null" json:"git"`                   // 集群应用绑定的git地址
	Level      uint16     `gorm:"column:level;type:smallint(6) unsigned;not null" json:"level"`       // 应用级别
	Language   uint16     `gorm:"column:language;type:smallint(5) unsigned;not null" json:"language"` // 开发语言
	Status     int8       `gorm:"column:status;type:tinyint(4);not null;default:0" json:"status"`     // 状态
	CreateTime *time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`                 // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`                 // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantApplication) TableName() string {
	return "sgr_tenant_application"
}

// SgrTenantApplicationColumns get sql column name.获取数据库列名
var SgrTenantApplicationColumns = struct {
	ID         string
	TenantID   string
	Name       string
	Labels     string
	Remarks    string
	Git        string
	Level      string
	Language   string
	Status     string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	TenantID:   "tenant_Id",
	Name:       "name",
	Labels:     "labels",
	Remarks:    "remarks",
	Git:        "git",
	Level:      "level",
	Language:   "language",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// SgrTenantCluster 集群信息
type SgrTenantCluster struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`   // ID
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint(20) unsigned;not null" json:"tenantId"` // 租户ID
	Nickname     string     `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`          // 别名
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                  // 集群名称
	Version      string     `gorm:"column:version;type:varchar(50)" json:"version"`                     // k8s 版本号
	Distribution string     `gorm:"column:distribution;type:varchar(30)" json:"distribution"`           // 来源
	Config       string     `gorm:"column:config;type:text;not null" json:"config"`                     // k8s config text
	Sort         int        `gorm:"column:sort;type:int(11);not null" json:"sort"`                      // 排序
	Status       int8       `gorm:"column:status;type:tinyint(4);not null" json:"status"`               // 状态
	CreateTime   *time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`                 // 创建时间
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`                 // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantCluster) TableName() string {
	return "sgr_tenant_cluster"
}

// SgrTenantClusterColumns get sql column name.获取数据库列名
var SgrTenantClusterColumns = struct {
	ID           string
	TenantID     string
	Nickname     string
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
	Nickname:     "nickname",
	Name:         "name",
	Version:      "version",
	Distribution: "distribution",
	Config:       "config",
	Sort:         "sort",
	Status:       "status",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// SgrTenantDeployments 集群部署
type SgrTenantDeployments struct {
	ID              uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`         // 部署ID
	Name            string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                        // 部署名称(英文唯一)
	Nickname        string     `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`                // 部署中文名称#
	TenantID        uint64     `gorm:"column:tenant_id;type:bigint(20) unsigned;not null" json:"tenantId"`       // 租户ID
	ClusterID       uint64     `gorm:"column:cluster_id;type:bigint(20) unsigned;not null" json:"clusterId"`     // 集群ID
	NamespaceID     uint64     `gorm:"column:namespace_id;type:bigint(20) unsigned;not null" json:"namespaceId"` // 命名空间ID
	AppID           uint64     `gorm:"column:app_id;type:bigint(20) unsigned" json:"appId"`                      // 应用ID
	AppName         string     `gorm:"column:app_name;type:varchar(50);not null" json:"appName"`                 // 应用名称(英文唯一)
	LastImage       string     `gorm:"column:last_image;type:varchar(150)" json:"lastImage"`
	Level           string     `gorm:"column:level;type:varchar(8);not null" json:"level"`                       // 环境级别 ( Prod , Test , Dev )
	ImageHub        string     `gorm:"column:image_hub;type:varchar(200)" json:"imageHub"`                       // 自动生成的镜像仓库地址( hub域名/apps/{应用名-部署名} , 如 http://hub.yoyogo.run/apps/demo-prod )
	Status          uint8      `gorm:"column:status;type:tinyint(3) unsigned;not null;default:1" json:"status"`  // 状态
	WorkloadType    string     `gorm:"column:workload_type;type:varchar(15);not null" json:"workloadType"`       // 部署类型(Deployment、DaemonSet、StatefulSet、CronJob)
	Replicas        int32      `gorm:"column:replicas;type:int(10) unsigned;not null;default:1" json:"replicas"` // 部署副本数#
	ServiceEnable   bool       `gorm:"column:service_enable;type:tinyint(1);not null" json:"serviceEnable"`      // 是否开启 Service
	ServiceName     string     `gorm:"column:service_name;type:varchar(150)" json:"serviceName"`                 // 服务名称
	ServiceAway     string     `gorm:"column:service_away;type:varchar(10)" json:"serviceAway"`                  // Service访问方式(NodePort、ClusterPort)
	ServicePortType string     `gorm:"column:service_port_type;type:varchar(8)" json:"servicePortType"`          // Service端口映射类型(TCP/UDP)
	ServicePort     *uint64    `gorm:"column:service_port;type:smallint(5) unsigned" json:"servicePort"`         // Service端口映射(容器端口->服务端口)
	CreateTime      *time.Time `gorm:"column:create_time;type:datetime;not null" json:"createTime"`              // 创建时间
	UpdateTime      *time.Time `gorm:"column:update_time;type:datetime;not null" json:"updateTime"`              // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantDeployments) TableName() string {
	return "sgr_tenant_deployments"
}

// SgrTenantDeploymentsColumns get sql column name.获取数据库列名
var SgrTenantDeploymentsColumns = struct {
	ID              string
	Name            string
	Nickname        string
	TenantID        string
	ClusterID       string
	NamespaceID     string
	AppID           string
	AppName         string
	LastImage       string
	Level           string
	ImageHub        string
	Status          string
	WorkloadType    string
	Replicas        string
	ServiceEnable   string
	ServiceName     string
	ServiceAway     string
	ServicePortType string
	ServicePort     string
	CreateTime      string
	UpdateTime      string
}{
	ID:              "id",
	Name:            "name",
	Nickname:        "nickname",
	TenantID:        "tenant_id",
	ClusterID:       "cluster_id",
	NamespaceID:     "namespace_id",
	AppID:           "app_id",
	AppName:         "app_name",
	LastImage:       "last_image",
	Level:           "level",
	ImageHub:        "image_hub",
	Status:          "status",
	WorkloadType:    "workload_type",
	Replicas:        "replicas",
	ServiceEnable:   "service_enable",
	ServiceName:     "service_name",
	ServiceAway:     "service_away",
	ServicePortType: "service_port_type",
	ServicePort:     "service_port",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// SgrTenantDeploymentsContainers [...]
type SgrTenantDeploymentsContainers struct {
	ID                uint64  `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`              // 容器ID
	Name              string  `gorm:"column:name;type:varchar(30);not null" json:"name"`                             // 容器名称：主容器默认为 app
	DeployID          uint64  `gorm:"column:deploy_id;type:bigint(20);not null" json:"deployId"`                     // 部署ID
	IsMain            bool    `gorm:"column:is_main;type:tinyint(1);not null" json:"isMain"`                         // 是否是主容器
	Image             string  `gorm:"column:image;type:varchar(150);not null" json:"image"`                          // 镜像
	ImageVersion      string  `gorm:"column:image_version;type:varchar(20);not null" json:"imageVersion"`            // 镜像版本
	ImagePullStrategy string  `gorm:"column:image_pull_strategy;type:varchar(20);not null" json:"imagePullStrategy"` // 镜像拉取策略(Always、IfNotPresent、Never)
	RequestCPU        float64 `gorm:"column:request_cpu;type:decimal(4,2) unsigned;not null" json:"requestCpu"`      // CPU限制Core - request
	RequestMemory     float64 `gorm:"column:request_memory;type:decimal(5,0);not null" json:"requestMemory"`         // 内存限制MiB - request
	LimitCPU          float64 `gorm:"column:limit_cpu;type:decimal(4,2) unsigned;not null" json:"limitCpu"`          // CPU限制Core - limit
	LimitMemory       float64 `gorm:"column:limit_memory;type:decimal(5,0) unsigned;not null" json:"limitMemory"`    // 内存限制MiB
	Environments      string  `gorm:"column:environments;type:varchar(200)" json:"environments"`                     // 环境变量JSON [{ key: value }]
	Workdir           string  `gorm:"column:workdir;type:varchar(200)" json:"workdir"`                               // 工作目录
	RunCmd            string  `gorm:"column:run_cmd;type:varchar(200)" json:"runCmd"`                                // 运行命令
	RunParams         string  `gorm:"column:run_params;type:varchar(50)" json:"runParams"`                           // 运行参数
	Podstop           string  `gorm:"column:podstop;type:varchar(100)" json:"podstop"`                               // 容器结束前执行
	Liveness          string  `gorm:"column:liveness;type:varchar(300)" json:"liveness"`                             // 存活检查
	Readness          string  `gorm:"column:readness;type:varchar(300)" json:"readness"`                             // 就绪检查
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantDeploymentsContainers) TableName() string {
	return "sgr_tenant_deployments_containers"
}

// SgrTenantDeploymentsContainersColumns get sql column name.获取数据库列名
var SgrTenantDeploymentsContainersColumns = struct {
	ID                string
	Name              string
	DeployID          string
	IsMain            string
	Image             string
	ImageVersion      string
	ImagePullStrategy string
	RequestCPU        string
	RequestMemory     string
	LimitCPU          string
	LimitMemory       string
	Environments      string
	Workdir           string
	RunCmd            string
	RunParams         string
	Podstop           string
	Liveness          string
	Readness          string
}{
	ID:                "id",
	Name:              "name",
	DeployID:          "deploy_id",
	IsMain:            "is_main",
	Image:             "image",
	ImageVersion:      "image_version",
	ImagePullStrategy: "image_pull_strategy",
	RequestCPU:        "request_cpu",
	RequestMemory:     "request_memory",
	LimitCPU:          "limit_cpu",
	LimitMemory:       "limit_memory",
	Environments:      "environments",
	Workdir:           "workdir",
	RunCmd:            "run_cmd",
	RunParams:         "run_params",
	Podstop:           "podstop",
	Liveness:          "liveness",
	Readness:          "readness",
}

// SgrTenantNamespace 集群_命名空间
type SgrTenantNamespace struct {
	ID         uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	TenantID   *uint64    `gorm:"column:tenant_id;type:bigint(20) unsigned" json:"tenantId"`   // 租户ID
	ClusterID  *uint64    `gorm:"column:cluster_id;type:bigint(20) unsigned" json:"clusterId"` // 集群ID
	Namespace  string     `gorm:"column:namespace;type:varchar(50);not null" json:"namespace"` // 命名空间名称
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;not null" json:"createTime"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;not null" json:"updateTime"` // 更新时间
	Status     int8       `gorm:"column:status;type:tinyint(4);not null" json:"status"`        // 状态
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantNamespace) TableName() string {
	return "sgr_tenant_namespace"
}

// SgrTenantNamespaceColumns get sql column name.获取数据库列名
var SgrTenantNamespaceColumns = struct {
	ID         string
	TenantID   string
	ClusterID  string
	Namespace  string
	CreateTime string
	UpdateTime string
	Status     string
}{
	ID:         "id",
	TenantID:   "tenant_id",
	ClusterID:  "cluster_id",
	Namespace:  "namespace",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Status:     "status",
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
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint(11) unsigned;not null" json:"tenantId"` // 租户
	UserName     string     `gorm:"column:user_name;type:varchar(50)" json:"userName"`                  // 用户名
	Account      string     `gorm:"column:account;type:varchar(50);not null" json:"account"`            // 账号
	Password     string     `gorm:"column:password;type:varchar(255);not null" json:"password"`         // 密码
	Mobile       string     `gorm:"column:mobile;type:varchar(20)" json:"mobile"`                       // 手机
	Email        string     `gorm:"column:email;type:varchar(50)" json:"email"`                         // 邮箱
	Status       int8       `gorm:"column:status;type:tinyint(3);not null;default:0" json:"status"`     // 状态
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
	UserID       uint64     `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"userId"` // 用户id
	RoleID       int64      `gorm:"column:role_id;type:bigint(20);not null" json:"roleId"`          // 角色id
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
