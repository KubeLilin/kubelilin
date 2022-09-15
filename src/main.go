package main

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/pkg/configuration"
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
	ShowLogo()
	global.GlobalLogger = xlog.GetXLogger("global")

	web.NewWebHostBuilder().
		UseConfiguration(configuration.LocalConfig("config")).
		Configure(func(app *web.ApplicationBuilder) {
			//app.UseMiddleware(middlewares.NewCORS())
			app.SetJsonSerializer(extension.CamelJson())
			// 在 api/init.go 中定义 api or mvc 路由
			app.UseEndpoints(api.ConfigureApi)
			app.UseMvc(api.ConfigureMvc)
			// and more ...
		}).
		ConfigureServices(Bootstrap).Build().
		Run()
}

func Bootstrap(container *dependencyinjection.ServiceCollection) {
	//nacos.UseServiceDiscovery(container)
}

func ShowLogo() {
	logo := `
	KubeLilin is starting......
                                       ##         .
                                 ## ## ##        ==
                              ## ## ## ## ##    ===
                           /""""""""""""""""\___/ ===
                      ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
                           \______ o          _,/
                            \      \       _,'
                              '--.._\..--''         KubeLilin An Cloud-Native application platform for Kubernetes.
	`
	fmt.Println(logo)
}
