package featuresdtos

type UpdateLoanRequest struct {
	ID           string  `json:"id"            binding:"required"`
	MobileNumber *string `json:"mobile_number"`
	TotalAmount  *string `json:"total_amount"`
	PaidAmount   *string `json:"paid_amount"`
	InterestRate *string `json:"interest_rate"`
	Term         *int    `json:"term"`
	ActiveSwitch *bool   `json:"active_switch"`
}
