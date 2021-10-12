package kubernetes

import (
	"context"
	"flag"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sgr/domain/dto"
	"time"
)

func NewClientSet(path string) (*kubernetes.Clientset, error) {
	var kubeConfig *string
	if path == "" {
		if home := homedir.HomeDir(); home != "" {
			// 如果没有输入kube config参数，就用默认路径~/.kube/config
			kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kube config file")
		} else {
			// 如果取不到当前用户的家目录，就没办法设置kube config的默认目录了，只能从入参中取
			kubeConfig = flag.String("kubeconfig", "", "absolute path to the kube config file")
		}
		flag.Parse()
	} else {
		*kubeConfig = path
	}

	// 从本机加载kube config配置文件，因此第一个参数为空字符串
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	// kube config加载失败就直接退出了
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}

func GetPodList(client *kubernetes.Clientset, namespace string) []dto.Pod {
	emptyOptions := v1.ListOptions{}
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
	emptyOptions := v1.ListOptions{}
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
