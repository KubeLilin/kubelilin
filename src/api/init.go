package api

import (
	"github.com/yoyofx/yoyogo/web/endpoints"
	"github.com/yoyofx/yoyogo/web/mvc"
	"github.com/yoyofx/yoyogo/web/router"
	"sgr/api/controllers"
)

func ConfigureApi(builder router.IRouterBuilder) {
	endpoints.UseHealth(builder)
	endpoints.UseReadiness(builder)
	endpoints.UseLiveness(builder)
	endpoints.UsePrometheus(builder)
	endpoints.UseViz(builder)
	endpoints.UseRouteInfo(builder)
}

func ConfigureMvc(builder *mvc.ControllerBuilder) {
	//builder.AddFilter("/v1/user/info", &contollers.TestActionFilter{})
	builder.AddController(controllers.NewDemoController)
	builder.AddController(controllers.NewTenantController)
	builder.AddController(controllers.NewSysMenuController)
	builder.AddController(controllers.NewUserController)
	builder.AddController(controllers.NewTenantRoleController)
	builder.AddController(controllers.NewUserRoleController)
}
