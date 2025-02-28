package dtos

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/shopspring/decimal"
)

type Loan struct {
	ID           string          `json:"id"`
	MobileNumber string          `json:"mobile_number"`
	LoanNumber   string          `json:"loan_number"`
	LoanType     enums.LoanType  `json:"loan_type"`
	TotalAmount  decimal.Decimal `json:"total_amount"`
	PaidAmount   decimal.Decimal `json:"paid_amount"`
	InterestRate decimal.Decimal `json:"interest_rate"`
	Term         int             `json:"term"`
	ActiveSwitch bool            `json:"active_switch"`
	CommonHistoryModelWithUpdate
}
