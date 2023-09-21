package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type PipelineStatusRequest struct {
	mvc.RequestBody

	Pid         string `json:"pid"`
	Appid       string `json:"appid"`
	Branch      string `json:"branch"`
	Image       string `json:"image"`
	BuildNumber string `json:"buildNumber"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
}
