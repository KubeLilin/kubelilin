package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type GatewayTeamRequest struct {
	mvc.RequestBody
	Id        uint64 `json:"id"`
	GatewayId uint64 `json:"gatewayId"`
	TenantId  uint64 `json:"tenantId"`
	TeamName  string `json:"name"`
	Level     string `json:"level"`
}
