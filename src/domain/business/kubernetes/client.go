package kubernetes

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sgr/domain/dto"
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
	}
	return kubernetes.NewForConfig(config)
}

func GetPodList(client *kubernetes.Clientset, namespace string, app string) []dto.Pod {
	emptyOptions := metav1.ListOptions{}

	if app != "" {
		emptyOptions.LabelSelector = "k8s-app=" + app
	}

	list, _ := client.CoreV1().Pods(namespace).List(context.TODO(), emptyOptions)

	var podList []dto.Pod
	for _, item := range list.Items {
		podCount := len(item.Status.ContainerStatuses)
		podReadyCount := 0
		podRestartCount := 0
		for _, containerStatus := range item.Status.ContainerStatuses {
			if containerStatus.Ready {
				podReadyCount++
			}
			podRestartCount = podRestartCount + int(containerStatus.RestartCount)
		}

		podInfo := dto.Pod{
			Namespace:   item.Namespace,
			PodName:     item.Name,
			PodIP:       item.Status.PodIP,
			HostIP:      item.Status.HostIP,
			ClusterName: item.ClusterName,
			Count:       podCount,
			Ready:       podReadyCount,
			Age:         time.Now().Sub(item.Status.StartTime.Time),
			Status:      string(item.Status.Phase),
			Restarts:    podRestartCount,
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
		}

		nodeList = append(nodeList, node)
	}
	return nodeList
}
