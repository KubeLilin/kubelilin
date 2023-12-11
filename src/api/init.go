package api

import (
	"github.com/yoyofx/yoyogo/web/endpoints"
	"github.com/yoyofx/yoyogo/web/mvc"
	"github.com/yoyofx/yoyogo/web/router"
	"kubelilin/api/controllers"
)

func ConfigureApi(builder router.IRouterBuilder) {
	endpoints.UseHealth(builder)
	endpoints.UseReadiness(builder)
	endpoints.UseLiveness(builder)
	endpoints.UsePrometheus(builder)
	endpoints.UseViz(builder)
	endpoints.UseRouteInfo(builder)
	endpoints.UseJwt(builder)
}

func ConfigureMvc(builder *mvc.ControllerBuilder) {
	//builder.AddFilter("/v1/user/info", &contollers.TestActionFilter{})
	builder.AddController(controllers.NewDemoController)
	builder.AddController(controllers.NewTenantController)
	builder.AddController(controllers.NewSysMenuController)
	builder.AddController(controllers.NewUserController)
	builder.AddController(controllers.NewTenantRoleController)
	builder.AddController(controllers.NewTenantUserRoleController)
	builder.AddController(controllers.NewRoleMenuController)
	builder.AddController(controllers.NewClusterController)
	builder.AddController(controllers.NewApplicationController)
	builder.AddController(controllers.NewDeploymentController)
	builder.AddController(controllers.NewPodController)
	builder.AddController(controllers.NewMetricsController)
	builder.AddController(controllers.NewServiceConnectionController)
	builder.AddController(controllers.NewServiceController)
	builder.AddController(controllers.NewDevopsController)
	builder.AddController(controllers.NewApiGatewayController)
	builder.AddController(controllers.NewConfigmapController)
	builder.AddController(controllers.NewDynamicController)
	builder.AddController(controllers.NewRuntimeController)
	builder.AddController(controllers.NewPipelineController)
	builder.AddController(controllers.NewDeliverablesController)
}
