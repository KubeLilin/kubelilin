package apply

import (
	"gorm.io/gorm"
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
)

type DeploymentApplyFuncContext struct {
	DeployConfiguration *appsapplyv1.DeploymentApplyConfiguration
	DbContext           *gorm.DB
}
