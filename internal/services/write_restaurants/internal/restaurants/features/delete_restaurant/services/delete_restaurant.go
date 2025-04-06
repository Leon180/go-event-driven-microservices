package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/dtos"
	deleteRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type DeleteRestaurant interface {
	DeleteRestaurant(ctx context.Context, req *featuresdtos.DeleteRestaurantRequest) error
}

func NewDeleteRestaurant(
	uuidGenerator uuid.UUIDGenerator,
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction],
	readRestaurantsRepository repositories.ReadRestaurants,
	deleteRestaurantMessageBuilder deleteRestaurantEvents.DeleteRestaurantMessageBuilder,
	messageSerializer serializers.MessageSerializer,
) DeleteRestaurant {
	return &deleteRestaurantImpl{
		uuidGenerator: uuidGenerator,
		updateRestaurantsWithTransactionRepository: updateRestaurantsWithTransactionRepository,
		readRestaurantsRepository:                  readRestaurantsRepository,
		deleteRestaurantMessageBuilder:             deleteRestaurantMessageBuilder,
		messageSerializer:                          messageSerializer,
	}
}

type deleteRestaurantImpl struct {
	uuidGenerator                              uuid.UUIDGenerator
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction]
	readRestaurantsRepository                  repositories.ReadRestaurants
	deleteRestaurantMessageBuilder             deleteRestaurantEvents.DeleteRestaurantMessageBuilder
	messageSerializer                          serializers.MessageSerializer
}

func (handle *deleteRestaurantImpl) DeleteRestaurant(
	ctx context.Context,
	req *featuresdtos.DeleteRestaurantRequest,
) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}

	existed, err := handle.readRestaurantsRepository.ReadRestaurantFullInfo(ctx, req.ID)
	if err != nil {
		return err
	}
	if existed == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if !existed.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}

	message := handle.deleteRestaurantMessageBuilder.Build(existed)
	serializationResult, err := handle.messageSerializer.Serialize(message)
	if err != nil {
		return err
	}
	t := time.Now()
	outboxMessage := entities.OutboxMessage{
		ID:        handle.uuidGenerator.GenerateUUID(),
		MessageID: message.ID(),
		Type:      message.Type(),
		Payload:   serializationResult.Data,
		Status:    enums.OutboxStatusPending,
		CreatedAt: t,
		UpdatedAt: t,
	}

	tx, err := handle.updateRestaurantsWithTransactionRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	updateRestaurant := entities.UpdateRestaurant{
		ID:           existed.ID,
		ActiveStatus: lo.ToPtr(false),
	}
	if err := tx.UpdateRestaurant(ctx, &updateRestaurant); err != nil {
		return err
	}

	// store outbox message
	if err := tx.CreateOutboxMessages(ctx, entities.OutboxMessages{outboxMessage}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
