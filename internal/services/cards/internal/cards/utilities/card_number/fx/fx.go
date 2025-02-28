package accountnumberutilitiesfx

import (
	cardnumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/utilities/card_number"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"cardnumberutilities",
	fx.Provide(cardnumberutilities.NewCardNumberGenerator),
)
