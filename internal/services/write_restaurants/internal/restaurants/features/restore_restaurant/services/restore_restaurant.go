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
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/dtos"
	restoreRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type RestoreRestaurant interface {
	RestoreRestaurant(ctx context.Context, req *featuresdtos.RestoreRestaurantRequest) error
}

func NewRestoreRestaurant(
	uuidGenerator uuid.UUIDGenerator,
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction],
	readRestaurantsRepository repositories.ReadRestaurants,
	restoreRestaurantMessageBuilder restoreRestaurantEvents.RestoreRestaurantMessageBuilder,
	messageSerializer serializers.MessageSerializer,
) RestoreRestaurant {
	return &restoreRestaurantImpl{
		uuidGenerator: uuidGenerator,
		updateRestaurantsWithTransactionRepository: updateRestaurantsWithTransactionRepository,
		readRestaurantsRepository:                  readRestaurantsRepository,
		restoreRestaurantMessageBuilder:            restoreRestaurantMessageBuilder,
		messageSerializer:                          messageSerializer,
	}
}

type restoreRestaurantImpl struct {
	uuidGenerator                              uuid.UUIDGenerator
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction]
	readRestaurantsRepository                  repositories.ReadRestaurants
	restoreRestaurantMessageBuilder            restoreRestaurantEvents.RestoreRestaurantMessageBuilder
	messageSerializer                          serializers.MessageSerializer
}

func (handle *restoreRestaurantImpl) RestoreRestaurant(
	ctx context.Context,
	req *featuresdtos.RestoreRestaurantRequest,
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
	if existed.IsActive() {
		return customizeerrors.AlreadyActiveError
	}

	message := handle.restoreRestaurantMessageBuilder.Build(existed)
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
		ActiveStatus: lo.ToPtr(true),
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
