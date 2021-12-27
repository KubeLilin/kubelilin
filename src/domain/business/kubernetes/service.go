package kubernetes

import (
	"context"
	"errors"
	"gorm.io/gorm"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"sgr/domain/database/models"
)

type ServiceSupervisor struct {
	db *gorm.DB
}

func NewServiceSupervisor(db *gorm.DB) ServiceSupervisor {
	return ServiceSupervisor{
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
	serviceInfo := v1.Service{}
	serviceInfo.Name = dp.ServiceName
	serviceInfo.APIVersion = APPS_V1
	serviceInfo.Kind = "Service"
	//匹配dp的label
	metaLabel := make(map[string]string)
	metaLabel["k8s-app"] = dp.Name
	spec := v1.ServiceSpec{}
	spec.Selector = metaLabel
	if dp.ServicePortType == CLUSTER_IP {
		spec.Type = v1.ServiceTypeClusterIP
	} else if dp.ServicePortType == NODE_PORT {
		spec.Type = v1.ServiceTypeNodePort
	}
	var ports []v1.ServicePort
	portNumber := int32(dp.ServicePort)
	port := v1.ServicePort{
		Protocol:   v1.ProtocolTCP,
		Port:       portNumber,
		TargetPort: intstr.FromInt(int(dp.ServicePort)),
		NodePort:   portNumber,
	}
	ports = append(ports, port)
	spec.Ports = ports
	serviceInfo.Spec = spec
	return k8sService.Create(context.TODO(), &serviceInfo, metav1.CreateOptions{})
}
