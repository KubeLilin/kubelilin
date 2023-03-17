package apply

import (
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	resultV1 "kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/utils"
)

func ApplyVolume(deployConfiguration *appsapplyv1.DeploymentApplyConfiguration, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers, context DeploymentApplyFuncContext) {
	if dp.Volumes != "" {
		var volumeApplyConfigurationList []*corev1.VolumeApplyConfiguration
		var volumes []resultV1.Volume
		utils.StringToJson(dp.Volumes, &volumes)
		for _, volume := range volumes {
			var volumeApplyConfiguration *corev1.VolumeApplyConfiguration
			switch volume.VolumeType {
			case "configmap":
				volumeApplyConfiguration = corev1.Volume().WithName(volume.VolumeName).WithConfigMap(corev1.ConfigMapVolumeSource().WithName(volume.Value))
				break
			case "emptydir":
				volumeApplyConfiguration = corev1.Volume().WithName(volume.VolumeName).WithEmptyDir(corev1.EmptyDirVolumeSource())
				break
			}
			volumeApplyConfigurationList = append(volumeApplyConfigurationList, volumeApplyConfiguration)
		}
		deployConfiguration.Spec.Template.Spec.WithVolumes(volumeApplyConfigurationList...)
	}

	if dpc.VolumeMounts != "" {
		var volumeMouteApplyConfigurationList []*corev1.VolumeMountApplyConfiguration
		var volumeMounts []resultV1.VolumeMount
		utils.StringToJson(dpc.VolumeMounts, &volumeMounts)
		var containerApplyConfig corev1.ContainerApplyConfiguration
		if len(deployConfiguration.Spec.Template.Spec.Containers) > 0 {
			containerApplyConfig = deployConfiguration.Spec.Template.Spec.Containers[0]
			for _, volumeMount := range volumeMounts {
				volumeMountApplyConfiguration := corev1.VolumeMount().
					WithName(volumeMount.Volume).
					WithMountPath(volumeMount.DesPath)
				if volumeMount.SubPathType == "subPath" {
					volumeMountApplyConfiguration.WithSubPath(volumeMount.SubPath)
				}
				if volumeMount.MountType == "readonly" {
					volumeMountApplyConfiguration.WithReadOnly(true)
				}
				volumeMouteApplyConfigurationList = append(volumeMouteApplyConfigurationList, volumeMountApplyConfiguration)
			}
			containerApplyConfig.WithVolumeMounts(volumeMouteApplyConfigurationList...)
		}
		deployConfiguration.Spec.Template.Spec.Containers[0] = containerApplyConfig
	}
}
