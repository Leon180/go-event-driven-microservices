package dtos

import (
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	"github.com/samber/lo"
)

type Account struct {
	ID              string                        `json:"id"`
	MobileNumber    string                        `json:"mobile_number"`
	AccountNumber   string                        `json:"account_number"`
	AccountTypeCode enumsaccounts.AccountTypeCode `json:"account_type_code"`
	BranchCode      enumsbanks.BanksBranchCode    `json:"branch_code"`
	ActiveSwitch    bool                          `json:"active_switch"`
	History         CommonHistoryModelWithUpdate  `json:"history"`
}

func (a *Account) IsActive() bool {
	return a.ActiveSwitch
}

type Accounts []Account

func (a Accounts) IncludeAccountTypeCode(mobileNumber string, accountTypeCode enumsaccounts.AccountTypeCode) bool {
	_, ok := lo.Find(a, func(a Account) bool {
		return a.MobileNumber == mobileNumber && a.AccountTypeCode == accountTypeCode
	})
	return ok
}

type AccountWithHistory struct {
	Account
	History CommonHistoryModelWithUpdate `json:"history"`
}

func (a *AccountWithHistory) IsActive() bool {
	return a.Account.ActiveSwitch
}

type AccountsWithHistory []AccountWithHistory

func (a AccountsWithHistory) IncludeAccountTypeCode(mobileNumber string, accountTypeCode enumsaccounts.AccountTypeCode) bool {
	_, ok := lo.Find(a, func(a AccountWithHistory) bool {
		return a.Account.MobileNumber == mobileNumber && a.Account.AccountTypeCode == accountTypeCode
	})
	return ok
}

type UpdateAccount struct {
	ID            string
	MobileNumber  *string
	AccountNumber *string
	BranchCode    *enumsbanks.BanksBranchCode
	ActiveSwitch  *bool
}

func (u *UpdateAccount) RemoveUnchangedFields(account Account) {
	if u.ID != account.ID {
		return
	}
	if u.MobileNumber != nil && *u.MobileNumber == account.MobileNumber {
		u.MobileNumber = nil
	}
	if u.AccountNumber != nil && *u.AccountNumber == account.AccountNumber {
		u.AccountNumber = nil
	}
	if u.BranchCode != nil && *u.BranchCode == account.BranchCode {
		u.BranchCode = nil
	}
	if u.ActiveSwitch != nil && *u.ActiveSwitch == account.ActiveSwitch {
		u.ActiveSwitch = nil
	}
}

func (u *UpdateAccount) NoUpdates() bool {
	return u.MobileNumber == nil && u.AccountNumber == nil && u.BranchCode == nil && u.ActiveSwitch == nil
}

func (u *UpdateAccount) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.MobileNumber != nil {
		updateMap["mobile_number"] = *u.MobileNumber
	}
	if u.AccountNumber != nil {
		updateMap["account_number"] = *u.AccountNumber
	}
	if u.BranchCode != nil {
		updateMap["branch_code"] = *u.BranchCode
	}
	if u.ActiveSwitch != nil {
		updateMap["active_switch"] = *u.ActiveSwitch
	}
	return updateMap
}
