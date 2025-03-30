package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/sync_categories/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
)

type SyncCategories interface {
	SyncCategories(ctx context.Context) error
}

func NewSyncCategories(
	uuidGenerator uuid.UUIDGenerator,
	listCategoriesRepository repositories.ListCategories,
	rabbitmqProducer producer.Producer,
	syncCategoriesMessageBuilder events.SyncCategoriesMessageBuilder,
) SyncCategories {
	return &syncCategoriesImpl{
		uuidGenerator:                uuidGenerator,
		listCategoriesRepository:     listCategoriesRepository,
		rabbitmqProducer:             rabbitmqProducer,
		syncCategoriesMessageBuilder: syncCategoriesMessageBuilder,
	}
}

type syncCategoriesImpl struct {
	uuidGenerator                uuid.UUIDGenerator
	listCategoriesRepository     repositories.ListCategories
	rabbitmqProducer             producer.Producer
	syncCategoriesMessageBuilder events.SyncCategoriesMessageBuilder
}

func (handle *syncCategoriesImpl) SyncCategories(ctx context.Context) error {
	// get current categories
	categories, err := handle.listCategoriesRepository.ListCategories(ctx)
	if err != nil {
		return err
	}

	// publish create book events
	aggregates := aggregates.CategoryEntities(categories).ToAggregate()
	message := handle.syncCategoriesMessageBuilder.Build(aggregates)
	if err := handle.rabbitmqProducer.PublishMessage(ctx, message, nil, nil); err != nil {
		return err
	}

	return nil
}
