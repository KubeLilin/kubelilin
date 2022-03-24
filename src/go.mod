module kubelilin

go 1.16

require (
	github.com/ahmetb/go-linq/v3 v3.2.0
	github.com/fasthttp/websocket v1.4.3
	github.com/gogs/go-gogs-client v0.0.0-20210131175652-1d7215cd8d85
	github.com/isbrick/http-client v0.0.0-20210321135403-0a5df00fdb84
	github.com/jinzhu/copier v0.3.2
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/yoyofx/yoyogo v1.7.11
	github.com/yoyofxteam/dependencyinjection v1.0.1
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.21.12
	k8s.io/api v0.22.0
	k8s.io/apimachinery v0.22.0
	k8s.io/cli-runtime v0.22.0
	k8s.io/client-go v0.22.0
	k8s.io/metrics v0.22.0

)

replace gorm.io/gorm v1.21.11 => github.com/go-gorm/gorm v1.21.11
