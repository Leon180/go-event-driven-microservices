package dtos

import (
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
)

type Account struct {
	ID              string                        `json:"id"`
	MobileNumber    string                        `json:"mobile_number"`
	AccountNumber   string                        `json:"account_number"`
	AccountTypeCode enumsaccounts.AccountTypeCode `json:"account_type_code"`
	BranchCode      enumsbanks.BanksBranchCode    `json:"branch_code"`
	ActiveSwitch    bool                          `json:"active_switch"`
	CommonHistoryModelWithUpdate
}
