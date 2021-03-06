package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	_ "github.com/yoyofx/yoyogo/pkg/datasources/mysql"
	_ "github.com/yoyofx/yoyogo/pkg/datasources/redis"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/actionresult/extension"
	"github.com/yoyofxteam/dependencyinjection"
	"kubelilin/api"
	_ "kubelilin/domain"
	"kubelilin/pkg/global"
)

func main() {
	global.GlobalLogger = xlog.GetXLogger("global")

	// 加载 nacos 远程配置
	//config := nacosConfig.RemoteConfig("config")
	config := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	web.NewWebHostBuilder().
		UseConfiguration(config).
		Configure(func(app *web.ApplicationBuilder) {
			// add http middlewares
			//app.UseMiddleware(middlewares.NewCORS())
			app.SetJsonSerializer(extension.CamelJson())
			app.UseEndpoints(api.ConfigureApi) // 在 api/init.go 中定义 api or mvc 路由
			app.UseMvc(api.ConfigureMvc)
			//app.UseMiddleware(middlewares.NewJwt())
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
	//nacos.UseServiceDiscovery(container)
}
