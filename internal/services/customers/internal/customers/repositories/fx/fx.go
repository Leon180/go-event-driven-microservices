package repositoriesfx

import (
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories/postgres"
	"go.uber.org/fx"
)

// ProvideModule is the module for the repositories
// It provides the repositories:
// dependencies:
// - *gorm.DB
var ProvideModule = fx.Module(
	"customersRepositoriesProvideFx",
	fx.Provide(
		repositoriespostgres.NewCreateCustomer,
		repositoriespostgres.NewReadCustomer,
		repositoriespostgres.NewReadCustomerByMobileNumberAndActiveSwitch,
		repositoriespostgres.NewUpdateCustomerByID,
	),
)
