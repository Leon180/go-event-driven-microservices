package json

import (
	"reflect"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

func NewMessageJSONSerializer(
	serializer serializers.Serializer,
	typeMaker types.TypeMaker,
) serializers.MessageSerializer {
	return &messageJSONSerializer{serializer: serializer, typeMaker: typeMaker}
}

type messageJSONSerializer struct {
	serializer serializers.Serializer
	typeMaker  types.TypeMaker
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
		return nil, err
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

	instance, err := m.typeMaker.GetTypeInstance(messageType)
	if err != nil {
		return nil, err
	}

	if reflect.TypeOf(instance).Kind() != reflect.Ptr {
		if err := m.serializer.Unmarshal(data, &instance); err != nil {
			return nil, err
		}
	} else {
		if err := m.serializer.Unmarshal(data, instance); err != nil {
			return nil, err
		}
	}

	if message, ok := instance.(types.Message); ok {
		return message, nil
	}

	return nil, customizeerrors.MessageTypeInvalidError
}

func (m *messageJSONSerializer) DeserializeObject(
	data []byte,
	messageType string,
	contentType enums.ContentType,
) (any, error) {
	if data == nil {
		return nil, nil
	}

	if contentType != m.ContentType() {
		return nil, customizeerrors.MessageTypeInvalidError
	}

	instance, err := m.typeMaker.GetTypeInstance(messageType)
	if err != nil {
		return nil, err
	}

	if reflect.TypeOf(instance).Kind() != reflect.Ptr {
		if err := m.serializer.Unmarshal(data, &instance); err != nil {
			return nil, err
		}
	} else {
		if err := m.serializer.Unmarshal(data, instance); err != nil {
			return nil, err
		}
	}

	return instance, nil
}

func (m *messageJSONSerializer) ContentType() enums.ContentType {
	return enums.ContentTypeJSON
}

func (m *messageJSONSerializer) Serializer() serializers.Serializer {
	return m.serializer
}
