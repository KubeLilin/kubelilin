package kubernetes

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/dynamic"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"kubelilin/api/dto/requests"
	"kubelilin/api/dto/responses"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/pkg/page"
	"kubelilin/utils"
	"strconv"
	"strings"
)

type ServiceSupervisor struct {
	db             *gorm.DB
	clusterService *ClusterService
}

var temp string

func NewServiceSupervisor(db *gorm.DB, clusterService *ClusterService) *ServiceSupervisor {
	return &ServiceSupervisor{
		db:             db,
		clusterService: clusterService,
	}
}

func (svc *ServiceSupervisor) ApplyService(client corev1.CoreV1Interface, dp *models.SgrTenantDeployments) (interface{}, error) {
	namespace := &models.SgrTenantNamespace{}
	dbErr := svc.db.Model(&models.SgrTenantNamespace{}).Where("id=?", dp.NamespaceID).First(namespace)
	if dbErr.Error != nil {
		return nil, errors.New("未找到命名空间信息")
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
	//构造label
	serviceInfo.Labels = map[string]string{
		"kubelilin-default": "true",
		"k8s-app":           dp.Name,
	}
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
	servicePortName := "default-" + strings.ToLower((string)(protocol))
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

func (svc *ServiceSupervisor) QueryServiceList(req requests.ServiceRequest) (*page.Page, error) {
	var svcList []dto.ServiceList
	var clusterID uint64 = req.ClusterId
	client, err := svc.clusterService.GetClusterClientByTenantAndId(req.TenantId, clusterID)
	if err != nil {
		return nil, err
	}
	services := client.CoreV1().Services(req.Namespace)
	options := metav1.ListOptions{Limit: int64(req.PageSize)}
	if req.ContinueStr != "" {
		options.Continue = req.ContinueStr
	}
	list, err := services.List(context.TODO(), options)
	if err != nil {
		return nil, err
	}
	//data, err := json.Marshal(&list.Items[0])
	for _, x := range list.Items {

		svc := dto.ServiceList{
			Namespace:       x.Namespace,
			Name:            x.Name,
			Labels:          x.Labels,
			Selector:        x.Spec.Selector,
			Type:            string(x.Spec.Type),
			ClusterIP:       x.Spec.ClusterIP,
			SessionAffinity: string(x.Spec.SessionAffinity),
			CreateTime:      x.GetCreationTimestamp().Time,
			ContinueStr:     list.Continue,
		}
		svcList = append(svcList, svc)
	}
	var res = page.Page{PageIndex: req.CurrentPage}
	count := list.RemainingItemCount
	if count == nil {
		res.Total = int64(len(svcList))
	} else {
		res.Total = int64(len(svcList)) + *count
	}
	res.Data = svcList
	return &res, nil
}

func (svc *ServiceSupervisor) QueryServiceInfo(req requests.ServiceRequest) (*responses.ServiceInfo, error) {
	if req.Namespace == "" {
		return nil, errors.New("请传入命名空间")
	}
	if req.Name == "" {
		return nil, errors.New("请传入服务名称")
	}
	var clusterID uint64 = req.ClusterId
	client, err := svc.clusterService.GetClusterClientByTenantAndId(req.TenantId, clusterID)
	if err != nil {
		return nil, err
	}
	services := client.CoreV1().Services(req.Namespace)
	svcInfo, err := services.Get(context.TODO(), req.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	selectorStr, err := json.Marshal(svcInfo.Spec.Selector)
	labelsStr, err := json.Marshal(svcInfo.Labels)
	service := responses.ServiceInfo{
		Name:       svcInfo.Name,
		Namespace:  svcInfo.Namespace,
		Selector:   string(selectorStr),
		Labels:     string(labelsStr),
		CreateTime: svcInfo.GetCreationTimestamp().Time,
		Type:       string(svcInfo.Spec.Type),
		Port:       svcInfo.Spec.Ports,
	}
	return &service, err
}

func (svc *ServiceSupervisor) QueryNameSpaceByTenant(tenantId uint64, clusterId uint64) []responses.NamespaceList {
	var data = make([]responses.NamespaceList, 0)
	query := svc.db.Model(&models.SgrTenantNamespace{}).Where("tenant_id=?", tenantId)
	if clusterId > 0 {
		query = query.Where("cluster_id=?", clusterId)
	}
	query.Find(&data)

	return data
}

func (svc *ServiceSupervisor) QueryNamespaceList(tenantId uint64, clusterId uint64) ([]dto.Namespace, error) {
	client, err := svc.clusterService.GetClusterClientByTenantAndId(tenantId, clusterId)
	return GetAllNamespaces(client), err
}

func (svc *ServiceSupervisor) ChangeService(svcReq *requests.ServiceInfoReq) error {
	if svcReq.Namespace == "" {
		return errors.New("请传入命名空间")
	}
	if svcReq.Name == "" {
		return errors.New("请传入服务名称")
	}
	if len(svcReq.Port) == 0 {
		return errors.New("请传入端口映射信息")
	}
	for i, x := range svcReq.Port {
		for j, y := range svcReq.Port {
			if x.Port == y.Port && i != j {
				return errors.New("服务端口不可重复")
			}
		}
	}

	client, err := svc.clusterService.GetClusterClientByTenantAndId(svcReq.TenantId, svcReq.ClusterId)
	if err != nil {
		return err
	}
	configuration := applycorev1.ServiceApplyConfiguration{}
	serviceInfo := configuration.WithName(svcReq.Name)
	services := client.CoreV1().Services(svcReq.Namespace)
	var apiVersion = "v1"
	var kind = "Service"
	var svcName = svcReq.Name
	serviceInfo.Name = &svcName
	serviceInfo.APIVersion = &apiVersion
	serviceInfo.Kind = &kind

	metaLabels := make(map[string]string)
	spec := applycorev1.ServiceSpecApplyConfiguration{}
	_ = json.Unmarshal([]byte(svcReq.Selector), &metaLabels)

	spec.Selector = metaLabels
	//构造端口数据
	var ports []applycorev1.ServicePortApplyConfiguration
	for _, x := range svcReq.Port {
		var portNumber int32
		var nodePort int32
		var targetPort intstr.IntOrString
		if x.Port.Type == intstr.Int {
			portNumber = x.Port.IntVal
			nodePort = x.NodePort.IntVal
			targetPort = x.TargetPort
		} else {
			protStrForInt, _ := strconv.Atoi(x.Port.StrVal)
			portNumber = int32(protStrForInt)
			nodePortStrForInt, _ := strconv.Atoi(x.NodePort.StrVal)
			nodePort = int32(nodePortStrForInt)
			targetPort = intstr.Parse(x.TargetPort.StrVal)
		}
		protocol := v1.ProtocolTCP

		name := x.Name
		port := applycorev1.ServicePortApplyConfiguration{
			Name:       &name,
			Protocol:   &protocol,
			Port:       &portNumber,
			TargetPort: &targetPort,
		}
		var specType v1.ServiceType
		if svcReq.Type == CLUSTER_IP {
			specType = v1.ServiceTypeClusterIP
			spec.Type = &specType
		} else if svcReq.Type == NODE_PORT {
			specType = v1.ServiceTypeNodePort
			spec.Type = &specType
			port.NodePort = &nodePort
		}
		ports = append(ports, port)
	}
	spec.Ports = ports
	serviceInfo.Spec = &spec
	_, err = services.Apply(context.TODO(), serviceInfo, metav1.ApplyOptions{Force: true, FieldManager: "service-apply-fields"})
	return err
}

func (svc *ServiceSupervisor) QueryServiceByLabel(clusterId uint64, namespace string, label string) (*dto.ServicePort, error) {
	client, err := svc.clusterService.GetClusterClientByTenantAndId(0, clusterId)
	if err != nil {
		return nil, err
	}
	// get service by label
	options := metav1.ListOptions{}
	options.LabelSelector = label // "k8s-app= deployment name"
	services, _ := client.CoreV1().Services(namespace).List(context.TODO(), options)
	if len(services.Items) > 0 {
		if len(services.Items[0].Spec.Ports) > 0 {
			return &dto.ServicePort{
				ServiceName: services.Items[0].Name,
				PortName:    services.Items[0].Spec.Ports[0].Name,
				Port:        services.Items[0].Spec.Ports[0].Port,
				Success:     true,
			}, nil
		}
	}
	return nil, errors.New("未找到服务")
}

func (svc *ServiceSupervisor) CreateOrUpdateServiceMonitorByK8sCluster(model models.ApplicationServiceMonitor) error {
	config, err := svc.clusterService.GetClusterConfig(0, model.ClusterID)
	if err != nil {
		return err
	}
	// Create a dynamic client
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return err
	}
	// Define the GVR for ServiceMonitor resources
	gvr := schema.GroupVersionResource{
		Group:    "monitoring.coreos.com",
		Version:  "v1",
		Resource: "servicemonitors",
	}

	endpoint := map[string]interface{}{
		"path":     model.Path,
		"interval": utils.ToString(model.Interval) + "s",
	}
	// model.Port is number ?
	if utils.IsNumeric(model.Port) {
		endpoint["targetPort"] = utils.ToInt(model.Port)
	} else {
		endpoint["port"] = model.Port
	}
	// Create a ServiceMonitor resource object
	serviceMonitor := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "monitoring.coreos.com/v1",
			"kind":       "ServiceMonitor",
			"metadata": map[string]interface{}{
				"name":      model.Name,
				"namespace": model.Namespace,
				"labels": map[string]interface{}{
					"k8s-app": model.DeploymentName,
				},
			},
			"spec": map[string]interface{}{
				"endpoints": []map[string]interface{}{
					{
						"path":     model.Path,
						"port":     model.Port,
						"interval": utils.ToString(model.Interval) + "s",
					},
				},
				"namespaceSelector": map[string]interface{}{
					"matchNames": []string{model.Namespace},
				},
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"k8s-app": model.DeploymentName,
					},
				},
			},
		},
	}

	// Create or update the ServiceMonitor resource
	_, err = dynamicClient.Resource(gvr).Namespace(model.Namespace).Create(context.Background(), serviceMonitor, metav1.CreateOptions{})
	return err
}

// CreateServiceMonitor create service monitor
func (svc *ServiceSupervisor) CreateServiceMonitor(model models.ApplicationServiceMonitor) error {
	err := svc.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&model).Error
		if err != nil {
			return err
		}
		if err = svc.CreateOrUpdateServiceMonitorByK8sCluster(model); err != nil {
			tx.Rollback()
		}
		return err
	})
	return err
}

// query service monitor
func (svc *ServiceSupervisor) QueryServiceMonitorByAppId(appId uint64) ([]models.ApplicationServiceMonitor, error) {
	var monitors []models.ApplicationServiceMonitor
	err := svc.db.Model(&models.ApplicationServiceMonitor{}).Where("app_id = ?", appId).Find(&monitors).Error
	return monitors, err
}
