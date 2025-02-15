package dtos

type CreateAccountRequest struct {
	MobileNumber  string `json:"mobile_number" binding:"required"`
	AccountNumber int64  `json:"account_number" binding:"required"`
	AccountType   string `json:"account_type" binding:"required"`
	BranchAddress string `json:"branch_address" binding:"required"`
}
