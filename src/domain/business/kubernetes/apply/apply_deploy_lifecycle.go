package apply

import (
	"k8s.io/apimachinery/pkg/util/intstr"
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"kubelilin/domain/database/models"
	"kubelilin/utils"
)

func ApplyLifecycle(deployConfiguration *appsapplyv1.DeploymentApplyConfiguration, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers) {
	if dpc.EnableLife != nil && *dpc.EnableLife > 1 {
		var terminationGracePeriodSeconds = int64(dp.TerminationGracePeriodSeconds)
		if terminationGracePeriodSeconds < 30 {
			terminationGracePeriodSeconds = 30
		}
		deployConfiguration.Spec.Template.Spec.TerminationGracePeriodSeconds = &terminationGracePeriodSeconds
		strMaxSurge := "25%"
		strMaxUnavailable := "25%"
		if dp.MaxSurge != nil && dp.MaxUnavailable != nil {
			if *dp.MaxSurge > 0 && *dp.MaxSurge < 100 {
				strMaxSurge = utils.ToString(dp.MaxSurge) + "%"
			}
			if *dp.MaxUnavailable > 0 && *dp.MaxUnavailable < 100 {
				strMaxUnavailable = utils.ToString(dp.MaxUnavailable) + "%"
			}
			deployConfiguration.Spec.Strategy = &appsapplyv1.DeploymentStrategyApplyConfiguration{
				RollingUpdate: &appsapplyv1.RollingUpdateDeploymentApplyConfiguration{
					MaxUnavailable: &intstr.IntOrString{Type: intstr.String, StrVal: strMaxUnavailable},
					MaxSurge:       &intstr.IntOrString{Type: intstr.String, StrVal: strMaxSurge},
				},
			}
		}

		//apply lifecycle configuration for the container ,that name is 'app' , is main container.
		var containerApplyConfig corev1.ContainerApplyConfiguration
		if len(deployConfiguration.Spec.Template.Spec.Containers) > 0 {
			containerApplyConfig = deployConfiguration.Spec.Template.Spec.Containers[0]
		}
		if dpc.Poststart != "" {
			containerApplyConfig.Lifecycle.WithPostStart(corev1.Handler().WithExec(corev1.ExecAction().WithCommand(dpc.Poststart)))
		}
		if dpc.Podstop != "" {
			containerApplyConfig.Lifecycle.WithPreStop(corev1.Handler().WithExec(corev1.ExecAction().WithCommand(dpc.Podstop)))
		}
	}
}
