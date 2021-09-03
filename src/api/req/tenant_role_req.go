package req

import "sgr/pkg/page"

type TenantRoleReq struct {
	ID     string `json:"id" uri:"id"`
	TName  string `json:"tName" uri:"tName"`
	TCode  string `json:"tCode" uri:"id"`
	Status string `json:"status" uri:"status"`
	page.PageRequest
}
