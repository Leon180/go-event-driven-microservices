package featuresdtos

type UpdateAccountRequest struct {
	ID            string  `json:"id" binding:"required"`
	MobileNumber  *string `json:"mobile_number"`
	AccountNumber *string `json:"account_number"`
	BranchAddress *string `json:"branch_address"`
}
