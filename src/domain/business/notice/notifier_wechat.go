package notice

import (
	"encoding/json"
	"kubelilin/utils"
	"log"
)

type wechatMarkdownMessage struct {
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
	Msgtype string `json:"msgtype" `
}

type Wechat struct {
	key string
}

func NewWechat(key string) Notifier {
	return &Wechat{key: key}
}

func (wechat Wechat) PostMessage(message Message) error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + wechat.key

	mkMessage := &wechatMarkdownMessage{
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: "# Kubelilin 部署通知 \n" +
				" > [应　用] : " + message.App + "\n" +
				" > [服　务] : " + message.Service + "\n" +
				" > [环　境] : " + message.Environment + "\n" +
				" > [版　本] : " + message.Version + "\n" +
				" > [分　支] : " + message.Branch + "\n" +
				" > [时　间] : " + message.Timestamp + "\n" +
				" > [结　果] : <font color=\"info\">" + message.Success + "</font>\n \n" +
				" > 详情请在Kubelilin平台应用中心查看 \n",
		},
		Msgtype: "markdown",
	}

	msg, err := json.Marshal(mkMessage)
	msgStr := string(msg)

	response := utils.PostHttpMessage(url, msgStr)
	log.Println(response)
	return err
}
