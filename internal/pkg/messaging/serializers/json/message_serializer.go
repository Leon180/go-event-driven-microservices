package json

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

func NewMessageJSONSerializer(
	serializer serializers.Serializer,
) serializers.MessageSerializer {
	return &messageJSONSerializer{serializer: serializer}
}

type messageJSONSerializer struct {
	serializer serializers.Serializer
}

func (m *messageJSONSerializer) Serialize(message types.Message) (*serializers.SerializationResult, error) {
	return m.SerializeObject(message)
}

func (m *messageJSONSerializer) SerializeObject(message any) (*serializers.SerializationResult, error) {
	if message == nil {
		return &serializers.SerializationResult{Data: nil, ContentType: m.ContentType()}, nil
	}

	data, err := m.serializer.Marshal(message)
	if err != nil {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	return &serializers.SerializationResult{Data: data, ContentType: m.ContentType()}, nil
}

func (m *messageJSONSerializer) Deserialize(
	data []byte,
	messageType string,
	contentType enums.ContentType,
) (types.Message, error) {
	if data == nil {
		return nil, nil
	}

	if contentType != m.ContentType() {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	targetMessagePointer := types.EmptyInstanceByTypeNameAndImplementedInterface[types.Message](messageType)

	if targetMessagePointer == nil {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	if contentType != m.ContentType() {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	if err := m.serializer.Unmarshal(data, targetMessagePointer); err != nil {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	return targetMessagePointer.(types.Message), nil
}

func (m *messageJSONSerializer) DeserializeObject(
	data []byte,
	messageType string,
	contentType enums.ContentType,
) (any, error) {
	if data == nil {
		return nil, nil
	}

	targetMessagePointer := types.InstanceByTypeName(messageType)

	if targetMessagePointer == nil {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	if contentType != m.ContentType() {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	if err := m.serializer.Unmarshal(data, targetMessagePointer); err != nil {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	return targetMessagePointer, nil
}

func (m *messageJSONSerializer) ContentType() enums.ContentType {
	return enums.ContentTypeJSON
}

func (m *messageJSONSerializer) Serializer() serializers.Serializer {
	return m.serializer
}
