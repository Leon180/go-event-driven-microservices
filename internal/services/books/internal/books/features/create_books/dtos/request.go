package featuresdtos

type CreateBooksRequest struct {
	BranchID  string `json:"branch_id"`
	StartDate string `json:"start_date"` // format: YYYY-MM-DD
	EndDate   string `json:"end_date"`   // format: YYYY-MM-DD
}
