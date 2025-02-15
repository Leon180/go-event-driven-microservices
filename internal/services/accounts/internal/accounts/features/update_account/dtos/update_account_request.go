package dtos

type UpdateAccountRequest struct {
	MobileNumber  string  `json:"mobile_number" binding:"required"`
	AccountNumber *int64  `json:"account_number"`
	AccountType   *string `json:"account_type"`
	BranchAddress *string `json:"branch_address"`
	ActiveSwitch  *bool   `json:"active_switch"`
}
