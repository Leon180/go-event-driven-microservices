package postgresdbmodels

import (
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
)

type Account struct {
	ID              string                        `gorm:"primaryKey;type:uuid" comment:"ID"`
	MobileNumber    string                        `gorm:"not null;type:varchar(20)" comment:"Mobile Number"`
	AccountNumber   string                        `gorm:"not null;type:varchar(20)" comment:"Account Number"`
	AccountTypeCode enumsaccounts.AccountTypeCode `gorm:"not null;type:int" comment:"Account Type Code"`
	BranchCode      enumsbanks.BanksBranchCode    `gorm:"not null;type:int" comment:"Branch Code"`
	ActiveSwitch    bool                          `gorm:"not null;type:boolean" comment:"Active Switch"`
	CommonHistoryModelWithUpdate
}

func (a *Account) ToDTOs() dtos.Account {
	return dtos.Account{
		ID:              a.ID,
		MobileNumber:    a.MobileNumber,
		AccountNumber:   a.AccountNumber,
		AccountTypeCode: a.AccountTypeCode,
		BranchCode:      a.BranchCode,
		ActiveSwitch:    a.ActiveSwitch,
	}
}

func (a *Account) ToDTOsWithHistory() dtos.AccountWithHistory {
	return dtos.AccountWithHistory{
		Account: a.ToDTOs(),
		History: a.CommonHistoryModelWithUpdate.ToDTOs(),
	}
}
