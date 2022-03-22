package notice

import (
	"encoding/json"
	"fmt"
	"kubelilin/utils"
	"log"
)

type dingTalkMarkdownMessage struct {
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	Msgtype string `json:"msgtype" `
}

type DingTalk struct {
	token string
}

func NewDingTalk(token string) Notifier {
	return &DingTalk{token}
}

func (ding DingTalk) PostMessage(message Message) error {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + ding.token

	mkMessage := &dingTalkMarkdownMessage{
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			Title: "Kubelilin 部署通知",
			Text: fmt.Sprintf(`
 > [应　用] :  %s

 > [流水线] :  %s

 > [环　境] :  %s

 > [版　本] :  %s

 > [分　支] :  %s

 > [时　间] :  %s

 > [结　果] :  **%s**


 > [详情请在Kubelilin平台应用中心查看](https://www.kubelilin.com)`,
				message.App, message.Pipeline, message.Environment,
				message.Version, message.Branch, message.Timestamp, message.Success),
		},
		Msgtype: "markdown",
	}
	msg, err := json.Marshal(mkMessage)
	msgStr := string(msg)

	response := utils.PostHttpMessage(url, msgStr)
	log.Println(response)
	return err
}
