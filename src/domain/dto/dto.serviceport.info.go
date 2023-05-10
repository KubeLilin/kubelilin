package dto

type ServicePort struct {
	ServiceName string `json:"serviceName"`
	PortName    string `json:"portName"`
	Success     bool   `json:"success"`
}
