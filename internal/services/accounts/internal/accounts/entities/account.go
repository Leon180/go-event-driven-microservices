package entities

import (
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/samber/lo"
)

type Account struct {
	ID              string                `gorm:"primaryKey;type:uuid" comment:"ID"`
	MobileNumber    string                `gorm:"not null;type:varchar(20)" comment:"Mobile Number"`
	AccountNumber   string                `gorm:"not null;type:varchar(20)" comment:"Account Number"`
	AccountTypeCode enums.AccountTypeCode `gorm:"not null;type:int" comment:"Account Type Code"`
	BranchCode      enums.BanksBranchCode `gorm:"not null;type:int" comment:"Branch Code"`
	ActiveSwitch    bool                  `gorm:"not null;type:boolean" comment:"Active Switch"`
	CommonHistoryModelWithUpdate
}

func (a *Account) IsActive() bool {
	return a.ActiveSwitch
}

type Accounts []Account

func (a Accounts) IncludeAccountTypeCode(mobileNumber string, accountTypeCode enums.AccountTypeCode) bool {
	_, ok := lo.Find(a, func(a Account) bool {
		return a.MobileNumber == mobileNumber && a.AccountTypeCode == accountTypeCode
	})
	return ok
}

type UpdateAccount struct {
	ID            string
	MobileNumber  *string
	AccountNumber *string
	BranchCode    *enums.BanksBranchCode
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
