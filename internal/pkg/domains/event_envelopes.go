package domain

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
)

type EventEnvelope struct {
	EventData any
	Metadata  metadatas.Metadata
}
