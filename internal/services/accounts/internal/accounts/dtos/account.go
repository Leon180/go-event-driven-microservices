package dtos

import (
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
)

type Account struct {
	ID            string
	MobileNumber  string
	AccountNumber string
	AccountType   enumsaccounts.AccountType
	Branch        enumsbanks.BanksBranch
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
	AccountNumber *string
	AccountType   *enumsaccounts.AccountType
	BranchAddress *enumsbanks.BanksBranch
	ActiveSwitch  *bool
}

func (u *UpdateAccount) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.AccountNumber != nil {
		updateMap["account_number"] = *u.AccountNumber
	}
	if u.AccountType != nil {
		accountTypeCode := u.AccountType.ToAccountTypeCode()
		updateMap["account_type"] = accountTypeCode
	}
	if u.BranchAddress != nil {
		branchCode := u.BranchAddress.ToBanksBranchCode()
		updateMap["branch_address"] = branchCode
	}
	if u.ActiveSwitch != nil {
		updateMap["active_switch"] = *u.ActiveSwitch
	}
	return updateMap
}
