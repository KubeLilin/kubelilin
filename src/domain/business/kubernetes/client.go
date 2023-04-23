package kubernetes

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	resourcev1 "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	memory "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
	"kubelilin/domain/dto"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

func NewClientSet(path string) (*kubernetes.Clientset, error) {
	var kubeConfig string
	if path == "" {
		if home := homedir.HomeDir(); home != "" {
			// 如果没有输入kube config参数，就用默认路径~/.kube/config
			kubeConfig = filepath.Join(home, ".kube", "config")
			flag.Parse()
		}
	} else {
		kubeConfig = path
	}

	// 从本机加载kube config配置文件，因此第一个参数为空字符串
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	// kube config加载失败就直接退出了
	if err != nil {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return kubernetes.NewForConfig(config)
}

func NewClientSetWithFileContent(fileContent string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(fileContent))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	client.ServerVersion()
	return client, err
}

func GetPodList(client *kubernetes.Clientset, namespace string, node string, app string) []dto.Pod {
	emptyOptions := metav1.ListOptions{}

	if app != "" {
		emptyOptions.LabelSelector = "k8s-app=" + app
	}
	if node != "" {
		emptyOptions.FieldSelector = "spec.nodeName=" + node
	}

	list, _ := client.CoreV1().Pods(namespace).List(context.TODO(), emptyOptions)

	var podList []dto.Pod
	for _, item := range list.Items {
		podCount := len(item.Status.ContainerStatuses)
		podReadyCount := 0
		podRestartCount := 0
		var containerList []dto.Container
		for _, containerStatus := range item.Status.ContainerStatuses {
			// add container list to pod item;
			containerInfo := dto.Container{
				Id:           containerStatus.ContainerID,
				Name:         containerStatus.Name,
				Image:        containerStatus.Image,
				State:        containerStatus.State.String(),
				Ready:        containerStatus.Ready,
				RestartCount: containerStatus.RestartCount,
				Started:      containerStatus.Started,
			}
			containerList = append(containerList, containerInfo)
			// add pod status
			if containerStatus.Ready {
				podReadyCount++
			}
			podRestartCount = podRestartCount + int(containerStatus.RestartCount)
		}
		st := ""
		var age time.Duration
		if item.Status.StartTime != nil {
			st = item.Status.StartTime.Time.Format("2006-01-02 15:04:05")
			age = time.Now().Sub(item.Status.StartTime.Time)
		}
		podInfo := dto.Pod{
			Namespace:     item.Namespace,
			PodName:       item.Name,
			PodIP:         item.Status.PodIP,
			HostIP:        item.Status.HostIP,
			ClusterName:   item.ClusterName,
			Count:         podCount,
			Ready:         podReadyCount,
			StartTime:     st,
			Age:           age,
			Status:        string(item.Status.Phase),
			Restarts:      podRestartCount,
			ContainerList: containerList,
		}
		podList = append(podList, podInfo)
	}
	return podList
}

func GetAllNamespaces(client *kubernetes.Clientset) []dto.Namespace {
	emptyOptions := metav1.ListOptions{}
	list, _ := client.CoreV1().Namespaces().List(context.TODO(), emptyOptions)
	var namespaceList []dto.Namespace
	for _, ns := range list.Items {
		info := dto.Namespace{
			Name:   ns.Name,
			Status: string(ns.Status.Phase),
		}
		namespaceList = append(namespaceList, info)
	}
	return namespaceList
}

func getNodeRole(node *v1.Node) string {
	if _, ok := node.Labels["node-role.kubernetes.io/master"]; ok {
		return "master"
	}
	return "<none>"
}

func GetNodeList(client *kubernetes.Clientset) []dto.Node {
	emptyOptions := metav1.ListOptions{}
	list, _ := client.CoreV1().Nodes().List(context.TODO(), emptyOptions)
	var nodeList []dto.Node
	for _, nd := range list.Items {
		var address []dto.NodeAddress
		for _, addr := range nd.Status.Addresses {
			address = append(address, dto.NodeAddress{Type: string(addr.Type), Address: addr.Address})
		}

		node := dto.Node{
			Uid:       string(nd.UID),
			Name:      nd.Name,
			PodCIDR:   nd.Spec.PodCIDR,
			Addresses: address,
			Role:      getNodeRole(&nd),
			Capacity: dto.NodeStatus{
				CPU:     nd.Status.Capacity.Cpu().AsApproximateFloat64(),
				Memory:  nd.Status.Capacity.Memory().AsApproximateFloat64(),
				Pods:    nd.Status.Capacity.Pods().Value(),
				Storage: nd.Status.Capacity.StorageEphemeral().AsApproximateFloat64(),
			},
			Allocatable: dto.NodeStatus{
				CPU:     nd.Status.Allocatable.Cpu().AsApproximateFloat64(),
				Memory:  nd.Status.Allocatable.Memory().AsApproximateFloat64(),
				Pods:    nd.Status.Allocatable.Pods().Value(),
				Storage: nd.Status.Allocatable.StorageEphemeral().AsApproximateFloat64(),
			},
			OSImage:                 nd.Status.NodeInfo.OSImage,
			ContainerRuntimeVersion: nd.Status.NodeInfo.ContainerRuntimeVersion,
			KubeletVersion:          nd.Status.NodeInfo.KubeletVersion,
			OperatingSystem:         nd.Status.NodeInfo.OperatingSystem,
			Architecture:            nd.Status.NodeInfo.Architecture,
			Status:                  string(nd.Status.Phase),
		}
		node.Status = "notready"
		for _, condition := range nd.Status.Conditions {
			if condition.Type == v1.NodeReady && condition.Status == v1.ConditionTrue {
				node.Status = "ready"
				break
			}
		}

		nodeList = append(nodeList, node)
	}
	return nodeList
}

func GetDeploymentList(client *kubernetes.Clientset, namespace string) []dto.Workload {
	emptyOptions := metav1.ListOptions{}
	var deploymentList []dto.Workload
	list, _ := client.AppsV1().Deployments(namespace).List(context.TODO(), emptyOptions)

	for _, deploy := range list.Items {
		item := dto.Workload{
			Name:                deploy.Name,
			Namespace:           deploy.Namespace,
			Labels:              deploy.Labels,
			Selectors:           deploy.Spec.Selector.MatchLabels,
			Replicas:            deploy.Status.Replicas,
			AvailableReplicas:   deploy.Status.AvailableReplicas,
			UpdatedReplicas:     deploy.Status.UpdatedReplicas,
			ReadyReplicas:       deploy.Status.ReadyReplicas,
			UnavailableReplicas: deploy.Status.UnavailableReplicas,
		}
		if len(deploy.Spec.Template.Spec.Containers) > 0 {
			item.Image = deploy.Spec.Template.Spec.Containers[0].Image
			item.RequestCPU = deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().AsApproximateFloat64()
			item.RequestMemory = deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().AsApproximateFloat64()
			item.LimitsCPU = deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().AsApproximateFloat64()
			item.LimitsMemory = deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().AsApproximateFloat64()
		}
		deploymentList = append(deploymentList, item)
	}
	return deploymentList
}

func GetStatefulSetList(client *kubernetes.Clientset, namespace string) []dto.Workload {
	emptyOptions := metav1.ListOptions{}
	var deploymentList []dto.Workload
	list, _ := client.AppsV1().StatefulSets(namespace).List(context.TODO(), emptyOptions)

	for _, deploy := range list.Items {
		item := dto.Workload{
			Name:              deploy.Name,
			Namespace:         deploy.Namespace,
			Labels:            deploy.Labels,
			Selectors:         deploy.Spec.Selector.MatchLabels,
			Replicas:          deploy.Status.Replicas,
			AvailableReplicas: deploy.Status.AvailableReplicas,
			UpdatedReplicas:   deploy.Status.UpdatedReplicas,
			ReadyReplicas:     deploy.Status.ReadyReplicas,
		}
		if len(deploy.Spec.Template.Spec.Containers) > 0 {
			item.Image = deploy.Spec.Template.Spec.Containers[0].Image
			item.RequestCPU = deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().AsApproximateFloat64()
			item.RequestMemory = deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().AsApproximateFloat64()
			item.LimitsCPU = deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().AsApproximateFloat64()
			item.LimitsMemory = deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().AsApproximateFloat64()
		}
		deploymentList = append(deploymentList, item)
	}
	return deploymentList
}

func GetDaemonSetList(client *kubernetes.Clientset, namespace string) []dto.Workload {
	emptyOptions := metav1.ListOptions{}
	var deploymentList []dto.Workload
	list, _ := client.AppsV1().DaemonSets(namespace).List(context.TODO(), emptyOptions)

	for _, deploy := range list.Items {
		item := dto.Workload{
			Name:              deploy.Name,
			Namespace:         deploy.Namespace,
			Labels:            deploy.Labels,
			Selectors:         deploy.Spec.Selector.MatchLabels,
			Replicas:          deploy.Status.NumberAvailable,
			AvailableReplicas: deploy.Status.CurrentNumberScheduled,
			UpdatedReplicas:   deploy.Status.UpdatedNumberScheduled,
			ReadyReplicas:     deploy.Status.NumberReady,
		}
		if len(deploy.Spec.Template.Spec.Containers) > 0 {
			item.Image = deploy.Spec.Template.Spec.Containers[0].Image
			item.RequestCPU = deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().AsApproximateFloat64()
			item.RequestMemory = deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().AsApproximateFloat64()
			item.LimitsCPU = deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().AsApproximateFloat64()
			item.LimitsMemory = deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().AsApproximateFloat64()
		}
		deploymentList = append(deploymentList, item)
	}
	return deploymentList
}

func SetReplicasNumber(client *kubernetes.Clientset, namespace string, deploymentName string, number int32) (bool, error) {
	emptyOptions := metav1.GetOptions{}
	deployment, getErr := client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, emptyOptions)
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		return false, getErr
	}
	if number >= 0 && number <= 20 {
		//replica数量降低到1
		deployment.Spec.Replicas = &number
		_, err := client.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
		if err != nil {
			return false, err
		}
	} else {
		err := errors.New("Replicas Number can be set 0-20.")
		panic(err)
		return false, err
	}
	return true, nil
}

func DestroyPod(client *kubernetes.Clientset, namespace string, podName string) error {
	return client.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
}

func GetLogs(client *kubernetes.Clientset, namespace string, podName string, containerName string, lines int64) ([]string, error) {
	options := &v1.PodLogOptions{Container: containerName, TailLines: &lines}
	request := client.CoreV1().Pods(namespace).GetLogs(podName, options)
	readCloser, err := request.Stream(context.TODO())
	if err != nil {
		return nil, err
	}
	defer readCloser.Close()
	var logLines []string
	r := bufio.NewReader(readCloser)
	for {
		bytes, err := r.ReadBytes('\n')
		logLines = append(logLines, string(bytes))
		if err != nil {
			break
		}
		bytes = nil
	}
	return logLines, nil
}

func GetEvents(client *kubernetes.Clientset, namespace string, deployment string) []dto.EventItemDto {
	var eventList []dto.EventItemDto
	var k8sEventList []v1.Event
	podList, _ := client.CoreV1().Pods(namespace).List(context.TODO(),
		metav1.ListOptions{LabelSelector: "k8s-app=" + deployment})

	deploymentEvents, _ := client.CoreV1().Events(namespace).List(context.TODO(),
		metav1.ListOptions{TypeMeta: metav1.TypeMeta{Kind: "Deployment"}, FieldSelector: "involvedObject.name=" + deployment})
	k8sEventList = append(k8sEventList, deploymentEvents.Items...)

	for _, item := range podList.Items {
		podEvents, _ := client.CoreV1().Events(namespace).List(context.TODO(),
			metav1.ListOptions{TypeMeta: metav1.TypeMeta{Kind: "Pod"}, FieldSelector: "involvedObject.name=" + item.Name})
		k8sEventList = append(k8sEventList, podEvents.Items...)
	}

	for _, event := range k8sEventList {
		eventItem := dto.EventItemDto{
			FirstTime:   event.FirstTimestamp.Time,
			LastTime:    event.LastTimestamp.Time,
			Name:        event.Name,
			Level:       event.Type,
			Reason:      event.Reason,
			Information: event.Message,
			Kind:        event.InvolvedObject.Kind,
		}

		eventList = append(eventList, eventItem)
	}
	sort.Slice(eventList, func(i, j int) bool {
		return eventList[i].FirstTime.After(eventList[j].FirstTime)
	})
	return eventList
}

func Exec(client *kubernetes.Clientset, cfg *rest.Config, terminal *WebTerminal, shell string, namespace string, podName string, containerName string) error {
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Container: containerName,
		Command:   []string{shell},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       terminal.Tty(),
	}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(cfg, "POST", req.URL())
	if err != nil {
		return err
	}
	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             terminal.Stdin(),
		Stdout:            terminal.Stdout(),
		Stderr:            terminal.Stderr(),
		TerminalSizeQueue: terminal,
		Tty:               terminal.Tty(),
	})
	return err
}

func CreateNamespace(client *kubernetes.Clientset, namespace string, lables map[string]string) error {
	nsClient := client.CoreV1().Namespaces()
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	if lables != nil {
		ns.Labels = lables
	}

	_, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{})
	return err
}

func MapResourceQuotas(resourceQuotas *v1.ResourceQuota) dto.ResourceQuotas {
	resourceQuotasInfo := dto.ResourceQuotas{Labels: resourceQuotas.Labels}
	var resourceQuotaItemList []dto.ResourceQuotasItem
	resource := "limits.cpu"
	Limit := resourceQuotas.Status.Hard.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	Used := resourceQuotas.Status.Used.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	resourceQuotaItemList = append(resourceQuotaItemList,
		dto.ResourceQuotasItem{Name: resource,
			DisplayValue: Limit.String(), DisplayUsedValue: Used.String(), LimitValue: Limit.AsApproximateFloat64(), UsedValue: Used.AsApproximateFloat64()})

	resource = "limits.memory"
	Limit = resourceQuotas.Status.Hard.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	Used = resourceQuotas.Status.Used.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	resourceQuotaItemList = append(resourceQuotaItemList,
		dto.ResourceQuotasItem{Name: resource,
			DisplayValue: Limit.String(), DisplayUsedValue: Used.String(), LimitValue: Limit.AsApproximateFloat64(), UsedValue: Used.AsApproximateFloat64()})

	resource = "pods"
	Limit = resourceQuotas.Status.Hard.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	Used = resourceQuotas.Status.Used.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	resourceQuotaItemList = append(resourceQuotaItemList,
		dto.ResourceQuotasItem{Name: resource,
			DisplayValue: Limit.String(), DisplayUsedValue: Used.String(), LimitValue: Limit.AsApproximateFloat64(), UsedValue: Used.AsApproximateFloat64()})

	resourceQuotasInfo.Items = resourceQuotaItemList
	return resourceQuotasInfo
}

func GetResourceQuotasByNamespace(client *kubernetes.Clientset, namespace string) ([]dto.ResourceQuotasItem, error) {
	quotaClient := client.CoreV1().ResourceQuotas(namespace)
	resourceQuotas, err := quotaClient.Get(context.TODO(), "quota-"+namespace, metav1.GetOptions{})
	return MapResourceQuotas(resourceQuotas).Items, err
}

func GetAllNamespaceResourceQuotas(client *kubernetes.Clientset) map[string][]dto.ResourceQuotasItem {
	quotaClient := client.CoreV1().ResourceQuotas("")
	resourceQuotasList, err := quotaClient.List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(map[string]string{"kubelilin-default": "true"}).String(),
	})
	mapResourceQuotasList := make(map[string][]dto.ResourceQuotasItem)
	if err == nil {
		for _, rqItem := range resourceQuotasList.Items {
			rq := MapResourceQuotas(&rqItem)
			rqNs := rq.Labels["namespace"]
			mapResourceQuotasList["ns-"+rqNs] = rq.Items
		}
	}
	return mapResourceQuotasList
}

func CreateResourceQuotasByNamespace(client *kubernetes.Clientset, quotas dto.QuotasSpec) error {
	quotaClient := client.CoreV1().ResourceQuotas(quotas.Namespace)
	resourceQuotas, err := quotaClient.Get(context.TODO(), "quota-"+quotas.Namespace, metav1.GetOptions{})

	resourceHard := map[v1.ResourceName]resourcev1.Quantity{
		v1.ResourceLimitsCPU:    resourcev1.MustParse(strconv.Itoa(quotas.LimitCpu)),
		v1.ResourceLimitsMemory: resourcev1.MustParse(strconv.Itoa(quotas.LimitMemory) + "Gi"),
		v1.ResourcePods:         resourcev1.MustParse(strconv.Itoa(quotas.LimitPods)),
	}
	if err != nil { // not found for create
		resourceQuotas = &v1.ResourceQuota{
			ObjectMeta: metav1.ObjectMeta{
				Name: "quota-" + quotas.Namespace,
				Labels: map[string]string{
					"kubelilin-default": "true",
					"tenantId":          strconv.FormatUint(quotas.TenantID, 10),
					"clusterId":         strconv.FormatUint(quotas.ClusterId, 10),
					"namespace":         quotas.Namespace},
			},
			Spec: v1.ResourceQuotaSpec{
				Hard: resourceHard,
			},
		}
		_, err = quotaClient.Create(context.TODO(), resourceQuotas, metav1.CreateOptions{})
	} else { // founded for update
		resourceQuotas.Labels = map[string]string{
			"kubelilin-default": "true",
			"tenantId":          strconv.FormatUint(quotas.TenantID, 10),
			"clusterId":         strconv.FormatUint(quotas.ClusterId, 10),
			"namespace":         quotas.Namespace}
		resourceQuotas.Spec.Hard = resourceHard
		_, err = quotaClient.Update(context.TODO(), resourceQuotas, metav1.UpdateOptions{})
	}

	return err
}

func IsInstallDAPRRuntime(client *kubernetes.Clientset) bool {
	emptyOptions := metav1.GetOptions{}
	_, err := client.CoreV1().Namespaces().Get(context.TODO(), "dapr-system", emptyOptions)
	return !k8sErrors.IsNotFound(err)
}

func CreateDynamicResource(ctx context.Context, cfg *rest.Config, codec runtime.Serializer, data []byte) error {
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return err
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return err
	}
	obj := &unstructured.Unstructured{}
	_, gvk, err := codec.Decode(data, nil, obj)
	if err != nil {
		return err
	}

	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return err
	}

	namesapce := ""
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		namesapce = obj.GetNamespace()
		if namesapce == "" {
			namesapce = "default"
		}
	}

	var dynamicResource dynamic.ResourceInterface = dynamicClient.Resource(mapping.Resource)
	dynamicResource = dynamicClient.Resource(mapping.Resource).Namespace(namesapce)
	if _, err := dynamicResource.Create(ctx, obj, metav1.CreateOptions{}); err != nil {
		return err
	}

	return nil
}

func GetDaprComponentResource(cfg *rest.Config, namespace string) (any, error) {
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	dynamicResource := dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "dapr.io",
		Version:  "v1alpha1",
		Resource: "components",
	})
	componentList, err := dynamicResource.Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	//for _, component := range componentList.Items {
	//	fmt.Printf("component name=%s\n", component.GetName())
	//	fmt.Printf("component spec=%+v\n", component.Object["spec"])
	//	fmt.Println("=====================================")
	//}
	return componentList.Items, nil
}

// CreateOrUpdateDaprComponentResource create or update dapr component resource
func CreateOrUpdateDaprComponentResource(cfg *rest.Config, namespace string, component *unstructured.Unstructured) error {
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return err
	}
	daprGVR := schema.GroupVersionResource{
		Group:    "dapr.io",
		Version:  "v1alpha1",
		Resource: "components",
	}

	existingComponent, err := dynamicClient.Resource(daprGVR).Namespace("default").Get(context.Background(), "<NAME>", metav1.GetOptions{})
	// if the resource doesn't exist, we'll create it
	if k8sErrors.IsNotFound(err) {
		_, err = dynamicClient.Resource(daprGVR).Namespace(namespace).Create(context.Background(), component, metav1.CreateOptions{})
		return err
	} else {
		// if the resource exists, we'll update it
		existingComponent.Object["spec"] = component.Object["spec"]
		_, err = dynamicClient.Resource(daprGVR).Namespace(namespace).Update(context.Background(), existingComponent, metav1.UpdateOptions{})
		return err
	}
}

// DeleteDaprComponentResource delete dapr component resource
func DeleteDaprComponentResource(cfg *rest.Config, namespace string, componentName string) error {
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return err
	}
	dynamicResource := dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "dapr.io",
		Version:  "v1alpha1",
		Resource: "components",
	})
	err = dynamicResource.Namespace(namespace).Delete(context.TODO(), componentName, metav1.DeleteOptions{})
	return err
}
