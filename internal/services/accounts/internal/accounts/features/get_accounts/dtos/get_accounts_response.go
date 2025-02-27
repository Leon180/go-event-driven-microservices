package featuresdtos

import (
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
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

type AccountEntity entities.Account

func (a *AccountEntity) ToGetAccountsResponse() *GetAccountsResponse {
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
		History: dtos.CommonHistoryModelWithUpdate{
			CommonHistoryModel: dtos.CommonHistoryModel{
				CreatedAt: a.CreatedAt,
				DeletedAt: a.DeletedAt,
			},
			UpdatedAt: a.UpdatedAt,
		},
	}
}

type AccountsEntities []entities.Account

func (a AccountsEntities) ToGetAccountsResponse() []GetAccountsResponse {
	if a == nil {
		return nil
	}
	return lo.Map(a, func(a entities.Account, _ int) GetAccountsResponse {
		accountResponse := AccountEntity(a)
		return *accountResponse.ToGetAccountsResponse()
	})
}
