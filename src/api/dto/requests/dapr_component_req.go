package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type DaprComponentRequest struct {
	mvc.RequestBody

	ClusterId uint64         `json:"cid"`
	Version   string         `json:"version"`
	Namespace string         `json:"namespace"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Metadata  map[string]any `json:"metadata"`
}
