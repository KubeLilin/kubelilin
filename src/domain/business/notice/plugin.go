package notice

type Plugin struct {
	Label string                `json:"label"`
	Value string                `json:"value"`
	New   func(string) Notifier `json:"-"`
}
