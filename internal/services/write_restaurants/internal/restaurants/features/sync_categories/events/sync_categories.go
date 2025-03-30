package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
)

type SyncCategories struct {
	*types.MessageImpl
	aggregates.Categories
}

type SyncCategoriesMessageBuilder interface {
	Build(categories aggregates.Categories) *SyncCategories
}

func NewSyncCategoriesMessageBuilder(uuidGenerator uuid.UUIDGenerator) SyncCategoriesMessageBuilder {
	return &syncCategoriesMessageBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type syncCategoriesMessageBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *syncCategoriesMessageBuilderImpl) Build(categories aggregates.Categories) *SyncCategories {
	return &SyncCategories{
		Categories:  categories,
		MessageImpl: types.NewMessageImpl(b.uuidGenerator.GenerateUUID(), "SyncCategories"),
	}
}
