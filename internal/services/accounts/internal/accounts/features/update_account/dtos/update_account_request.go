package featuresdtos

type UpdateAccountRequest struct {
	MobileNumber  string  `json:"-"`
	AccountNumber *string `json:"account_number"`
	AccountType   *string `json:"account_type"`
	BranchAddress *string `json:"branch_address"`
	ActiveSwitch  *bool   `json:"active_switch"`
}
