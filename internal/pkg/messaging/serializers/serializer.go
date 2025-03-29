package serializers

import "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"

type Serializer interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
	UnmarshalFromJson(data string, v any) error
	DecodeWithMapStructure(input any, output any) error
	UnmarshalToMap(data []byte, v *map[string]any) error
	UnmarshalToMapFromJson(data string, v *map[string]any) error
	MarshalIndent(v any) string
	ColoredMarshalIndent(v any) string
}

type SerializationResult struct {
	Data        []byte
	ContentType enums.ContentType
}
