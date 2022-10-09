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
	err := apigateway.CreateOrUpdateRoute("ccc121", models.ApplicationAPIGatewayRouters{
		Name:         "test1",
		Desc:         "test1234efwefwf",
		TeamID:       10000000,
		Host:         "wiofjiowefuhu22.kubelilin.com",
		URI:          "/*",
		Websocket:    1,
		UpstreamType: "service",
		Loadbalance:  "roundrobin",
		Nodes:        "127.0.0.1:9080",
		Timeout:      3,
		DeploymentID: nil,
		Status:       1,
	})

	assert.Equal(t, err == nil, true)
}

func TestGatewayDeleteRoute(t *testing.T) {
	apigateway := getApiGateway()
	err := apigateway.DeleteRoute("ccc121")
	assert.Equal(t, err == nil, true)
}
