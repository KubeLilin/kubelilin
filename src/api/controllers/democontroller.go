package controllers

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"gorm.io/gorm"
	dbmodels "sgr/domain/database/models"
)

type DemoController struct {
	mvc.ApiController

	db *gorm.DB
}

func NewDemoController(database *gorm.DB) *DemoController {
	return &DemoController{db: database}
}

func (controller DemoController) GetHello() mvc.ApiResult {
	return controller.OK("hello")
}

/*
	GetTestSQL test sql with that get tenant by id
	SQL: SELECT * FROM sgr_tenant WHERE id=1
	URL: http://localhost:8080/v1/demo/testsql
*/
func (controller DemoController) GetTestSQL() mvc.ApiResult {
	// get db object and then take a tenant manager
	tenantMgr := dbmodels.SgrTenantMgr(controller.db)
	// get tenant by manager and then that by get from id
	adminTenant, _ := tenantMgr.GetFromID(1)
	// return tenant to response
	return controller.OK(adminTenant)
}
