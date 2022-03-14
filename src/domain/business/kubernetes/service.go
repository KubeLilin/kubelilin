package kubernetes

import (
	"context"
	"errors"
	"gorm.io/gorm"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"kubelilin/domain/database/models"
	"strconv"
	"strings"
)

type ServiceSupervisor struct {
	db *gorm.DB
}

func NewServiceSupervisor(db *gorm.DB) *ServiceSupervisor {
	return &ServiceSupervisor{
		db: db,
	}
}

func (svc *ServiceSupervisor) ApplyService(client corev1.CoreV1Interface, dp *models.SgrTenantDeployments) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := svc.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return errors.New("未找到命名空间信息"), nil
	}
	k8sService := client.Services(namespace.Namespace)
	configuration := applycorev1.ServiceApplyConfiguration{}
	serviceInfo := configuration.WithName(dp.ServiceName)
	var apiVersion = "v1"
	var kind = "Service"
	var svcName = dp.ServiceName
	serviceInfo.Name = &svcName
	serviceInfo.APIVersion = &apiVersion
	serviceInfo.Kind = &kind
	//匹配dp的label
	//metaLabel := make(map[string]string)
	//metaLabel["k8s-app"] = dp.Name
	metaLabels := map[string]string{
		"kubelilin-default": "true",
		"appId":             strconv.FormatUint(dp.AppID, 10),
		"tenantId":          strconv.FormatUint(dp.TenantID, 10),
		"clusterId":         strconv.FormatUint(dp.ClusterID, 10),
		"namespaceId":       strconv.FormatUint(dp.NamespaceID, 10),
		"namespace":         namespace.Namespace,
		"k8s-app":           dp.Name,
		"profileLevel":      dp.Level,
	}
	spec := applycorev1.ServiceSpecApplyConfiguration{}
	spec.Selector = metaLabels
	//构造端口数据
	var ports []applycorev1.ServicePortApplyConfiguration
	portNumber := int32(dp.ServicePort)
	protocol := v1.ProtocolTCP
	targetPort := intstr.FromInt(int(dp.ServicePort))

	servicePortName := strings.ToLower((string)(protocol))
	port := applycorev1.ServicePortApplyConfiguration{
		Name:       &servicePortName,
		Protocol:   &protocol,
		Port:       &portNumber,
		TargetPort: &targetPort,
	}
	var specType v1.ServiceType
	if dp.ServicePortType == CLUSTER_IP {
		specType = v1.ServiceTypeClusterIP
		spec.Type = &specType
	} else if dp.ServicePortType == NODE_PORT {
		specType = v1.ServiceTypeNodePort
		spec.Type = &specType
		port.NodePort = &portNumber
	}
	ports = append(ports, port)
	spec.Ports = ports
	serviceInfo.Spec = &spec
	return k8sService.Apply(context.TODO(), serviceInfo, metav1.ApplyOptions{Force: true, FieldManager: "service-apply-fields"})
}
