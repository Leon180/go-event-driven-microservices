package featuresdtos

type UpdateAccountRequest struct {
	MobileNumber  *string `json:"mobile_number"`
	AccountNumber *string `json:"account_number"`
	BranchAddress *string `json:"branch_address"`
}
