package responses

type LabelValues struct {
	Label string      `json:"label,omitempty"`
	Value interface{} `json:"value,omitempty"`
}
