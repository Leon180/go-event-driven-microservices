package serializers

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

type MessageSerializer interface {
	Serialize(message types.Message) (*SerializationResult, error)
	SerializeObject(message any) (*SerializationResult, error)
	Deserialize(data []byte, messageType string, contentType enums.ContentType) (types.Message, error)
	DeserializeObject(data []byte, messageType string, contentType enums.ContentType) (any, error)
	ContentType() enums.ContentType
	Serializer() Serializer
}
