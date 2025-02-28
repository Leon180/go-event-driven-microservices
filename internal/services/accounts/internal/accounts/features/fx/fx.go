package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/gin_endpoints"
	createAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/services"
	deleteAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/gin_endpoints"
	deleteAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/services"
	getAccountsGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/gin_endpoints"
	getAccountsServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/services"
	restoreAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/restore_account/gin_endpoints"
	restoreAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/restore_account/services"
	updateAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/gin_endpoints"
	updateAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the accounts features
// It provides the services and customizeginendpoints.Endpoint for the accounts features:
// - createAccountServices.CreateAccount
// - getAccountServices.GetAccount
// - updateAccountServices.UpdateAccount
// - deleteAccountServices.DeleteAccount
// - customizeginendpoints.Endpoint(createAccountGinEndpoints.NewCreateAccount)
// - customizeginendpoints.Endpoint(getAccountGinEndpoints.NewGetAccount)
// - customizeginendpoints.Endpoint(updateAccountGinEndpoints.NewUpdateAccount)
// - customizeginendpoints.Endpoint(deleteAccountGinEndpoints.NewDeleteAccount)
// dependencies:
// - accountnumberutilities.AccountNumberGenerator
// - uuid.UUIDGenerator
// - repositories.CreateAccount
// - repositories.GetAccountWithHistoryByMobileNumber
// - repositories.UpdateAccountByID
// - repositories.DeleteAccount
var ProvideModule = fx.Module(
	"accountsFeaturesProvideFx",

	// services
	fx.Provide(
		createAccountServices.NewCreateAccount,
		getAccountsServices.NewGetAccountsByMobileNumber,
		updateAccountServices.NewUpdateAccount,
		deleteAccountServices.NewDeleteAccount,
		restoreAccountServices.NewRestoreAccount,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			createAccountGinEndpoints.NewCreateAccount,
			getAccountsGinEndpoints.NewGetAccountsByMobileNumber,
			updateAccountGinEndpoints.NewUpdateAccount,
			deleteAccountGinEndpoints.NewDeleteAccount,
			restoreAccountGinEndpoints.NewRestoreAccount,
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
