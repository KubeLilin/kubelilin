package controllers

import "github.com/yoyofx/yoyogo/web/mvc"

type DemoController struct {
	mvc.ApiController
}

func NewDemoController() *DemoController {
	return &DemoController{}
}

func (controller DemoController) GetHello() mvc.ApiResult {
	return controller.OK("hello")
}
