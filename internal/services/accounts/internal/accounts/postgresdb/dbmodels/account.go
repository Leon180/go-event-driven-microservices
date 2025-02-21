package postgresdbmodels

import "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"

type Account struct {
	ID            string `gorm:"primaryKey" comment:"ID"`
	MobileNumber  string `gorm:"not null" comment:"Mobile Number"`
	AccountNumber int64  `gorm:"not null" comment:"Account Number"`
	AccountType   string `gorm:"not null" comment:"Account Type"`
	BranchAddress string `gorm:"not null" comment:"Branch Address"`
	ActiveSwitch  bool   `gorm:"not null" comment:"Active Switch"`
	CommonHistoryModelWithUpdate
}

func (a *Account) ToDTOs() dtos.Account {
	return dtos.Account{
		ID:            a.ID,
		MobileNumber:  a.MobileNumber,
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		BranchAddress: a.BranchAddress,
		ActiveSwitch:  a.ActiveSwitch,
	}
}

func (a *Account) ToDTOsWithHistory() dtos.AccountWithHistory {
	return dtos.AccountWithHistory{
		Account: a.ToDTOs(),
		History: a.CommonHistoryModelWithUpdate.ToDTOs(),
	}
}
