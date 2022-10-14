package networks

import (
	"fmt"
	"github.com/guonaihong/gout"
	"kubelilin/domain/database/models"
	"kubelilin/utils"
)

type APISIXProxy struct {
	adminUrl string
	token    string
}

func NewAPISIXProxy(url string, token string) APISIXProxy {
	return APISIXProxy{adminUrl: url, token: token}
}

func (proxy APISIXProxy) GetRoutes() map[string]interface{} {
	successRes := new(map[string]interface{})
	url := proxy.adminUrl + "/apisix/admin/routes"
	_ = gout.
		GET(url).
		BindJSON(&successRes).
		SetHeader(gout.H{"X-API-KEY": proxy.token}).
		Do()
	return *successRes
}

func (proxy APISIXProxy) GetRouteById(id string) map[string]interface{} {
	successRes := new(map[string]interface{})
	url := fmt.Sprintf("%s/apisix/admin/routes/%s", proxy.adminUrl, id)
	_ = gout.
		GET(url).
		BindJSON(&successRes).
		SetHeader(gout.H{"X-API-KEY": proxy.token}).
		Do()
	return *successRes
}

/*
CreateOrUpdateRoute 创建或更新路由
*/
func (proxy APISIXProxy) CreateOrUpdateRoute(id string, router models.ApplicationAPIGatewayRouters) error {
	url := fmt.Sprintf("%s/apisix/admin/routes/%s", proxy.adminUrl, id)
	m := gout.H{
		"name":             router.Name,
		"desc":             router.Desc,
		"uri":              router.URI,
		"enable_websocket": utils.Uint8ToBool(router.Websocket),
		"upstream": gout.H{
			"type": router.Loadbalance,
			"nodes": gout.H{
				router.Nodes: 1,
			},
		},
	}
	if router.URI == "" {
		m["uri"] = "/*"
	}
	if router.Host != "" {
		m["host"] = router.Host
	}
	if router.Rewrite > 0 {
		m["plugins"] = gout.H{
			"proxy-rewrite": gout.H{
				"regex_uri": gout.A{
					router.RegexURI, router.RegexTmp}},
		}
	}
	if router.Label != "" {
		m["labels"] = gout.H{"API_ENV": router.Label}
	}
	return gout.
		PUT(url).
		SetHeader(gout.H{"X-API-KEY": proxy.token}).
		SetJSON(m).
		Do()
}

func (proxy APISIXProxy) DeleteRoute(id string) error {
	url := fmt.Sprintf("%s/apisix/admin/routes/%s", proxy.adminUrl, id)
	return gout.
		DELETE(url).
		SetHeader(gout.H{"X-API-KEY": proxy.token}).
		Do()
}
