package messagingfx

import (
	jsonserializers "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers/json"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"messagingProvidefx",
	fx.Provide(
		jsonserializers.NewJSONSerializer,
		jsonserializers.NewMessageJSONSerializer,
		jsonserializers.NewMetadataJSONSerializer,
	),
)
