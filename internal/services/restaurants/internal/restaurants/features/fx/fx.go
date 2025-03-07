package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/create_restaurant/gin_endpoints"
	createRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/create_restaurant/services"
	deleteRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/delete_restaurant/gin_endpoints"
	deleteRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/delete_restaurant/services"
	getRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/get_restaurant/gin_endpoints"
	getRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/get_restaurant/services"
	restoreRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/restore_restaurant/gin_endpoints"
	restoreRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/restore_restaurant/services"
	searchRestaurantsGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/search_restaurants/gin_endpoints"
	searchRestaurantsServices "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/search_restaurants/services"
	updateRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/update_restaurant/gin_endpoints"
	updateRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/update_restaurant/services"
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
// dependencies:
// - uuid.UUIDGenerator
// - repositories.SearchRestaurantsFullInfo
// - repositories.ReadRestaurant
// - postgres.Transactor[repositories.UpdateRestaurantsWithTransaction]
// - contextloggers.ContextLogger
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
