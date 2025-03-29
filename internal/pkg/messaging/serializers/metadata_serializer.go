package serializers

import "github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"

type MetadataSerializer interface {
	Serialize(meta metadatas.Metadata) ([]byte, error)
	Deserialize(bytes []byte) (metadatas.Metadata, error)
}
