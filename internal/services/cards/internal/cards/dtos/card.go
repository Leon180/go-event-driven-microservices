package dtos

import "github.com/shopspring/decimal"

type CreditCard struct {
	ID           string           `json:"id"`
	CardNumber   string           `json:"card_number"`
	MobileNumber string           `json:"mobile_number"`
	TotalLimit   *decimal.Decimal `json:"total_limit"`
	AmountUsed   decimal.Decimal  `json:"amount_used"`
	ActiveSwitch bool             `json:"active_switch"`
	CommonHistoryModelWithUpdate
}
