package dto

type ServicePort struct {
	ServiceName string `json:"serviceName"`
	PortName    string `json:"portName"`
	Port        int32  `json:"port"`
	Success     bool   `json:"success"`
}
