package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createBookEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_book/events"
	createBookGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_book/gin_endpoints"
	createBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_book/services"
	createRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_restaurant/events"
	createRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_restaurant/gin_endpoints"
	createRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_restaurant/services"
	deleteBookEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_book/events"
	deleteBookGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_book/gin_endpoints"
	deleteBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_book/services"
	deleteRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/events"
	deleteRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/gin_endpoints"
	deleteRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/services"
	getRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/get_restaurant/gin_endpoints"
	getRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/get_restaurant/services"
	listCategoriesGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/list_categories/gin_endpoints"
	listCategoriesServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/list_categories/services"
	restoreRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/events"
	restoreRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/gin_endpoints"
	restoreRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/services"
	searchBooksGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/search_books/gin_endpoints"
	searchBooksServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/search_books/services"
	searchRestaurantsGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/search_restaurants/gin_endpoints"
	searchRestaurantsServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/search_restaurants/services"
	updateRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/update_restaurant/events"
	updateRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/update_restaurant/gin_endpoints"
	updateRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/update_restaurant/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the accounts features
// It provides the services and customizeginendpoints.Endpoint for the restaurants features:
// - createRestaurantServices.CreateRestaurant
// - deleteRestaurantServices.DeleteRestaurant
// - getRestaurantServices.GetRestaurant
// - searchRestaurantsServices.SearchRestaurants
// - customizeginendpoints.Endpoint(createRestaurantGinEndpoints.NewCreateRestaurant)
// - customizeginendpoints.Endpoint(deleteRestaurantGinEndpoints.NewDeleteRestaurant)
// - customizeginendpoints.Endpoint(getRestaurantGinEndpoints.NewGetRestaurant)
// - customizeginendpoints.Endpoint(restoreRestaurantGinEndpoints.NewRestoreRestaurant)
// - customizeginendpoints.Endpoint(searchRestaurantsGinEndpoints.NewSearchRestaurants)
// - customizeginendpoints.Endpoint(updateRestaurantGinEndpoints.NewUpdateRestaurant)
// - customizeginendpoints.Endpoint(createBookGinEndpoints.NewCreateBook)
// - customizeginendpoints.Endpoint(deleteBookGinEndpoints.NewDeleteBook)
// - customizeginendpoints.Endpoint(searchBooksGinEndpoints.NewSearchBooks)
// dependencies:
// - uuid.UUIDGenerator
// - repositories.SearchRestaurantsFullInfo
// - repositories.ReadRestaurant
// - repositories.UpdateRestaurantsWithTransaction
// - repositories.SearchBooksFullInfo
// - repositories.ReadBooks
// - postgres.Transactor[repositories.UpdateRestaurantsWithTransaction]
// - postgres.Transactor[repositories.UpdateBooksWithTransaction]
// - contextloggers.ContextLogger
var ProvideModule = fx.Module(
	"restaurantsFeaturesProvideFx",

	// events
	fx.Provide(
		createBookEvents.NewCreateBookMessageBuilder,
		createRestaurantEvents.NewCreateRestaurantMessageBuilder,
		deleteBookEvents.NewDeleteBookMessageBuilder,
		deleteRestaurantEvents.NewDeleteRestaurantMessageBuilder,
		restoreRestaurantEvents.NewRestoreRestaurantMessageBuilder,
		updateRestaurantEvents.NewUpdateRestaurantMessageBuilder,
	),

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
