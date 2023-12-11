package models

// ApplicationAPIGateway 集群网关(APISIX)
type ApplicationAPIGateway struct {
	ID          uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`         // 网关ID
	Name        string `gorm:"column:name;type:varchar(100);not null" json:"name"`                   // 网关名称
	Desc        string `gorm:"column:desc;type:varchar(255);not null" json:"desc"`                   // 网关描述
	Vip         string `gorm:"column:vip;type:varchar(50);not null;default:''" json:"vip"`           // 内网vip
	ExportIP    string `gorm:"column:export_ip;type:varchar(50);default:''" json:"exportIp"`         // 出口IP
	ClusterID   uint64 `gorm:"column:cluster_id;type:bigint unsigned;not null" json:"clusterId"`     // 集群ID
	DefaultHost string `gorm:"column:default_host;type:varchar(255);not null" json:"defaultHost"`    // 默认域名
	AdminURI    string `gorm:"column:admin_uri;type:varchar(255);not null" json:"adminUri"`          // 网关admin api
	AccessToken string `gorm:"column:access_token;type:varchar(255);not null" json:"accessToken"`    // 网关 admin api访问token
	Status      uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"` // 网关状态
}

// TableName get sql table name.获取数据库表名
func (m *ApplicationAPIGateway) TableName() string {
	return "application_api_gateway"
}

// ApplicationAPIGatewayColumns get sql column name.获取数据库列名
var ApplicationAPIGatewayColumns = struct {
	ID          string
	Name        string
	Desc        string
	Vip         string
	ExportIP    string
	ClusterID   string
	DefaultHost string
	AdminURI    string
	AccessToken string
	Status      string
}{
	ID:          "id",
	Name:        "name",
	Desc:        "desc",
	Vip:         "vip",
	ExportIP:    "export_ip",
	ClusterID:   "cluster_id",
	DefaultHost: "default_host",
	AdminURI:    "admin_uri",
	AccessToken: "access_token",
	Status:      "status",
}

// ApplicationAPIGatewayRouters 集群网关团队路由
type ApplicationAPIGatewayRouters struct {
	ID            uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                        // 路由ID
	Name          string `gorm:"index:name;column:name;type:varchar(50);not null;default:''" json:"name"`             // 路由名称
	Desc          string `gorm:"column:desc;type:varchar(255);not null;default:''" json:"desc"`                       // 路由描述
	TeamID        uint64 `gorm:"column:team_id;type:bigint unsigned;not null;default:0" json:"teamId"`                // 团队目录ID
	Host          string `gorm:"column:host;type:varchar(150);not null;default:''" json:"host"`                       // 路由域名
	URI           string `gorm:"column:uri;type:varchar(255);not null;default:/*" json:"uri"`                         // 路由路径
	Websocket     uint8  `gorm:"column:websocket;type:tinyint unsigned;not null;default:1" json:"websocket"`          // 是否开启websocket
	UpstreamType  string `gorm:"column:upstream_type;type:varchar(20);not null;default:service" json:"upstreamType"`  // service | node
	Loadbalance   string `gorm:"column:loadbalance;type:varchar(255);not null;default:roundrobin" json:"loadbalance"` // roundrobin | chash | ewma | least_conn
	Nodes         string `gorm:"column:nodes;type:varchar(255);not null;default:''" json:"nodes"`                     // {,   "127.0.0.1:1980": 1,,   "127.0.0.1:1981": 1,}
	Timeout       uint   `gorm:"column:timeout;type:int unsigned;not null;default:60" json:"timeout"`                 // 超时时间
	ApplicationID uint64 `gorm:"column:application_id;type:bigint unsigned;not null;default:0" json:"applicationId"`  // 应用ID
	DeploymentID  uint64 `gorm:"column:deployment_id;type:bigint unsigned;not null;default:0" json:"deploymentId"`    // 应用部署ID
	Rewrite       uint8  `gorm:"column:rewrite;type:tinyint unsigned;not null;default:0" json:"rewrite"`              // 是否重写
	RegexURI      string `gorm:"column:regex_uri;type:varchar(150);not null;default:''" json:"regexUri"`              // 匹配正则表达式,
	RegexTmp      string `gorm:"column:regex_tmp;type:varchar(100);not null;default:''" json:"regexTmp"`              // 转发路径模版,
	Liveness      string `gorm:"column:liveness;type:varchar(255);not null;default:''" json:"liveness"`               // 存活探针
	Label         string `gorm:"column:label;type:varchar(50);not null;default:''" json:"label"`                      // 标签
	Status        uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`                // 状态
}

// TableName get sql table name.获取数据库表名
func (m *ApplicationAPIGatewayRouters) TableName() string {
	return "application_api_gateway_routers"
}

// ApplicationAPIGatewayRoutersColumns get sql column name.获取数据库列名
var ApplicationAPIGatewayRoutersColumns = struct {
	ID            string
	Name          string
	Desc          string
	TeamID        string
	Host          string
	URI           string
	Websocket     string
	UpstreamType  string
	Loadbalance   string
	Nodes         string
	Timeout       string
	ApplicationID string
	DeploymentID  string
	Rewrite       string
	RegexURI      string
	RegexTmp      string
	Liveness      string
	Label         string
	Status        string
}{
	ID:            "id",
	Name:          "name",
	Desc:          "desc",
	TeamID:        "team_id",
	Host:          "host",
	URI:           "uri",
	Websocket:     "websocket",
	UpstreamType:  "upstream_type",
	Loadbalance:   "loadbalance",
	Nodes:         "nodes",
	Timeout:       "timeout",
	ApplicationID: "application_id",
	DeploymentID:  "deployment_id",
	Rewrite:       "rewrite",
	RegexURI:      "regex_uri",
	RegexTmp:      "regex_tmp",
	Liveness:      "liveness",
	Label:         "label",
	Status:        "status",
}

// ApplicationAPIGatewayTeams 集群网关团队目录
type ApplicationAPIGatewayTeams struct {
	ID        uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`     // 网关团队目录ID
	Name      string `gorm:"column:name;type:varchar(100);not null" json:"name"`               // 网关团队目录名称
	GatewayID uint64 `gorm:"column:gateway_id;type:bigint unsigned;not null" json:"gatewayId"` // 网关ID
	TenantID  uint64 `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`   // 租户ID
	Level     string `gorm:"column:level;type:varchar(50);not null;default:P1" json:"level"`   // 服务级别 P0 - P4
	Status    uint8  `gorm:"column:status;type:tinyint unsigned;not null" json:"status"`       // 状态
}

// TableName get sql table name.获取数据库表名
func (m *ApplicationAPIGatewayTeams) TableName() string {
	return "application_api_gateway_teams"
}

// ApplicationAPIGatewayTeamsColumns get sql column name.获取数据库列名
var ApplicationAPIGatewayTeamsColumns = struct {
	ID        string
	Name      string
	GatewayID string
	TenantID  string
	Level     string
	Status    string
}{
	ID:        "id",
	Name:      "name",
	GatewayID: "gateway_id",
	TenantID:  "tenant_id",
	Level:     "level",
	Status:    "status",
}

// ApplicationDaprCoponentsTemplete Dapr 运行时组件模板
type ApplicationDaprCoponentsTemplete struct {
	ID            uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                          // 组件模板ID
	Name          string     `gorm:"column:name;type:varchar(30);not null" json:"name"`                                     // 组件模板名称
	ComponentType string     `gorm:"column:component_type;type:varchar(20);not null" json:"componentType"`                  // 组件模板类型
	Doc           string     `gorm:"column:doc;type:varchar(150);not null;default:''" json:"doc"`                           // 组件文档
	Template      string     `gorm:"column:template;type:text;not null" json:"template"`                                    // 组件模板yaml
	CreateTime    *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	UpdateTime    *time.Time `gorm:"column:update_time;type:datetime;not null" json:"updateTime"`                           // 更新时间,
}

// TableName get sql table name.获取数据库表名
func (m *ApplicationDaprCoponentsTemplete) TableName() string {
	return "application_dapr_coponents_templete"
}

// ApplicationDaprCoponentsTempleteColumns get sql column name.获取数据库列名
var ApplicationDaprCoponentsTempleteColumns = struct {
	ID            string
	Name          string
	ComponentType string
	Doc           string
	Template      string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	Name:          "name",
	ComponentType: "component_type",
	Doc:           "doc",
	Template:      "template",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// ApplicationLanguageCompile CI流水线编译环境容器镜像
type ApplicationLanguageCompile struct {
	ID           uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`         // 编译环境ID
	LanguageID   uint64 `gorm:"column:language_id;type:bigint unsigned;not null" json:"languageId"`   // 语言ID
	CompileImage string `gorm:"column:compile_image;type:varchar(120);not null" json:"compileImage"`  // 编译镜像
	AliasName    string `gorm:"column:alias_name;type:varchar(100);not null" json:"aliasName"`        // 别名
	Sort         uint   `gorm:"column:sort;type:int unsigned;not null" json:"sort"`                   // 排序
	Status       uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"` // 状态
}

// TableName get sql table name.获取数据库表名
func (m *ApplicationLanguageCompile) TableName() string {
	return "application_language_compile"
}

// ApplicationLanguageCompileColumns get sql column name.获取数据库列名
var ApplicationLanguageCompileColumns = struct {
	ID           string
	LanguageID   string
	CompileImage string
	AliasName    string
	Sort         string
	Status       string
}{
	ID:           "id",
	LanguageID:   "language_id",
	CompileImage: "compile_image",
	AliasName:    "alias_name",
	Sort:         "sort",
	Status:       "status",
}

// ApplicationServiceMonitor [...]
type ApplicationServiceMonitor struct {
	ID             uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Name           string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                       // ServiceMonitor名称
	AppID          uint64     `gorm:"column:app_id;type:bigint unsigned;not null" json:"appId"`                // 应用ID
	ClusterID      uint64     `gorm:"column:cluster_id;type:bigint unsigned;not null" json:"clusterId"`        // 集群ID
	Namespace      string     `gorm:"column:namespace;type:varchar(50);not null" json:"namespace"`             // 目标命名空间
	DeploymentID   uint64     `gorm:"column:deployment_id;type:bigint unsigned;not null" json:"deploymentId"`  // 部署ID
	DeploymentName string     `gorm:"column:deployment_name;type:varchar(100);not null" json:"deploymentName"` // 部署名称
	Interval       uint       `gorm:"column:interval;type:int unsigned;not null" json:"interval"`              // 采集间隔时间
	Port           string     `gorm:"column:port;type:varchar(50);not null" json:"port"`                       // 采集服务端口名称,
	Path           string     `gorm:"column:path;type:varchar(200);not null" json:"path"`                      // 采集指标端点
	CreateTime     *time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`                      // 创建时间
	UpdateTime     *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`                      // 更新时间
	Status         uint8      `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`    // 状态
}

// TableName get sql table name.获取数据库表名
func (m *ApplicationServiceMonitor) TableName() string {
	return "application_service_monitor"
}

// ApplicationServiceMonitorColumns get sql column name.获取数据库列名
var ApplicationServiceMonitorColumns = struct {
	ID             string
	Name           string
	AppID          string
	ClusterID      string
	Namespace      string
	DeploymentID   string
	DeploymentName string
	Interval       string
	Port           string
	Path           string
	CreateTime     string
	UpdateTime     string
	Status         string
}{
	ID:             "id",
	Name:           "name",
	AppID:          "app_id",
	ClusterID:      "cluster_id",
	Namespace:      "namespace",
	DeploymentID:   "deployment_id",
	DeploymentName: "deployment_name",
	Interval:       "interval",
	Port:           "port",
	Path:           "path",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	Status:         "status",
}

// CodeApplicationRuntime 运行时字典
type CodeApplicationRuntime struct {
	ID   uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Name string `gorm:"column:name;type:varchar(20)" json:"name"`
	Desc string `gorm:"column:desc;type:varchar(100)" json:"desc"`
}

// TableName get sql table name.获取数据库表名
func (m *CodeApplicationRuntime) TableName() string {
	return "code_application_runtime"
}

// CodeApplicationRuntimeColumns get sql column name.获取数据库列名
var CodeApplicationRuntimeColumns = struct {
	ID   string
	Name string
	Desc string
}{
	ID:   "id",
	Name: "name",
	Desc: "desc",
}

// DeploymentContainerLifecycleCheck 容器生命周期健康检查
type DeploymentContainerLifecycleCheck struct {
	ID                  uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                                 // 生命周期
	DeploymentID        uint64     `gorm:"column:deployment_id;type:bigint unsigned;not null" json:"deploymentId"`                       // 部署id
	ContainerID         uint64     `gorm:"column:container_id;type:bigint unsigned;not null" json:"containerId"`                         // 容器id
	Type                string     `gorm:"column:type;type:varchar(50);not null" json:"type"`                                            // 检查类型readinessProbe/livenessProbe
	Scheme              string     `gorm:"column:scheme;type:varchar(20);not null" json:"scheme"`                                        // HTTP/TCP
	Path                string     `gorm:"column:path;type:varchar(500);not null" json:"path"`                                           // 请求路径
	Port                uint       `gorm:"column:port;type:int unsigned;not null" json:"port"`                                           // 检查端口
	SuccessThreshold    uint       `gorm:"column:success_threshold;type:int unsigned;not null;default:1" json:"successThreshold"`        // 成功阈值(次数)
	FailureThreshold    uint       `gorm:"column:failure_threshold;type:int unsigned;not null;default:3" json:"failureThreshold"`        // 错误阀值(次数)
	InitialDelaySeconds uint       `gorm:"column:initial_delay_seconds;type:int unsigned;not null;default:4" json:"initialDelaySeconds"` // 启动延时(秒)
	PeriodSeconds       uint       `gorm:"column:period_seconds;type:int unsigned;not null;default:10" json:"periodSeconds"`             // 间隔时间(秒)
	TimeoutSeconds      uint       `gorm:"column:timeout_seconds;type:int unsigned;not null;default:3" json:"timeoutSeconds"`            // 响应超时(秒)
	Enable              uint8      `gorm:"column:enable;type:tinyint unsigned;not null;default:0" json:"enable"`                         // 是否启用 1是0否
	CreationTime        *time.Time `gorm:"column:creation_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime          *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *DeploymentContainerLifecycleCheck) TableName() string {
	return "deployment_container_lifecycle_check"
}

// DeploymentContainerLifecycleCheckColumns get sql column name.获取数据库列名
var DeploymentContainerLifecycleCheckColumns = struct {
	ID                  string
	DeploymentID        string
	ContainerID         string
	Type                string
	Scheme              string
	Path                string
	Port                string
	SuccessThreshold    string
	FailureThreshold    string
	InitialDelaySeconds string
	PeriodSeconds       string
	TimeoutSeconds      string
	Enable              string
	CreationTime        string
	UpdateTime          string
}{
	ID:                  "id",
	DeploymentID:        "deployment_id",
	ContainerID:         "container_id",
	Type:                "type",
	Scheme:              "scheme",
	Path:                "path",
	Port:                "port",
	SuccessThreshold:    "success_threshold",
	FailureThreshold:    "failure_threshold",
	InitialDelaySeconds: "initial_delay_seconds",
	PeriodSeconds:       "period_seconds",
	TimeoutSeconds:      "timeout_seconds",
	Enable:              "enable",
	CreationTime:        "creation_time",
	UpdateTime:          "update_time",
}

// DevopsProjects devops 项目管理
type DevopsProjects struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                                         // 项目名称
	TenantID     uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`                            // 租户ID
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"creationTime"` // 创建时间
	SoftDel      uint8      `gorm:"column:soft_del;type:tinyint unsigned;not null;default:0" json:"softDel"`                   // 删除标记
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
	SoftDel      string
}{
	ID:           "id",
	Name:         "name",
	TenantID:     "tenant_id",
	CreationTime: "creation_time",
	SoftDel:      "soft_del",
}

// DevopsProjectsApps devops 项目应用对应表
type DevopsProjectsApps struct {
	ID            uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
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

// PromethusClusterConfig Promethus 集群配置
type PromethusClusterConfig struct {
	ID        uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	ClusterID int64  `gorm:"column:cluster_id;type:bigint;not null" json:"clusterId"` // 集群ID
	URL       string `gorm:"column:url;type:varchar(255)" json:"url"`                 // Promethus访问地址
}

// TableName get sql table name.获取数据库表名
func (m *PromethusClusterConfig) TableName() string {
	return "promethus_cluster_config"
}

// PromethusClusterConfigColumns get sql column name.获取数据库列名
var PromethusClusterConfigColumns = struct {
	ID        string
	ClusterID string
	URL       string
}{
	ID:        "id",
	ClusterID: "cluster_id",
	URL:       "url",
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
	Type         int        `gorm:"column:type;type:int;not null" json:"type"`                  // 连接类型 vcs type:,1: 'github',,2: 'gitlab',,3: 'gogs',,4: 'gitee',
	Detail       string     `gorm:"column:detail;type:varchar(500)" json:"detail"`
	Enable       *uint8     `gorm:"column:enable;type:tinyint unsigned;default:1" json:"enable"`
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
	Enable       string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	MainID:       "main_id",
	Type:         "type",
	Detail:       "detail",
	Enable:       "enable",
	CreationTime: "creation_time",
	UpdateTime:   "update_time",
}

// ServiceConnectionTypeCode 服务连接类型
type ServiceConnectionTypeCode struct {
	ID          uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	ServiceType uint   `gorm:"column:service_type;type:int unsigned;not null" json:"serviceType"`
	TypeCode    string `gorm:"column:type_code;type:varchar(20);not null" json:"typeCode"`
	TypeName    string `gorm:"column:type_name;type:varchar(255);not null" json:"typeName"`
}

// TableName get sql table name.获取数据库表名
func (m *ServiceConnectionTypeCode) TableName() string {
	return "service_connection_type_code"
}

// ServiceConnectionTypeCodeColumns get sql column name.获取数据库列名
var ServiceConnectionTypeCodeColumns = struct {
	ID          string
	ServiceType string
	TypeCode    string
	TypeName    string
}{
	ID:          "id",
	ServiceType: "service_type",
	TypeCode:    "type_code",
	TypeName:    "type_name",
}

// ServiceConnectionTypeList 服务连接类别(git , docker, jenkins , system)
type ServiceConnectionTypeList struct {
	ID          uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	ServiceType uint   `gorm:"column:service_type;type:int unsigned;not null" json:"serviceType"`
	Value       uint   `gorm:"column:value;type:int unsigned;not null" json:"value"`
	Name        string `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Default     uint8  `gorm:"column:default;type:tinyint unsigned;not null" json:"default"`
}

// TableName get sql table name.获取数据库表名
func (m *ServiceConnectionTypeList) TableName() string {
	return "service_connection_type_list"
}

// ServiceConnectionTypeListColumns get sql column name.获取数据库列名
var ServiceConnectionTypeListColumns = struct {
	ID          string
	ServiceType string
	Value       string
	Name        string
	Default     string
}{
	ID:          "id",
	ServiceType: "service_type",
	Value:       "value",
	Name:        "name",
	Default:     "default",
}

// SgrCodeApplicationLanguage 字典-应用开发语言
type SgrCodeApplicationLanguage struct {
	ID             uint16 `gorm:"primaryKey;column:id;type:smallint unsigned;not null" json:"id"`
	Code           string `gorm:"column:code;type:varchar(8)" json:"code"`                                 // 编号
	Name           string `gorm:"column:name;type:varchar(50);not null" json:"name"`                       // 语言名称
	Alias          string `gorm:"column:alias;type:varchar(50);not null" json:"alias"`                     // 别名
	Icon           string `gorm:"column:icon;type:varchar(255);not null" json:"icon"`                      // 图标
	Content        string `gorm:"column:content;type:varchar(255);not null" json:"content"`                // 描述
	Sort           uint16 `gorm:"column:sort;type:smallint unsigned;not null;default:0" json:"sort"`       // 排序
	CompileScripts string `gorm:"column:compile_scripts;type:varchar(255);not null" json:"compileScripts"` // 默认编译脚本
}

// TableName get sql table name.获取数据库表名
func (m *SgrCodeApplicationLanguage) TableName() string {
	return "sgr_code_application_language"
}

// SgrCodeApplicationLanguageColumns get sql column name.获取数据库列名
var SgrCodeApplicationLanguageColumns = struct {
	ID             string
	Code           string
	Name           string
	Alias          string
	Icon           string
	Content        string
	Sort           string
	CompileScripts string
}{
	ID:             "id",
	Code:           "code",
	Name:           "name",
	Alias:          "alias",
	Icon:           "icon",
	Content:        "content",
	Sort:           "sort",
	CompileScripts: "compile_scripts",
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
	LastCommit   string     `gorm:"column:last_commit;type:text" json:"lastCommit"`                           // git最后一次提交记录
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
	LastCommit   string
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
	LastCommit:   "last_commit",
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

// SgrTenantConfigMap k8sconfig_map
type SgrTenantConfigMap struct {
	ID           uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                    // configmap 名称
	ClusterID    uint64     `gorm:"column:cluster_id;type:bigint unsigned;not null" json:"clusterId"`     // 集群id
	NamespaceID  uint64     `gorm:"column:namespace_id;type:bigint unsigned;not null" json:"namespaceId"` // 命名空间
	DeploymentID uint64     `gorm:"column:deployment_id;type:bigint unsigned;not null" json:"deploymentId"`
	AppID        uint64     `gorm:"column:app_id;type:bigint unsigned;not null" json:"appId"`
	Data         string     `gorm:"column:data;type:text;not null" json:"data"` // 配置内容
	CreationTime *time.Time `gorm:"column:creation_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"creationTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantConfigMap) TableName() string {
	return "sgr_tenant_config_map"
}

// SgrTenantConfigMapColumns get sql column name.获取数据库列名
var SgrTenantConfigMapColumns = struct {
	ID           string
	Name         string
	ClusterID    string
	NamespaceID  string
	DeploymentID string
	AppID        string
	Data         string
	CreationTime string
	UpdateTime   string
}{
	ID:           "id",
	Name:         "name",
	ClusterID:    "cluster_id",
	NamespaceID:  "namespace_id",
	DeploymentID: "deployment_id",
	AppID:        "app_id",
	Data:         "data",
	CreationTime: "creation_time",
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
	ID                            uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID                      uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`
	Name                          string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                                                                  // 部署名称(英文唯一)
	Nickname                      string     `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`                                                          // 部署中文名称
	ClusterID                     uint64     `gorm:"column:cluster_id;type:bigint unsigned;not null" json:"clusterId"`                                                   // 集群ID
	NamespaceID                   uint64     `gorm:"column:namespace_id;type:bigint unsigned;not null" json:"namespaceId"`                                               // 命名空间ID
	AppID                         uint64     `gorm:"column:app_id;type:bigint unsigned;not null" json:"appId"`                                                           // 应用ID
	Status                        uint8      `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`                                               // 状态
	ImageHub                      string     `gorm:"column:image_hub;type:varchar(200);not null;default:''" json:"imageHub"`                                             // 镜像仓库地址
	AppName                       string     `gorm:"column:app_name;type:varchar(50)" json:"appName"`                                                                    // 应用名称
	WorkloadType                  string     `gorm:"column:workload_type;type:varchar(25);not null;default:''" json:"workloadType"`                                      // 工作负载类型
	Replicas                      uint       `gorm:"column:replicas;type:int unsigned;not null;default:1" json:"replicas"`                                               // 副本数
	ServiceEnable                 *uint8     `gorm:"column:service_enable;type:tinyint unsigned;default:1" json:"serviceEnable"`                                         // 是否开启网络服务
	ServiceName                   string     `gorm:"column:service_name;type:varchar(150);not null;default:''" json:"serviceName"`                                       // 服务名称
	ServiceAway                   string     `gorm:"column:service_away;type:varchar(30)" json:"serviceAway"`                                                            // 服务类型(clusterIP nodeIP ...)
	ServicePort                   uint       `gorm:"column:service_port;type:int unsigned;not null;default:0" json:"servicePort"`                                        // 服务端口
	ServicePortType               string     `gorm:"column:service_port_type;type:varchar(8);not null;default:''" json:"servicePortType"`                                // 服务端口类型 http tcp
	LastImage                     string     `gorm:"column:last_image;type:varchar(350)" json:"lastImage"`                                                               // 最终部署镜像
	Level                         string     `gorm:"index:levev_idx;column:level;type:varchar(8)" json:"level"`                                                          // 应用部署级别
	MaxUnavailable                *uint      `gorm:"column:max_unavailable;type:int unsigned;default:25" json:"maxUnavailable"`                                          // 最大不可用
	MaxSurge                      *uint      `gorm:"column:max_surge;type:int unsigned;default:25" json:"maxSurge"`                                                      // 额外Pod
	TerminationGracePeriodSeconds uint       `gorm:"column:termination_grace_period_seconds;type:int unsigned;not null;default:30" json:"terminationGracePeriodSeconds"` //  最大容忍pod销毁时间(默认30s)
	Volumes                       string     `gorm:"column:volumes;type:varchar(255)" json:"volumes"`                                                                    // 卷(json)
	RuntimeEngine                 string     `gorm:"column:runtime_engine;type:varchar(20)" json:"runtimeEngine"`                                                        // 运行时引擎 (dapr, istio)
	CreateTime                    *time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`                                                                 // 创建时间
	UpdateTime                    *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`                                                                 // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantDeployments) TableName() string {
	return "sgr_tenant_deployments"
}

// SgrTenantDeploymentsColumns get sql column name.获取数据库列名
var SgrTenantDeploymentsColumns = struct {
	ID                            string
	TenantID                      string
	Name                          string
	Nickname                      string
	ClusterID                     string
	NamespaceID                   string
	AppID                         string
	Status                        string
	ImageHub                      string
	AppName                       string
	WorkloadType                  string
	Replicas                      string
	ServiceEnable                 string
	ServiceName                   string
	ServiceAway                   string
	ServicePort                   string
	ServicePortType               string
	LastImage                     string
	Level                         string
	MaxUnavailable                string
	MaxSurge                      string
	TerminationGracePeriodSeconds string
	Volumes                       string
	RuntimeEngine                 string
	CreateTime                    string
	UpdateTime                    string
}{
	ID:                            "id",
	TenantID:                      "tenant_id",
	Name:                          "name",
	Nickname:                      "nickname",
	ClusterID:                     "cluster_id",
	NamespaceID:                   "namespace_id",
	AppID:                         "app_id",
	Status:                        "status",
	ImageHub:                      "image_hub",
	AppName:                       "app_name",
	WorkloadType:                  "workload_type",
	Replicas:                      "replicas",
	ServiceEnable:                 "service_enable",
	ServiceName:                   "service_name",
	ServiceAway:                   "service_away",
	ServicePort:                   "service_port",
	ServicePortType:               "service_port_type",
	LastImage:                     "last_image",
	Level:                         "level",
	MaxUnavailable:                "max_unavailable",
	MaxSurge:                      "max_surge",
	TerminationGracePeriodSeconds: "termination_grace_period_seconds",
	Volumes:                       "volumes",
	RuntimeEngine:                 "runtime_engine",
	CreateTime:                    "create_time",
	UpdateTime:                    "update_time",
}

// SgrTenantDeploymentsContainers 应用部署容器配置
type SgrTenantDeploymentsContainers struct {
	ID                uint64  `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	Name              string  `gorm:"column:name;type:varchar(30);not null" json:"name"`                             // 容器名称
	DeployID          uint64  `gorm:"column:deploy_id;type:bigint unsigned;not null" json:"deployId"`                // 部署ID
	IsMain            uint8   `gorm:"column:is_main;type:tinyint unsigned;not null" json:"isMain"`                   // 是否是主容器
	Image             string  `gorm:"column:image;type:varchar(255);not null" json:"image"`                          // 镜像地址
	ImageVersion      string  `gorm:"column:image_version;type:varchar(40);not null" json:"imageVersion"`            // 镜像版本tag
	ImagePullStrategy string  `gorm:"column:image_pull_strategy;type:varchar(20);not null" json:"imagePullStrategy"` // 镜像拉取策略
	RequestCPU        float64 `gorm:"column:request_cpu;type:decimal(4,2);not null" json:"requestCpu"`               // 请求CPU使用量
	RequestMemory     float64 `gorm:"column:request_memory;type:decimal(5,0);not null" json:"requestMemory"`         // 请求内存使用量
	LimitCPU          float64 `gorm:"column:limit_cpu;type:decimal(4,2);not null" json:"limitCpu"`                   // 限制CPU
	LimitMemory       float64 `gorm:"column:limit_memory;type:decimal(5,0);not null" json:"limitMemory"`             // 限制内存
	Environments      string  `gorm:"column:environments;type:varchar(255);not null;default:''" json:"environments"` // 容器环境变量
	VolumeMounts      string  `gorm:"column:volume_mounts;type:varchar(255)" json:"volumeMounts"`                    // 卷挂载(json)
	Workdir           string  `gorm:"column:workdir;type:varchar(200)" json:"workdir"`                               // 工作目录
	RunCmd            string  `gorm:"column:run_cmd;type:varchar(200)" json:"runCmd"`                                // 启动命令行
	RunParams         string  `gorm:"column:run_params;type:varchar(100)" json:"runParams"`                          // 启动命令行参数
	Podstop           string  `gorm:"column:podstop;type:varchar(255)" json:"podstop"`                               // Pod生命周期-销毁前执行命令
	Poststart         string  `gorm:"column:poststart;type:varchar(255)" json:"poststart"`                           // Pod生命周期-创建前执行命令
	EnableLife        *uint8  `gorm:"column:enable_life;type:tinyint unsigned;default:0" json:"enableLife"`          // 是否开启生命周期
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
	VolumeMounts      string
	Workdir           string
	RunCmd            string
	RunParams         string
	Podstop           string
	Poststart         string
	EnableLife        string
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
	VolumeMounts:      "volume_mounts",
	Workdir:           "workdir",
	RunCmd:            "run_cmd",
	RunParams:         "run_params",
	Podstop:           "podstop",
	Poststart:         "poststart",
	EnableLife:        "enable_life",
}

// SgrTenantNamespace 集群_命名空间
type SgrTenantNamespace struct {
	ID            uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID      uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`   // 租户ID
	ClusterID     uint64     `gorm:"column:cluster_id;type:bigint unsigned;not null" json:"clusterId"` // 集群ID
	Namespace     string     `gorm:"column:namespace;type:varchar(50);not null" json:"namespace"`      // 命名空间名称
	Status        int8       `gorm:"column:status;type:tinyint;not null;default:1" json:"status"`      // 状态
	EnableRuntime uint8      `gorm:"column:enable_runtime;type:tinyint unsigned;not null;default:0" json:"enableRuntime"`
	RuntimeName   string     `gorm:"column:runtime_name;type:varchar(20);default:''" json:"runtimeName"`
	CreateTime    *time.Time `gorm:"column:create_time;type:datetime;not null" json:"createTime"` // 创建时间
	UpdateTime    *time.Time `gorm:"column:update_time;type:datetime;not null" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SgrTenantNamespace) TableName() string {
	return "sgr_tenant_namespace"
}

// SgrTenantNamespaceColumns get sql column name.获取数据库列名
var SgrTenantNamespaceColumns = struct {
	ID            string
	TenantID      string
	ClusterID     string
	Namespace     string
	Status        string
	EnableRuntime string
	RuntimeName   string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	TenantID:      "tenant_id",
	ClusterID:     "cluster_id",
	Namespace:     "namespace",
	Status:        "status",
	EnableRuntime: "enable_runtime",
	RuntimeName:   "runtime_name",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
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

// TenantDeliverablesProject 租户CI可交付物项目
type TenantDeliverablesProject struct {
	ID              uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	TenantID        uint64     `gorm:"column:tenant_id;type:bigint unsigned;not null" json:"tenantId"`                // 租户id
	ProjectName     string     `gorm:"column:project_name;type:varchar(50);not null" json:"projectName"`              // 项目名称
	HarborProjectID uint64     `gorm:"column:harbor_project_id;type:bigint unsigned;not null" json:"harborProjectId"` // harbor项目id
	CreateTime      *time.Time `gorm:"column:create_time;type:datetime;not null" json:"createTime"`                   // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *TenantDeliverablesProject) TableName() string {
	return "tenant_deliverables_project"
}

// TenantDeliverablesProjectColumns get sql column name.获取数据库列名
var TenantDeliverablesProjectColumns = struct {
	ID              string
	TenantID        string
	ProjectName     string
	HarborProjectID string
	CreateTime      string
}{
	ID:              "id",
	TenantID:        "tenant_id",
	ProjectName:     "project_name",
	HarborProjectID: "harbor_project_id",
	CreateTime:      "create_time",
}
