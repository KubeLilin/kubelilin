package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type GatewayRouterListRequest struct {
	mvc.RequestBody
	Name   string `json:"name" uri:"name"`
	Host   string `json:"host" uri:"host"`
	Uri    string `json:"uri" uri:"uri"`
	Desc   string `json:"desc" uri:"desc"`
	TeamId uint64 `json:"teamId" uri:"teamId"`
}
type GatewayRouterRequest struct {
	mvc.RequestBody

	ID            uint64 `json:"id"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Loadbalance   string `json:"loadbalance"`
	UpstreamType  string `json:"upstreamType"`
	ApplicationID uint64 `json:"applicationId"`
	DeploymentID  uint64 `json:"deploymentId"`
	TeamID        uint64 `json:"teamId"`
	GatewayID     uint64 `json:"gatewayId"`
	Host          string `json:"host"`
	URI           string `json:"uri"`
	Rewrite       uint8  `json:"rewrite"`
	RegexURI      string `json:"regexUri"`
	RegexTmp      string `json:"regexTmp"`
	Liveness      string `json:"liveness"`
}
