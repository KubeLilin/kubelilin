package kubernetes

import (
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	"kubelilin/domain/business/kubernetes/apply"
	"kubelilin/domain/database/models"
)

type DeploymentApplyFunc func(deployConfiguration *appsapplyv1.DeploymentApplyConfiguration, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers)

var DeploymentApplyFuncList []DeploymentApplyFunc

func init() {
	DeploymentApplyFuncList = []DeploymentApplyFunc{apply.ApplyLifecycle, apply.ApplyVolume, apply.ApplyProbe}
}
