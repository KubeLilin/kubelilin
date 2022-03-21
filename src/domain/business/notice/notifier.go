package notice

type Message struct {
	// 应用名称
	App         string `json:"app"`
	Pipeline    string `json:"pipeline"`    // 流水线名称
	Environment string `json:"environment"` // 部署环境
	Version     string `json:"version"`     // 版本号
	Branch      string `json:"branch"`
	Timestamp   string `json:"timestamp"` // 发布时间
	Success     string `json:"success"`   // 成功、失败
}

type Notifier interface {
	PostMessage(message Message) error
}
