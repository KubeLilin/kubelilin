package models

import "time"

// CodeServiceConnection 服务连接类型
type CodeServiceConnection struct {
	ID   uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Code string `gorm:"column:code;type:varchar(20);not null" json:"code"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

// TableName get sql table name.获取数据库表名
func (m *CodeServiceConnection) TableName() string {
	return "code_service_connection"
}

// CodeServiceConnectionColumns get sql column name.获取数据库列名
var CodeServiceConnectionColumns = struct {
	ID   string
	Code string
	Name string
}{
	ID:   "id",
	Code: "code",
	Name: "name",
}

// DevopsProjects devops 项目管理
type DevopsProjects struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned zerofill;not null" json:"id"`
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                                         // 项目名称
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`                            // 租户ID
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"creationTime"` // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *DevopsProjects) TableName() string {
	return "devops_projects"
}

// DevopsProjectsColumns get sql column name.获取数据库列名
var DevopsProjectsColumns = struct {
	ID           string
	Name         string
	TenantID     string
	CreationTime string
}{
	ID:           "id",
	Name:         "name",
	TenantID:     "tenant_id",
	CreationTime: "creation_time",
}

// DevopsProjectsApps devops 项目应用对应表
type DevopsProjectsApps struct {
	ID            uint64 `gorm:"primaryKey;column:id;type:bigint(20) unsigned zerofill;not null" json:"id"`
	ProjectID     uint64 `gorm:"column:project_id;type:bigint unsigned;not null" json:"projectId"`         // 项目ID
	ApplicationID uint64 `gorm:"column:application_id;type:bigint unsigned;not null" json:"applicationId"` // 应用 ID
}

// TableName get sql table name.获取数据库表名
func (m *DevopsProjectsApps) TableName() string {
	return "devops_projects_apps"
}

// DevopsProjectsAppsColumns get sql column name.获取数据库列名
var DevopsProjectsAppsColumns = struct {
	ID            string
	ProjectID     string
	ApplicationID string
}{
	ID:            "id",
	ProjectID:     "project_id",
	ApplicationID: "application_id",
}

// ServiceConnection 用于保存其他服务或者第三方组件所依赖的资源，例如连接字符串，ssh秘钥，git连接等等
type ServiceConnection struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"` // 租户id
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`              // 连接名称
	ServiceURL   string     `gorm:"column:service_url;type:varchar(255)" json:"serviceUrl"`         // 服务连接地址
	ServiceType  int        `gorm:"column:service_type;type:int;not null" json:"serviceType"`       // 连接类型: 1: vcs 2. hub 3. pipline 4.url
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"updateTime"`
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP" json:"creationTime"`
}

// TableName get sql table name.获取数据库表名
func (m *ServiceConnection) TableName() string {
	return "service_connection"
}

// ServiceConnectionColumns get sql column name.获取数据库列名
var ServiceConnectionColumns = struct {
	ID           string
	TenantID     string
	Name         string
	ServiceURL   string
	ServiceType  string
	UpdateTime   string
	CreationTime string
}{
	ID:           "id",
	TenantID:     "tenant_id",
	Name:         "name",
	ServiceURL:   "service_url",
	ServiceType:  "service_type",
	UpdateTime:   "update_time",
	CreationTime: "creation_time",
}

// ServiceConnectionCredentials 常用的连接凭证，例如token
type ServiceConnectionCredentials struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Name         string     `gorm:"column:name;type:varchar(20);not null" json:"name"`         // 凭据名称
	Type         int        `gorm:"column:type;type:int;not null" json:"type"`                 // 凭证类型 1. 用户密码 2.token
	Username     string     `gorm:"column:username;type:varchar(50);not null" json:"username"` // 凭证用户名
	Password     string     `gorm:"column:password;type:varchar(50);not null" json:"password"` // 凭证密码
	Token        string     `gorm:"column:token;type:varchar(255);not null" json:"token"`      // 凭证TOKEN
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *ServiceConnectionCredentials) TableName() string {
	return "service_connection_credentials"
}

// ServiceConnectionCredentialsColumns get sql column name.获取数据库列名
var ServiceConnectionCredentialsColumns = struct {
	ID           string
	Name         string
	Type         string
	Username     string
	Password     string
	Token        string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	Name:         "name",
	Type:         "type",
	Username:     "username",
	Password:     "password",
	Token:        "token",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// ServiceConnectionDetails 连接的详细信息，例如mysql连接字符串
type ServiceConnectionDetails struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	MainID       uint64     `gorm:"column:main_id;type:bigint unsigned;not null" json:"mainId"` // 主数据id
	Type         int        `gorm:"column:type;type:int;not null" json:"type"`                  // 连接类型
	Detail       string     `gorm:"column:detail;type:varchar(500)" json:"detail"`
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *ServiceConnectionDetails) TableName() string {
	return "service_connection_details"
}

// ServiceConnectionDetailsColumns get sql column name.获取数据库列名
var ServiceConnectionDetailsColumns = struct {
	ID           string
	MainID       string
	Type         string
	Detail       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	MainID:       "main_id",
	Type:         "type",
	Detail:       "detail",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrCodeApplicationLanguage 字典-应用开发语言
type SgrCodeApplicationLanguage struct {
	ID   uint16 `gorm:"primaryKey;column:id;type:smallint unsigned;not null" json:"id"`
	Code string `gorm:"column:code;type:varchar(8)" json:"code"`
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Sort uint16 `gorm:"column:sort;type:smallint unsigned;not null;default:0" json:"sort"`
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
	ID   uint16 `gorm:"primaryKey;column:id;type:smallint unsigned;not null" json:"id"`
	Code string `gorm:"column:code;type:varchar(8)" json:"code"`
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Sort uint16 `gorm:"column:sort;type:smallint unsigned;not null;default:0" json:"sort"`
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

// SgrCodeDeploymentLevel 部署环境
type SgrCodeDeploymentLevel struct {
	ID   uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Code string `gorm:"column:code;type:varchar(8);not null" json:"code"`
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Sort int16  `gorm:"column:sort;type:smallint;not null" json:"sort"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrCodeDeploymentLevel) TableName() string {
	return "sgr_code_deployment_level"
}

// SgrCodeDeploymentLevelColumns get sql column name.获取数据库列名
var SgrCodeDeploymentLevelColumns = struct {
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	RoleID       uint64     `gorm:"column:role_id;type:bigint unsigned;not null" json:"roleId"` // 角色ID
	MenuID       uint64     `gorm:"column:menu_id;type:bigint unsigned;not null" json:"menuId"` // 菜单ID
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID     int64      `gorm:"column:tenant_id;type:bigint;not null" json:"tenantId"`                    // 租户
	MenuCode     string     `gorm:"column:menu_code;type:varchar(100);not null" json:"menuCode"`              // 编码
	MenuName     string     `gorm:"column:menu_name;type:varchar(50);not null" json:"menuName"`               // 目录名称
	Icon         string     `gorm:"column:icon;type:varchar(50);not null" json:"icon"`                        // 图标
	Path         string     `gorm:"column:path;type:varchar(100);not null" json:"path"`                       // 路由路径
	Component    string     `gorm:"column:component;type:varchar(100);not null" json:"component"`             // react组件路径
	IsRoot       int8       `gorm:"column:is_root;type:tinyint;not null;default:0" json:"isRoot"`             // 是否是根目录
	ParentID     uint64     `gorm:"column:parent_id;type:bigint unsigned;not null;default:0" json:"parentId"` // 父层级id
	Sort         int        `gorm:"column:sort;type:int;not null;default:0" json:"sort"`                      // 权重，正序排序
	Status       int8       `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`              // 状态
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TName        string     `gorm:"column:t_name;type:varchar(50);not null" json:"tName"`        // 租户名称
	TCode        string     `gorm:"unique;column:t_code;type:varchar(16);not null" json:"tCode"` // 租户编码
	Status       int8       `gorm:"column:status;type:tinyint;not null;default:0" json:"status"` // 状态
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime" json:"creationTime"`      // 创建时间
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`          // 修改时间
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
	ID         uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID   uint64     `gorm:"column:tenant_Id;type:bigint unsigned;not null" json:"tenantId"`        // 租户ID
	Name       string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                     // 集群应用名称(英文唯一)
	Nickname   string     `gorm:"column:nickname;type:varchar(50);not null;default:''" json:"nickname"`  // 应用中文名称
	Remarks    string     `gorm:"column:remarks;type:varchar(200);not null;default:''" json:"remarks"`   // 集群应用备注
	Git        string     `gorm:"column:git;type:varchar(500);not null" json:"git"`                      // 集群应用绑定的git地址
	Imagehub   string     `gorm:"column:imagehub;type:varchar(500);not null;default:''" json:"imagehub"` // 集群应用绑定镜像仓库地址
	Level      uint16     `gorm:"column:level;type:smallint unsigned;not null" json:"level"`             // 应用级别
	Language   uint16     `gorm:"column:language;type:smallint unsigned;not null" json:"language"`       // 开发语言
	Status     int8       `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`           // 状态
	CreateTime *time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`                    // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`                    // 更新时间
	Labels     string     `gorm:"column:labels;type:varchar(100);not null;default:''" json:"labels"`     // 应用标签
	GitType    string     `gorm:"column:git_type;type:varchar(20);not null" json:"gitType"`              // git类型 github/ gitee/ gogs/gitlab
	ScID       *uint64    `gorm:"column:sc_id;type:bigint unsigned;default:0" json:"scId"`               // 服务连接git类型的凭据ID
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
	Nickname   string
	Remarks    string
	Git        string
	Imagehub   string
	Level      string
	Language   string
	Status     string
	CreateTime string
	UpdateTime string
	Labels     string
	GitType    string
	ScID       string
}{
	ID:         "id",
	TenantID:   "tenant_Id",
	Name:       "name",
	Nickname:   "nickname",
	Remarks:    "remarks",
	Git:        "git",
	Imagehub:   "imagehub",
	Level:      "level",
	Language:   "language",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Labels:     "labels",
	GitType:    "git_type",
	ScID:       "sc_id",
}

// SgrTenantApplicationPipelines 应用流水线
type SgrTenantApplicationPipelines struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`             // Pipeline ID
	Appid        uint64     `gorm:"column:appid;type:bigint unsigned;not null" json:"appid"`                  // 应用ID
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                        // 流水线名称, appid 下唯一
	Dsl          string     `gorm:"column:dsl;type:text;not null" json:"dsl"`                                 // 流水线DSL
	TaskStatus   *uint      `gorm:"column:taskStatus;type:int unsigned" json:"taskStatus"`                    // 流水线任务状态( ready=0 , running=1, success=2, fail=3,  )
	LastTaskID   string     `gorm:"column:lastTaskId;type:varchar(15);not null;default:''" json:"lastTaskId"` // 最后一次任务执行ID
	Status       uint8      `gorm:"column:status;type:tinyint unsigned;not null" json:"status"`
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantApplicationPipelines) TableName() string {
	return "sgr_tenant_application_pipelines"
}

// SgrTenantApplicationPipelinesColumns get sql column name.获取数据库列名
var SgrTenantApplicationPipelinesColumns = struct {
	ID           string
	Appid        string
	Name         string
	Dsl          string
	TaskStatus   string
	LastTaskID   string
	Status       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	Appid:        "appid",
	Name:         "name",
	Dsl:          "dsl",
	TaskStatus:   "taskStatus",
	LastTaskID:   "lastTaskId",
	Status:       "status",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// SgrTenantCluster 集群信息
type SgrTenantCluster struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                 // ID
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`               // 租户ID
	Nickname     string     `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`                    // 别名
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                            // 集群名称
	Version      string     `gorm:"column:version;type:varchar(50)" json:"version"`                               // k8s 版本号
	Distribution string     `gorm:"column:distribution;type:varchar(30)" json:"distribution"`                     // 来源
	Config       string     `gorm:"column:config;type:text;not null" json:"config"`                               // k8s config text
	Sort         *int       `gorm:"column:sort;type:int" json:"sort"`                                             // 排序
	Status       int8       `gorm:"column:status;type:tinyint;not null" json:"status"`                            // 状态
	CreateTime   *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"updateTime"` // 更新时间
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

// SgrTenantDeploymentRecord 部署发布记录
type SgrTenantDeploymentRecord struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	AppID        uint64     `gorm:"column:app_id;type:bigint unsigned;not null" json:"appId"`
	DeploymentID uint64     `gorm:"column:deployment_id;type:bigint unsigned;not null" json:"deploymentId"`
	ApplyImage   string     `gorm:"column:apply_image;type:varchar(255);not null" json:"applyImage"`
	OpsType      string     `gorm:"column:ops_type;type:char(20);not null" json:"opsType"`
	Operator     *uint64    `gorm:"column:operator;type:bigint unsigned" json:"operator"`
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"creationTime"`
	State        string     `gorm:"column:state;type:varchar(20)" json:"state"`
	Remark       string     `gorm:"column:remark;type:varchar(500)" json:"remark"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantDeploymentRecord) TableName() string {
	return "sgr_tenant_deployment_record"
}

// SgrTenantDeploymentRecordColumns get sql column name.获取数据库列名
var SgrTenantDeploymentRecordColumns = struct {
	ID           string
	AppID        string
	DeploymentID string
	ApplyImage   string
	OpsType      string
	Operator     string
	CreationTime string
	State        string
	Remark       string
	UpdateTime   string
}{
	ID:           "id",
	AppID:        "app_id",
	DeploymentID: "deployment_id",
	ApplyImage:   "apply_image",
	OpsType:      "ops_type",
	Operator:     "operator",
	CreationTime: "creation_time",
	State:        "state",
	Remark:       "remark",
	UpdateTime:   "update_time",
}

// SgrTenantDeployments 集群部署
type SgrTenantDeployments struct {
	ID              uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID        uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`
	Name            string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                    // 部署名称(英文唯一)
	Nickname        string     `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`            // 部署中文名称
	ClusterID       uint64     `gorm:"column:cluster_id;type:bigint unsigned;not null" json:"clusterId"`     // 集群ID
	NamespaceID     uint64     `gorm:"column:namespace_id;type:bigint unsigned;not null" json:"namespaceId"` // 命名空间ID
	AppID           uint64     `gorm:"column:app_id;type:bigint unsigned;not null" json:"appId"`             // 应用ID
	Status          uint8      `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"` // 状态
	CreateTime      *time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`                   // 创建时间
	UpdateTime      *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`                   // 更新时间
	ImageHub        string     `gorm:"column:image_hub;type:varchar(200);not null;default:''" json:"imageHub"`
	AppName         string     `gorm:"column:app_name;type:varchar(50)" json:"appName"`
	WorkloadType    string     `gorm:"column:workload_type;type:varchar(25);not null;default:''" json:"workloadType"`
	Replicas        uint       `gorm:"column:replicas;type:int unsigned;not null;default:1" json:"replicas"`
	ServiceEnable   *uint8     `gorm:"column:service_enable;type:tinyint unsigned" json:"serviceEnable"`
	ServiceName     string     `gorm:"column:service_name;type:varchar(150);not null;default:''" json:"serviceName"`
	ServiceAway     string     `gorm:"column:service_away;type:varchar(30)" json:"serviceAway"`
	ServicePort     uint       `gorm:"column:service_port;type:int unsigned;not null;default:0" json:"servicePort"`
	ServicePortType string     `gorm:"column:service_port_type;type:varchar(8);not null;default:''" json:"servicePortType"`
	LastImage       string     `gorm:"column:last_image;type:varchar(350)" json:"lastImage"`
	Level           string     `gorm:"index:levev_idx;column:level;type:varchar(8)" json:"level"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantDeployments) TableName() string {
	return "sgr_tenant_deployments"
}

// SgrTenantDeploymentsColumns get sql column name.获取数据库列名
var SgrTenantDeploymentsColumns = struct {
	ID              string
	TenantID        string
	Name            string
	Nickname        string
	ClusterID       string
	NamespaceID     string
	AppID           string
	Status          string
	CreateTime      string
	UpdateTime      string
	ImageHub        string
	AppName         string
	WorkloadType    string
	Replicas        string
	ServiceEnable   string
	ServiceName     string
	ServiceAway     string
	ServicePort     string
	ServicePortType string
	LastImage       string
	Level           string
}{
	ID:              "id",
	TenantID:        "tenant_id",
	Name:            "name",
	Nickname:        "nickname",
	ClusterID:       "cluster_id",
	NamespaceID:     "namespace_id",
	AppID:           "app_id",
	Status:          "status",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	ImageHub:        "image_hub",
	AppName:         "app_name",
	WorkloadType:    "workload_type",
	Replicas:        "replicas",
	ServiceEnable:   "service_enable",
	ServiceName:     "service_name",
	ServiceAway:     "service_away",
	ServicePort:     "service_port",
	ServicePortType: "service_port_type",
	LastImage:       "last_image",
	Level:           "level",
}

// SgrTenantDeploymentsContainers 应用部署容器配置
type SgrTenantDeploymentsContainers struct {
	ID                uint64  `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Name              string  `gorm:"column:name;type:varchar(30);not null" json:"name"`
	DeployID          uint64  `gorm:"column:deploy_id;type:bigint unsigned;not null" json:"deployId"`
	IsMain            uint8   `gorm:"column:is_main;type:tinyint unsigned;not null" json:"isMain"`
	Image             string  `gorm:"column:image;type:varchar(255);not null" json:"image"`
	ImageVersion      string  `gorm:"column:image_version;type:varchar(40);not null" json:"imageVersion"`
	ImagePullStrategy string  `gorm:"column:image_pull_strategy;type:varchar(20);not null" json:"imagePullStrategy"`
	RequestCPU        float64 `gorm:"column:request_cpu;type:decimal(4,2);not null" json:"requestCpu"`
	RequestMemory     float64 `gorm:"column:request_memory;type:decimal(5,0);not null" json:"requestMemory"`
	LimitCPU          float64 `gorm:"column:limit_cpu;type:decimal(4,2);not null" json:"limitCpu"`
	LimitMemory       float64 `gorm:"column:limit_memory;type:decimal(5,0);not null" json:"limitMemory"`
	Environments      string  `gorm:"column:environments;type:varchar(255);not null;default:''" json:"environments"`
	Workdir           string  `gorm:"column:workdir;type:varchar(200)" json:"workdir"`
	RunCmd            string  `gorm:"column:run_cmd;type:varchar(200)" json:"runCmd"`
	RunParams         string  `gorm:"column:run_params;type:varchar(100)" json:"runParams"`
	Podstop           string  `gorm:"column:podstop;type:varchar(100)" json:"podstop"`
	Liveness          string  `gorm:"column:liveness;type:varchar(300)" json:"liveness"`
	Readness          string  `gorm:"column:readness;type:varchar(300)" json:"readness"`
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
	ID         uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID   uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`   // 租户ID
	ClusterID  uint64     `gorm:"column:cluster_id;type:bigint unsigned;not null" json:"clusterId"` // 集群ID
	Namespace  string     `gorm:"column:namespace;type:varchar(50);not null" json:"namespace"`      // 命名空间名称
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;not null" json:"createTime"`      // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;not null" json:"updateTime"`      // 更新时间
	Status     int8       `gorm:"column:status;type:tinyint;not null" json:"status"`                // 状态
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	RoleCode     string     `gorm:"uniqueIndex:un_role_code_name;column:role_code;type:varchar(30);not null" json:"roleCode"` // 角色编码
	RoleName     string     `gorm:"uniqueIndex:un_role_code_name;column:role_name;type:varchar(50);not null" json:"roleName"` // 角色名称
	Description  string     `gorm:"column:description;type:varchar(50)" json:"description"`                                   // 角色描述
	Status       int8       `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`                              // 状态
	TenantID     int64      `gorm:"column:tenant_id;type:bigint;not null" json:"tenantId"`                                    // 租户
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"` // 租户
	UserName     string     `gorm:"column:user_name;type:varchar(50);not null" json:"userName"`     // 用户名
	Account      string     `gorm:"column:account;type:varchar(50);not null" json:"account"`        // 账号
	Password     string     `gorm:"column:password;type:varchar(255);not null" json:"password"`     // 密码
	Mobile       string     `gorm:"column:mobile;type:varchar(20);not null" json:"mobile"`          // 手机
	Email        string     `gorm:"column:email;type:varchar(50);not null" json:"email"`            // 邮箱
	Status       int8       `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`    // 状态
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
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	UserID       uint64     `gorm:"column:user_id;type:bigint unsigned;not null" json:"userId"` // 用户id
	RoleID       int64      `gorm:"column:role_id;type:bigint;not null" json:"roleId"`          // 角色id
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
