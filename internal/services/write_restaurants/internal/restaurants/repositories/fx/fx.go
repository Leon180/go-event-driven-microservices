package repositoriesfx

import (
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories/postgres"
	"go.uber.org/fx"
)

// ProvideModule is the module for the repositories
// It provides the repositories:
// - repositories.SearchRestaurantsFullInfo
// - repositories.ReadRestaurants
// - repositories.UpdateRestaurants
// - repositories.UpdateRestaurantsWithTransaction
// - repositories.ListCategories
// dependencies:
// - *gorm.DB
// - contextloggers.ContextLogger
var ProvideModule = fx.Module(
	"restaurantsRepositoriesProvideFx",
	fx.Provide(
		repositoriespostgres.NewSearchRestaurantsFullInfo,
		repositoriespostgres.NewReadRestaurants,
		repositoriespostgres.NewUpdateRestaurants,
		repositoriespostgres.NewUpdateRestaurantsWithTransaction,
		repositoriespostgres.NewSearchBooksFullInfo,
		repositoriespostgres.NewReadBooks,
		repositoriespostgres.NewUpdateBooks,
		repositoriespostgres.NewUpdateBooksWithTransaction,
		repositoriespostgres.NewListCategories,
	),
)
