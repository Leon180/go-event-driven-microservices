package contextloggersfx

import (
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"go.uber.org/fx"
)

// contextLoggers provide module provide context logger
// dependencies:
// - loggers.Logger
var ProvideModule = fx.Module(
	"contextLoggersProvideFx",
	fx.Provide(
		contextloggers.NewContextLogger,
	),
)
