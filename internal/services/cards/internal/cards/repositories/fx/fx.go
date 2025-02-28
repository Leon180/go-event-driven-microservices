package repositoriesfx

import (
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories/postgres"
	"go.uber.org/fx"
)

// ProvideModule is the module for the repositories
// It provides the repositories:
// dependencies:
// - *gorm.DB
var ProvideModule = fx.Module(
	"cardsRepositoriesProvideFx",
	fx.Provide(
		repositoriespostgres.NewCreateCreditCard,
		repositoriespostgres.NewReadCreditCard,
		repositoriespostgres.NewReadCreditCardByMobileNumberAndActiveSwitch,
		repositoriespostgres.NewUpdateCreditCardByID,
	),
)
