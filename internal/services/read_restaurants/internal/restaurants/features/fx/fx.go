package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createBookCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/commands"
	createBookGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/gin_endpoints"
	createBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/services"
	createRestaurantCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/commands"
	createRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/gin_endpoints"
	createRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/services"
	deleteBookCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/commands"
	deleteBookGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/gin_endpoints"
	deleteBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/services"
	deleteRestaurantCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/commands"
	deleteRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/gin_endpoints"
	deleteRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/services"
	getRestaurantCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/commands"
	getRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/gin_endpoints"
	getRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/services"
	listCategoriesCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/commands"
	listCategoriesGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/gin_endpoints"
	listCategoriesServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/services"
	restoreRestaurantCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/commands"
	restoreRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/gin_endpoints"
	restoreRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/services"
	searchBooksCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_books/commands"
	searchBooksGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_books/gin_endpoints"
	searchBooksServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_books/services"
	searchRestaurantsCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/commands"
	searchRestaurantsGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/gin_endpoints"
	searchRestaurantsServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/services"
	updateRestaurantCommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/commands"
	updateRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/gin_endpoints"
	updateRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/services"

	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the restaurants features services, commands and endpoints
var ProvideModule = fx.Module(
	"restaurantsFeaturesProvideFx",

	// services
	fx.Provide(
		createRestaurantServices.NewCreateRestaurant,
		deleteRestaurantServices.NewDeleteRestaurant,
		getRestaurantServices.NewGetRestaurant,
		restoreRestaurantServices.NewRestoreRestaurant,
		searchRestaurantsServices.NewSearchRestaurants,
		updateRestaurantServices.NewUpdateRestaurant,
		createBookServices.NewCreateBook,
		deleteBookServices.NewDeleteBook,
		searchBooksServices.NewSearchBooks,
		listCategoriesServices.NewListCategories,
	),

	// commands
	fx.Provide(
		createRestaurantCommands.NewCreateRestaurantHandler,
		deleteRestaurantCommands.NewDeleteRestaurantHandler,
		getRestaurantCommands.NewGetRestaurantHandler,
		restoreRestaurantCommands.NewRestoreRestaurantHandler,
		searchRestaurantsCommands.NewSearchRestaurantsHandler,
		updateRestaurantCommands.NewUpdateRestaurantHandler,
		createBookCommands.NewCreateBookHandler,
		deleteBookCommands.NewDeleteBookHandler,
		searchBooksCommands.NewSearchBooksHandler,
		listCategoriesCommands.NewListCategoriesHandler,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			createRestaurantGinEndpoints.NewCreateRestaurant,
			deleteRestaurantGinEndpoints.NewDeleteRestaurant,
			getRestaurantGinEndpoints.NewGetRestaurant,
			restoreRestaurantGinEndpoints.NewRestoreRestaurant,
			searchRestaurantsGinEndpoints.NewSearchRestaurants,
			updateRestaurantGinEndpoints.NewUpdateRestaurant,
			createBookGinEndpoints.NewCreateBook,
			deleteBookGinEndpoints.NewDeleteBook,
			searchBooksGinEndpoints.NewSearchBooks,
			listCategoriesGinEndpoints.NewListCategories,
		)...,
	),
)

// fxTagEndpoints will tag the endpoints with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagEndpoints(handlers ...any) []any {
	return lo.Map(handlers, func(handler any, _ int) any {
		return fxTagEndpoint(handler)
	})
}

// fxTagEndpoint will tag the endpoint with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagEndpoint(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(customizegin.Endpoint)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupEndpoints.ToString())),
	)
}
