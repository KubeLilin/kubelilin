package dto

type ServiceConnectionDetails struct {
	Name     string `json:"name"`
	Repo     string `json:"repo"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Type     int    `json:"type"`
}
