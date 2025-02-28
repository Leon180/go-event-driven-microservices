package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createCreditCardGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/gin_endpoints"
	createCreditCardServices "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/services"
	deleteCreditCardGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/delete_card/gin_endpoints"
	deleteCreditCardServices "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/delete_card/services"
	getCreditCardsGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/gin_endpoints"
	getCreditCardServices "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/services"
	updateCreditCardGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/gin_endpoints"
	updateCreditCardServices "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the cards features
// It provides the services and customizegin.Endpoint for the cards features:
// - createCreditCardServices.CreateCreditCard
// - getCreditCardServices.GetCreditCardsByMobileNumberAndActiveSwitch
// - updateCreditCardServices.UpdateCreditCard
// - deleteCreditCardServices.DeleteCreditCard
// - customizegin.Endpoint(createCreditCardGinEndpoints.NewCreateCreditCard)
// - customizeginendpoints.Endpoint(getCreditCardsGinEndpoints.NewGetCreditCardsByMobileNumberAndActiveSwitch)
// - customizeginendpoints.Endpoint(updateCreditCardGinEndpoints.NewUpdateCreditCard)
// - customizeginendpoints.Endpoint(deleteCreditCardGinEndpoints.NewDeleteCreditCard)
// dependencies:
// - cardnumberutilities.CardNumberGenerator
// - uuid.UUIDGenerator
// - repositories.CreateCreditCard
// - repositories.ReadCreditCardByMobileNumberAndActiveSwitch
// - repositories.ReadCreditCard
// - repositories.UpdateCreditCardByID
// - repositories.DeleteCreditCard
var ProvideModule = fx.Module(
	"cardsFeaturesProvideFx",

	// services
	fx.Provide(
		createCreditCardServices.NewCreateCreditCard,
		getCreditCardServices.NewGetCreditCardsByMobileNumberAndActiveSwitch,
		updateCreditCardServices.NewUpdateCreditCard,
		deleteCreditCardServices.NewDeleteCreditCard,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			createCreditCardGinEndpoints.NewCreateCreditCard,
			getCreditCardsGinEndpoints.NewGetCreditCardsByMobileNumberAndActiveSwitch,
			updateCreditCardGinEndpoints.NewUpdateCreditCard,
			deleteCreditCardGinEndpoints.NewDeleteCreditCard,
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
