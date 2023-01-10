package apply

import (
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"kubelilin/domain/database/models"
)

func ApplyProbe(deployConfiguration *appsapplyv1.DeploymentApplyConfiguration, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) {
	//apply lifecycle configuration for the container ,that name is 'app' , is main container.
	var containerApplyConfig corev1.ContainerApplyConfiguration
	if len(deployConfiguration.Spec.Template.Spec.Containers) > 0 {
		containerApplyConfig = deployConfiguration.Spec.Template.Spec.Containers[0]
	}
	if dpc.EnableLife != nil && *dpc.EnableLife > 1 {
		_ = containerApplyConfig
		//containerApplyConfig.LivenessProbe
		//containerApplyConfig.ReadinessProbe
	}
}
