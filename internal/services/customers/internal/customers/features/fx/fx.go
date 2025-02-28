package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createCustomerGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/gin_endpoints"
	createCustomerServices "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/services"
	deleteCustomerGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/delete_customer/gin_endpoints"
	deleteCustomerServices "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/delete_customer/services"
	getCustomerGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/gin_endpoints"
	getCustomerServices "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/services"
	updateCustomerGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/gin_endpoints"
	updateCustomerServices "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the customers features
// It provides the services and customizegin.Endpoint for the customers features:
// - createCustomerServices.CreateCustomer
// - getCustomerServices.GetCustomerByMobileNumberAndActiveSwitch
// - updateCustomerServices.UpdateCustomer
// - deleteCustomerServices.DeleteCustomer
// - customizegin.Endpoint(createCustomerGinEndpoints.NewCreateCustomer)
// - customizeginendpoints.Endpoint(getCustomerGinEndpoints.NewGetCustomerByMobileNumberAndActiveSwitch)
// - customizeginendpoints.Endpoint(updateCustomerGinEndpoints.NewUpdateCustomer)
// - customizeginendpoints.Endpoint(deleteCustomerGinEndpoints.NewDeleteCustomer)
// dependencies:
// - uuid.UUIDGenerator
// - repositories.CreateCustomer
// - repositories.ReadCustomerByMobileNumberAndActiveSwitch
// - repositories.ReadCustomer
// - repositories.UpdateCustomerByID
// - repositories.DeleteCustomer
var ProvideModule = fx.Module(
	"customersFeaturesProvideFx",

	// services
	fx.Provide(
		createCustomerServices.NewCreateCustomer,
		getCustomerServices.NewGetCustomerByMobileNumberAndActiveSwitch,
		updateCustomerServices.NewUpdateCustomer,
		deleteCustomerServices.NewDeleteCustomer,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			createCustomerGinEndpoints.NewCreateCustomer,
			getCustomerGinEndpoints.NewGetCustomerByMobileNumberAndActiveSwitch,
			updateCustomerGinEndpoints.NewUpdateCustomer,
			deleteCustomerGinEndpoints.NewDeleteCustomer,
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
