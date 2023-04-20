package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type ApplyYAML struct {
	mvc.RequestBody
	ClusterId uint64 `json:"clusterId"`
	YAML      string `json:"yaml"`
}
