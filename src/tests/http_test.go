package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"kubelilin/domain/business/networks"
	"kubelilin/domain/database/models"
	"testing"
)

func getApiGateway() networks.APISIXProxy {
	gatewayService := networks.NewApiGatewayService(InitDb())
	gatewayEntity, _ := gatewayService.GetById(1)
	apisix := networks.NewAPISIXProxy(gatewayEntity.AdminURI, gatewayEntity.AccessToken)
	return apisix
}

func TestGatewayGetRoutes(t *testing.T) {
	apigateway := getApiGateway()
	s := apigateway.GetRoutes()
	fmt.Printf("html size = %d\n", len(s))
	assert.Equal(t, len(s) > 0, true)
}

func TestGatewayGetRouteById(t *testing.T) {
	apigateway := getApiGateway()
	s := apigateway.GetRouteById("428846874171016082")
	fmt.Printf("html size = %d\n", len(s))
	assert.Equal(t, len(s) > 0, true)
}

func TestGatewayCreateRoute(t *testing.T) {
	apigateway := getApiGateway()
	err := apigateway.CreateOrUpdateRoute("r10000001", models.ApplicationAPIGatewayRouters{
		Name:         "dev-nginx-cls-hbktlqm5_r10000001",
		Desc:         "dev-nginx-cls-hbktlqm5",
		TeamID:       10000000,
		Host:         "proxy.kubelilin.com",
		URI:          "/dev-nginx-cls-hbktlqm5/*",
		Websocket:    1,
		UpstreamType: "service",
		Loadbalance:  "roundrobin",
		Nodes:        "dev-nginx-cls-hbktlqm5-svc-cluster-sgr.yoyogo:80",
		Timeout:      3,
		DeploymentID: 0,
		Rewirte:      1,
		RegexURI:     "^/dev-nginx-cls-hbktlqm5/(.*)",
		RegexTmp:     "/$1",
		Label:        "deployment",
		Status:       1,
	})

	assert.Equal(t, err == nil, true)
}

func TestGatewayDeleteRoute(t *testing.T) {
	apigateway := getApiGateway()
	err := apigateway.DeleteRoute("ccc121")
	assert.Equal(t, err == nil, true)
}
