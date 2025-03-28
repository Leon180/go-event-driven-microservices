package json

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
)

func NewMetadataJSONSerializer(serializer serializers.Serializer) serializers.MetadataSerializer {
	return &metadataJSONSerializer{serializer: serializer}
}

type metadataJSONSerializer struct {
	serializer serializers.Serializer
}

func (s *metadataJSONSerializer) Serialize(meta metadatas.Metadata) ([]byte, error) {
	if meta == nil {
		return nil, nil
	}
	return s.serializer.Marshal(meta)
}

func (s *metadataJSONSerializer) Deserialize(bytes []byte) (metadatas.Metadata, error) {
	if bytes == nil {
		return nil, nil
	}
	var meta metadatas.Metadata
	if err := s.serializer.Unmarshal(bytes, &meta); err != nil {
		return nil, err
	}
	return meta, nil
}
