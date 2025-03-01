package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createLoanGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/gin_endpoints"
	createLoanServices "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/services"
	deleteLoanGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/delete_loan/gin_endpoints"
	deleteLoanServices "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/delete_loan/services"
	getLoansGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/gin_endpoints"
	getLoansServices "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/services"
	updateLoanGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/gin_endpoints"
	updateLoanServices "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the cards features
// It provides the services and customizegin.Endpoint for the loans features:
// - createLoanServices.CreateLoan
// - getLoansServices.GetLoansByMobileNumberAndActiveSwitch
// - updateLoanServices.UpdateLoan
// - deleteLoanServices.DeleteLoan
// - customizegin.Endpoint(createLoanGinEndpoints.NewCreateLoan)
// - customizeginendpoints.Endpoint(getLoansGinEndpoints.NewGetLoansByMobileNumberAndActiveSwitch)
// - customizeginendpoints.Endpoint(updateLoanGinEndpoints.NewUpdateLoan)
// - customizeginendpoints.Endpoint(deleteLoanGinEndpoints.NewDeleteLoan)
// dependencies:
// - loannumberutilities.LoanNumberGenerator
// - uuid.UUIDGenerator
// - repositories.CreateLoan
// - repositories.ReadLoanByMobileNumberAndActiveSwitch
// - repositories.ReadLoan
// - repositories.UpdateLoanByID
// - repositories.DeleteLoan
var ProvideModule = fx.Module(
	"cardsFeaturesProvideFx",

	// services
	fx.Provide(
		createLoanServices.NewCreateLoan,
		getLoansServices.NewGetLoansByMobileNumberAndActiveSwitch,
		updateLoanServices.NewUpdateLoan,
		deleteLoanServices.NewDeleteLoan,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			createLoanGinEndpoints.NewCreateLoan,
			getLoansGinEndpoints.NewGetLoansByMobileNumberAndActiveSwitch,
			updateLoanGinEndpoints.NewUpdateLoan,
			deleteLoanGinEndpoints.NewDeleteLoan,
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
