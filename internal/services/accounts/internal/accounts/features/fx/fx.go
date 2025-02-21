package featuresfx

import (
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	createAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/gin_endpoints"
	createAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/services"
	deleteAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/gin_endpoints"
	deleteAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/services"
	getAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/gin_endpoints"
	getAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/services"
	updateAccountGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/gin_endpoints"
	updateAccountServices "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/services"
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
		getAccountServices.NewGetAccount,
		updateAccountServices.NewUpdateAccount,
		deleteAccountServices.NewDeleteAccount,
	),

	// endpoints
	fx.Provide(
		customizeginendpoints.FxTagEndpoint(
			createAccountGinEndpoints.NewCreateAccount,
		),
		customizeginendpoints.FxTagEndpoint(
			getAccountGinEndpoints.NewGetAccount,
		),
		customizeginendpoints.FxTagEndpoint(
			updateAccountGinEndpoints.NewUpdateAccount,
		),
		customizeginendpoints.FxTagEndpoint(
			deleteAccountGinEndpoints.NewDeleteAccount,
		),
	),
)
