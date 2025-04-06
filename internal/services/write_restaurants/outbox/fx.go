package outbox

import (
	"context"

	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"provideoutbox",
	fx.Provide(
		NewOutboxConfig,
		NewOutboxOperator,
	),
)

var InvokeModule = fx.Module(
	"invokeoutbox",
	fx.Invoke(func(outboxOperator *OutboxOperator) {
		go outboxOperator.Start(context.Background())
	}),
)
