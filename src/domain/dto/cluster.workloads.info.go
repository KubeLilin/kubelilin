package dto

type Workload struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`

	Labels    map[string]string `json:"labels"`
	Selectors map[string]string `json:"selectors"`

	Replicas          int32 `json:"replicas"`
	AvailableReplicas int32 `json:"availableReplicas"`

	Image string `json:"image"`

	RequestCPU    float64 `json:"requestCPU"`
	RequestMemory float64 `json:"requestMemory"`
	LimitsCPU     float64 `json:"limitsCPU"`
	LimitsMemory  float64 `json:"limitsMemory"`

	UpdatedReplicas     int32 `json:"updatedReplicas"`
	ReadyReplicas       int32 `json:"readyReplicas"`
	UnavailableReplicas int32 `json:"unavailableReplicas"`

	// job
	JobSchedule           string `json:"jobSchedule"`
	JobParallelism        int32  `json:"jobParallelism"`
	JobCompletions        int32  `json:"jobCompletions"`
	JobBackoffLimit       int32  `json:"jobBackoffLimit"`
	JobActive             int    `json:"jobActive"`
	JobLastScheduleTime   string `json:"lastScheduleTime"`
	JobLastSuccessfulTime string `json:"lastSuccessfulTime"`
}
