package postgresgorm

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type PaginationResponse struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}
