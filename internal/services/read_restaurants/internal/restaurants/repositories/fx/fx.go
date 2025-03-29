package repositoriesfx

import (
	repositoriesmongo "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories/mongo"
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories/postgres"
	repositoriesredis "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories/redis"
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

		repositoriesredis.NewReadRestaurantsRedis,
		repositoriesredis.NewSetRestaurantsRedis,
		repositoriesredis.NewListCategoriesRedis,
		repositoriesredis.NewSetCategoriesRedis,
		repositoriesredis.NewReadBooksRedis,
		repositoriesredis.NewSetBookRedis,

		repositoriesmongo.NewSearchRestaurantsMongo,
		repositoriesmongo.NewReadRestaurantsMongo,
		repositoriesmongo.NewUpdateRestaurantsMongo,
		repositoriesmongo.NewListCategoriesMongo,
		repositoriesmongo.NewSearchBooksMongo,
		repositoriesmongo.NewReadBooksMongo,
		repositoriesmongo.NewUpdateBooksMongo,
		repositoriesmongo.NewSyncCategoriesMongo,
	),
)
