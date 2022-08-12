package dto

type ServiceConnectionDTO struct {
	Detail      string `gorm:"column:detail;"`
	ServiceType uint   `gorm:"column:service_type;"` //服务连接器类型
	Value       uint   `gorm:"column:value;"`        //服务连接器ID
}

type ServiceConnectionInfo struct {
	Name        string `json:"name"`
	Repo        string `json:"repo"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	Type        uint   `json:"type"`
	ServiceType uint
}

const (
	SC_IMAGEHUB = 2
	SC_PIPELINE = 3
	SC_CALLBACK = 4
)
