package page

type Page struct {
	PageIndex int
	PageSize  int
	Total     int64
	Data      interface{}
}

type PageRequest struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

func (pg PageRequest) OffSet() int {
	return (pg.PageIndex - 1) * pg.PageSize
}
