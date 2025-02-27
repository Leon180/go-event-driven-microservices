package featuresdtos

import (
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	"github.com/samber/lo"
)

type GetAccountsResponse struct {
	ID            string                            `json:"id"`
	MobileNumber  string                            `json:"mobile_number"`
	AccountNumber string                            `json:"account_number"`
	AccountType   enumsaccounts.AccountType         `json:"account_type"`
	Branch        enumsbanks.BanksBranch            `json:"branch"`
	ActiveSwitch  bool                              `json:"active_switch"`
	History       dtos.CommonHistoryModelWithUpdate `json:"history"`
}

type AccountWithHistory dtos.AccountWithHistory

func (a *AccountWithHistory) ToGetAccountsResponse() *GetAccountsResponse {
	if a == nil {
		return nil
	}
	return &GetAccountsResponse{
		ID:            a.ID,
		MobileNumber:  a.MobileNumber,
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountTypeCode.ToAccountType(),
		Branch:        a.BranchCode.ToBanksBranch(),
		ActiveSwitch:  a.ActiveSwitch,
		History:       a.History,
	}
}

type AccountsWithHistory []dtos.AccountWithHistory

func (a AccountsWithHistory) ToGetAccountsResponse() []GetAccountsResponse {
	if a == nil {
		return nil
	}
	return lo.Map(a, func(a dtos.AccountWithHistory, _ int) GetAccountsResponse {
		accountResponse := AccountWithHistory(a)
		return *accountResponse.ToGetAccountsResponse()
	})
}
