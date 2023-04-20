package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type RuntimeReq struct {
	mvc.RequestBody
	ID            uint64 `json:"id" uri:"id"`
	ComponentType string `json:"componentType" uri:"componentType"`
	Template      string `json:"template" uri:"template"`
	Doc           string `json:"doc" uri:"doc"`
}
