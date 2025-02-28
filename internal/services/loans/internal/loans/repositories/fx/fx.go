package repositoriesfx

import (
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories/postgres"
	"go.uber.org/fx"
)

// ProvideModule is the module for the repositories
// It provides the repositories:
// dependencies:
// - *gorm.DB
var ProvideModule = fx.Module(
	"cardsRepositoriesProvideFx",
	fx.Provide(
		repositoriespostgres.NewCreateLoan,
		repositoriespostgres.NewReadLoan,
		repositoriespostgres.NewReadLoanByMobileNumberAndActiveSwitch,
		repositoriespostgres.NewUpdateLoanByID,
	),
)
