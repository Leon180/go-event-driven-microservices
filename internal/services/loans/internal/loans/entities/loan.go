package entities

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/shopspring/decimal"
)

type Loan struct {
	ID           string             `gorm:"primaryKey;type:uuid" comment:"ID"`
	MobileNumber string             `gorm:"not null;type:varchar(20)" comment:"Mobile Number"`
	LoanNumber   string             `gorm:"not null;type:varchar(20)" comment:"Loan Number"`
	LoanTypeCode enums.LoanTypeCode `gorm:"not null;type:integer" comment:"Loan Type Code"`
	TotalAmount  decimal.Decimal    `gorm:"type:NUMERIC(20,6)" comment:"Total Amount"`
	PaidAmount   decimal.Decimal    `gorm:"not null;type:NUMERIC(20,6)" comment:"Paid Amount"`
	InterestRate decimal.Decimal    `gorm:"type:NUMERIC(20,6)" comment:"Interest Rate"`
	Term         int                `gorm:"not null;type:integer" comment:"Term"`
	ActiveSwitch bool               `gorm:"not null;type:boolean" comment:"Active Switch"`
	CommonHistoryModelWithUpdate
}

func (l *Loan) IsActive() bool {
	return l.ActiveSwitch
}

type Loans []Loan

type UpdateLoan struct {
	ID           string
	MobileNumber *string
	TotalAmount  *decimal.Decimal
	PaidAmount   *decimal.Decimal
	InterestRate *decimal.Decimal
	Term         *int
	ActiveSwitch *bool
}

func (u *UpdateLoan) RemoveUnchangedFields(loan Loan) {
	if u.ID != loan.ID {
		return
	}
	if u.MobileNumber != nil && *u.MobileNumber == loan.MobileNumber {
		u.MobileNumber = nil
	}
	if u.TotalAmount != nil && u.TotalAmount.Equal(loan.TotalAmount) {
		u.TotalAmount = nil
	}
	if u.PaidAmount != nil && u.PaidAmount.Equal(loan.PaidAmount) {
		u.PaidAmount = nil
	}
	if u.InterestRate != nil && u.InterestRate.Equal(loan.InterestRate) {
		u.InterestRate = nil
	}
	if u.Term != nil && *u.Term == loan.Term {
		u.Term = nil
	}
	if u.ActiveSwitch != nil && *u.ActiveSwitch == loan.ActiveSwitch {
		u.ActiveSwitch = nil
	}
}

func (u *UpdateLoan) NoUpdates() bool {
	return u.MobileNumber == nil && u.TotalAmount == nil && u.PaidAmount == nil && u.InterestRate == nil && u.Term == nil && u.ActiveSwitch == nil
}

func (u *UpdateLoan) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.MobileNumber != nil {
		updateMap["mobile_number"] = *u.MobileNumber
	}
	if u.TotalAmount != nil {
		updateMap["total_amount"] = *u.TotalAmount
	}
	if u.PaidAmount != nil {
		updateMap["paid_amount"] = *u.PaidAmount
	}
	if u.InterestRate != nil {
		updateMap["interest_rate"] = *u.InterestRate
	}
	if u.Term != nil {
		updateMap["term"] = *u.Term
	}
	if u.ActiveSwitch != nil {
		updateMap["active_switch"] = *u.ActiveSwitch
	}
	return updateMap
}
