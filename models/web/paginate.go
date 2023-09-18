package web

type Pagination struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	CurrentPage int         `json:"current_page"`
	TotalPage   int64       `json:"total_page"`
	Data        interface{} `json:"data"`
}
