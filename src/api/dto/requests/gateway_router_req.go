package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type GatewayRouterRequest struct {
	mvc.RequestBody
	Name   string `json:"name" uri:"name"`
	Host   string `json:"host" uri:"host"`
	Uri    string `json:"uri" uri:"uri"`
	Desc   string `json:"desc" uri:"desc"`
	TeamId uint64 `json:"teamId" uri:"teamId"`
}
