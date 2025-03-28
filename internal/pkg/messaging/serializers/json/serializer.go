package json

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/TylerBrock/colorjson"
	"github.com/goccy/go-json"
	"github.com/mitchellh/mapstructure"
)

func NewJSONSerializer() serializers.Serializer {
	return &jsonSerializer{}
}

type jsonSerializer struct{}

func (s *jsonSerializer) Marshal(v any) ([]byte, error) {
	return marshal(v)
}

func (s *jsonSerializer) Unmarshal(data []byte, v any) error {
	return unmarshal(data, v)
}

func (s *jsonSerializer) UnmarshalFromJson(data string, v any) error {
	return unmarshalFromJSON(data, v)
}

func (s *jsonSerializer) DecodeWithMapStructure(input, output any) error {
	return decodeWithMapStructure(input, output)
}

func (s *jsonSerializer) UnmarshalToMap(data []byte, v *map[string]any) error {
	return unmarshalToMap(data, v)
}

func (s *jsonSerializer) UnmarshalToMapFromJson(data string, v *map[string]any) error {
	return unmarshalToMapFromJSON(data, v)
}

func (s *jsonSerializer) MarshalIndent(data any) string {
	return marshalIndent(data)
}

func (s *jsonSerializer) ColoredMarshalIndent(data any) string {
	return coloredMarshalIndent(data)
}

func marshalIndent(data any) string {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return ""
	}
	return string(val)
}

func coloredMarshalIndent(data any) string {
	var obj map[string]any
	err := json.Unmarshal([]byte(marshalIndent(data)), &obj)
	if err != nil {
		return ""
	}
	f := colorjson.NewFormatter()
	f.Indent = 4
	val, err := f.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(val)
}

func marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func unmarshalFromJSON(data string, v any) error {
	return unmarshal([]byte(data), v)
}

func decodeWithMapStructure(input any, output any) error {
	return mapstructure.Decode(input, output)
}

func unmarshalToMapFromJSON(data string, v *map[string]any) error {
	return unmarshalToMap([]byte(data), v)
}

func unmarshalToMap(data []byte, v *map[string]any) error {
	return json.Unmarshal(data, v)
}
