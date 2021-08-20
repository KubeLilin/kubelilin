package main

import (
	nacosConfig "github.com/yoyofx/yoyogo/pkg/configuration/nacos"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/nacos"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/middlewares"
	"github.com/yoyofxteam/dependencyinjection"
	"sgr/api"
	_ "sgr/domain"
)

func main() {
	// 加载 nacos 远程配置
	config := nacosConfig.RemoteConfig("config")

	web.NewWebHostBuilder().
		UseConfiguration(config).
		Configure(func(app *web.ApplicationBuilder) {
			app.UseEndpoints(api.ConfigureApi) // 在 api/init.go 中定义 api or mvc 路由
			app.UseMvc(api.ConfigureMvc)
			// add http middlewares
			app.UseMiddleware(middlewares.NewCORS())
			// and more ...
		}).
		ConfigureServices(func(container *dependencyinjection.ServiceCollection) {
			Bootstrap(container)
		}).
		Build().
		Run()

}

func Bootstrap(container *dependencyinjection.ServiceCollection) {
	// 注册 Nacos 服务发现
	nacos.UseServiceDiscovery(container)
}
