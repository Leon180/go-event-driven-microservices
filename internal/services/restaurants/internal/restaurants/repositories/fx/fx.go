package repositoriesfx

import (
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories/postgres"
	"go.uber.org/fx"
)

// ProvideModule is the module for the repositories
// It provides the repositories:
// - repositories.SearchRestaurantsFullInfo
// - repositories.Restaurants
// - repositories.Branches
// - repositories.Addresses
// - repositories.PriceRanges
// - repositories.BranchCategoryRelations
// - repositories.Tables
// - repositories.TableAvailables
// dependencies:
// - *gorm.DB
// - contextloggers.ContextLogger
var ProvideModule = fx.Module(
	"restaurantsRepositoriesProvideFx",
	fx.Provide(
		repositoriespostgres.NewSearchRestaurantsFullInfoRepository,
		repositoriespostgres.NewRestaurantsRepository,
		repositoriespostgres.NewBranchesRepository,
		repositoriespostgres.NewAddressesRepository,
		repositoriespostgres.NewPriceRangesRepository,
		repositoriespostgres.NewBranchCategoryRelationsRepository,
		repositoriespostgres.NewTablesRepository,
		repositoriespostgres.NewTableAvailablesRepository,
	),
)
