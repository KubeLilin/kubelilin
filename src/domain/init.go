package domain

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/pkg/configuration"
	"github.com/yoyofxteam/dependencyinjection"
	"kubelilin/domain/business/app"
	"kubelilin/domain/business/kubernetes"
	"kubelilin/domain/business/tenant"
	"kubelilin/domain/conf"
	pipelineV1 "kubelilin/pkg/pipeline"
)

// init 所有业务对象的IOC容器注入入口
func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			// 程序启动时，引用init 的模块会自动注入IOC容器，此模块可访问配置文件
			// 将所有domain的业务对象全部由此入口进行容器注入，以便外部访问，如controller.

			//serviceCollection.AddSingleton(........)

			// 加载配置文件对象
			serviceCollection.AddTransient(tenant.NewUser)
			serviceCollection.AddTransient(tenant.NewTenantService)
			serviceCollection.AddTransient(tenant.NewSysMenuService)
			serviceCollection.AddTransient(tenant.NewTenantRoleService)
			serviceCollection.AddTransient(tenant.NewTenantUserRoleService)
			serviceCollection.AddTransient(tenant.NewRoleMenuService)
			serviceCollection.AddTransient(kubernetes.NewMetricsServer)
			serviceCollection.AddTransient(kubernetes.NewDeploymentSupervisor)
			serviceCollection.AddTransient(kubernetes.NewServiceSupervisor)
			serviceCollection.AddTransient(kubernetes.NewClusterService)
			serviceCollection.AddTransient(app.NewApplicationService)
			serviceCollection.AddTransient(app.NewDeploymentService)
			serviceCollection.AddTransient(app.NewVcsService)
			serviceCollection.AddTransient(app.NewPipelineService)
			serviceCollection.AddTransient(app.NewServiceConnectionService)
			serviceCollection.AddTransient(app.NewDevopsService)

			configuration.AddConfiguration(serviceCollection, conf.NewDbConfig)
			injectionJenkinsBuilder(config, serviceCollection)

		})
}

func injectionJenkinsBuilder(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
	jenkinsUrl := config.GetString("pipeline.jenkins.url")
	jenkinsToken := config.GetString("pipeline.jenkins.token")
	jenkinsUser := config.GetString("pipeline.jenkins.username")
	jenkinsNamespace := config.GetString("pipeline.jenkins.k8s-namespace")

	serviceCollection.AddSingleton(func() *pipelineV1.Builder {
		return pipelineV1.NewBuilder().UseJenkins(jenkinsUrl, jenkinsUser, jenkinsToken).
			UseKubernetes(jenkinsNamespace)
	})
}
