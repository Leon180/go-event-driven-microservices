package repositoriesfx

import (
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories/postgres"
	"go.uber.org/fx"
)

// ProvideModule is the module for the repositories
// It provides the repositories:
// - repositories.CreateAccount
// - repositories.GetAccountWithHistoryByMobileNumber
// - repositories.GetAccountWithHistory
// - repositories.UpdateAccountByID
// - repositories.DeleteAccountByID
// dependencies:
// - *gorm.DB
// - contextloggers.ContextLogger
var ProvideModule = fx.Module(
	"accountsRepositoriesProvideFx",
	fx.Provide(
		repositoriespostgres.NewCreateAccount,
		repositoriespostgres.NewGetAccountWithHistoryByMobileNumber,
		repositoriespostgres.NewGetAccountWithHistory,
		repositoriespostgres.NewUpdateAccountByID,
		repositoriespostgres.NewDeleteAccountByID,
	),
)
