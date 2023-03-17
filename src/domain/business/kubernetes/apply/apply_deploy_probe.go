package apply

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"kubelilin/domain/database/models"
)

func ApplyProbe(deployConfiguration *appsapplyv1.DeploymentApplyConfiguration, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers, context DeploymentApplyFuncContext) {
	var lifecycleCheckList []models.DeploymentContainerLifecycleCheck
	context.DbContext.Model(&models.DeploymentContainerLifecycleCheck{}).Find(&lifecycleCheckList, "deployment_id=?", dp.ID)
	if len(lifecycleCheckList) == 0 {
		return
	}
	var liveness *models.DeploymentContainerLifecycleCheck = nil
	var readiness *models.DeploymentContainerLifecycleCheck = nil
	if lifecycleCheckList[0].Type == "LIVENESS" {
		liveness = &lifecycleCheckList[0]
	} else {
		readiness = &lifecycleCheckList[0]
	}
	if len(lifecycleCheckList) > 1 {
		if liveness == nil {
			liveness = &lifecycleCheckList[1]
		} else {
			readiness = &lifecycleCheckList[1]
		}
	}

	//apply lifecycle configuration for the container ,that name is 'app' , is main container.
	var containerApplyConfig corev1.ContainerApplyConfiguration
	if len(deployConfiguration.Spec.Template.Spec.Containers) > 0 {
		containerApplyConfig = deployConfiguration.Spec.Template.Spec.Containers[0]
	}
	if dpc.EnableLife != nil && *dpc.EnableLife > 0 {
		_ = containerApplyConfig
		if liveness != nil && liveness.Enable > 0 {
			containerApplyConfig.LivenessProbe = corev1.Probe()
			containerApplyConfig.LivenessProbe.
				WithHTTPGet(corev1.HTTPGetAction().WithScheme(v1.URIScheme(liveness.Scheme)).
					WithPort(intstr.FromInt(int(liveness.Port))).WithPath(liveness.Path)).
				WithSuccessThreshold(int32(liveness.SuccessThreshold)).
				WithFailureThreshold(int32(liveness.FailureThreshold)).
				WithInitialDelaySeconds(int32(liveness.InitialDelaySeconds)).
				WithPeriodSeconds(int32(liveness.PeriodSeconds)).
				WithTimeoutSeconds(int32(liveness.TimeoutSeconds))
		}
		if readiness != nil && readiness.Enable > 0 {
			containerApplyConfig.ReadinessProbe = corev1.Probe()
			containerApplyConfig.ReadinessProbe.
				WithHTTPGet(corev1.HTTPGetAction().WithScheme(v1.URIScheme(readiness.Scheme)).
					WithPort(intstr.FromInt(int(readiness.Port))).WithPath(readiness.Path)).
				WithSuccessThreshold(int32(readiness.SuccessThreshold)).
				WithFailureThreshold(int32(readiness.FailureThreshold)).
				WithInitialDelaySeconds(int32(readiness.InitialDelaySeconds)).
				WithPeriodSeconds(int32(readiness.PeriodSeconds)).
				WithTimeoutSeconds(int32(readiness.TimeoutSeconds))
		}
	}
	deployConfiguration.Spec.Template.Spec.Containers[0] = containerApplyConfig
}
