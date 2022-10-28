package dto

type DevOpsProjectResource struct {
	Level     string  `json:"level" gorm:"level"`
	Replicas  uint32  `json:"replicas" gorm:"replicas"`
	SumCpu    float64 `json:"sum_cpu" gorm:"sum_cpu"`
	SumMemory float64 `json:"sum_memory" gorm:"sum_memory"`
}

type DevOpsProjectResourceTotals struct {
	TotalCpu    float64 `json:"total_cpu" gorm:"total_cpu"`
	TotalMemory float64 `json:"total_memory" gorm:"total_memory"`

	DevMetrics     DevOpsProjectResource `json:"dev"`
	TestMetrics    DevOpsProjectResource `json:"test"`
	ProdMetrics    DevOpsProjectResource `json:"prod"`
	ReleaseMetrics DevOpsProjectResource `json:"release"`
}
