package dtos

type Account struct {
	ID            string
	MobileNumber  string
	AccountNumber int64
	AccountType   string
	BranchAddress string
	ActiveSwitch  bool
}

func (a *Account) IsActive() bool {
	return a.ActiveSwitch
}

type Accounts []Account

type AccountWithHistory struct {
	Account
	History CommonHistoryModelWithUpdate
}

func (a *AccountWithHistory) IsActive() bool {
	return a.Account.ActiveSwitch
}

type AccountsWithHistory []AccountWithHistory

type UpdateAccount struct {
	ID            string
	MobileNumber  string
	AccountNumber *int64
	AccountType   *string
	BranchAddress *string
	ActiveSwitch  *bool
}

func (u *UpdateAccount) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.AccountNumber != nil {
		updateMap["account_number"] = *u.AccountNumber
	}
	if u.AccountType != nil {
		updateMap["account_type"] = *u.AccountType
	}
	if u.BranchAddress != nil {
		updateMap["branch_address"] = *u.BranchAddress
	}
	if u.ActiveSwitch != nil {
		updateMap["active_switch"] = *u.ActiveSwitch
	}
	return updateMap
}
