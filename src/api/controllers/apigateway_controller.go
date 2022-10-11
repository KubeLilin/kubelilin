package controllers

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/domain/business/networks"
)

type ApiGatewayController struct {
	mvc.ApiController
	service *networks.ApiGatewayService
}

func NewApiGatewayController(service *networks.ApiGatewayService) *ApiGatewayController {
	return &ApiGatewayController{service: service}
}
