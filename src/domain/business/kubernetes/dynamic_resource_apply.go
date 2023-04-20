package kubernetes

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	yamlCodec "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

type DynamicResourceSupervisor struct {
	clusterService *ClusterService
}

func NewDynamicResourceSupervisor(clusterService *ClusterService) *DynamicResourceSupervisor {
	return &DynamicResourceSupervisor{clusterService: clusterService}
}

func (dynamicResource *DynamicResourceSupervisor) CreateFromYAML(clusterId uint64, resource string) error {
	restConfig, err := dynamicResource.clusterService.GetClusterConfig(0, clusterId)
	if err != nil {
		return err
	}
	codec := yamlCodec.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	return CreateDynamicResource(context.TODO(), restConfig, codec, []byte(resource))
}
