package accountnumberutilitiesfx

import (
	loannumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/utilities/loan_number"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"loannumberutilities",
	fx.Provide(loannumberutilities.NewLoanNumberGenerator),
)
