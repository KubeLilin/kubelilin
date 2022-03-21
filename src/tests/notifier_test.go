package tests

import (
	"github.com/stretchr/testify/assert"
	"kubelilin/domain/business/notice"
	"testing"
)

func Test_WechatNotification(t *testing.T) {
	//NewWechat()

	notifier := notice.NewWechat("428c53f6-261a-404d-8315-7ca368598a06")
	message := notice.Message{
		App:         "ApiServer",
		Pipeline:    "pipeline-1-app-13",
		Environment: "dev-apiserver-microk8s",
		Version:     "v1.0.3",
		Branch:      "dev",
		Timestamp:   "2022-3-21 20:33:00",
		Success:     "成功",
	}
	err := notifier.PostMessage(message)

	assert.Equal(t, err == nil, true)
}
