package metadatas

import (
	"github.com/goccy/go-json"
)

func MapToMetadata(data map[string]any) Metadata {
	return Metadata(data)
}

func MetadataToMap(meta Metadata) map[string]any {
	return meta
}

type Metadata map[string]any

func (m Metadata) ExistsKey(key string) bool {
	_, ok := m[key]
	return ok
}

func (m Metadata) Get(key string) any {
	val, ok := m[key]
	if !ok {
		return nil
	}
	return val
}

func (m Metadata) Set(key string, value any) {
	m[key] = value
}

func (m Metadata) Keys() []string {
	i := 0
	r := make([]string, len(m))
	for k := range m {
		r[i] = k
		i++
	}
	return r
}

func (m Metadata) ToJson() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(marshal)
}
