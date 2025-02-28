package dtos

import (
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type Account struct {
	ID              string                `json:"id"`
	MobileNumber    string                `json:"mobile_number"`
	AccountNumber   string                `json:"account_number"`
	AccountTypeCode enums.AccountTypeCode `json:"account_type_code"`
	BranchCode      enums.BanksBranchCode `json:"branch_code"`
	ActiveSwitch    bool                  `json:"active_switch"`
	CommonHistoryModelWithUpdate
}
