package apply

import (
	appsapplyv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	"kubelilin/domain/database/models"
	"kubelilin/utils"
)

func ApplyRuntime(deployConfiguration *appsapplyv1.DeploymentApplyConfiguration, dp *models.SgrTenantDeployments, dpc *models.SgrTenantDeploymentsContainers, context DeploymentApplyFuncContext) {
	if dp.RuntimeEngine != "" {
		switch dp.RuntimeEngine {
		case "Dapr":
			if deployConfiguration.Spec.Template.Annotations == nil {
				deployConfiguration.Spec.Template.Annotations = make(map[string]string)
			}
			//annotations:
			//	dapr.io/enabled: "true"
			//	dapr.io/app-id: "myapp"
			//	dapr.io/app-port: "8223"
			//	dapr.io/enable-api-logging: "true"
			//	dapr.io/enable-metrics: "true"
			//	dapr.io/metrics-port: "9090"
			deployConfiguration.Spec.Template.Annotations["dapr.io/enabled"] = "true"
			deployConfiguration.Spec.Template.Annotations["dapr.io/app-id"] = dp.Name
			deployConfiguration.Spec.Template.Annotations["dapr.io/app-port"] = utils.ToString(dp.ServicePort)
			deployConfiguration.Spec.Template.Annotations["dapr.io/enable-api-logging"] = "true"
			// Enable metrics : https://docs.dapr.io/operations/monitoring/metrics/metrics-overview/
			deployConfiguration.Spec.Template.Annotations["dapr.io/enable-metrics"] = "true"
			deployConfiguration.Spec.Template.Annotations["dapr.io/metrics-port"] = "9090"
			//# This enables JSON-formatted logging
			//	dapr.io/log-as-json: "true"
			deployConfiguration.Spec.Template.Annotations["dapr.io/log-as-json"] = "true"
			// Enables tracing : https://docs.dapr.io/operations/monitoring/tracing/zipkin/
			//deployConfiguration.Spec.Template.Annotations["dapr.io/config"] = "tracing"
			//  Sidecar Resources:
			//  Cpu Limit: 300m, Request: 100m
			//  Memory Limit: 1000Mi, Request: 250Mi
			//	dapr.io/sidecar-cpu-limit
			//	dapr.io/sidecar-memory-limit
			//	dapr.io/sidecar-cpu-request
			//	dapr.io/sidecar-memory-request
			deployConfiguration.Spec.Template.Annotations["dapr.io/sidecar-cpu-limit"] = "300m"
			deployConfiguration.Spec.Template.Annotations["dapr.io/sidecar-memory-limit"] = "1000Mi"
			deployConfiguration.Spec.Template.Annotations["dapr.io/sidecar-cpu-request"] = "100m"
			deployConfiguration.Spec.Template.Annotations["dapr.io/sidecar-memory-request"] = "250Mi"

		}
	}
}
