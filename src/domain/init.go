package domain

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/pkg/configuration"
	"github.com/yoyofxteam/dependencyinjection"
	"sgr/domain/conf"
)

// init 所有业务对象的IOC容器注入入口
func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			// 程序启动时，引用init 的模块会自动注入IOC容器，此模块可访问配置文件
			// 将所有domain的业务对象全部由此入口进行容器注入，以便外部访问，如controller.

			//serviceCollection.AddSingleton(........)

			// 加载配置文件对象
			configuration.AddConfiguration(serviceCollection, conf.NewDbConfig)

		})
}
