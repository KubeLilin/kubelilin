package dto

type (
	StageInfo struct {
		Name  string     `json:"name"`
		Steps []StepInfo `json:"steps"`
	}

	StepInfo struct {
		Name    string                 `json:"name"`
		Key     string                 `json:"key"`
		Content map[string]interface{} `json:"content"`
	}
)
