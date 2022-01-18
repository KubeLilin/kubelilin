package kubernetes

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	resourcev1 "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sgr/domain/dto"
	"sort"
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

		podInfo := dto.Pod{
			Namespace:     item.Namespace,
			PodName:       item.Name,
			PodIP:         item.Status.PodIP,
			HostIP:        item.Status.HostIP,
			ClusterName:   item.ClusterName,
			Count:         podCount,
			Ready:         podReadyCount,
			StartTime:     item.Status.StartTime.Time.Format("2006-01-02 15:04:05"),
			Age:           time.Now().Sub(item.Status.StartTime.Time),
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
			Capacity: dto.NodeStatus{
				CPU:    nd.Status.Capacity.Cpu().AsApproximateFloat64(),
				Memory: nd.Status.Capacity.Memory().AsApproximateFloat64(),
				Pods:   nd.Status.Capacity.Pods().Size(),
			},
			Allocatable: dto.NodeStatus{
				CPU:    nd.Status.Allocatable.Cpu().AsApproximateFloat64(),
				Memory: nd.Status.Allocatable.Memory().AsApproximateFloat64(),
				Pods:   nd.Status.Allocatable.Pods().Size(),
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

func GetDeploymentList(client *kubernetes.Clientset, namespace string) []dto.Deployment {
	emptyOptions := metav1.ListOptions{}
	var deploymentList []dto.Deployment
	list, _ := client.AppsV1().Deployments(namespace).List(context.TODO(), emptyOptions)

	for _, deploy := range list.Items {
		item := dto.Deployment{
			Name:                deploy.Name,
			Namespace:           deploy.Namespace,
			Labels:              deploy.Labels,
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

func Exec(client *kubernetes.Clientset, cfg *rest.Config, terminal *WebTerminal, namespace string, podName string, containerName string) error {
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Container: containerName,
		Command:   []string{"/bin/bash"},
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

func CreateNamespace(client *kubernetes.Clientset, namespace string) error {
	nsClient := client.CoreV1().Namespaces()
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	_, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{})
	return err
}

func GetResourceQuotasByNamespace(client *kubernetes.Clientset, namespace string) ([]dto.ResourceQuotas, error) {
	quotaClient := client.CoreV1().ResourceQuotas(namespace)
	resourceQuotas, err := quotaClient.Get(context.TODO(), "quota-"+namespace, metav1.GetOptions{})
	var resourceQuotaInfo []dto.ResourceQuotas

	resource := "limits.cpu"
	Limit := resourceQuotas.Status.Hard.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	Used := resourceQuotas.Status.Used.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	resourceQuotaInfo = append(resourceQuotaInfo,
		dto.ResourceQuotas{Name: resource,
			DisplayValue: Limit.String(), DisplayUsedValue: Used.String(), LimitValue: Limit.Value(), UsedValue: Used.Value()})

	resource = "limits.memory"
	Limit = resourceQuotas.Status.Hard.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	Used = resourceQuotas.Status.Used.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	resourceQuotaInfo = append(resourceQuotaInfo,
		dto.ResourceQuotas{Name: resource,
			DisplayValue: Limit.String(), DisplayUsedValue: Used.String(), LimitValue: Limit.Value(), UsedValue: Used.Value()})

	resource = "count.pods"
	Limit = resourceQuotas.Status.Hard.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	Used = resourceQuotas.Status.Used.Name(v1.ResourceName(resource), resourcev1.DecimalExponent)
	resourceQuotaInfo = append(resourceQuotaInfo,
		dto.ResourceQuotas{Name: resource,
			DisplayValue: Limit.String(), DisplayUsedValue: Used.String(), LimitValue: Limit.Value(), UsedValue: Used.Value()})

	//resourceQuotaInfo = append(resourceQuotaInfo,
	//	dto.ResourceQuotas{
	//		Name:  "limits.memory",
	//		Value: resourceQuotas.Status.Hard.Name("limits.memory", resourcev1.DecimalExponent).String()})
	//resourceQuotaInfo = append(resourceQuotaInfo,
	//	dto.ResourceQuotas{
	//		Name:  "count.pods",
	//		Value: resourceQuotas.Status.Hard.Name("count.pods", resourcev1.DecimalExponent).String()})

	return resourceQuotaInfo, err
}
