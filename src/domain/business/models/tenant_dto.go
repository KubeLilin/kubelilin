package models

type TenantDTO struct {
	//properties
	ID     uint64 `json:"id"`
	TName  string `json:"name"`
	TCode  string `json:"code"`   // 租户编码
	Status int8   `json:"status"` // 状态

	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
}
