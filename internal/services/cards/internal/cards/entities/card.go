package entities

import (
	"github.com/shopspring/decimal"
)

type CreditCard struct {
	ID           string          `gorm:"primaryKey;type:uuid" comment:"ID"`
	CardNumber   string          `gorm:"not null;type:varchar(20)" comment:"Card Number"`
	MobileNumber string          `gorm:"not null;type:varchar(20)" comment:"Mobile Number"`
	TotalLimit   decimal.Decimal `gorm:"type:NUMERIC(20,6)" comment:"Total Limit"`
	AmountUsed   decimal.Decimal `gorm:"not null;type:NUMERIC(20,6)" comment:"Amount Used"`
	ActiveSwitch bool            `gorm:"not null;type:boolean" comment:"Active Switch"`
	CommonHistoryModelWithUpdate
}

func (c *CreditCard) IsActive() bool {
	return c.ActiveSwitch
}

type CreditCards []CreditCard

type UpdateCreditCard struct {
	ID           string
	MobileNumber *string
	TotalLimit   *decimal.Decimal
	AmountUsed   *decimal.Decimal
	ActiveSwitch *bool
}

func (u *UpdateCreditCard) RemoveUnchangedFields(creditCard CreditCard) {
	if u.ID != creditCard.ID {
		return
	}
	if u.MobileNumber != nil && *u.MobileNumber == creditCard.MobileNumber {
		u.MobileNumber = nil
	}
	if u.TotalLimit != nil && u.TotalLimit.Equal(creditCard.TotalLimit) {
		u.TotalLimit = nil
	}
	if u.AmountUsed != nil && u.AmountUsed.Equal(creditCard.AmountUsed) {
		u.AmountUsed = nil
	}
	if u.ActiveSwitch != nil && *u.ActiveSwitch == creditCard.ActiveSwitch {
		u.ActiveSwitch = nil
	}
}

func (u *UpdateCreditCard) NoUpdates() bool {
	return u.MobileNumber == nil && u.TotalLimit == nil && u.AmountUsed == nil && u.ActiveSwitch == nil
}

func (u *UpdateCreditCard) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.MobileNumber != nil {
		updateMap["mobile_number"] = *u.MobileNumber
	}
	if u.TotalLimit != nil {
		updateMap["total_limit"] = *u.TotalLimit
	}
	if u.AmountUsed != nil {
		updateMap["amount_used"] = *u.AmountUsed
	}
	if u.ActiveSwitch != nil {
		updateMap["active_switch"] = *u.ActiveSwitch
	}
	return updateMap
}
