package dto

type Deployment struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`

	Image string `json:"image"`

	Replicas          int32 `json:"replicas"`
	AvailableReplicas int32 `json:"availableReplicas"`

	RequestCPU    float64 `json:"requestCPU"`
	RequestMemory float64 `json:"requestMemory"`
	LimitsCPU     float64 `json:"limitsCPU"`
	LimitsMemory  float64 `json:"limitsMemory"`

	UpdatedReplicas     int32 `json:"updatedReplicas"`
	ReadyReplicas       int32 `json:"readyReplicas"`
	UnavailableReplicas int32 `json:"unavailableReplicas"`
}
