package notice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type markdownMessage struct {
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
	Msgtype string `json:"msgtype" `
}

type Wechat struct {
	key string
}

func NewWechat(key string) Notifier {
	//428c53f6-261a-404d-8315-7ca368598a06
	return &Wechat{key: key}
}

func (w Wechat) PostMessage(message Message) error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + w.key

	mkMessage := &markdownMessage{
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: "# Kubelilin 部署通知 \n" +
				" > [应　用] : " + message.App + "\n" +
				" > [流水线] : " + message.Pipeline + "\n" +
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

	response := postWechatMessage(url, msgStr)
	log.Println(response)
	return err
}

func postWechatMessage(sendUrl, msg string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", sendUrl, bytes.NewBuffer([]byte(msg)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	fmt.Println("response Body:", strBody)
	return strBody
}
