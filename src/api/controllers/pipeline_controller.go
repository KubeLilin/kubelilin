package controllers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/binding"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"gorm.io/gorm"
)

type PipelineController struct {
	mvc.ApiController

	db *gorm.DB
}

func NewPipelineController(database *gorm.DB) *PipelineController {
	return &PipelineController{db: database}
}

func (controller PipelineController) PostWebHook(ctx *context.HttpContext) mvc.ApiResult {

	req := map[string]interface{}{}
	_ = ctx.BindWith(&req, binding.JSON)
	fmt.Println(req)
	return controller.OK(req)
}
