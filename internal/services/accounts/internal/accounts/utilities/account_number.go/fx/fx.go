package accountnumberutilitiesfx

import (
	accountnumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/utilities/account_number.go"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"accountnumberutilities",
	fx.Provide(accountnumberutilities.NewAccountNumberGenerator),
)
