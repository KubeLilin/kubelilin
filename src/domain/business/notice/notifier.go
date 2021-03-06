package notice

type Message struct {
	// 应用名称
	App         string `json:"app"`
	Level       string `json:"level"`
	Service     string `json:"service"`     // 网络名称
	Environment string `json:"environment"` // 部署环境
	Version     string `json:"version"`     // 版本号
	Branch      string `json:"branch"`
	Timestamp   string `json:"timestamp"` // 发布时间
	Success     string `json:"success"`   // 成功、失败
}

type Notifier interface {
	PostMessage(message Message) error
}

type NotifierFactory interface {
}

// Plugins 扩展通知接口
var Plugins = []Plugin{
	{Label: "企业微信机器人", Value: "wechat", New: NewWechat},
	{Label: "钉钉机器人", Value: "dingtalk", New: NewDingTalk},
}
